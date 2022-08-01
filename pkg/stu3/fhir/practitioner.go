package fhir

import (
	"bytes"
	"encoding/json"
)

// Practitioner is documented here http://hl7.org/fhir/StructureDefinition/Practitioner
type Practitioner struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage            `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource                  `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []*Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active            *bool                        `bson:"active,omitempty" json:"active,omitempty"`
	Name              []*HumanName                 `bson:"name,omitempty" json:"name,omitempty"`
	Telecom           []*ContactPoint              `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address           []*Address                   `bson:"address,omitempty" json:"address,omitempty"`
	Gender            *AdministrativeGender        `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate         *string                      `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Photo             []*Attachment                `bson:"photo,omitempty" json:"photo,omitempty"`
	Qualification     []*PractitionerQualification `bson:"qualification,omitempty" json:"qualification,omitempty"`
	Communication     []*CodeableConcept           `bson:"communication,omitempty" json:"communication,omitempty"`
}
type PractitionerQualification struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []*Identifier   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code              CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Period            *Period         `bson:"period,omitempty" json:"period,omitempty"`
	Issuer            *Reference      `bson:"issuer,omitempty" json:"issuer,omitempty"`
}

// OtherPractitioner is a helper type to use the default implementations of Marshall and Unmarshal
type OtherPractitioner Practitioner

// MarshalJSON marshals the given Practitioner as JSON into a byte slice
func (r Practitioner) MarshalJSON() ([]byte, error) {
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
		OtherPractitioner
	}{
		OtherPractitioner: OtherPractitioner(r),
		ResourceType:      "Practitioner",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into Practitioner
func (r *Practitioner) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherPractitioner)(r)); err != nil {
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
func (r Practitioner) GetResourceType() ResourceType {
	return ResourceTypePractitioner
}
