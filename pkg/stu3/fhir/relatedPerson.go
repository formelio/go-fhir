package fhir

import (
	"bytes"
	"encoding/json"
)

// RelatedPerson is documented here http://hl7.org/fhir/StructureDefinition/RelatedPerson
type RelatedPerson struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage     `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource           `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []*Identifier         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active            *bool                 `bson:"active,omitempty" json:"active,omitempty"`
	Patient           Reference             `bson:"patient,omitempty" json:"patient,omitempty"`
	Relationship      *CodeableConcept      `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Name              []*HumanName          `bson:"name,omitempty" json:"name,omitempty"`
	Telecom           []*ContactPoint       `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Gender            *AdministrativeGender `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate         *string               `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Address           []*Address            `bson:"address,omitempty" json:"address,omitempty"`
	Photo             []*Attachment         `bson:"photo,omitempty" json:"photo,omitempty"`
	Period            *Period               `bson:"period,omitempty" json:"period,omitempty"`
}

// OtherRelatedPerson is a helper type to use the default implementations of Marshall and Unmarshal
type OtherRelatedPerson RelatedPerson

// MarshalJSON marshals the given RelatedPerson as JSON into a byte slice
func (r RelatedPerson) MarshalJSON() ([]byte, error) {
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
		OtherRelatedPerson
	}{
		OtherRelatedPerson: OtherRelatedPerson(r),
		ResourceType:       "RelatedPerson",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into RelatedPerson
func (r *RelatedPerson) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherRelatedPerson)(r)); err != nil {
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
func (r RelatedPerson) GetResourceType() ResourceType {
	return ResourceTypeRelatedPerson
}
