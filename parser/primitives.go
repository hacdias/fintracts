package parser

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"go.uber.org/multierr"
)

type partyValidator func(string) error

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
	Currency string      `parser:"@Ident" json:"currency"`
	Amount   MoneyAmount `parser:"@Money" json:"amount"`
}

// Date is a date in the format '1st of September 2022'. It marshals
// in JSON in the RFC3339 format.
type Date struct {
	date  time.Time
	Day   int    `parser:"@Integer" json:"day"`
	Month string `parser:"('th' | 'rd' | 'st') 'of' @Ident" json:"month"`
	Year  int    `parser:"@Integer" json:"year"`
}

func (d *Date) Validate() error {
	day := fmt.Sprintf("%02d", d.Day)
	year := strconv.Itoa(d.Year)

	date, err := time.Parse("02 January 2006", day+" "+d.Month+" "+year)
	d.date = date
	return err
}

func (d *Date) MarshalJSON() ([]byte, error) {
	str := d.date.Format(time.RFC3339)
	str = fmt.Sprintf("\"%s\"", str)
	return []byte(str), nil
}

// Signature represents a signature parsed from the format
// 'Signed by <Party>, [<Party>, ...] and <Party> on <Date>.'
//
// The signature does not need the parties per se. They are
// only used to semantically validate the English text.
// The marshalled output is a Date object.
type Signature struct {
	Parties []string `parser:"'Signed' 'by'  @Ident (',' @Ident)* 'and' @Ident" json:"parties"`
	Date    *Date    `parser:"'on' 'the' @@ '.'" json:"date"`
}

func (s *Signature) Validate() error {
	sort.Strings(s.Parties)
	return s.Date.Validate()
}

func (s *Signature) MarshalJSON() ([]byte, error) {
	return s.Date.MarshalJSON()
}

// Party represents a party and its identifier.
type Party struct {
	Name       LongIdent `parser:"@(~',')+" json:"name"`
	Identifier string    `parser:"',' 'undermentioned' 'as' @Ident" json:"identifier"`
}

// Contract represents a contract with parties, agreements and a signature.
type Contract struct {
	Parties    []Party      `parser:"'The' 'parties' ':' @@ ';' 'and' (@@ ';' 'and')* @@ '.'" json:"parties"`
	Agreements []*Agreement `parser:"@@+" json:"agreements"`
	SignedOn   Signature    `parser:"@@" json:"signedOn"`
}

func (c *Contract) Validate() error {
	err := c.SignedOn.Validate()
	if err != nil {
		return err
	}

	parties := getIdentifiers(c.Parties)
	sigParties := c.SignedOn.Parties

	if len(parties) != len(sigParties) {
		return fmt.Errorf("mentioned parties do not match signature")
	}

	if !equal(parties, sigParties) {
		return fmt.Errorf("parties and signing parties must be the same")
	}

	validateParty := func(v string) error {
		for _, p := range parties {
			if p == v {
				return nil
			}
		}

		return fmt.Errorf("party %s not found", v)
	}

	for _, agreement := range c.Agreements {
		err = multierr.Append(err, agreement.Validate(validateParty))
	}

	err = multierr.Append(err, c.SignedOn.Validate())
	return err
}

type Coupons struct {
	Rate  float64 `parser:"'with' 'an' 'interest' 'rate' 'of' (@Float | @Integer) '%'" json:"rate"`
	Dates []*Date `parser:"'paid' 'on' 'the' 'following' 'dates' ':' (@@ ',' | @@ | 'and' @@)+ '.'" json:"dates"`
}

func (c *Coupons) Validate() error {
	var err error
	for _, date := range c.Dates {
		err = multierr.Append(err, date.Validate())
	}
	return err
}

type ExchangeRate struct {
	BaseCurrency    string  `parser:"@Ident" json:"baseCurrency"`
	CounterCurrency string  `parser:"'/' @Ident" json:"counterCurrency"`
	Rate            float64 `parser:"@Float" json:"rate"`
}

type Agreement struct {
	BondPurchase     *BondPurchase     `parser:"'Hereby' 'enter' 'in' ( 'a' 'Bond' 'Purchase' 'Agreement' 'defined' 'as' 'follows' ':' @@ " json:"bondPurchase,omitempty"`
	InterestRateSwap *InterestRateSwap `parser:"| 'an' 'Interest' 'Rate' 'Swap' 'Transaction' 'Agreement' 'defined' 'as' 'follows' ':' @@" json:"interestRateSwap,omitempty"`
	CurrencySwap     *CurrencySwap     `parser:"| 'a' 'Currency' 'Swap' 'Transaction' 'Agreement' 'defined' 'as' 'follows' ':' @@ )" json:"currencySwap,omitempty"`
}

func (a *Agreement) Validate(validateParty partyValidator) error {
	if a.BondPurchase != nil {
		return a.BondPurchase.Validate(validateParty)
	}
	if a.InterestRateSwap != nil {
		return a.InterestRateSwap.Validate(validateParty)
	}
	return a.CurrencySwap.Validate(validateParty)
}

type InterestPayment struct {
	Payer              string    `parser:"@Ident 'will' 'pay' 'a'" json:"payer"`
	FixedRate          float64   `parser:"( 'fixed' 'rate' 'interest' 'of' (@Float | @Integer) '%' " json:"fixedRate"`
	InitialRate        float64   `parser:"| 'floating' 'rate' 'interest' ',' 'initially' 'defined' 'as' (@Float | @Integer) '%' ',' ) " json:"initialRate"`
	Dates              []*Date   `parser:"'over' 'the' 'notational' 'amount' 'on' 'the' 'following' 'dates' ':' (@@ ',' | @@ | 'and' @@)+ '.'" json:"dates"`
	InterestRateOption LongIdent `parser:"('The' 'floating' 'rate' 'option' 'is' @(~'.')+ '.')?" json:"interestRateOption"`
}

func (i *InterestPayment) Validate(validateParty partyValidator) error {
	if i.FixedRate != 0 && i.InterestRateOption != "" {
		return fmt.Errorf("fixed rate cannot be used with an interest rate option")
	}

	if i.InitialRate != 0 && i.InterestRateOption == "" {
		return fmt.Errorf("floating rate must have an interest rate option attached")
	}

	var err error
	for _, date := range i.Dates {
		err = multierr.Append(err, date.Validate())
	}

	return multierr.Append(err, validateParty(i.Payer))
}

func getIdentifiers(parties []Party) []string {
	ids := []string{}
	for _, party := range parties {
		ids = append(ids, party.Identifier)
	}
	sort.Strings(ids)
	return ids
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
