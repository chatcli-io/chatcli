package control

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	OpenAIKey     string
	OpenAIModel   string
	PreInjection  string
	PostInjection string
}

func NewConfig(env *Env, defaultConfig *DefaultConfig) *Config {
	configPath := defaultConfig.ConfigPath
	var cfg Config
	cfgEnv := Config{
		OpenAIKey:     env.OpenAIKey,
		OpenAIModel:   env.OpenAIModel,
		PreInjection:  env.PreInjection,
		PostInjection: env.PostInjection,
	}
	configExists := false
	if _, err := os.Stat(configPath); err == nil {
		configExists = true
	}

	if configExists {
		file, err := os.Open(configPath)
		if err != nil {
			fmt.Printf("Error reading config file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Printf("Error reading config file: %v\n", err)
			os.Exit(1)
		}

		if err = yaml.Unmarshal(data, &cfg); err != nil {
			fmt.Printf("Error unmarshalling config file: %v\n", err)
			os.Exit(1)
		}

		// openai api key
		if cfgEnv.OpenAIKey != "" {
			cfg.OpenAIKey = cfgEnv.OpenAIKey
		}

		// openai model
		if cfgEnv.OpenAIModel != "" {
			cfg.OpenAIModel = cfgEnv.OpenAIModel
		}

		// pre injection
		if cfgEnv.PreInjection != "" {
			cfg.PreInjection = cfgEnv.PreInjection
		}

		// post injection
		if cfgEnv.PostInjection != "" {
			cfg.PostInjection = cfgEnv.PostInjection
		}

	} else {
		cfg = cfgEnv
	}

	return &cfg
}

type DefaultConfig struct {
	ConfigPath        string
	DefaultAIModel    string
	DefaultPreInject  string
	DefaultPostInject string
}

func NewDefaultConfig(env *Env) *DefaultConfig {
	cfg := &DefaultConfig{
		ConfigPath:        env.ConfigPath,
		DefaultAIModel:    "gpt-3.5-turbo",
		DefaultPreInject:  "Just respond with a command that can be used in a bash based terminal and achieves a result that matches the following description:",
		DefaultPostInject: "If this description does not make sense as a command, reply with 'Can't generate a command from that.'",
	}

	if cfg.ConfigPath == "" {
		cfg.ConfigPath = "~/.config/chatcli/config.yml"
	}
	if cfg.ConfigPath[0] == '~' {
		cfg.ConfigPath = filepath.Join(os.Getenv("HOME"), cfg.ConfigPath[1:])
	}
	if !filepath.IsAbs(cfg.ConfigPath) {
		panic(errors.New("You have to use an absolute path. Using ~ for home is allowed."))
	}

	return cfg
}
