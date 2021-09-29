package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println(parser.String())
	file, err := ioutil.ReadFile("./contract2.txt")
	if err != nil {
		log.Fatal(err)
	}

	contract, err := Parse(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = Validate(contract)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bytes, err := json.MarshalIndent(contract, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(bytes))
}
