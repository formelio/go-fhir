package fhir

import "encoding/json"

// SupplyDelivery is documented here http://hl7.org/fhir/StructureDefinition/SupplyDelivery
type SupplyDelivery struct {
	Id                 *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	RawContained       []json.RawMessage           `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained          []IResource                 `bson:"-,omitempty" json:"-,omitempty"`
	Extension          []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         *Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn            []Reference                 `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf             []Reference                 `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status             *SupplyDeliveryStatus       `bson:"status,omitempty" json:"status,omitempty"`
	Patient            *Reference                  `bson:"patient,omitempty" json:"patient,omitempty"`
	Type               *CodeableConcept            `bson:"type,omitempty" json:"type,omitempty"`
	SuppliedItem       *SupplyDeliverySuppliedItem `bson:"suppliedItem,omitempty" json:"suppliedItem,omitempty"`
	OccurrenceDateTime *string                     `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period                     `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing                     `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	Supplier           *Reference                  `bson:"supplier,omitempty" json:"supplier,omitempty"`
	Destination        *Reference                  `bson:"destination,omitempty" json:"destination,omitempty"`
	Receiver           []Reference                 `bson:"receiver,omitempty" json:"receiver,omitempty"`
}
type SupplyDeliverySuppliedItem struct {
	Id                  *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Quantity            *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
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
