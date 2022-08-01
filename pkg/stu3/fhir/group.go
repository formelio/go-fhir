package fhir

import (
	"bytes"
	"encoding/json"
)

// Group is documented here http://hl7.org/fhir/StructureDefinition/Group
type Group struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage      `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource            `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []*Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active            *bool                  `bson:"active,omitempty" json:"active,omitempty"`
	Type              GroupType              `bson:"type,omitempty" json:"type,omitempty"`
	Actual            bool                   `bson:"actual,omitempty" json:"actual,omitempty"`
	Code              *CodeableConcept       `bson:"code,omitempty" json:"code,omitempty"`
	Name              *string                `bson:"name,omitempty" json:"name,omitempty"`
	Quantity          *int                   `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Characteristic    []*GroupCharacteristic `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	Member            []*GroupMember         `bson:"member,omitempty" json:"member,omitempty"`
}
type GroupCharacteristic struct {
	Id                   *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code                 CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	Exclude              bool             `bson:"exclude,omitempty" json:"exclude,omitempty"`
	Period               *Period          `bson:"period,omitempty" json:"period,omitempty"`
}
type GroupMember struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Entity            Reference    `bson:"entity,omitempty" json:"entity,omitempty"`
	Period            *Period      `bson:"period,omitempty" json:"period,omitempty"`
	Inactive          *bool        `bson:"inactive,omitempty" json:"inactive,omitempty"`
}

// OtherGroup is a helper type to use the default implementations of Marshall and Unmarshal
type OtherGroup Group

// MarshalJSON marshals the given Group as JSON into a byte slice
func (r Group) MarshalJSON() ([]byte, error) {
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
		OtherGroup
	}{
		OtherGroup:   OtherGroup(r),
		ResourceType: "Group",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into Group
func (r *Group) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherGroup)(r)); err != nil {
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
func (r Group) GetResourceType() ResourceType {
	return ResourceTypeGroup
}
