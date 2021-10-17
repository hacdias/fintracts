package fintracts

import (
	"encoding/json"
)

// FromJSON converts a JSON byte array into a contract.
func FromJSON(data []byte) (*Contract, error) {
	contract := &Contract{}
	return contract, json.Unmarshal(data, contract)
}
