package main

import (
	"fmt"
	"os"

	fintracts "github.com/hacdias/fintracts/cli"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.Flags().StringP("input", "i", "", "indicates the input file")
	validateCmd.Flags().BoolP("fix", "f", false, "fix auto-fixable errors")
	validateCmd.MarkFlagRequired("input")
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validates a JSON financial contract",
	Run: func(cmd *cobra.Command, args []string) {
		input := mustGetString(cmd, "input")
		fix := mustGetBool(cmd, "fix")

		data, err := os.ReadFile(input)
		checkErr(err)

		contract, err := fintracts.FromJSON(data)
		checkErr(err)

		err = fintracts.Validate(contract, fix)
		checkErr(err)

		str, err := contract.String()
		checkErr(err)

		fmt.Println(str)
	},
}
