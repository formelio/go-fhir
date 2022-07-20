package fhir

import "encoding/json"

// Specimen is documented here http://hl7.org/fhir/StructureDefinition/Specimen
type Specimen struct {
	Id                  *string              `bson:"id" json:"id"`
	Meta                *Meta                `bson:"meta" json:"meta"`
	ImplicitRules       *string              `bson:"implicitRules" json:"implicitRules"`
	Language            *string              `bson:"language" json:"language"`
	Text                *Narrative           `bson:"text" json:"text"`
	Contained           []json.RawMessage    `bson:"contained" json:"contained"`
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

// UnmarshalSpecimen unmarshalls a Specimen.
func UnmarshalSpecimen(b []byte) (Specimen, error) {
	var specimen Specimen
	if err := json.Unmarshal(b, &specimen); err != nil {
		return specimen, err
	}
	return specimen, nil
}
