package parser

import (
	"fmt"

	"go.uber.org/multierr"
)

type BondPurchase struct {
	Issuer       string   `parser:"@Ident 'agrees' 'on' 'issuing' 'and' 'selling'" json:"issuer"`
	FaceValue    Money    `parser:"'a' 'bond' 'of' @@" json:"faceValue"`
	Underwriter  string   `parser:"'to' @Ident" json:"underwriter"`
	IssuePrice   Money    `parser:"'for' @@ '.'" json:"issuePrice"`
	MaturityDate *Date    `parser:"'The' 'aforementioned' 'bond' 'reaches' 'maturity' 'on' 'the' @@ '.'" json:"maturityDate"`
	Coupons      *Coupons `parser:"('The' 'bond' 'pays' 'coupons' @@)?" json:"coupons,omitempty"`
}

func (b *BondPurchase) Validate(validateParty partyValidator) error {
	err := multierr.Combine(
		validateParty(b.Issuer),
		validateParty(b.Underwriter),
		b.MaturityDate.Validate(),
		ensureDifferentParties(b.Issuer, b.Underwriter),
	)

	if b.Coupons != nil {
		err = multierr.Append(err, b.Coupons.Validate())
	}

	return err
}

type InterestRateSwap struct {
	NotationalAmount Money              `parser:"'The' 'parties' 'agree' 'on' 'an' 'interest' 'rate' 'swap' 'transaction' 'over' 'the' 'notational' 'principal' 'of' @@ ','" json:"notationalAmount"`
	EffectiveDate    *Date              `parser:"'with' 'an' 'effective' 'date' 'as' 'of' 'the' @@" json:"effectiveDate"`
	MaturityDate     *Date              `parser:"'and' 'termination' 'on' 'the' @@ '.'" json:"maturityDate"`
	Interest         []*InterestPayment `parser:"@@+" json:"interest"`
}

func (i *InterestRateSwap) Validate(validateParty partyValidator) error {
	err := multierr.Combine(
		i.EffectiveDate.Validate(),
		i.MaturityDate.Validate(),
	)

	for _, payment := range i.Interest {
		err = multierr.Append(err, payment.Validate(validateParty))
	}

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

func (c *CurrencySwap) Validate(validateParty partyValidator) error {
	c.ImpliedExchangeRate.BaseCurrency = c.PrincipalA.Currency
	c.ImpliedExchangeRate.CounterCurrency = c.PrincipalB.Currency
	c.ImpliedExchangeRate.Rate = float64(c.PrincipalB.Amount) / float64(c.PrincipalA.Amount)

	err := multierr.Combine(
		validateParty(c.PayerA),
		validateParty(c.PayerB),
		c.EffectiveDate.Validate(),
		c.MaturityDate.Validate(),
		ensureDifferentParties(c.PayerA, c.PayerB),
	)

	for _, payment := range c.Interest {
		err = multierr.Append(err, payment.Validate(validateParty))
	}

	return err
}

func ensureDifferentParties(a, b string) error {
	if a != b {
		return nil
	}

	return fmt.Errorf("expected '%s' and '%s' to be different parties", a, b)
}
