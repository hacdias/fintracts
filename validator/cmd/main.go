package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/hacdias/fintracts"
	"github.com/hacdias/fintracts/validator"
	"github.com/spf13/cobra"
	"go.uber.org/multierr"
)

var cmd = &cobra.Command{
	Use:   "validate [< file.json]",
	Short: "Validates a contract in the JSON specification",
	Run: func(cmd *cobra.Command, args []string) {
		data := inputFlagOrStdin(cmd)

		contract, err := fintracts.FromJSON(data)
		checkErr(err)

		err = validator.Validate(contract)
		checkErr(err)

		str, err := contract.String()
		checkErr(err)

		fmt.Println(str)
	},
}

func init() {
	cmd.Flags().StringP("input", "i", "", "indicates the input file")
}

func main() {
	cmd.Execute()
}

// Utils

func mustGetString(cmd *cobra.Command, flag string) string {
	s, err := cmd.Flags().GetString(flag)
	checkErr(err)
	return s
}

func checkErr(err error) {
	if err != nil {
		errs := multierr.Errors(err)
		for _, err := range errs {
			fmt.Println(err)
		}
		os.Exit(1)
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
