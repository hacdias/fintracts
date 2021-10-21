package validator

import (
	"fmt"
	"sort"
	"time"

	"github.com/hacdias/fintracts"
	"go.uber.org/multierr"
)

type validator struct {
	c *fintracts.Contract

	partiesIds    []string
	lastSignature time.Time
}

// // Validate validates a contract.
func Validate(c *fintracts.Contract) error {
	v := &validator{c: c}

	return v.validate()
}

func (v *validator) validate() error {
	parties := v.getPartiesIds()
	uniqueIds := unique(parties)

	if len(uniqueIds) != len(parties) {
		return fmt.Errorf("different parties cannot have the same identifier")
	}

	sigParties := v.getSigningParties()

	if !equal(parties, sigParties) {
		return fmt.Errorf("parties do not match signing parties")
	}

	var err error

	for _, party := range v.c.Parties {
		err = multierr.Append(err, v.validateParty(party))
	}

	for _, sig := range v.c.Signatures {
		err = multierr.Append(err, v.validateSignature(sig))
	}

	for _, agreement := range v.c.Agreements {
		err = multierr.Append(err, v.validateAgreement(agreement))
	}

	return err
}

func (v *validator) getPartiesIds() []string {
	ids := []string{}
	for _, party := range v.c.Parties {
		ids = append(ids, party.Identifier)
	}
	sort.Strings(ids)
	return ids
}

func (v *validator) getSigningParties() []string {
	ids := []string{}
	for _, sig := range v.c.Signatures {
		ids = append(ids, sig.Parties...)
	}
	sort.Strings(ids)
	return ids
}

func (v *validator) ensureLastSignature() {
	if v.lastSignature.IsZero() {
		date := time.Time{}

		for _, sig := range v.c.Signatures {
			if time.Time(sig.Date).After(date) {
				date = time.Time(sig.Date)
			}
		}

		v.lastSignature = date
	}
}

func (v *validator) ensurePartiesIds() {
	if v.partiesIds == nil {
		v.partiesIds = v.getPartiesIds()
	}
}

func (v *validator) validateAfterSignatures(date fintracts.Date) error {
	v.ensureLastSignature()

	if time.Time(date).Before(v.lastSignature) {
		return fmt.Errorf("date occurs before the last signature of the contract: %s", time.Time(date))
	}

	return nil
}

func (v *validator) validatePartyExists(p string) error {
	v.ensurePartiesIds()

	for _, v := range v.partiesIds {
		if v == p {
			return nil
		}
	}

	return fmt.Errorf("party %s not found", p)
}

func (v *validator) validateDifferentParties(a, b string) error {
	if a != b {
		return nil
	}

	return fmt.Errorf("expected '%s' and '%s' to be different parties", a, b)
}

func (v *validator) validateDateInRange(date, start, end time.Time) error {
	var err error

	if !start.IsZero() && date.After(end) && !date.Equal(end) {
		err = multierr.Append(err, fmt.Errorf("date %s outside of range: [%s, %s]", date, start, end))
	}

	if !end.IsZero() && date.Before(start) && !date.Equal(start) {
		err = multierr.Append(err, fmt.Errorf("date %s outside of range: [%s, %s]", date, start, end))
	}

	return err
}
