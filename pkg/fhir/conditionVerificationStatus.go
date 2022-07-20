package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ConditionVerificationStatus is documented here http://hl7.org/fhir/ValueSet/condition-ver-status
type ConditionVerificationStatus int

const (
	ConditionVerificationStatusProvisional ConditionVerificationStatus = iota
	ConditionVerificationStatusDifferential
	ConditionVerificationStatusConfirmed
	ConditionVerificationStatusRefuted
	ConditionVerificationStatusEnteredInError
	ConditionVerificationStatusUnknown
)

func (code ConditionVerificationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ConditionVerificationStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "provisional":
		*code = ConditionVerificationStatusProvisional
	case "differential":
		*code = ConditionVerificationStatusDifferential
	case "confirmed":
		*code = ConditionVerificationStatusConfirmed
	case "refuted":
		*code = ConditionVerificationStatusRefuted
	case "entered-in-error":
		*code = ConditionVerificationStatusEnteredInError
	case "unknown":
		*code = ConditionVerificationStatusUnknown
	default:
		return fmt.Errorf("unknown ConditionVerificationStatus code `%s`", s)
	}
	return nil
}
func (code ConditionVerificationStatus) String() string {
	return code.Code()
}
func (code ConditionVerificationStatus) Code() string {
	switch code {
	case ConditionVerificationStatusProvisional:
		return "provisional"
	case ConditionVerificationStatusDifferential:
		return "differential"
	case ConditionVerificationStatusConfirmed:
		return "confirmed"
	case ConditionVerificationStatusRefuted:
		return "refuted"
	case ConditionVerificationStatusEnteredInError:
		return "entered-in-error"
	case ConditionVerificationStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code ConditionVerificationStatus) Display() string {
	switch code {
	case ConditionVerificationStatusProvisional:
		return "Provisional"
	case ConditionVerificationStatusDifferential:
		return "Differential"
	case ConditionVerificationStatusConfirmed:
		return "Confirmed"
	case ConditionVerificationStatusRefuted:
		return "Refuted"
	case ConditionVerificationStatusEnteredInError:
		return "Entered In Error"
	case ConditionVerificationStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code ConditionVerificationStatus) Definition() string {
	switch code {
	case ConditionVerificationStatusProvisional:
		return "This is a tentative diagnosis - still a candidate that is under consideration."
	case ConditionVerificationStatusDifferential:
		return "One of a set of potential (and typically mutually exclusive) diagnoses asserted to further guide the diagnostic process and preliminary treatment."
	case ConditionVerificationStatusConfirmed:
		return "There is sufficient diagnostic and/or clinical evidence to treat this as a confirmed condition."
	case ConditionVerificationStatusRefuted:
		return "This condition has been ruled out by diagnostic and clinical evidence."
	case ConditionVerificationStatusEnteredInError:
		return "The statement was entered in error and is not valid."
	case ConditionVerificationStatusUnknown:
		return "The condition status is unknown.  Note that \"unknown\" is a value of last resort and every attempt should be made to provide a meaningful value other than \"unknown\"."
	}
	return "<unknown>"
}
