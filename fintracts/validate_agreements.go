package fintracts

import (
	"fmt"
	"time"

	"go.uber.org/multierr"
)

func (a *Agreement) validate(c *Contract) error {
	var (
		err    error
		nonNil = 0
	)

	if a.BondPurchase != nil {
		nonNil++
		err = multierr.Append(err, a.BondPurchase.validate(c))
	}

	if a.InterestRateSwap != nil {
		nonNil++
		err = multierr.Append(err, a.InterestRateSwap.validate(c))
	}

	if a.CurrencySwap != nil {
		nonNil++
		err = multierr.Append(err, a.CurrencySwap.validate(c))
	}

	if nonNil != 1 {
		return fmt.Errorf("agreement must be of a single type")
	}

	return err
}

func (b *BondPurchase) validate(c *Contract) error {
	err := multierr.Combine(
		c.validatePartyExists(b.Issuer),
		c.validatePartyExists(b.Underwriter),
		b.FaceValue.validate(),
		b.IssuePrice.validate(),
		b.MaturityDate.validate(),
		c.validateAfterSignatures(b.MaturityDate),
		validateDifferentParties(b.Issuer, b.Underwriter),
	)

	if b.Coupons != nil {
		err = multierr.Append(err, b.Coupons.validate(c))

		for _, date := range b.Coupons.Dates {
			err = multierr.Append(err, validateDateInRange(time.Time(date), time.Time{}, time.Time(b.MaturityDate)))
		}
	}

	return err
}

func (c *Coupons) validate(co *Contract) error {
	var err error

	if c.Rate <= 0 {
		err = multierr.Append(err, fmt.Errorf("expected coupon rate to be larger than 0: %f", c.Rate))
	}

	for _, date := range c.Dates {
		err = multierr.Append(err, date.validate())
		err = multierr.Append(err, co.validateAfterSignatures(date))
	}

	return err
}

func (i *InterestRateSwap) validate(c *Contract) error {
	err := multierr.Combine(
		i.NotationalAmount.validate(),
		i.EffectiveDate.validate(),
		i.MaturityDate.validate(),
		c.validateAfterSignatures(i.MaturityDate),
		c.validateAfterSignatures(i.EffectiveDate),
	)

	if i.Interest == nil {
		err = multierr.Append(err, fmt.Errorf("interest cannot be non-existent for interest rate swap agreements"))
	}

	if len(i.Interest) < 2 {
		err = multierr.Append(err, fmt.Errorf("interest rate swap agreements must have 2 or more interest payments"))
	}

	for _, payment := range i.Interest {
		err = multierr.Append(err, payment.validate(c))

		for _, date := range payment.Dates {
			err = multierr.Append(err, validateDateInRange(time.Time(date), time.Time(i.EffectiveDate), time.Time(i.MaturityDate)))
		}
	}

	return err
}

func (s *CurrencySwap) validate(c *Contract) error {
	err := multierr.Combine(
		c.validatePartyExists(s.PayerA),
		c.validatePartyExists(s.PayerB),
		s.PrincipalA.validate(),
		s.PrincipalB.validate(),
		s.EffectiveDate.validate(),
		s.MaturityDate.validate(),
		c.validateAfterSignatures(s.MaturityDate),
		c.validateAfterSignatures(s.EffectiveDate),
		validateDifferentParties(s.PayerA, s.PayerB),
	)

	if s.EndExchangeRate != nil {
		err = multierr.Append(err, s.EndExchangeRate.validate())
	}

	for _, payment := range s.Interest {
		err = multierr.Append(err, payment.validate(c))

		for _, date := range payment.Dates {
			err = multierr.Append(err, validateDateInRange(time.Time(date), time.Time(s.EffectiveDate), time.Time(s.MaturityDate)))
		}
	}

	return err
}