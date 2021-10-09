package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/hacdias/fintracts/parser"
)

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Printf("usage: parser <input file> or parser < /input/file.txt ")
		os.Exit(1)
	}

	var (
		bytes []byte
		err   error
	)

	if len(args) == 1 {
		bytes, err = os.ReadFile(args[0])
	} else {
		bytes, err = readStdin()
	}

	checkErr(err)

	contract, err := parser.Parse(bytes)
	checkErr(err)

	bytes, err = json.MarshalIndent(contract, "", "  ")
	checkErr(err)

	fmt.Println(string(bytes))
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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
