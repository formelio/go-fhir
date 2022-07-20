package fhir

import "encoding/json"

// SupplyRequest is documented here http://hl7.org/fhir/StructureDefinition/SupplyRequest
type SupplyRequest struct {
	Id                    *string                   `bson:"id" json:"id"`
	Meta                  *Meta                     `bson:"meta" json:"meta"`
	ImplicitRules         *string                   `bson:"implicitRules" json:"implicitRules"`
	Language              *string                   `bson:"language" json:"language"`
	Text                  *Narrative                `bson:"text" json:"text"`
	Contained             []json.RawMessage         `bson:"contained" json:"contained"`
	Extension             []Extension               `bson:"extension" json:"extension"`
	ModifierExtension     []Extension               `bson:"modifierExtension" json:"modifierExtension"`
	Identifier            *Identifier               `bson:"identifier" json:"identifier"`
	Status                *SupplyRequestStatus      `bson:"status" json:"status"`
	Category              *CodeableConcept          `bson:"category" json:"category"`
	Priority              *RequestPriority          `bson:"priority" json:"priority"`
	OrderedItem           *SupplyRequestOrderedItem `bson:"orderedItem" json:"orderedItem"`
	OccurrenceDateTime    *string                   `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod      *Period                   `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceTiming      *Timing                   `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	AuthoredOn            *string                   `bson:"authoredOn" json:"authoredOn"`
	Requester             *SupplyRequestRequester   `bson:"requester" json:"requester"`
	Supplier              []Reference               `bson:"supplier" json:"supplier"`
	ReasonCodeableConcept *CodeableConcept          `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference                `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	DeliverFrom           *Reference                `bson:"deliverFrom" json:"deliverFrom"`
	DeliverTo             *Reference                `bson:"deliverTo" json:"deliverTo"`
}
type SupplyRequestOrderedItem struct {
	Id                  *string          `bson:"id" json:"id"`
	Extension           []Extension      `bson:"extension" json:"extension"`
	ModifierExtension   []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Quantity            Quantity         `bson:"quantity,omitempty" json:"quantity,omitempty"`
	ItemCodeableConcept *CodeableConcept `bson:"itemCodeableConcept,omitempty" json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `bson:"itemReference,omitempty" json:"itemReference,omitempty"`
}
type SupplyRequestRequester struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Agent             Reference   `bson:"agent,omitempty" json:"agent,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf" json:"onBehalfOf"`
}
type OtherSupplyRequest SupplyRequest

// MarshalJSON marshals the given SupplyRequest as JSON into a byte slice
func (r SupplyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSupplyRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherSupplyRequest: OtherSupplyRequest(r),
		ResourceType:       "SupplyRequest",
	})
}

// UnmarshalSupplyRequest unmarshalls a SupplyRequest.
func UnmarshalSupplyRequest(b []byte) (SupplyRequest, error) {
	var supplyRequest SupplyRequest
	if err := json.Unmarshal(b, &supplyRequest); err != nil {
		return supplyRequest, err
	}
	return supplyRequest, nil
}
