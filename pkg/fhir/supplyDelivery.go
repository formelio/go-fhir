package fhir

import "encoding/json"

// SupplyDelivery is documented here http://hl7.org/fhir/StructureDefinition/SupplyDelivery
type SupplyDelivery struct {
	Id                 *string                     `bson:"id" json:"id"`
	Meta               *Meta                       `bson:"meta" json:"meta"`
	ImplicitRules      *string                     `bson:"implicitRules" json:"implicitRules"`
	Language           *string                     `bson:"language" json:"language"`
	Text               *Narrative                  `bson:"text" json:"text"`
	Contained          []json.RawMessage           `bson:"contained" json:"contained"`
	Extension          []Extension                 `bson:"extension" json:"extension"`
	ModifierExtension  []Extension                 `bson:"modifierExtension" json:"modifierExtension"`
	Identifier         *Identifier                 `bson:"identifier" json:"identifier"`
	BasedOn            []Reference                 `bson:"basedOn" json:"basedOn"`
	PartOf             []Reference                 `bson:"partOf" json:"partOf"`
	Status             *SupplyDeliveryStatus       `bson:"status" json:"status"`
	Patient            *Reference                  `bson:"patient" json:"patient"`
	Type               *CodeableConcept            `bson:"type" json:"type"`
	SuppliedItem       *SupplyDeliverySuppliedItem `bson:"suppliedItem" json:"suppliedItem"`
	OccurrenceDateTime *string                     `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period                     `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing                     `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	Supplier           *Reference                  `bson:"supplier" json:"supplier"`
	Destination        *Reference                  `bson:"destination" json:"destination"`
	Receiver           []Reference                 `bson:"receiver" json:"receiver"`
}
type SupplyDeliverySuppliedItem struct {
	Id                  *string          `bson:"id" json:"id"`
	Extension           []Extension      `bson:"extension" json:"extension"`
	ModifierExtension   []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Quantity            *Quantity        `bson:"quantity" json:"quantity"`
	ItemCodeableConcept *CodeableConcept `bson:"itemCodeableConcept,omitempty" json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `bson:"itemReference,omitempty" json:"itemReference,omitempty"`
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

// UnmarshalSupplyDelivery unmarshalls a SupplyDelivery.
func UnmarshalSupplyDelivery(b []byte) (SupplyDelivery, error) {
	var supplyDelivery SupplyDelivery
	if err := json.Unmarshal(b, &supplyDelivery); err != nil {
		return supplyDelivery, err
	}
	return supplyDelivery, nil
}
