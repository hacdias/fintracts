package english

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	fintracts "github.com/hacdias/fintracts/cli"
)

var basicLexer = lexer.MustSimple([]lexer.Rule{
	{"Money", `\d{1,3}(?:,?\d{3})+\.\d{2}`, nil},
	{"Float", `\d+[.]\d+`, nil},
	{"Integer", `[-+]?(\d*\.)?\d+`, nil},
	{"Ident", `[a-zA-Z_]\w*`, nil},
	{"Punct", `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`, nil},

	// Elided lexical elements.
	{"eol", `[\n\r]+`, nil},
	{"whitespace", `[ \t]+`, nil},
})

var parser = participle.MustBuild(&Contract{},
	participle.Lexer(basicLexer),
	participle.UseLookahead(100),
)

// Parse parses the contract in English to the JSON format.
func Parse(data []byte) (*fintracts.Contract, error) {
	ast := &Contract{}
	err := parser.ParseBytes("", data, ast)
	if err != nil {
		return nil, err
	}

	contract, err := ast.convert()
	if err != nil {
		return nil, err
	}

	err = fintracts.Validate(contract, true)
	if err != nil {
		return nil, err
	}

	return contract, err
}
