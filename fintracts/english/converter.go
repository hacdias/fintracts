package english

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	fintracts "github.com/hacdias/fintracts/cli"
)

func (c *Contract) convert() (*fintracts.Contract, error) {
	fc := &fintracts.Contract{
		Parties:    []fintracts.Party{},
		Agreements: []fintracts.Agreement{},
		Signatures: []fintracts.Signature{},
	}

	for _, party := range c.Parties {
		fc.Parties = append(fc.Parties, party.convert())
	}

	for _, a := range c.Agreements {
		fa, err := a.convert()
		if err != nil {
			return nil, err
		}
		fc.Agreements = append(fc.Agreements, *fa)
	}

	for _, sig := range c.Signatures {
		fsig, err := sig.convert()
		if err != nil {
			return nil, err
		}
		fc.Signatures = append(fc.Signatures, *fsig)
	}

	return fc, nil
}

func (p Party) convert() fintracts.Party {
	return fintracts.Party{
		Name:       string(p.Name),
		Identifier: p.Identifier,
	}
}

func (m Money) convert() fintracts.Money {
	return fintracts.Money{
		Currency: fintracts.Currency(m.Currency),
		Amount:   float64(m.Amount),
	}
}

func (e ExchangeRate) convert() fintracts.ExchangeRate {
	return fintracts.ExchangeRate{
		BaseCurrency:    fintracts.Currency(e.BaseCurrency),
		CounterCurrency: fintracts.Currency(e.CounterCurrency),
		Rate:            e.Rate,
	}
}

func (i InterestPayment) convert() (*fintracts.InterestPayment, error) {
	fi := &fintracts.InterestPayment{
		Payer:       i.Payer,
		FixedRate:   i.FixedRate,
		InitialRate: i.InitialRate,
		RateOption:  string(i.RateOption),
		Dates:       []fintracts.Date{},
	}

	for _, date := range i.Dates {
		fd, err := date.convert()
		if err != nil {
			return nil, err
		}
		fi.Dates = append(fi.Dates, fd)
	}

	return fi, nil
}

func (a Agreement) convert() (*fintracts.Agreement, error) {
	fa := &fintracts.Agreement{}

	if a.BondPurchase != nil {
		bp, err := a.BondPurchase.convert()
		if err != nil {
			return nil, err
		}
		fa.BondPurchase = bp
	}

	if a.InterestRateSwap != nil {
		irs, err := a.InterestRateSwap.convert()
		if err != nil {
			return nil, err
		}
		fa.InterestRateSwap = irs
	}

	if a.CurrencySwap != nil {
		cs, err := a.CurrencySwap.convert()
		if err != nil {
			return nil, err
		}
		fa.CurrencySwap = cs
	}

	return fa, nil
}

func (a BondPurchase) convert() (*fintracts.BondPurchase, error) {
	fa := &fintracts.BondPurchase{
		Issuer:      a.Issuer,
		Underwriter: a.Underwriter,
		FaceValue:   a.FaceValue.convert(),
		IssuePrice:  a.IssuePrice.convert(),
	}

	md, err := a.MaturityDate.convert()
	if err != nil {
		return nil, err
	}
	fa.MaturityDate = md

	if a.Coupons != nil {
		fa.Coupons, err = a.Coupons.convert()
		if err != nil {
			return nil, err
		}
	}

	return fa, nil
}

func (c Coupons) convert() (*fintracts.Coupons, error) {
	fc := &fintracts.Coupons{
		Rate:  c.Rate,
		Dates: []fintracts.Date{},
	}

	for _, date := range c.Dates {
		fd, err := date.convert()
		if err != nil {
			return nil, err
		}
		fc.Dates = append(fc.Dates, fd)
	}

	return fc, nil
}

func (a InterestRateSwap) convert() (*fintracts.InterestRateSwap, error) {
	fa := &fintracts.InterestRateSwap{
		NotationalAmount: a.NotationalAmount.convert(),
		Interest:         []fintracts.InterestPayment{},
	}

	md, err := a.MaturityDate.convert()
	if err != nil {
		return nil, err
	}
	fa.MaturityDate = md

	ed, err := a.EffectiveDate.convert()
	if err != nil {
		return nil, err
	}
	fa.EffectiveDate = ed

	for _, ip := range a.Interest {
		fip, err := ip.convert()
		if err != nil {
			return nil, err
		}
		fa.Interest = append(fa.Interest, *fip)
	}

	return fa, nil
}

func (a CurrencySwap) convert() (*fintracts.CurrencySwap, error) {
	fa := &fintracts.CurrencySwap{
		PayerA:              a.PayerA,
		PayerB:              a.PayerB,
		PrincipalA:          a.PrincipalA.convert(),
		PrincipalB:          a.PrincipalB.convert(),
		ImpliedExchangeRate: a.ImpliedExchangeRate.convert(),
	}

	if a.EndExchangeRate != nil {
		er := a.EndExchangeRate.convert()
		fa.EndExchangeRate = &er
	}

	md, err := a.MaturityDate.convert()
	if err != nil {
		return nil, err
	}
	fa.MaturityDate = md

	ed, err := a.EffectiveDate.convert()
	if err != nil {
		return nil, err
	}
	fa.EffectiveDate = ed

	if a.Interest != nil {
		for _, ip := range a.Interest {
			fip, err := ip.convert()
			if err != nil {
				return nil, err
			}
			fa.Interest = append(fa.Interest, *fip)
		}
	}

	return fa, nil
}

func (d Date) convert() (fintracts.Date, error) {
	day := fmt.Sprintf("%02d", d.Day)
	year := strconv.Itoa(d.Year)

	date, err := time.Parse("02 January 2006", day+" "+d.Month+" "+year)
	if err != nil {
		return fintracts.Date{}, err
	}

	return fintracts.Date(date), nil
}

func (s Signature) convert() (*fintracts.Signature, error) {
	date, err := s.Date.convert()
	if err != nil {
		return nil, err
	}

	sort.Strings(s.Parties)

	return &fintracts.Signature{
		Parties: s.Parties,
		Date:    date,
	}, nil
}
