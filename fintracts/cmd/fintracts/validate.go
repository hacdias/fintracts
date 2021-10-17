package main

import (
	"fmt"

	fintracts "github.com/hacdias/fintracts/cli"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.Flags().StringP("input", "i", "", "indicates the input file")
	validateCmd.Flags().BoolP("fix", "f", false, "fix auto-fixable errors")
}

var validateCmd = &cobra.Command{
	Use:   "validate [< file.json]",
	Short: "Validates a contract in the JSON specification",
	Run: func(cmd *cobra.Command, args []string) {
		data := inputFlagOrStdin(cmd)
		fix := mustGetBool(cmd, "fix")

		contract, err := fintracts.FromJSON(data)
		checkErr(err)

		err = fintracts.Validate(contract, fix)
		checkErr(err)

		str, err := contract.String()
		checkErr(err)

		fmt.Println(str)
	},
}
