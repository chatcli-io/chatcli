package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chatcli",
	Short: "ChatCLI interacts with ChatGPT and returns a response",
	Long:  `ChatCLI is a command line tool that interacts with ChatGPT and returns a response.`,
}

func Execute() error {
	return rootCmd.Execute()
}