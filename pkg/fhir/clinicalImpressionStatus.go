package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ClinicalImpressionStatus is documented here http://hl7.org/fhir/ValueSet/clinical-impression-status
type ClinicalImpressionStatus int

const (
	ClinicalImpressionStatusDraft ClinicalImpressionStatus = iota
	ClinicalImpressionStatusCompleted
	ClinicalImpressionStatusEnteredInError
)

func (code ClinicalImpressionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ClinicalImpressionStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "draft":
		*code = ClinicalImpressionStatusDraft
	case "completed":
		*code = ClinicalImpressionStatusCompleted
	case "entered-in-error":
		*code = ClinicalImpressionStatusEnteredInError
	default:
		return fmt.Errorf("unknown ClinicalImpressionStatus code `%s`", s)
	}
	return nil
}
func (code ClinicalImpressionStatus) String() string {
	return code.Code()
}
func (code ClinicalImpressionStatus) Code() string {
	switch code {
	case ClinicalImpressionStatusDraft:
		return "draft"
	case ClinicalImpressionStatusCompleted:
		return "completed"
	case ClinicalImpressionStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code ClinicalImpressionStatus) Display() string {
	switch code {
	case ClinicalImpressionStatusDraft:
		return "In progress"
	case ClinicalImpressionStatusCompleted:
		return "Completed"
	case ClinicalImpressionStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code ClinicalImpressionStatus) Definition() string {
	switch code {
	case ClinicalImpressionStatusDraft:
		return "The assessment is still on-going and results are not yet final."
	case ClinicalImpressionStatusCompleted:
		return "The assessment is done and the results are final."
	case ClinicalImpressionStatusEnteredInError:
		return "This assessment was never actually done and the record is erroneous (e.g. Wrong patient)."
	}
	return "<unknown>"
}
