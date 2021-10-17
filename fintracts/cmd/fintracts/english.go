package main

import (
	"fmt"
	"os"

	"github.com/hacdias/fintracts/cli/english"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(englishCmd)
	englishCmd.Flags().StringP("input", "i", "", "indicates the input file")
	englishCmd.MarkFlagRequired("input")
}

var englishCmd = &cobra.Command{
	Use:   "english",
	Short: "parses and validates an English financial contract",
	Run: func(cmd *cobra.Command, args []string) {
		input := mustGetString(cmd, "input")

		data, err := os.ReadFile(input)
		checkErr(err)

		contract, err := english.Parse(data)
		checkErr(err)

		str, err := contract.String()
		checkErr(err)

		fmt.Println(str)
	},
}
