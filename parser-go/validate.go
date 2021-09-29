package main

type Validator interface {
	Validate() error
}

// Validate semantically validates the contract and fills
// implicit values that may be required depending on the contract.
func Validate(contract *Contract) error {
	// TODO
	return nil
}
