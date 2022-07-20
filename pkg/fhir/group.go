package fhir

import "encoding/json"

// Group is documented here http://hl7.org/fhir/StructureDefinition/Group
type Group struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active            *bool                 `bson:"active,omitempty" json:"active,omitempty"`
	Type              string                `bson:"type" json:"type"`
	Actual            bool                  `bson:"actual" json:"actual"`
	Code              *CodeableConcept      `bson:"code,omitempty" json:"code,omitempty"`
	Name              *string               `bson:"name,omitempty" json:"name,omitempty"`
	Quantity          *int                  `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Characteristic    []GroupCharacteristic `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	Member            []GroupMember         `bson:"member,omitempty" json:"member,omitempty"`
}
type GroupCharacteristic struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept `bson:"code" json:"code"`
	Exclude           bool            `bson:"exclude" json:"exclude"`
	Period            *Period         `bson:"period,omitempty" json:"period,omitempty"`
}
type GroupMember struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Period            *Period     `bson:"period,omitempty" json:"period,omitempty"`
	Inactive          *bool       `bson:"inactive,omitempty" json:"inactive,omitempty"`
}
type OtherGroup Group

// MarshalJSON marshals the given Group as JSON into a byte slice
func (r Group) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGroup
		ResourceType string `json:"resourceType"`
	}{
		OtherGroup:   OtherGroup(r),
		ResourceType: "Group",
	})
}

// UnmarshalGroup unmarshals a Group.
func UnmarshalGroup(b []byte) (Group, error) {
	var group Group
	if err := json.Unmarshal(b, &group); err != nil {
		return group, err
	}
	return group, nil
}
