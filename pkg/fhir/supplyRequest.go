package fhir

import "encoding/json"

// SupplyRequest is documented here http://hl7.org/fhir/StructureDefinition/SupplyRequest
type SupplyRequest struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                   `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            *string                   `bson:"status,omitempty" json:"status,omitempty"`
	Category          *CodeableConcept          `bson:"category,omitempty" json:"category,omitempty"`
	Priority          *string                   `bson:"priority,omitempty" json:"priority,omitempty"`
	OrderedItem       *SupplyRequestOrderedItem `bson:"orderedItem,omitempty" json:"orderedItem,omitempty"`
	AuthoredOn        *string                   `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester         *SupplyRequestRequester   `bson:"requester,omitempty" json:"requester,omitempty"`
	Supplier          []Reference               `bson:"supplier,omitempty" json:"supplier,omitempty"`
}
type SupplyRequestOrderedItem struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Quantity          Quantity    `bson:"quantity" json:"quantity"`
}
type SupplyRequestRequester struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
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

// UnmarshalSupplyRequest unmarshals a SupplyRequest.
func UnmarshalSupplyRequest(b []byte) (SupplyRequest, error) {
	var supplyRequest SupplyRequest
	if err := json.Unmarshal(b, &supplyRequest); err != nil {
		return supplyRequest, err
	}
	return supplyRequest, nil
}
