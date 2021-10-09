package parser

import "go.uber.org/multierr"

type BondPurchase struct {
	Seller       string   `parser:"@Ident 'agrees' 'on' 'issuing' 'and' 'selling'" json:"seller"`
	FaceValue    Money    `parser:"'a' 'bond' 'of' @@" json:"faceValue"`
	Buyer        string   `parser:"'to' @Ident" json:"buyer"`
	IssuePrice   Money    `parser:"'for' @@ '.'" json:"issuePrice"`
	MaturityDate *Date    `parser:"'The' 'aforementioned' 'bond' 'reaches' 'maturity' 'on' 'the' @@ '.'" json:"maturityDate"`
	Coupons      *Coupons `parser:"('The' 'bond' 'pays' 'coupons' @@)?" json:"coupons,omitempty"`
}

func (b *BondPurchase) Validate() error {
	// TODO: check seller and buyer validity

	err := multierr.Combine(
		b.MaturityDate.Validate(),
	)

	return err
}

type InterestRateSwap struct {
	NotationalAmount Money              `parser:"'The' 'parties' 'agree' 'on' 'an' 'interest' 'rate' 'swap' 'transaction' 'over' 'the' 'notational' 'principal' 'of' @@ ','" json:"notationalAmount"`
	EffectiveDate    *Date              `parser:"'with' 'an' 'effective' 'date' 'as' 'of' 'the' @@" json:"effectiveDate"`
	MaturityDate     *Date              `parser:"'and' 'termination' 'on' 'the' @@ '.'" json:"maturityDate"`
	Interest         []*InterestPayment `parser:"@@+" json:"interest"`
}

func (i *InterestRateSwap) Validate() error {
	// TODO: check seller and buyer validity

	err := multierr.Combine(
		i.EffectiveDate.Validate(),
		i.MaturityDate.Validate(),
	)

	return err
}

type CurrencySwap struct {
	EffectiveDate       *Date              `parser:"'The' 'parties' 'agree' 'on' 'a' 'currency' 'swap' 'transaction' 'effective' 'as' 'of' 'the' @@" json:"effectiveDate"`
	MaturityDate        *Date              `parser:"'and' 'termination' 'on' 'the' @@ '.'" json:"maturityDate"`
	PayerA              string             `parser:"@Ident 'will' 'pay' 'a'" json:"payerA"`
	PrincipalA          Money              `parser:"'principal' 'amount' 'of' @@ ','" json:"principalA"`
	PayerB              string             `parser:"'and' 'the' @Ident 'will' 'pay' 'a'" json:"payerB"`
	PrincipalB          Money              `parser:"'principal' 'amount' 'of' @@ '.'" json:"principalB"`
	ImpliedExchangeRate ExchangeRate       `parser:"" json:"impliedExchangeRate"`
	EndExchangeRate     *ExchangeRate      `parser:"('At' 'maturity' ',' 'the' 'principal' 'amounts' 'shall' 'be' 'exchanged' 'back' 'with' 'an' 'interest' 'rate' 'of' @@ '.')?" json:"endExchangeRate,omitempty"`
	Interest            []*InterestPayment `parser:"@@*" json:"interest,omitempty"`
}

func (c *CurrencySwap) Validate() error {
	// TODO: check payerA, payerB

	c.ImpliedExchangeRate.BaseCurrency = c.PrincipalA.Currency
	c.ImpliedExchangeRate.CounterCurrency = c.PrincipalB.Currency
	c.ImpliedExchangeRate.Rate = float64(c.PrincipalB.Amount) / float64(c.PrincipalA.Amount)

	err := multierr.Combine(
		c.EffectiveDate.Validate(),
		c.MaturityDate.Validate(),
	)

	return err
}
