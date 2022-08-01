package fhir

import "encoding/json"

// Specimen is documented here http://hl7.org/fhir/StructureDefinition/Specimen
type Specimen struct {
	Id                  *string              `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string              `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	RawContained        []json.RawMessage    `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained           []IResource          `bson:"-,omitempty" json:"-,omitempty"`
	Extension           []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []Identifier         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	AccessionIdentifier *Identifier          `bson:"accessionIdentifier,omitempty" json:"accessionIdentifier,omitempty"`
	Status              *SpecimenStatus      `bson:"status,omitempty" json:"status,omitempty"`
	Type                *CodeableConcept     `bson:"type,omitempty" json:"type,omitempty"`
	Subject             Reference            `bson:"subject,omitempty" json:"subject,omitempty"`
	ReceivedTime        *string              `bson:"receivedTime,omitempty" json:"receivedTime,omitempty"`
	Parent              []Reference          `bson:"parent,omitempty" json:"parent,omitempty"`
	Request             []Reference          `bson:"request,omitempty" json:"request,omitempty"`
	Collection          *SpecimenCollection  `bson:"collection,omitempty" json:"collection,omitempty"`
	Processing          []SpecimenProcessing `bson:"processing,omitempty" json:"processing,omitempty"`
	Container           []SpecimenContainer  `bson:"container,omitempty" json:"container,omitempty"`
	Note                []Annotation         `bson:"note,omitempty" json:"note,omitempty"`
}
type SpecimenCollection struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Collector         *Reference       `bson:"collector,omitempty" json:"collector,omitempty"`
	CollectedDateTime *string          `bson:"collectedDateTime,omitempty" json:"collectedDateTime,omitempty"`
	CollectedPeriod   *Period          `bson:"collectedPeriod,omitempty" json:"collectedPeriod,omitempty"`
	Quantity          *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Method            *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	BodySite          *CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
}
type SpecimenProcessing struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string          `bson:"description,omitempty" json:"description,omitempty"`
	Procedure         *CodeableConcept `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Additive          []Reference      `bson:"additive,omitempty" json:"additive,omitempty"`
	TimeDateTime      *string          `bson:"timeDateTime,omitempty" json:"timeDateTime,omitempty"`
	TimePeriod        *Period          `bson:"timePeriod,omitempty" json:"timePeriod,omitempty"`
}
type SpecimenContainer struct {
	Id                      *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension               []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier              []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Description             *string          `bson:"description,omitempty" json:"description,omitempty"`
	Type                    *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Capacity                *Quantity        `bson:"capacity,omitempty" json:"capacity,omitempty"`
	SpecimenQuantity        *Quantity        `bson:"specimenQuantity,omitempty" json:"specimenQuantity,omitempty"`
	AdditiveCodeableConcept *CodeableConcept `bson:"additiveCodeableConcept,omitempty" json:"additiveCodeableConcept,omitempty"`
	AdditiveReference       *Reference       `bson:"additiveReference,omitempty" json:"additiveReference,omitempty"`
}

// OtherSpecimen is a helper type to use the default implementations of Marshall and Unmarshal
type OtherSpecimen Specimen

// MarshalJSON marshals the given Specimen as JSON into a byte slice
func (r Specimen) MarshalJSON() ([]byte, error) {
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
		OtherSpecimen
		ResourceType string `json:"resourceType"`
	}{
		OtherSpecimen: OtherSpecimen(r),
		ResourceType:  "Specimen",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Specimen
func (r *Specimen) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherSpecimen)(r)); err != nil {
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
func (r Specimen) GetResourceType() ResourceType {
	return ResourceTypeSpecimen
}
