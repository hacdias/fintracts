package validator

import (
	"fmt"
	"time"

	"github.com/hacdias/fintracts"
	"go.uber.org/multierr"
)

func (v *validator) validateDate(d fintracts.Date) error {
	date := time.Time(d)

	if date.IsZero() {
		return fmt.Errorf("signature has invalid date")
	}

	return nil
}

func (v *validator) validateParty(p fintracts.Party) error {
	var err error

	if p.Name == "" {
		err = multierr.Append(err, fmt.Errorf("party name cannot be empty"))
	}

	if p.Identifier == "" {
		err = multierr.Append(err, fmt.Errorf("party identifier cannot be empty"))
	}

	return err
}

func (v *validator) validateSignature(s fintracts.Signature) error {
	var err error

	if len(s.Parties) == 0 {
		err = multierr.Append(err, fmt.Errorf("signature must have one or more parties"))
	}

	return multierr.Append(err, v.validateDate(s.Date))
}

func (v *validator) validateCurrency(c fintracts.Currency) error {
	cur := string(c)

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

func (v *validator) validateMoney(m fintracts.Money) error {
	var err error

	if m.Amount <= 0 {
		err = multierr.Append(err, fmt.Errorf("amount must be larger than 0"))
	}

	return multierr.Append(err, v.validateCurrency(m.Currency))
}

func (v *validator) validateExchangeRate(e fintracts.ExchangeRate) error {
	err := multierr.Combine(
		v.validateCurrency(e.BaseCurrency),
		v.validateCurrency(e.CounterCurrency),
	)

	if e.Rate <= 0 {
		err = multierr.Append(err, fmt.Errorf("exchange rate cannot be below or equal to 0"))
	}

	return err
}

func (v *validator) validateInterestPayment(i fintracts.InterestPayment) error {
	var err error

	if i.FixedRate != 0 && i.RateOption != "" {
		err = multierr.Append(err, fmt.Errorf("fixed rate cannot be used with an interest rate option"))
	}

	if i.InitialRate != 0 && i.RateOption == "" {
		err = multierr.Append(err, fmt.Errorf("floating rate must have an interest rate option attached"))
	}

	for _, date := range i.Dates {
		err = multierr.Append(err, v.validateDate(date))
		err = multierr.Append(err, v.validateAfterSignatures(date))
	}

	err = multierr.Append(err, v.validatePartyExists(i.Payer))
	err = multierr.Append(err, v.validatePartyExists(i.Receiver))
	err = multierr.Append(err, v.validateDifferentParties(i.Payer, i.Receiver))
	return err
}
