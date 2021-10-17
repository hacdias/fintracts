package main

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "fintracts",
	Short: "Fintracts CLI Tool for Parsing and Validation",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func main() {
	rootCmd.Execute()
}
