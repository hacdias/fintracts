package main

import (
	"log"

	"github.com/spf13/cobra"
)

func mustGetString(cmd *cobra.Command, flag string) string {
	s, err := cmd.Flags().GetString(flag)
	checkErr(err)
	return s
}

func mustGetBool(cmd *cobra.Command, flag string) bool {
	s, err := cmd.Flags().GetBool(flag)
	checkErr(err)
	return s
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
