package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// AccountStatus is documented here http://hl7.org/fhir/ValueSet/account-status
type AccountStatus int

const (
	AccountStatusActive AccountStatus = iota
	AccountStatusInactive
	AccountStatusEnteredInError
)

func (code AccountStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AccountStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = AccountStatusActive
	case "inactive":
		*code = AccountStatusInactive
	case "entered-in-error":
		*code = AccountStatusEnteredInError
	default:
		return fmt.Errorf("unknown AccountStatus code `%s`", s)
	}
	return nil
}
func (code AccountStatus) String() string {
	return code.Code()
}
func (code AccountStatus) Code() string {
	switch code {
	case AccountStatusActive:
		return "active"
	case AccountStatusInactive:
		return "inactive"
	case AccountStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code AccountStatus) Display() string {
	switch code {
	case AccountStatusActive:
		return "Active"
	case AccountStatusInactive:
		return "Inactive"
	case AccountStatusEnteredInError:
		return "Entered in error"
	}
	return "<unknown>"
}
func (code AccountStatus) Definition() string {
	switch code {
	case AccountStatusActive:
		return "This account is active and may be used."
	case AccountStatusInactive:
		return "This account is inactive and should not be used to track financial information."
	case AccountStatusEnteredInError:
		return "This instance should not have been part of this patient's medical record."
	}
	return "<unknown>"
}
