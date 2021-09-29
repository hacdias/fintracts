package main

import (
	"strconv"
	"strings"
)

// LongIdentifier is an identifier that spans across multiple token yieldings.
type LongIdentifier string

func (en *LongIdentifier) Capture(values []string) error {
	if *en != "" {
		values = append([]string{string(*en)}, values...)
	}
	*en = LongIdentifier(strings.Join(values, " "))
	return nil
}

// MoneyAmount is the the type of a monetary amount in the format '10,000,000.00'.
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

// Date is a date in the format '1st of September 2022'.
// TODO: possibly replace by time.Parse.
type Date struct {
	Day   int    `parser:"@Integer" json:"day"`
	Month string `parser:"('th' | 'rd' | 'st') 'of' @Ident" json:"month"`
	Year  int    `parser:"@Integer" json:"year"`
}

// Money represents a monetary amount of a certain currency parsed
// from the format 'EUR 10,000.00'.
type Money struct {
	Currency string      `parser:"@Ident" json:"currency"`
	Amount   MoneyAmount `parser:"@Money" json:"amount"`
}

// Signature represents a signature parsed from the format
// 'Signed by <Party>, [<Party>, ...] and <Party> on <Date>.'
type Signature struct {
	Parties []string `parser:"'Signed' 'by'  @Ident (',' @Ident)* 'and' @Ident" json:"parties"`
	Date    Date     `parser:"'on' 'the' @@ '.'" json:"date"`
}

// Party represents a party and its identifiers parsed from the format
// '<Name>, undermentioned as <Identifier>'
type Party struct {
	Name       LongIdentifier `parser:"@(~',')+" json:"name"`
	Identifier string         `parser:"',' 'undermentioned' 'as' @Ident" json:"identifier"`
}

// Contract represents a contract with parties, agreements and a signature.
type Contract struct {
	Parties    []Party     `parser:"'The' 'parties' ':' @@ ';' 'and' (@@ ';' 'and')* @@ '.'" json:"parties"`
	Agreements []Agreement `parser:"@@+" json:"agreements"`
	Signature  Signature   `parser:"@@" json:"signature"`
}
