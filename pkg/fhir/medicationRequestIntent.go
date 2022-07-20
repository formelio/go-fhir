package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MedicationRequestIntent is documented here http://hl7.org/fhir/ValueSet/medication-request-intent
type MedicationRequestIntent int

const (
	MedicationRequestIntentProposal MedicationRequestIntent = iota
	MedicationRequestIntentPlan
	MedicationRequestIntentOrder
	MedicationRequestIntentInstanceOrder
)

func (code MedicationRequestIntent) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *MedicationRequestIntent) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "proposal":
		*code = MedicationRequestIntentProposal
	case "plan":
		*code = MedicationRequestIntentPlan
	case "order":
		*code = MedicationRequestIntentOrder
	case "instance-order":
		*code = MedicationRequestIntentInstanceOrder
	default:
		return fmt.Errorf("unknown MedicationRequestIntent code `%s`", s)
	}
	return nil
}
func (code MedicationRequestIntent) String() string {
	return code.Code()
}
func (code MedicationRequestIntent) Code() string {
	switch code {
	case MedicationRequestIntentProposal:
		return "proposal"
	case MedicationRequestIntentPlan:
		return "plan"
	case MedicationRequestIntentOrder:
		return "order"
	case MedicationRequestIntentInstanceOrder:
		return "instance-order"
	}
	return "<unknown>"
}
func (code MedicationRequestIntent) Display() string {
	switch code {
	case MedicationRequestIntentProposal:
		return "Proposal"
	case MedicationRequestIntentPlan:
		return "Plan"
	case MedicationRequestIntentOrder:
		return "Order"
	case MedicationRequestIntentInstanceOrder:
		return "Instance Order"
	}
	return "<unknown>"
}
func (code MedicationRequestIntent) Definition() string {
	switch code {
	case MedicationRequestIntentProposal:
		return "The request is a suggestion made by someone/something that doesn't have an intention to ensure it occurs and without providing an authorization to act"
	case MedicationRequestIntentPlan:
		return "The request represents an intension to ensure something occurs without providing an authorization for others to act"
	case MedicationRequestIntentOrder:
		return "The request represents a request/demand and authorization for action"
	case MedicationRequestIntentInstanceOrder:
		return "The request represents an instance for the particular order, for example a medication administration record."
	}
	return "<unknown>"
}
