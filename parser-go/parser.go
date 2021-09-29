package main

import (
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var basicLexer = lexer.MustSimple([]lexer.Rule{
	// {"String", `"(\\"|[^"])*"`, nil},
	{"Money", `\d{1,3}(?:,?\d{3})*\.\d{2}`, nil},
	{"Float", `[0-9]*[.][0-9]+`, nil},
	{"Number", `[-+]?(\d*\.)?\d+`, nil},
	{"Ident", `[a-zA-Z_]\w*`, nil},
	{"Punct", `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`, nil},
	{"EOL", `[\n\r]+`, nil},

	// {"Comment", `#[^\n]+`, nil},
	{"Whitespace", `[ \t]+`, nil},
})

var parser = participle.MustBuild(&Contract{},
	participle.Lexer(basicLexer),
	participle.Elide("Whitespace", "EOL"),
	// participle.Unquote("String"),
	participle.UseLookahead(20),
)

// type MoneyAmount float64

// func (b *MoneyAmount) Capture(values []string) error {
// 	if len(values) != 1 {
// 		panic("expected values to have length 1")
// 	}

// 	v := values[0]
// 	v = strings.ReplaceAll(v, ",", "")

// 	vv, err := strconv.ParseFloat(v, 64)
// 	*b = MoneyAmount(vv)
// 	return err
// }

type LongName string

func (en *LongName) Capture(values []string) error {
	if *en != "" {
		values = append([]string{string(*en)}, values...)
	}
	*en = LongName(strings.Join(values, " "))
	return nil
}

type Contract struct {
	Parties      []*Party      `parser:"'The' 'parties' ':' @@ ';' 'and' @@ '.'" json:"parties"`
	ContractType LongName      `parser:"'Hereby' 'enter' 'in' 'a' @(~'defined')+ 'defined' 'as' 'follows' ':'" json:"contractType"`
	BondPurchase *BondPurchase `parser:"@@" json:"bondPurchase,omitempty"`
	SignedOn     SignedOn      `parser:"@@" json:"signedOn"`
}

type Party struct {
	Name       LongName `parser:"@(~',')+" json:"name"`
	Identifier string   `parser:"',' 'undermentioned' 'as' @Ident" json:"identifier"`
}

type SignedOn struct {
	PartyA string `parser:"'Signed' 'by' @Ident" json:"partyA"`
	PartyB string `parser:"'and' @Ident" json:"partyB"`
	Date   Date   `parser:"'on' 'the' @@ '.'" json:"date"`
}

type Date struct {
	Day   int    `parser:"@Number" json:"day"`
	Month string `parser:"('th' | 'rd' | 'st') 'of' @Ident" json:"month"`
	Year  int    `parser:"@Number" json:"year"`
}

type Money struct {
	Currency string `parser:"@Ident" json:"currency"`
	Amount   string `parser:"@Money" json:"amount"`
}

type Coupons struct {
	Rate  float32 `parser:"'interest' 'rate' 'of' (@Float | @Number) '%'" json:"rate"`
	Dates []Date  `parser:"'on' 'the' 'following' 'dates' ':' (@@ ',' | @@ | 'and' @@)+ '.'" json:"dates"`
}

type BondPurchase struct {
	Seller       string   `parser:"'Purchase' 'of' 'a' 'bond' ',' 'provided' 'by' @Ident" json:"seller"`
	IssuePrice   Money    `parser:"',' 'by' 'a' 'monetary' 'amount' 'of' @@" json:"issuePrice"`
	Payer        string   `parser:"',' 'provided' 'by' @Ident '.'" json:"payer"`
	FaceValue    Money    `parser:"'The' 'bond' 'aforementioned' 'has' 'a' 'face' 'value' 'of' @@" json:"faceValue"`
	MaturityDate Date     `parser:"'and' 'reaches' 'maturity' 'on' 'the' @@ '.'" json:"maturityDate"`
	Coupons      *Coupons `parser:"('This' 'bond' 'shall' 'pay' 'coupons' 'with' 'an' @@)?" json:"coupons,omitempty"`
}

func Parse(contract []byte) (*Contract, error) {
	ast := &Contract{}
	err := parser.ParseBytes("", contract, ast)
	return ast, err
}
