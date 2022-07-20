package fhir

import "encoding/json"

// SupplyDelivery is documented here http://hl7.org/fhir/StructureDefinition/SupplyDelivery
type SupplyDelivery struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn           []Reference                 `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Status            *string                     `bson:"status,omitempty" json:"status,omitempty"`
	Patient           *Reference                  `bson:"patient,omitempty" json:"patient,omitempty"`
	Type              *CodeableConcept            `bson:"type,omitempty" json:"type,omitempty"`
	SuppliedItem      *SupplyDeliverySuppliedItem `bson:"suppliedItem,omitempty" json:"suppliedItem,omitempty"`
	Destination       *Reference                  `bson:"destination,omitempty" json:"destination,omitempty"`
	Receiver          []Reference                 `bson:"receiver,omitempty" json:"receiver,omitempty"`
}
type SupplyDeliverySuppliedItem struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Quantity          *Quantity   `bson:"quantity,omitempty" json:"quantity,omitempty"`
}
type OtherSupplyDelivery SupplyDelivery

// MarshalJSON marshals the given SupplyDelivery as JSON into a byte slice
func (r SupplyDelivery) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSupplyDelivery
		ResourceType string `json:"resourceType"`
	}{
		OtherSupplyDelivery: OtherSupplyDelivery(r),
		ResourceType:        "SupplyDelivery",
	})
}

// UnmarshalSupplyDelivery unmarshals a SupplyDelivery.
func UnmarshalSupplyDelivery(b []byte) (SupplyDelivery, error) {
	var supplyDelivery SupplyDelivery
	if err := json.Unmarshal(b, &supplyDelivery); err != nil {
		return supplyDelivery, err
	}
	return supplyDelivery, nil
}
