package main

type Agreement struct {
	Type            LongIdentifier   `parser:"'Hereby' 'enter' 'in' 'a' @(~'defined')+ 'defined' 'as' 'follows' ':'" json:"contractType"`
	BondPurchase    *BondPurchase    `parser:"(   @@" json:"bondPurchase,omitempty"`
	ExamplePurchase *ExamplePurchase `parser:"	| @@ )" json:"examplePurchase,omitempty"`
}

type Coupons struct {
	Rate  float32 `parser:"'interest' 'rate' 'of' (@Float | @Integer) '%'" json:"rate"`
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

// Accepts a contract in the format (example on how easy it is to add a new contract type.)
// 	"The buyer <buyer> buys <something> from <someone> for <something>."
type ExamplePurchase struct {
	Buyer string `parser:"'The' 'buyer' @Ident" json:"buyer"`
	Buys  string `parser:"'buys' @Ident" json:"buys"`
	From  string `parser:"'from' @Ident" json:"from"`
	For   string `parser:"'for' @Ident '.'" json:"for"`
}
