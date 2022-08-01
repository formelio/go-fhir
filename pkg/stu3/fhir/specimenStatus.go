package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SpecimenStatus is documented here http://hl7.org/fhir/ValueSet/specimen-status
type SpecimenStatus int

const (
	SpecimenStatusAvailable SpecimenStatus = iota
	SpecimenStatusUnavailable
	SpecimenStatusUnsatisfactory
	SpecimenStatusEnteredInError
)

func (code SpecimenStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SpecimenStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "available":
		*code = SpecimenStatusAvailable
	case "unavailable":
		*code = SpecimenStatusUnavailable
	case "unsatisfactory":
		*code = SpecimenStatusUnsatisfactory
	case "entered-in-error":
		*code = SpecimenStatusEnteredInError
	default:
		return fmt.Errorf("unknown SpecimenStatus code `%s`", s)
	}
	return nil
}
func (code SpecimenStatus) String() string {
	return code.Code()
}
func (code SpecimenStatus) Code() string {
	switch code {
	case SpecimenStatusAvailable:
		return "available"
	case SpecimenStatusUnavailable:
		return "unavailable"
	case SpecimenStatusUnsatisfactory:
		return "unsatisfactory"
	case SpecimenStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code SpecimenStatus) Display() string {
	switch code {
	case SpecimenStatusAvailable:
		return "Available"
	case SpecimenStatusUnavailable:
		return "Unavailable"
	case SpecimenStatusUnsatisfactory:
		return "Unsatisfactory"
	case SpecimenStatusEnteredInError:
		return "Entered-in-error"
	}
	return "<unknown>"
}
func (code SpecimenStatus) Definition() string {
	switch code {
	case SpecimenStatusAvailable:
		return "The physical specimen is present and in good condition."
	case SpecimenStatusUnavailable:
		return "There is no physical specimen because it is either lost, destroyed or consumed."
	case SpecimenStatusUnsatisfactory:
		return "The specimen cannot be used because of a quality issue such as a broken container, contamination, or too old."
	case SpecimenStatusEnteredInError:
		return "The specimen was entered in error and therefore nullified."
	}
	return "<unknown>"
}
