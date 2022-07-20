package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// FlagStatus is documented here http://hl7.org/fhir/ValueSet/flag-status
type FlagStatus int

const (
	FlagStatusActive FlagStatus = iota
	FlagStatusInactive
	FlagStatusEnteredInError
)

func (code FlagStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *FlagStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = FlagStatusActive
	case "inactive":
		*code = FlagStatusInactive
	case "entered-in-error":
		*code = FlagStatusEnteredInError
	default:
		return fmt.Errorf("unknown FlagStatus code `%s`", s)
	}
	return nil
}
func (code FlagStatus) String() string {
	return code.Code()
}
func (code FlagStatus) Code() string {
	switch code {
	case FlagStatusActive:
		return "active"
	case FlagStatusInactive:
		return "inactive"
	case FlagStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code FlagStatus) Display() string {
	switch code {
	case FlagStatusActive:
		return "Active"
	case FlagStatusInactive:
		return "Inactive"
	case FlagStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code FlagStatus) Definition() string {
	switch code {
	case FlagStatusActive:
		return "A current flag that should be displayed to a user. A system may use the category to determine which roles should view the flag."
	case FlagStatusInactive:
		return "The flag does not need to be displayed any more."
	case FlagStatusEnteredInError:
		return "The flag was added in error, and should no longer be displayed."
	}
	return "<unknown>"
}
