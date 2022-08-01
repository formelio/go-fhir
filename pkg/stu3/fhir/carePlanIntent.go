package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// CarePlanIntent is documented here http://hl7.org/fhir/ValueSet/care-plan-intent
type CarePlanIntent int

const (
	CarePlanIntentProposal CarePlanIntent = iota
	CarePlanIntentPlan
	CarePlanIntentOrder
	CarePlanIntentOption
)

func (code CarePlanIntent) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CarePlanIntent) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "proposal":
		*code = CarePlanIntentProposal
	case "plan":
		*code = CarePlanIntentPlan
	case "order":
		*code = CarePlanIntentOrder
	case "option":
		*code = CarePlanIntentOption
	default:
		return fmt.Errorf("unknown CarePlanIntent code `%s`", s)
	}
	return nil
}
func (code CarePlanIntent) String() string {
	return code.Code()
}
func (code CarePlanIntent) Code() string {
	switch code {
	case CarePlanIntentProposal:
		return "proposal"
	case CarePlanIntentPlan:
		return "plan"
	case CarePlanIntentOrder:
		return "order"
	case CarePlanIntentOption:
		return "option"
	}
	return "<unknown>"
}
func (code CarePlanIntent) Display() string {
	switch code {
	case CarePlanIntentProposal:
		return "Proposal"
	case CarePlanIntentPlan:
		return "Plan"
	case CarePlanIntentOrder:
		return "Order"
	case CarePlanIntentOption:
		return "Option"
	}
	return "<unknown>"
}
func (code CarePlanIntent) Definition() string {
	switch code {
	case CarePlanIntentProposal:
		return "The care plan is a suggestion made by someone/something that doesn't have an intention to ensure it occurs and without providing an authorization to act"
	case CarePlanIntentPlan:
		return "The care plan represents an intention to ensure something occurs without providing an authorization for others to act"
	case CarePlanIntentOrder:
		return "The care plan represents a request/demand and authorization for action"
	case CarePlanIntentOption:
		return "The care plan represents a component or option for a RequestGroup that establishes timing, conditionality and/or other constraints among a set of requests.\n\nRefer to [[[RequestGroup]]] for additional information on how this status is used"
	}
	return "<unknown>"
}
