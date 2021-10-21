package main

import (
	"fmt"

	"github.com/hacdias/fintracts/fintracts"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.Flags().StringP("input", "i", "", "indicates the input file")
}

var validateCmd = &cobra.Command{
	Use:   "validate [< file.json]",
	Short: "Validates a contract in the JSON specification",
	Run: func(cmd *cobra.Command, args []string) {
		data := inputFlagOrStdin(cmd)

		contract, err := fintracts.FromJSON(data)
		checkErr(err)

		err = fintracts.Validate(contract)
		checkErr(err)

		str, err := contract.String()
		checkErr(err)

		fmt.Println(str)
	},
}
