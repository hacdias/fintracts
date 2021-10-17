package fintracts

import (
	"fmt"
	"sort"
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

func validateDifferentParties(a, b string) error {
	if a != b {
		return nil
	}

	return fmt.Errorf("expected '%s' and '%s' to be different parties", a, b)
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
