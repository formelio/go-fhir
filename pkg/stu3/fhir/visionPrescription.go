package fhir

import (
	"bytes"
	"encoding/json"
)

// VisionPrescription is documented here http://hl7.org/fhir/StructureDefinition/VisionPrescription
type VisionPrescription struct {
	Id                    *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                         `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                       `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                       `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative                    `bson:"text,omitempty" json:"text,omitempty"`
	RawContained          []json.RawMessage             `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained             []IResource                   `bson:"-,omitempty" json:"-,omitempty"`
	Extension             []*Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []*Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []*Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                *string                       `bson:"status,omitempty" json:"status,omitempty"`
	Patient               *Reference                    `bson:"patient,omitempty" json:"patient,omitempty"`
	Encounter             *Reference                    `bson:"encounter,omitempty" json:"encounter,omitempty"`
	DateWritten           *string                       `bson:"dateWritten,omitempty" json:"dateWritten,omitempty"`
	Prescriber            *Reference                    `bson:"prescriber,omitempty" json:"prescriber,omitempty"`
	ReasonCodeableConcept *CodeableConcept              `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference                    `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Dispense              []*VisionPrescriptionDispense `bson:"dispense,omitempty" json:"dispense,omitempty"`
}
type VisionPrescriptionDispense struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Product           *CodeableConcept `bson:"product,omitempty" json:"product,omitempty"`
	Eye               *VisionEyes      `bson:"eye,omitempty" json:"eye,omitempty"`
	Sphere            *float64         `bson:"sphere,omitempty" json:"sphere,omitempty"`
	Cylinder          *float64         `bson:"cylinder,omitempty" json:"cylinder,omitempty"`
	Axis              *int             `bson:"axis,omitempty" json:"axis,omitempty"`
	Prism             *float64         `bson:"prism,omitempty" json:"prism,omitempty"`
	Base              *VisionBase      `bson:"base,omitempty" json:"base,omitempty"`
	Add               *float64         `bson:"add,omitempty" json:"add,omitempty"`
	Power             *float64         `bson:"power,omitempty" json:"power,omitempty"`
	BackCurve         *float64         `bson:"backCurve,omitempty" json:"backCurve,omitempty"`
	Diameter          *float64         `bson:"diameter,omitempty" json:"diameter,omitempty"`
	Duration          *Quantity        `bson:"duration,omitempty" json:"duration,omitempty"`
	Color             *string          `bson:"color,omitempty" json:"color,omitempty"`
	Brand             *string          `bson:"brand,omitempty" json:"brand,omitempty"`
	Note              []*Annotation    `bson:"note,omitempty" json:"note,omitempty"`
}

// OtherVisionPrescription is a helper type to use the default implementations of Marshall and Unmarshal
type OtherVisionPrescription VisionPrescription

// MarshalJSON marshals the given VisionPrescription as JSON into a byte slice
func (r VisionPrescription) MarshalJSON() ([]byte, error) {
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherVisionPrescription
	}{
		OtherVisionPrescription: OtherVisionPrescription(r),
		ResourceType:            "VisionPrescription",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into VisionPrescription
func (r *VisionPrescription) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherVisionPrescription)(r)); err != nil {
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
func (r VisionPrescription) GetResourceType() ResourceType {
	return ResourceTypeVisionPrescription
}
