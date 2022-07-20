package fhir

import "encoding/json"

// VisionPrescription is documented here http://hl7.org/fhir/StructureDefinition/VisionPrescription
type VisionPrescription struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            *string                      `bson:"status,omitempty" json:"status,omitempty"`
	Patient           *Reference                   `bson:"patient,omitempty" json:"patient,omitempty"`
	Encounter         *Reference                   `bson:"encounter,omitempty" json:"encounter,omitempty"`
	DateWritten       *string                      `bson:"dateWritten,omitempty" json:"dateWritten,omitempty"`
	Prescriber        *Reference                   `bson:"prescriber,omitempty" json:"prescriber,omitempty"`
	Dispense          []VisionPrescriptionDispense `bson:"dispense,omitempty" json:"dispense,omitempty"`
}
type VisionPrescriptionDispense struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Product           *CodeableConcept `bson:"product,omitempty" json:"product,omitempty"`
	Eye               *string          `bson:"eye,omitempty" json:"eye,omitempty"`
	Sphere            *string          `bson:"sphere,omitempty" json:"sphere,omitempty"`
	Cylinder          *string          `bson:"cylinder,omitempty" json:"cylinder,omitempty"`
	Axis              *int             `bson:"axis,omitempty" json:"axis,omitempty"`
	Prism             *string          `bson:"prism,omitempty" json:"prism,omitempty"`
	Base              *string          `bson:"base,omitempty" json:"base,omitempty"`
	Add               *string          `bson:"add,omitempty" json:"add,omitempty"`
	Power             *string          `bson:"power,omitempty" json:"power,omitempty"`
	BackCurve         *string          `bson:"backCurve,omitempty" json:"backCurve,omitempty"`
	Diameter          *string          `bson:"diameter,omitempty" json:"diameter,omitempty"`
	Duration          *Quantity        `bson:"duration,omitempty" json:"duration,omitempty"`
	Color             *string          `bson:"color,omitempty" json:"color,omitempty"`
	Brand             *string          `bson:"brand,omitempty" json:"brand,omitempty"`
	Note              []Annotation     `bson:"note,omitempty" json:"note,omitempty"`
}
type OtherVisionPrescription VisionPrescription

// MarshalJSON marshals the given VisionPrescription as JSON into a byte slice
func (r VisionPrescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherVisionPrescription
		ResourceType string `json:"resourceType"`
	}{
		OtherVisionPrescription: OtherVisionPrescription(r),
		ResourceType:            "VisionPrescription",
	})
}

// UnmarshalVisionPrescription unmarshals a VisionPrescription.
func UnmarshalVisionPrescription(b []byte) (VisionPrescription, error) {
	var visionPrescription VisionPrescription
	if err := json.Unmarshal(b, &visionPrescription); err != nil {
		return visionPrescription, err
	}
	return visionPrescription, nil
}
