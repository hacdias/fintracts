package fintracts

import (
	"fmt"

	"go.uber.org/multierr"
)

func (a *Agreement) validate(c *Contract, fix bool) error {
	var (
		err    error
		nonNil = 0
	)

	if a.BondPurchase != nil {
		nonNil++
		err = multierr.Append(err, a.BondPurchase.validate(c, fix))
	}

	if a.InterestRateSwap != nil {
		nonNil++
		err = multierr.Append(err, a.InterestRateSwap.validate(c, fix))
	}

	if a.CurrencySwap != nil {
		nonNil++
		err = multierr.Append(err, a.CurrencySwap.validate(c, fix))
	}

	if nonNil != 1 {
		return fmt.Errorf("agreement must be of a single type")
	}

	return err
}

func (b *BondPurchase) validate(c *Contract, fix bool) error {
	err := multierr.Combine(
		c.validatePartyExists(b.Issuer),
		c.validatePartyExists(b.Underwriter),
		b.FaceValue.validate(fix),
		b.IssuePrice.validate(fix),
		b.MaturityDate.validate(),
		validateDifferentParties(b.Issuer, b.Underwriter),
	)

	if b.Coupons != nil {
		err = multierr.Append(err, b.Coupons.validate())
	}

	return err
}

func (c *Coupons) validate() error {
	var err error

	if c.Rate <= 0 {
		err = multierr.Append(err, fmt.Errorf("expected coupon rate to be larger than 0: %f", c.Rate))
	}

	for _, date := range c.Dates {
		err = multierr.Append(err, date.validate())
	}
	return err
}

func (i *InterestRateSwap) validate(c *Contract, fix bool) error {
	err := multierr.Combine(
		i.NotationalAmount.validate(fix),
		i.EffectiveDate.validate(),
		i.MaturityDate.validate(),
	)

	if i.Interest == nil {
		err = multierr.Append(err, fmt.Errorf("interest cannot be non-existent for interest rate swap agreements"))
	}

	if len(i.Interest) < 2 {
		err = multierr.Append(err, fmt.Errorf("interest rate swap agreements must have 2 or more interest payments"))
	}

	for _, payment := range i.Interest {
		err = multierr.Append(err, payment.validate(c))
	}

	return err
}

func (s *CurrencySwap) validate(c *Contract, fix bool) error {
	err := multierr.Combine(
		c.validatePartyExists(s.PayerA),
		c.validatePartyExists(s.PayerB),
		s.PrincipalA.validate(fix),
		s.PrincipalB.validate(fix),
		s.EffectiveDate.validate(),
		s.MaturityDate.validate(),
		validateDifferentParties(s.PayerA, s.PayerB),
	)

	if s.ImpliedExchangeRate.BaseCurrency != s.PrincipalA.Currency {
		if fix {
			s.ImpliedExchangeRate.BaseCurrency = s.PrincipalA.Currency
		} else {
			err = multierr.Append(err, fmt.Errorf("implied exchange rate base currency should be %s", s.PrincipalA.Currency))
		}
	}

	if s.ImpliedExchangeRate.CounterCurrency != s.PrincipalB.Currency {
		if fix {
			s.ImpliedExchangeRate.CounterCurrency = s.PrincipalB.Currency
		} else {
			err = multierr.Append(err, fmt.Errorf("implied exchange rate counter currency should be %s", s.PrincipalB.Currency))
		}
	}

	rate := s.PrincipalB.Amount / s.PrincipalA.Amount
	if s.ImpliedExchangeRate.Rate != rate {
		if fix {
			s.ImpliedExchangeRate.Rate = rate
		} else {
			err = multierr.Append(err, fmt.Errorf("implied exchange rate should be %f", rate))
		}
	}

	err = multierr.Append(err, s.ImpliedExchangeRate.validate(fix))

	if s.EndExchangeRate != nil {
		err = multierr.Append(err, s.EndExchangeRate.validate(fix))
	}

	for _, payment := range s.Interest {
		err = multierr.Append(err, payment.validate(c))
	}

	return err
}
