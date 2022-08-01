package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Use is documented here http://hl7.org/fhir/ValueSet/claim-use
type Use int

const (
	UseComplete Use = iota
	UseProposed
	UseExploratory
	UseOther
)

func (code Use) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *Use) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "complete":
		*code = UseComplete
	case "proposed":
		*code = UseProposed
	case "exploratory":
		*code = UseExploratory
	case "other":
		*code = UseOther
	default:
		return fmt.Errorf("unknown Use code `%s`", s)
	}
	return nil
}
func (code Use) String() string {
	return code.Code()
}
func (code Use) Code() string {
	switch code {
	case UseComplete:
		return "complete"
	case UseProposed:
		return "proposed"
	case UseExploratory:
		return "exploratory"
	case UseOther:
		return "other"
	}
	return "<unknown>"
}
func (code Use) Display() string {
	switch code {
	case UseComplete:
		return "Complete"
	case UseProposed:
		return "Proposed"
	case UseExploratory:
		return "Exploratory"
	case UseOther:
		return "Other"
	}
	return "<unknown>"
}
func (code Use) Definition() string {
	switch code {
	case UseComplete:
		return "The treatment is complete and this represents a Claim for the services."
	case UseProposed:
		return "The treatment is proposed and this represents a Pre-authorization for the services."
	case UseExploratory:
		return "The treatment is proposed and this represents a Pre-determination for the services."
	case UseOther:
		return "A locally defined or otherwise resolved status."
	}
	return "<unknown>"
}
