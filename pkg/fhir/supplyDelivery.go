package fhir

import "encoding/json"

// SupplyDelivery is documented here http://hl7.org/fhir/StructureDefinition/SupplyDelivery
type SupplyDelivery struct {
	Id                 *string                     `bson:"id" json:"id"`
	Meta               *Meta                       `bson:"meta" json:"meta"`
	ImplicitRules      *string                     `bson:"implicitRules" json:"implicitRules"`
	Language           *string                     `bson:"language" json:"language"`
	Text               *Narrative                  `bson:"text" json:"text"`
	RawContained       []json.RawMessage           `bson:"contained" json:"contained"`
	Contained          []IResource                 `bson:"-" json:"-"`
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

// OtherSupplyDelivery is a helper type to use the default implementations of Marshall and Unmarshal
type OtherSupplyDelivery SupplyDelivery

// MarshalJSON marshals the given SupplyDelivery as JSON into a byte slice
func (r SupplyDelivery) MarshalJSON() ([]byte, error) {
	// If the field has contained resources, we need to marshal them individually and store them in .RawContained
	if len(r.Contained) > 0 {
		var err error
		r.RawContained = make([]json.RawMessage, len(r.Contained))
		for i, contained := range r.Contained {
			r.RawContained[i], err = json.Marshal(contained)
			if err != nil {
				return nil, err
			}
		}
	}
	return json.Marshal(struct {
		OtherSupplyDelivery
		ResourceType string `json:"resourceType"`
	}{
		OtherSupplyDelivery: OtherSupplyDelivery(r),
		ResourceType:        "SupplyDelivery",
	})
}

// UnmarshalJSON unmarshals the given byte slice into SupplyDelivery
func (r *SupplyDelivery) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherSupplyDelivery)(r)); err != nil {
		return err
	}
	// If the field has contained resources, we need to unmarshal them individually and store them in .Contained
	if len(r.RawContained) > 0 {
		var err error
		r.Contained = make([]IResource, len(r.RawContained))
		for i, rawContained := range r.RawContained {
			r.Contained[i], err = UnmarshalResource(rawContained)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Returns the resourceType of the resource, makes this resource an instance of IResource
func (r SupplyDelivery) GetResourceType() ResourceType {
	return ResourceTypeSupplyDelivery
}
