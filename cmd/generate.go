package cmd

import (
	"bytes"
	"chatcli/control"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate [prompt]",
	Aliases: []string{"gen"},
	Short:   "Send a prompt to ChatGPT and receive a response",
	Long: `Send a prompt to ChatGPT and receive a response that respects your pre- and post-injection.
Example usage: chatcli generate "Find a file with a specific name"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prompt := args[0]

		env := control.NewEnv()
		defaultCfg := control.NewDefaultConfig(env)
		cfg := control.NewConfig(env, defaultCfg)

		injectedPrompt := fmt.Sprintf("%s %s \n%s", cfg.PreInjection, prompt, cfg.PostInjection)

		// fmt.Print(injectedPrompt)

		payload := map[string]interface{}{
			"model": cfg.OpenAIModel,
			"messages": []map[string]string{
				{"role": "user", "content": injectedPrompt},
			},
			"temperature": 0.7,
		}

		data, err := json.Marshal(payload)
		if err != nil {
			fmt.Printf("Error marshalling payload: %v\n", err)
			os.Exit(1)
		}

		url := "https://api.openai.com/v1/chat/completions"

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			os.Exit(1)
		}

		if cfg.OpenAIKey == "" {
			fmt.Println("Can't send any requests without API token run:")
			fmt.Println("$ chatcli config")
			os.Exit(1)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+cfg.OpenAIKey)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error sending request: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %v\n", err)
			os.Exit(1)
		}

		var jsonResponse map[string]interface{}
		if err := json.Unmarshal(body, &jsonResponse); err != nil {
			fmt.Printf("Error unmarshalling JSON: %v\n", err)
			os.Exit(1)
		}

		choices := jsonResponse["choices"].([]interface{})
		choice := choices[0].(map[string]interface{})
		message := choice["message"].(map[string]interface{})
		content := message["content"].(string)

		fmt.Printf("$ %s\n", strings.TrimSpace(content))

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
