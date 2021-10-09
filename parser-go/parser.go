package parser

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var basicLexer = lexer.MustSimple([]lexer.Rule{
	{"Money", `\d{1,3}(?:,?\d{3})*\.\d{2}`, nil},
	{"Float", `[0-9]*[.][0-9]+`, nil},
	{"Integer", `[-+]?(\d*\.)?\d+`, nil},
	{"Ident", `[a-zA-Z_]\w*`, nil},
	{"Punct", `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`, nil},

	// Elided lexical elements.
	{"eol", `[\n\r]+`, nil},
	{"whitespace", `[ \t]+`, nil},
})

var parser = participle.MustBuild(&Contract{},
	participle.Lexer(basicLexer),
	participle.UseLookahead(20),
)

func Parse(contract []byte) (*Contract, error) {
	ast := &Contract{}
	err := parser.ParseBytes("", contract, ast)
	return ast, err
}
