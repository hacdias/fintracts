package main

import (
	"fmt"

	"github.com/hacdias/fintracts/fintracts/english"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(englishCmd)
	englishCmd.Flags().StringP("input", "i", "", "indicates the input file")
}

var englishCmd = &cobra.Command{
	Use:   "english [< file.txt]",
	Short: "Parses and validates contracts in the English specification",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data := inputFlagOrStdin(cmd)

		contract, err := english.Parse(data)
		checkErr(err)

		str, err := contract.String()
		checkErr(err)

		fmt.Println(str)
	},
}
