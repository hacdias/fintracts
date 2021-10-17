package fintracts

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/multierr"
)

func (c *Contract) validate(fix bool) error {
	parties := c.getPartiesIds()
	uniqueIds := unique(parties)

	if len(uniqueIds) != len(parties) {
		return fmt.Errorf("different parties cannot have the same identifier")
	}

	sigParties := c.getSigningParties()

	if !equal(parties, sigParties) {
		return fmt.Errorf("parties do not match signing parties")
	}

	var err error

	for _, party := range c.Parties {
		err = multierr.Append(err, party.validate())
	}

	for _, sig := range c.Signatures {
		err = multierr.Append(err, sig.validate())
	}

	for _, agreement := range c.Agreements {
		err = multierr.Append(err, agreement.validate(c, fix))
	}

	return err
}

func (d *Date) validate() error {
	date := time.Time(*d)

	if date.IsZero() {
		return fmt.Errorf("signature has invalid date")
	}

	return nil
}

func (p *Party) validate() error {
	var err error

	if p.Name == "" {
		err = multierr.Append(err, fmt.Errorf("party name cannot be empty"))
	}

	if p.Identifier == "" {
		err = multierr.Append(err, fmt.Errorf("party identifier cannot be empty"))
	}

	return err
}

func (s *Signature) validate() error {
	var err error

	if len(s.Parties) == 0 {
		err = multierr.Append(err, fmt.Errorf("signature must have one or more parties"))
	}

	return multierr.Append(err, s.Date.validate())
}

func (c *Currency) validate(fix bool) error {
	cur := string(*c)
	if fix {
		cur = strings.ToUpper(cur)
		*c = Currency(cur)
	}

	found := false
	for _, c := range currencies {
		if c == cur {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("unknown currency: %s", cur)
	}

	return nil
}

// Source: https://developers.google.com/public-data/docs/canonical/currencies_csv
var currencies = []string{
	"AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN", "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD",
	"BND", "BOB BOV", "BRL", "BSD", "BWP", "BYR", "BZD", "CAD", "CDF", "CHF", "CLP CLF", "CNY", "COP COU", "CRC", "CUP CUC",
	"CVE", "CZK", "DJF", "DKK", "DOP", "DZD", "EEK", "EGP", "ERN", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GHS", "GIP", "GMD",
	"GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG USD", "HUF", "IDR", "ILS", "INR", "INR BTN", "IQD", "IRR", "ISK", "JMD", "JOD",
	"JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LTL", "LVL", "LYD", "MAD", "MDL",
	"MGA", "MKD", "MMK", "MNT", "MOP", "MRO", "MUR", "MVR", "MWK", "MXN MXV", "MYR", "MZN", "NGN", "NIO", "NOK", "NPR", "NZD", "OMR",
	"PAB USD", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR", "SBD", "SCR", "SDG", "SEK", "SGD",
	"SHP", "SLL", "SOS", "SRD", "STD", "SVC USD", "SYP", "SZL", "THB", "TJS", "TMT", "TND", "TOP", "TRY", "TTD", "TWD", "TZS", "UAH",
	"UGX", "USD", "UYU UYI", "UZS", "VEF", "VND", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XFU",
	"XOF", "XPD", "XPF", "XPT", "XTS", "YER", "ZAR", "ZAR LSL", "ZAR NAD", "ZMK", "ZWL",
}

func (m *Money) validate(fix bool) error {
	var err error

	if m.Amount <= 0 {
		err = multierr.Append(err, fmt.Errorf("amount must be larger than 0"))
	}

	return multierr.Append(err, m.Currency.validate(fix))
}

func (e *ExchangeRate) validate(fix bool) error {
	err := multierr.Combine(
		e.BaseCurrency.validate(fix),
		e.CounterCurrency.validate(fix),
	)

	if e.Rate <= 0 {
		err = multierr.Append(err, fmt.Errorf("exchange rate cannot be below or equal to 0"))
	}

	return err
}

func (i *InterestPayment) validate(c *Contract) error {
	var err error

	if i.FixedRate != 0 && i.RateOption != "" {
		err = multierr.Append(err, fmt.Errorf("fixed rate cannot be used with an interest rate option"))
	}

	if i.InitialRate != 0 && i.RateOption == "" {
		err = multierr.Append(err, fmt.Errorf("floating rate must have an interest rate option attached"))
	}

	for _, date := range i.Dates {
		err = multierr.Append(err, date.validate())
	}

	err = multierr.Append(err, c.validatePartyExists(i.Payer))
	return err
}
