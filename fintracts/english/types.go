package english

import (
	"strconv"
	"strings"
)

// LongIdent is an identifier that spans across multiple token yieldings.
type LongIdent string

func (en *LongIdent) Capture(values []string) error {
	if *en != "" {
		values = append([]string{string(*en)}, values...)
	}
	*en = LongIdent(strings.Join(values, " "))
	return nil
}

// MoneyAmount is the the type of a monetary amount.
type MoneyAmount float64

func (b *MoneyAmount) Capture(values []string) error {
	if len(values) != 1 {
		panic("expected values to have length 1")
	}

	v := values[0]
	v = strings.ReplaceAll(v, ",", "")

	vv, err := strconv.ParseFloat(v, 64)
	*b = MoneyAmount(vv)
	return err
}

// Money represents a monetary amount of a certain currency.
type Money struct {
	Currency string      `parser:"@Ident"`
	Amount   MoneyAmount `parser:"(@Money | @Float)"`
}

// Date is a date in the format '1st of September 2022'. It marshals
// in JSON in the RFC3339 format.
type Date struct {
	Day   int    `parser:"@Integer"`
	Month string `parser:"('th' | 'rd' | 'st') 'of' @Ident"`
	Year  int    `parser:"@Integer"`
}

// Signature represents a signature parsed from the format
// 'Signed by <Party>, [<Party>, ...] and <Party> on <Date>.'
//
// The signature does not need the parties per se. They are
// only used to semantically validate the English text.
// The marshalled output is a Date object.
type Signature struct {
	Parties []string `parser:"'Signed' 'by'  @Ident (',' @Ident)* ('and' @Ident)*"`
	Date    *Date    `parser:"'on' 'the' @@ '.'"`
}

// Party represents a party and its identifier.
type Party struct {
	Name       LongIdent `parser:"@(~',')+"`
	Identifier string    `parser:"',' 'undermentioned' 'as' @Ident"`
}

// Contract represents a contract with parties, agreements and a signature.
type Contract struct {
	Parties    []Party      `parser:"'The' 'parties' ':' @@ ';' 'and' (@@ ';' 'and')* @@ '.'"`
	Agreements []*Agreement `parser:"@@+"`
	Signatures []*Signature `parser:"@@+"`
}

type ExchangeRate struct {
	BaseCurrency    string  `parser:"@Ident"`
	CounterCurrency string  `parser:"'/' @Ident"`
	Rate            float64 `parser:"@Float"`
}

type Agreement struct {
	BondPurchase     *BondPurchase     `parser:"'Hereby' 'enter' 'in' ( 'a' 'Bond' 'Purchase' 'Agreement' 'defined' 'as' 'follows' ':' @@ "`
	InterestRateSwap *InterestRateSwap `parser:"| 'an' 'Interest' 'Rate' 'Swap' 'Transaction' 'Agreement' 'defined' 'as' 'follows' ':' @@"`
	CurrencySwap     *CurrencySwap     `parser:"| 'a' 'Currency' 'Swap' 'Transaction' 'Agreement' 'defined' 'as' 'follows' ':' @@ )"`
}

type InterestPayment struct {
	Payer       string    `parser:"@Ident 'will' 'pay'"`
	Receiver    string    `parser:"@Ident 'a'"`
	FixedRate   float64   `parser:"( 'fixed' 'rate' 'interest' 'of' (@Float | @Integer) '%' "`
	InitialRate float64   `parser:"| 'floating' 'rate' 'interest' ',' 'initially' 'defined' 'as' (@Float | @Integer) '%' ',' ) "`
	Dates       []*Date   `parser:"'over' 'the' 'notational' 'amount' 'on' 'the' 'following' 'dates' ':' (@@ ',' | @@ | 'and' @@)+ '.'"`
	RateOption  LongIdent `parser:"('The' 'floating' 'rate' 'option' 'is' @(~'.')+ '.')?"`
}

type BondPurchase struct {
	Issuer       string   `parser:"@Ident 'agrees' 'on' 'issuing' 'and' 'selling'"`
	FaceValue    Money    `parser:"'a' 'bond' 'of' @@"`
	Underwriter  string   `parser:"'to' @Ident"`
	IssuePrice   Money    `parser:"'for' @@ '.'"`
	MaturityDate *Date    `parser:"'The' 'aforementioned' 'bond' 'reaches' 'maturity' 'on' 'the' @@ '.'"`
	Coupons      *Coupons `parser:"('The' 'bond' 'pays' 'coupons' @@)?"`
}

type Coupons struct {
	Rate  float64 `parser:"'with' 'an' 'interest' 'rate' 'of' (@Float | @Integer) '%'"`
	Dates []*Date `parser:"'paid' 'on' 'the' 'following' 'dates' ':' (@@ ',' | @@ | 'and' @@)+ '.'"`
}

type InterestRateSwap struct {
	NotationalAmount Money              `parser:"'The' 'parties' 'agree' 'on' 'an' 'interest' 'rate' 'swap' 'transaction' 'over' 'the' 'notational' 'principal' 'of' @@ ','"`
	EffectiveDate    *Date              `parser:"'with' 'an' 'effective' 'date' 'as' 'of' 'the' @@"`
	MaturityDate     *Date              `parser:"'and' 'termination' 'on' 'the' @@ '.'"`
	Interest         []*InterestPayment `parser:"@@+"`
}

type CurrencySwap struct {
	EffectiveDate   *Date              `parser:"'The' 'parties' 'agree' 'on' 'a' 'currency' 'swap' 'transaction' 'effective' 'as' 'of' 'the' @@"`
	MaturityDate    *Date              `parser:"'and' 'termination' 'on' 'the' @@ '.'"`
	PayerA          string             `parser:"@Ident 'will' 'pay' 'a'"`
	PrincipalA      Money              `parser:"'principal' 'amount' 'of' @@ ','"`
	PayerB          string             `parser:"'and' 'the' @Ident 'will' 'pay' 'a'"`
	PrincipalB      Money              `parser:"'principal' 'amount' 'of' @@ '.'"`
	EndExchangeRate *ExchangeRate      `parser:"('At' 'maturity' ',' 'the' 'principal' 'amounts' 'shall' 'be' 'exchanged' 'back' 'with' 'an' 'interest' 'rate' 'of' @@ '.')?"`
	Interest        []*InterestPayment `parser:"@@*"`
}
