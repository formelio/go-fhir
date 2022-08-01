package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// AdverseEventCausality is documented here http://hl7.org/fhir/ValueSet/adverse-event-causality
type AdverseEventCausality int

const (
	AdverseEventCausalityCausality1 AdverseEventCausality = iota
	AdverseEventCausalityCausality2
)

func (code AdverseEventCausality) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AdverseEventCausality) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "causality1":
		*code = AdverseEventCausalityCausality1
	case "causality2":
		*code = AdverseEventCausalityCausality2
	default:
		return fmt.Errorf("unknown AdverseEventCausality code `%s`", s)
	}
	return nil
}
func (code AdverseEventCausality) String() string {
	return code.Code()
}
func (code AdverseEventCausality) Code() string {
	switch code {
	case AdverseEventCausalityCausality1:
		return "causality1"
	case AdverseEventCausalityCausality2:
		return "causality2"
	}
	return "<unknown>"
}
func (code AdverseEventCausality) Display() string {
	switch code {
	case AdverseEventCausalityCausality1:
		return "causality1 placeholder"
	case AdverseEventCausalityCausality2:
		return "causality2 placeholder"
	}
	return "<unknown>"
}
func (code AdverseEventCausality) Definition() string {
	switch code {
	}
	return "<unknown>"
}
