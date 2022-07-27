package fhir

import "encoding/json"

// VisionPrescription is documented here http://hl7.org/fhir/StructureDefinition/VisionPrescription
type VisionPrescription struct {
	Id                    *string                      `bson:"id" json:"id"`
	Meta                  *Meta                        `bson:"meta" json:"meta"`
	ImplicitRules         *string                      `bson:"implicitRules" json:"implicitRules"`
	Language              *string                      `bson:"language" json:"language"`
	Text                  *Narrative                   `bson:"text" json:"text"`
	RawContained          []json.RawMessage            `bson:"contained" json:"contained"`
	Contained             []IResource                  `bson:"-" json:"-"`
	Extension             []Extension                  `bson:"extension" json:"extension"`
	ModifierExtension     []Extension                  `bson:"modifierExtension" json:"modifierExtension"`
	Identifier            []Identifier                 `bson:"identifier" json:"identifier"`
	Status                *string                      `bson:"status" json:"status"`
	Patient               *Reference                   `bson:"patient" json:"patient"`
	Encounter             *Reference                   `bson:"encounter" json:"encounter"`
	DateWritten           *string                      `bson:"dateWritten" json:"dateWritten"`
	Prescriber            *Reference                   `bson:"prescriber" json:"prescriber"`
	ReasonCodeableConcept *CodeableConcept             `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference                   `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Dispense              []VisionPrescriptionDispense `bson:"dispense" json:"dispense"`
}
type VisionPrescriptionDispense struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Product           *CodeableConcept `bson:"product" json:"product"`
	Eye               *VisionEyes      `bson:"eye" json:"eye"`
	Sphere            *float64         `bson:"sphere" json:"sphere"`
	Cylinder          *float64         `bson:"cylinder" json:"cylinder"`
	Axis              *int             `bson:"axis" json:"axis"`
	Prism             *float64         `bson:"prism" json:"prism"`
	Base              *VisionBase      `bson:"base" json:"base"`
	Add               *float64         `bson:"add" json:"add"`
	Power             *float64         `bson:"power" json:"power"`
	BackCurve         *float64         `bson:"backCurve" json:"backCurve"`
	Diameter          *float64         `bson:"diameter" json:"diameter"`
	Duration          *Quantity        `bson:"duration" json:"duration"`
	Color             *string          `bson:"color" json:"color"`
	Brand             *string          `bson:"brand" json:"brand"`
	Note              []Annotation     `bson:"note" json:"note"`
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
	return json.Marshal(struct {
		OtherVisionPrescription
		ResourceType string `json:"resourceType"`
	}{
		OtherVisionPrescription: OtherVisionPrescription(r),
		ResourceType:            "VisionPrescription",
	})
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
