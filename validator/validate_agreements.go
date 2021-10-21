package validator

import (
	"fmt"
	"time"

	"github.com/hacdias/fintracts"
	"go.uber.org/multierr"
)

func (v *validator) validateAgreement(a fintracts.Agreement) error {
	var (
		err    error
		nonNil = 0
	)

	if a.BondPurchase != nil {
		nonNil++
		err = multierr.Append(err, v.validateBondPurchase(a.BondPurchase))
	}

	if a.InterestRateSwap != nil {
		nonNil++
		err = multierr.Append(err, v.validateInterestRateSwap(a.InterestRateSwap))
	}

	if a.CurrencySwap != nil {
		nonNil++
		err = multierr.Append(err, v.validateCurrencySwap(a.CurrencySwap))
	}

	if nonNil != 1 {
		return fmt.Errorf("agreement must be of a single type")
	}

	return err
}

func (v *validator) validateBondPurchase(b *fintracts.BondPurchase) error {
	err := multierr.Combine(
		v.validatePartyExists(b.Issuer),
		v.validatePartyExists(b.Underwriter),
		v.validateMoney(b.FaceValue),
		v.validateMoney(b.IssuePrice),
		v.validateDate(b.MaturityDate),
		v.validateAfterSignatures(b.MaturityDate),
		v.validateDifferentParties(b.Issuer, b.Underwriter),
	)

	if b.Coupons != nil {
		err = multierr.Append(err, v.validateCoupons(b.Coupons))

		for _, date := range b.Coupons.Dates {
			err = multierr.Append(err, v.validateDateInRange(time.Time(date), time.Time{}, time.Time(b.MaturityDate)))
		}
	}

	return err
}

func (v *validator) validateCoupons(c *fintracts.Coupons) error {
	var err error

	if c.Rate <= 0 {
		err = multierr.Append(err, fmt.Errorf("expected coupon rate to be larger than 0: %f", c.Rate))
	}

	for _, date := range c.Dates {
		err = multierr.Append(err, v.validateDate(date))
		err = multierr.Append(err, v.validateAfterSignatures(date))
	}

	return err
}

func (v *validator) validateInterestRateSwap(i *fintracts.InterestRateSwap) error {
	err := multierr.Combine(
		v.validateMoney(i.NotationalAmount),
		v.validateDate(i.EffectiveDate),
		v.validateDate(i.MaturityDate),
		v.validateAfterSignatures(i.MaturityDate),
		v.validateAfterSignatures(i.EffectiveDate),
	)

	if i.Interest == nil {
		err = multierr.Append(err, fmt.Errorf("interest cannot be non-existent for interest rate swap agreements"))
	}

	if len(i.Interest) < 2 {
		err = multierr.Append(err, fmt.Errorf("interest rate swap agreements must have 2 or more interest payments"))
	}

	for _, payment := range i.Interest {
		err = multierr.Append(err, v.validateInterestPayment(payment))

		for _, date := range payment.Dates {
			err = multierr.Append(err, v.validateDateInRange(time.Time(date), time.Time(i.EffectiveDate), time.Time(i.MaturityDate)))
		}
	}

	return err
}

func (v *validator) validateCurrencySwap(s *fintracts.CurrencySwap) error {
	err := multierr.Combine(
		v.validatePartyExists(s.PayerA),
		v.validatePartyExists(s.PayerB),
		v.validateMoney(s.PrincipalA),
		v.validateMoney(s.PrincipalB),
		v.validateDate(s.EffectiveDate),
		v.validateDate(s.MaturityDate),
		v.validateAfterSignatures(s.MaturityDate),
		v.validateAfterSignatures(s.EffectiveDate),
		v.validateDifferentParties(s.PayerA, s.PayerB),
	)

	if s.EndExchangeRate != nil {
		err = multierr.Append(err, v.validateExchangeRate(*s.EndExchangeRate))
	}

	return err
}
