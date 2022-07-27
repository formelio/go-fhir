package fhir

import "encoding/json"

// Specimen is documented here http://hl7.org/fhir/StructureDefinition/Specimen
type Specimen struct {
	Id                  *string              `bson:"id" json:"id"`
	Meta                *Meta                `bson:"meta" json:"meta"`
	ImplicitRules       *string              `bson:"implicitRules" json:"implicitRules"`
	Language            *string              `bson:"language" json:"language"`
	Text                *Narrative           `bson:"text" json:"text"`
	RawContained        []json.RawMessage    `bson:"contained" json:"contained"`
	Contained           []IResource          `bson:"-" json:"-"`
	Extension           []Extension          `bson:"extension" json:"extension"`
	ModifierExtension   []Extension          `bson:"modifierExtension" json:"modifierExtension"`
	Identifier          []Identifier         `bson:"identifier" json:"identifier"`
	AccessionIdentifier *Identifier          `bson:"accessionIdentifier" json:"accessionIdentifier"`
	Status              *SpecimenStatus      `bson:"status" json:"status"`
	Type                *CodeableConcept     `bson:"type" json:"type"`
	Subject             Reference            `bson:"subject,omitempty" json:"subject,omitempty"`
	ReceivedTime        *string              `bson:"receivedTime" json:"receivedTime"`
	Parent              []Reference          `bson:"parent" json:"parent"`
	Request             []Reference          `bson:"request" json:"request"`
	Collection          *SpecimenCollection  `bson:"collection" json:"collection"`
	Processing          []SpecimenProcessing `bson:"processing" json:"processing"`
	Container           []SpecimenContainer  `bson:"container" json:"container"`
	Note                []Annotation         `bson:"note" json:"note"`
}
type SpecimenCollection struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Collector         *Reference       `bson:"collector" json:"collector"`
	CollectedDateTime *string          `bson:"collectedDateTime,omitempty" json:"collectedDateTime,omitempty"`
	CollectedPeriod   *Period          `bson:"collectedPeriod,omitempty" json:"collectedPeriod,omitempty"`
	Quantity          *Quantity        `bson:"quantity" json:"quantity"`
	Method            *CodeableConcept `bson:"method" json:"method"`
	BodySite          *CodeableConcept `bson:"bodySite" json:"bodySite"`
}
type SpecimenProcessing struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Description       *string          `bson:"description" json:"description"`
	Procedure         *CodeableConcept `bson:"procedure" json:"procedure"`
	Additive          []Reference      `bson:"additive" json:"additive"`
	TimeDateTime      *string          `bson:"timeDateTime,omitempty" json:"timeDateTime,omitempty"`
	TimePeriod        *Period          `bson:"timePeriod,omitempty" json:"timePeriod,omitempty"`
}
type SpecimenContainer struct {
	Id                      *string          `bson:"id" json:"id"`
	Extension               []Extension      `bson:"extension" json:"extension"`
	ModifierExtension       []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Identifier              []Identifier     `bson:"identifier" json:"identifier"`
	Description             *string          `bson:"description" json:"description"`
	Type                    *CodeableConcept `bson:"type" json:"type"`
	Capacity                *Quantity        `bson:"capacity" json:"capacity"`
	SpecimenQuantity        *Quantity        `bson:"specimenQuantity" json:"specimenQuantity"`
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
