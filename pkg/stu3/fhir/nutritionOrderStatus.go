package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// NutritionOrderStatus is documented here http://hl7.org/fhir/ValueSet/nutrition-request-status
type NutritionOrderStatus int

const (
	NutritionOrderStatusProposed NutritionOrderStatus = iota
	NutritionOrderStatusDraft
	NutritionOrderStatusPlanned
	NutritionOrderStatusRequested
	NutritionOrderStatusActive
	NutritionOrderStatusOnHold
	NutritionOrderStatusCompleted
	NutritionOrderStatusCancelled
	NutritionOrderStatusEnteredInError
)

func (code NutritionOrderStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *NutritionOrderStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "proposed":
		*code = NutritionOrderStatusProposed
	case "draft":
		*code = NutritionOrderStatusDraft
	case "planned":
		*code = NutritionOrderStatusPlanned
	case "requested":
		*code = NutritionOrderStatusRequested
	case "active":
		*code = NutritionOrderStatusActive
	case "on-hold":
		*code = NutritionOrderStatusOnHold
	case "completed":
		*code = NutritionOrderStatusCompleted
	case "cancelled":
		*code = NutritionOrderStatusCancelled
	case "entered-in-error":
		*code = NutritionOrderStatusEnteredInError
	default:
		return fmt.Errorf("unknown NutritionOrderStatus code `%s`", s)
	}
	return nil
}
func (code NutritionOrderStatus) String() string {
	return code.Code()
}
func (code NutritionOrderStatus) Code() string {
	switch code {
	case NutritionOrderStatusProposed:
		return "proposed"
	case NutritionOrderStatusDraft:
		return "draft"
	case NutritionOrderStatusPlanned:
		return "planned"
	case NutritionOrderStatusRequested:
		return "requested"
	case NutritionOrderStatusActive:
		return "active"
	case NutritionOrderStatusOnHold:
		return "on-hold"
	case NutritionOrderStatusCompleted:
		return "completed"
	case NutritionOrderStatusCancelled:
		return "cancelled"
	case NutritionOrderStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code NutritionOrderStatus) Display() string {
	switch code {
	case NutritionOrderStatusProposed:
		return "Proposed"
	case NutritionOrderStatusDraft:
		return "Draft"
	case NutritionOrderStatusPlanned:
		return "Planned"
	case NutritionOrderStatusRequested:
		return "Requested"
	case NutritionOrderStatusActive:
		return "Active"
	case NutritionOrderStatusOnHold:
		return "On-Hold"
	case NutritionOrderStatusCompleted:
		return "Completed"
	case NutritionOrderStatusCancelled:
		return "Cancelled"
	case NutritionOrderStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code NutritionOrderStatus) Definition() string {
	switch code {
	case NutritionOrderStatusProposed:
		return "The request has been proposed."
	case NutritionOrderStatusDraft:
		return "The request is in preliminary form prior to being sent."
	case NutritionOrderStatusPlanned:
		return "The request has been planned."
	case NutritionOrderStatusRequested:
		return "The request has been placed."
	case NutritionOrderStatusActive:
		return "The request is 'actionable', but not all actions that are implied by it have occurred yet."
	case NutritionOrderStatusOnHold:
		return "Actions implied by the request have been temporarily halted, but are expected to continue later. May also be called \"suspended\"."
	case NutritionOrderStatusCompleted:
		return "All actions that are implied by the order have occurred and no continuation is planned (this will rarely be made explicit)."
	case NutritionOrderStatusCancelled:
		return "The request has been withdrawn and is no longer actionable."
	case NutritionOrderStatusEnteredInError:
		return "The request was entered in error and voided."
	}
	return "<unknown>"
}
