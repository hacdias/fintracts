package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var basicLexer = lexer.MustSimple([]lexer.Rule{
	{"Integer", `[-+]?(\d*\.)?\d+`, nil},
	{"Ident", `[a-zA-Z_]\w*`, nil},
	{"Punct", `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`, nil},
	{"eol", `[\n\r]+`, nil},
	{"whitespace", `[ \t]+`, nil},
})

var parser = participle.MustBuild(&Signature{},
	participle.Lexer(basicLexer),
	participle.UseLookahead(20),
)

func Parse(signature []byte) (*Signature, error) {
	ast := &Signature{}
	err := parser.ParseBytes("", signature, ast)
	return ast, err
}

type Date struct {
	Day   int    `parser:"@Integer" json:"day"`
	Month string `parser:"('th' | 'rd' | 'st') 'of' @Ident" json:"month"`
	Year  int    `parser:"@Integer" json:"year"`
}

type Signature struct {
	Parties []string `parser:"'Signed' 'by'  @Ident (',' @Ident)* 'and' @Ident" json:"parties"`
	Date    Date     `parser:"'on' 'the' @@ '.'" json:"date"`
}

func main() {
	file, err := ioutil.ReadFile("../signature.txt")
	if err != nil {
		log.Fatal(err)
	}

	signature, err := Parse(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bytes, err := json.MarshalIndent(signature, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(bytes))
}
