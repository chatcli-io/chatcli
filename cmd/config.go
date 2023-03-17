package cmd

import (
	"bufio"
	"chatcli/control"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the tool.",
	Long:  `Configure the ChatCLI tool by setting the OpenAI API key, OpenAI model, pre-injection, and post-injection.`,
	Run: func(cmd *cobra.Command, args []string) {
		env := control.NewEnv()
		defaultCfg := control.NewDefaultConfig(env)
		cfg := control.NewConfig(env, defaultCfg)

		// Prompt for inputs
		reader := bufio.NewReader(os.Stdin)

		// API token
		fmt.Printf("Your OpenAI API token - enter to keep (current: %s): ", cfg.OpenAIKey)
		openaiToken, _ := reader.ReadString('\n')
		openaiToken = strings.TrimSpace(openaiToken)
		if openaiToken != "" {
			cfg.OpenAIKey = openaiToken
		}

		// model
		user_prompt := ""
		if cfg.OpenAIModel == "" {
			user_prompt = "The OpenAI model to use - enter for default (default: %s): "
			fmt.Printf(user_prompt, defaultCfg.DefaultAIModel)
		} else {
			user_prompt = "The OpenAI model to use - enter to keep (current: %s): "
			fmt.Printf(user_prompt, cfg.OpenAIModel)
		}

		openaiModel, _ := reader.ReadString('\n')
		openaiModel = strings.TrimSpace(openaiModel)
		if openaiModel != "" {
			cfg.OpenAIModel = openaiModel
		}

		if cfg.OpenAIModel == "" {
			cfg.OpenAIModel = defaultCfg.DefaultAIModel
		}

		// pre injection string
		if cfg.PreInjection == "" {
			user_prompt = "Type your pre-injection string - enter for default (default: %s): "
			fmt.Printf(user_prompt, defaultCfg.DefaultPreInject)
		} else {
			user_prompt = "Type your pre-injection string - enter to keep (current: %s): "
			fmt.Printf(user_prompt, cfg.PreInjection)
		}

		preInjection, _ := reader.ReadString('\n')
		preInjection = strings.TrimSpace(preInjection)
		if preInjection != "" {
			cfg.PreInjection = preInjection
		}

		if cfg.PreInjection == "" {
			cfg.PreInjection = defaultCfg.DefaultPreInject
		}

		// post injection string
		if cfg.PostInjection == "" {
			user_prompt = "Type your post-injection string - enter for default (default: %s): "
			fmt.Printf(user_prompt, defaultCfg.DefaultPostInject)
		} else {
			user_prompt = "Type your post-injection string - enter to keep (current: %s): "
			fmt.Printf(user_prompt, cfg.PostInjection)
		}

		postInjection, _ := reader.ReadString('\n')
		postInjection = strings.TrimSpace(postInjection)
		if postInjection != "" {
			cfg.PostInjection = postInjection
		}

		if cfg.PostInjection == "" {
			cfg.PostInjection = defaultCfg.DefaultPostInject
		}

		// Save the new configuration to the config file
		data, err := yaml.Marshal(cfg)
		if err != nil {
			fmt.Printf("Error marshalling config: %v\n", err)
			os.Exit(1)
		}

		configDirPath := path.Dir(defaultCfg.ConfigPath)

		// Check if the directory exists
		_, err = os.Stat(configDirPath)
		if os.IsNotExist(err) {
			// Create the directory if it doesn't exist yet
			errDir := os.MkdirAll(configDirPath, 0755)
			if errDir != nil {
				fmt.Printf("Error creating directory: %v\n", err)
				os.Exit(1)
			}
		}

		if err = ioutil.WriteFile(defaultCfg.ConfigPath, data, 0644); err != nil {
			fmt.Printf("Error writing config file: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
