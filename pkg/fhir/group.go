package fhir

import "encoding/json"

// Group is documented here http://hl7.org/fhir/StructureDefinition/Group
type Group struct {
	Id                *string               `bson:"id" json:"id"`
	Meta              *Meta                 `bson:"meta" json:"meta"`
	ImplicitRules     *string               `bson:"implicitRules" json:"implicitRules"`
	Language          *string               `bson:"language" json:"language"`
	Text              *Narrative            `bson:"text" json:"text"`
	Contained         []json.RawMessage     `bson:"contained" json:"contained"`
	Extension         []Extension           `bson:"extension" json:"extension"`
	ModifierExtension []Extension           `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier          `bson:"identifier" json:"identifier"`
	Active            *bool                 `bson:"active" json:"active"`
	Type              GroupType             `bson:"type,omitempty" json:"type,omitempty"`
	Actual            bool                  `bson:"actual,omitempty" json:"actual,omitempty"`
	Code              *CodeableConcept      `bson:"code" json:"code"`
	Name              *string               `bson:"name" json:"name"`
	Quantity          *int                  `bson:"quantity" json:"quantity"`
	Characteristic    []GroupCharacteristic `bson:"characteristic" json:"characteristic"`
	Member            []GroupMember         `bson:"member" json:"member"`
}
type GroupCharacteristic struct {
	Id                   *string          `bson:"id" json:"id"`
	Extension            []Extension      `bson:"extension" json:"extension"`
	ModifierExtension    []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Code                 CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	Exclude              bool             `bson:"exclude,omitempty" json:"exclude,omitempty"`
	Period               *Period          `bson:"period" json:"period"`
}
type GroupMember struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Entity            Reference   `bson:"entity,omitempty" json:"entity,omitempty"`
	Period            *Period     `bson:"period" json:"period"`
	Inactive          *bool       `bson:"inactive" json:"inactive"`
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

// UnmarshalGroup unmarshalls a Group.
func UnmarshalGroup(b []byte) (Group, error) {
	var group Group
	if err := json.Unmarshal(b, &group); err != nil {
		return group, err
	}
	return group, nil
}
