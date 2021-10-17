package fintracts

import (
	"fmt"
	"sort"
	"time"

	"go.uber.org/multierr"
)

func Validate(c *Contract, fix bool) error {
	return c.validate(fix)
}

func (c *Contract) getPartiesIds() []string {
	ids := []string{}
	for _, party := range c.Parties {
		ids = append(ids, party.Identifier)
	}
	sort.Strings(ids)
	return ids
}

func (c *Contract) getSigningParties() []string {
	ids := []string{}
	for _, sig := range c.Signatures {
		ids = append(ids, sig.Parties...)
	}
	sort.Strings(ids)
	return ids
}

func (c *Contract) validatePartyExists(p string) error {
	if c.partiesIds == nil {
		c.partiesIds = c.getPartiesIds()
	}

	for _, v := range c.partiesIds {
		if v == p {
			return nil
		}
	}

	return fmt.Errorf("party %s not found", p)
}

func (c *Contract) validateAfterSignatures(date Date) error {
	if c.lastSignature.IsZero() {
		date := time.Time{}

		for _, sig := range c.Signatures {
			if time.Time(sig.Date).After(date) {
				date = time.Time(sig.Date)
			}
		}

		c.lastSignature = date
	}

	if time.Time(date).Before(c.lastSignature) {
		return fmt.Errorf("date occurs before the last signature of the contract: %s", time.Time(date))
	}

	return nil
}

func validateDifferentParties(a, b string) error {
	if a != b {
		return nil
	}

	return fmt.Errorf("expected '%s' and '%s' to be different parties", a, b)
}

func validateDateInRange(date, start, end time.Time) error {
	var err error

	if !start.IsZero() && date.After(end) && !date.Equal(end) {
		err = multierr.Append(err, fmt.Errorf("date %s outside of range: [%s, %s]", date, start, end))
	}

	if !end.IsZero() && date.Before(start) && !date.Equal(start) {
		err = multierr.Append(err, fmt.Errorf("date %s outside of range: [%s, %s]", date, start, end))
	}

	return err
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

func unique(arr []string) []string {
	occurred := map[string]bool{}
	result := []string{}
	for e := range arr {

		if !occurred[arr[e]] {
			occurred[arr[e]] = true
			result = append(result, arr[e])
		}
	}

	return result
}
