package fhir

import "encoding/json"

// Specimen is documented here http://hl7.org/fhir/StructureDefinition/Specimen
type Specimen struct {
	Id                  *string              `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string              `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	Extension           []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []Identifier         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	AccessionIdentifier *Identifier          `bson:"accessionIdentifier,omitempty" json:"accessionIdentifier,omitempty"`
	Status              *string              `bson:"status,omitempty" json:"status,omitempty"`
	Type                *CodeableConcept     `bson:"type,omitempty" json:"type,omitempty"`
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
}
type SpecimenContainer struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Description       *string          `bson:"description,omitempty" json:"description,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Capacity          *Quantity        `bson:"capacity,omitempty" json:"capacity,omitempty"`
	SpecimenQuantity  *Quantity        `bson:"specimenQuantity,omitempty" json:"specimenQuantity,omitempty"`
}
type OtherSpecimen Specimen

// MarshalJSON marshals the given Specimen as JSON into a byte slice
func (r Specimen) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSpecimen
		ResourceType string `json:"resourceType"`
	}{
		OtherSpecimen: OtherSpecimen(r),
		ResourceType:  "Specimen",
	})
}

// UnmarshalSpecimen unmarshals a Specimen.
func UnmarshalSpecimen(b []byte) (Specimen, error) {
	var specimen Specimen
	if err := json.Unmarshal(b, &specimen); err != nil {
		return specimen, err
	}
	return specimen, nil
}
