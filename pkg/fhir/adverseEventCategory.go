package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// AdverseEventCategory is documented here http://hl7.org/fhir/ValueSet/adverse-event-category
type AdverseEventCategory int

const (
	AdverseEventCategoryAE AdverseEventCategory = iota
	AdverseEventCategoryPAE
)

func (code AdverseEventCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AdverseEventCategory) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "AE":
		*code = AdverseEventCategoryAE
	case "PAE":
		*code = AdverseEventCategoryPAE
	default:
		return fmt.Errorf("unknown AdverseEventCategory code `%s`", s)
	}
	return nil
}
func (code AdverseEventCategory) String() string {
	return code.Code()
}
func (code AdverseEventCategory) Code() string {
	switch code {
	case AdverseEventCategoryAE:
		return "AE"
	case AdverseEventCategoryPAE:
		return "PAE"
	}
	return "<unknown>"
}
func (code AdverseEventCategory) Display() string {
	switch code {
	case AdverseEventCategoryAE:
		return "Adverse Event"
	case AdverseEventCategoryPAE:
		return "Potential Adverse Event"
	}
	return "<unknown>"
}
func (code AdverseEventCategory) Definition() string {
	switch code {
	}
	return "<unknown>"
}
