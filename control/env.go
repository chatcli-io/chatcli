package control

import (
	"os"
)

type Env struct {
	OpenAIKey     string
	OpenAIModel   string
	ConfigPath    string
	PreInjection  string
	PostInjection string
}

func NewEnv() *Env {
	return &Env{
		OpenAIKey:     os.Getenv("OPENAI_API_KEY"),
		OpenAIModel:   os.Getenv("OPENAI_MODEL"),
		ConfigPath:    os.Getenv("CHATCLI_CONFIG"),
		PreInjection:  os.Getenv("CHATCLI_PRE_INJECTION"),
		PostInjection: os.Getenv("CHATCLI_POST_INJECTION"),
	}
}
