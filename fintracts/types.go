package fintracts

import (
	"encoding/json"
	"time"
)

type Contract struct {
	Parties    []Party     `json:"parties"`
	Agreements []Agreement `json:"agreements"`
	Signatures []Signature `json:"signatures"`

	partiesIds    []string
	lastSignature time.Time
}

// String returns a JSON-string representation of a contract.
func (c *Contract) String() (string, error) {
	bytes, err := json.MarshalIndent(c, "", "  ")
	return string(bytes), err
}

type Date time.Time

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	date, err := time.Parse(`"`+time.RFC3339+`"`, string(data))
	*d = Date(date)
	return
}

func (d Date) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(time.RFC3339)+2)
	b = append(b, '"')
	b = time.Time(d).AppendFormat(b, time.RFC3339)
	b = append(b, '"')
	return b, nil
}

type Party struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
}

type Signature struct {
	Parties []string `json:"parties"`
	Date    Date     `json:"date"`
}

type Currency string

type Money struct {
	Currency Currency `json:"currency"`
	Amount   float64  `json:"amount"`
}

type ExchangeRate struct {
	BaseCurrency    Currency `json:"baseCurrency"`
	CounterCurrency Currency `json:"counterCurrency"`
	Rate            float64  `json:"rate"`
}

type InterestPayment struct {
	Payer string `json:"payer"`
	Dates []Date `json:"dates"`

	// Fixed-rate only properties:
	FixedRate   float64 `json:"fixedRate,omitempty"`
	InitialRate float64 `json:"initialRate,omitempty"`

	// Floating-rate only properties:
	RateOption string `json:"rateOption,omitempty"`
}

type Agreement struct {
	BondPurchase     *BondPurchase     `json:"bondPurchase,omitempty"`
	InterestRateSwap *InterestRateSwap `json:"interestRateSwap,omitempty"`
	CurrencySwap     *CurrencySwap     `json:"currencySwap,omitempty"`
}

type BondPurchase struct {
	Issuer       string   `json:"issuer"`
	Underwriter  string   `json:"underwriter"`
	FaceValue    Money    `json:"faceValue"`
	IssuePrice   Money    `json:"issuePrice"`
	MaturityDate Date     `json:"maturityDate"`
	Coupons      *Coupons `json:"coupons,omitempty"`
}

type Coupons struct {
	Rate  float64 `json:"rate"`
	Dates []Date  `json:"dates"`
}

type InterestRateSwap struct {
	NotationalAmount Money             `json:"notationalAmount"`
	EffectiveDate    Date              `json:"effectiveDate"`
	MaturityDate     Date              `json:"maturityDate"`
	Interest         []InterestPayment `json:"interest"`
}

type CurrencySwap struct {
	PayerA              string            `json:"payerA"`
	PayerB              string            `json:"payerB"`
	PrincipalA          Money             `json:"principalA"`
	PrincipalB          Money             `json:"principalB"`
	ImpliedExchangeRate ExchangeRate      `json:"impliedExchangeRate"`
	EndExchangeRate     *ExchangeRate     `json:"endExchangeRate,omitempty"`
	EffectiveDate       Date              `json:"effectiveDate"`
	MaturityDate        Date              `json:"maturityDate"`
	Interest            []InterestPayment `json:"interest,omitempty"`
}
