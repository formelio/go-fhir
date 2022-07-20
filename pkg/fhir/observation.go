package fhir

import "encoding/json"

// Observation is documented here http://hl7.org/fhir/StructureDefinition/Observation
type Observation struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            string                      `bson:"status" json:"status"`
	Category          []CodeableConcept           `bson:"category,omitempty" json:"category,omitempty"`
	Code              CodeableConcept             `bson:"code" json:"code"`
	Issued            *string                     `bson:"issued,omitempty" json:"issued,omitempty"`
	DataAbsentReason  *CodeableConcept            `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	Interpretation    *CodeableConcept            `bson:"interpretation,omitempty" json:"interpretation,omitempty"`
	Comment           *string                     `bson:"comment,omitempty" json:"comment,omitempty"`
	BodySite          *CodeableConcept            `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Method            *CodeableConcept            `bson:"method,omitempty" json:"method,omitempty"`
	Specimen          *Reference                  `bson:"specimen,omitempty" json:"specimen,omitempty"`
	ReferenceRange    []ObservationReferenceRange `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
	Related           []ObservationRelated        `bson:"related,omitempty" json:"related,omitempty"`
	Component         []ObservationComponent      `bson:"component,omitempty" json:"component,omitempty"`
}
type ObservationReferenceRange struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Low               *Quantity         `bson:"low,omitempty" json:"low,omitempty"`
	High              *Quantity         `bson:"high,omitempty" json:"high,omitempty"`
	Type              *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	AppliesTo         []CodeableConcept `bson:"appliesTo,omitempty" json:"appliesTo,omitempty"`
	Age               *Range            `bson:"age,omitempty" json:"age,omitempty"`
	Text              *string           `bson:"text,omitempty" json:"text,omitempty"`
}
type ObservationRelated struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *string     `bson:"type,omitempty" json:"type,omitempty"`
}
type ObservationComponent struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept             `bson:"code" json:"code"`
	DataAbsentReason  *CodeableConcept            `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	Interpretation    *CodeableConcept            `bson:"interpretation,omitempty" json:"interpretation,omitempty"`
	ReferenceRange    []ObservationReferenceRange `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
}
type OtherObservation Observation

// MarshalJSON marshals the given Observation as JSON into a byte slice
func (r Observation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherObservation
		ResourceType string `json:"resourceType"`
	}{
		OtherObservation: OtherObservation(r),
		ResourceType:     "Observation",
	})
}

// UnmarshalObservation unmarshals a Observation.
func UnmarshalObservation(b []byte) (Observation, error) {
	var observation Observation
	if err := json.Unmarshal(b, &observation); err != nil {
		return observation, err
	}
	return observation, nil
}
