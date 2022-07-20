package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ConsentState is documented here http://hl7.org/fhir/ValueSet/consent-state-codes
type ConsentState int

const (
	ConsentStateDraft ConsentState = iota
	ConsentStateProposed
	ConsentStateActive
	ConsentStateRejected
	ConsentStateInactive
	ConsentStateEnteredInError
)

func (code ConsentState) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ConsentState) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "draft":
		*code = ConsentStateDraft
	case "proposed":
		*code = ConsentStateProposed
	case "active":
		*code = ConsentStateActive
	case "rejected":
		*code = ConsentStateRejected
	case "inactive":
		*code = ConsentStateInactive
	case "entered-in-error":
		*code = ConsentStateEnteredInError
	default:
		return fmt.Errorf("unknown ConsentState code `%s`", s)
	}
	return nil
}
func (code ConsentState) String() string {
	return code.Code()
}
func (code ConsentState) Code() string {
	switch code {
	case ConsentStateDraft:
		return "draft"
	case ConsentStateProposed:
		return "proposed"
	case ConsentStateActive:
		return "active"
	case ConsentStateRejected:
		return "rejected"
	case ConsentStateInactive:
		return "inactive"
	case ConsentStateEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code ConsentState) Display() string {
	switch code {
	case ConsentStateDraft:
		return "Pending"
	case ConsentStateProposed:
		return "Proposed"
	case ConsentStateActive:
		return "Active"
	case ConsentStateRejected:
		return "Rejected"
	case ConsentStateInactive:
		return "Inactive"
	case ConsentStateEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code ConsentState) Definition() string {
	switch code {
	case ConsentStateDraft:
		return "The consent is in development or awaiting use but is not yet intended to be acted upon."
	case ConsentStateProposed:
		return "The consent has been proposed but not yet agreed to by all parties. The negotiation stage."
	case ConsentStateActive:
		return "The consent is to be followed and enforced."
	case ConsentStateRejected:
		return "The consent has been rejected by one or more of the parties."
	case ConsentStateInactive:
		return "The consent is terminated or replaced."
	case ConsentStateEnteredInError:
		return "The consent was created wrongly (e.g. wrong patient) and should be ignored"
	}
	return "<unknown>"
}
