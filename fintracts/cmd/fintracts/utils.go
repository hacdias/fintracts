package main

import (
	"bufio"
	"io"
	"log"
	"os"

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

func inputFlagOrStdin(cmd *cobra.Command) []byte {
	inputFlag := mustGetString(cmd, "input")

	var (
		data []byte
		err  error
	)

	if inputFlag != "" {
		data, err = os.ReadFile(inputFlag)
	} else {
		data, err = readStdin()
	}

	checkErr(err)
	return data
}

func readStdin() ([]byte, error) {
	bytes := []byte{}
	in := bufio.NewReader(os.Stdin)
	for {
		b, err := in.ReadByte()
		if err != nil {
			// io.EOF is expected, anything else
			// should be handled/reported
			if err != io.EOF {
				return nil, err
			}
			break
		}
		bytes = append(bytes, b)
	}

	return bytes, nil
}
