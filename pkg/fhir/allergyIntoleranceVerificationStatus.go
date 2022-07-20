package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// AllergyIntoleranceVerificationStatus is documented here http://hl7.org/fhir/ValueSet/allergy-verification-status
type AllergyIntoleranceVerificationStatus int

const (
	AllergyIntoleranceVerificationStatusUnconfirmed AllergyIntoleranceVerificationStatus = iota
	AllergyIntoleranceVerificationStatusConfirmed
	AllergyIntoleranceVerificationStatusRefuted
	AllergyIntoleranceVerificationStatusEnteredInError
)

func (code AllergyIntoleranceVerificationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AllergyIntoleranceVerificationStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "unconfirmed":
		*code = AllergyIntoleranceVerificationStatusUnconfirmed
	case "confirmed":
		*code = AllergyIntoleranceVerificationStatusConfirmed
	case "refuted":
		*code = AllergyIntoleranceVerificationStatusRefuted
	case "entered-in-error":
		*code = AllergyIntoleranceVerificationStatusEnteredInError
	default:
		return fmt.Errorf("unknown AllergyIntoleranceVerificationStatus code `%s`", s)
	}
	return nil
}
func (code AllergyIntoleranceVerificationStatus) String() string {
	return code.Code()
}
func (code AllergyIntoleranceVerificationStatus) Code() string {
	switch code {
	case AllergyIntoleranceVerificationStatusUnconfirmed:
		return "unconfirmed"
	case AllergyIntoleranceVerificationStatusConfirmed:
		return "confirmed"
	case AllergyIntoleranceVerificationStatusRefuted:
		return "refuted"
	case AllergyIntoleranceVerificationStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code AllergyIntoleranceVerificationStatus) Display() string {
	switch code {
	case AllergyIntoleranceVerificationStatusUnconfirmed:
		return "Unconfirmed"
	case AllergyIntoleranceVerificationStatusConfirmed:
		return "Confirmed"
	case AllergyIntoleranceVerificationStatusRefuted:
		return "Refuted"
	case AllergyIntoleranceVerificationStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}
func (code AllergyIntoleranceVerificationStatus) Definition() string {
	switch code {
	case AllergyIntoleranceVerificationStatusUnconfirmed:
		return "A low level of certainty about the propensity for a reaction to the identified substance."
	case AllergyIntoleranceVerificationStatusConfirmed:
		return "A high level of certainty about the propensity for a reaction to the identified substance, which may include clinical evidence by testing or rechallenge."
	case AllergyIntoleranceVerificationStatusRefuted:
		return "A propensity for a reaction to the identified substance has been disproven with a high level of clinical certainty, which may include testing or rechallenge, and is refuted."
	case AllergyIntoleranceVerificationStatusEnteredInError:
		return "The statement was entered in error and is not valid."
	}
	return "<unknown>"
}
