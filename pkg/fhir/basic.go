package fhir

import "encoding/json"

// Basic is documented here http://hl7.org/fhir/StructureDefinition/Basic
type Basic struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string         `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative      `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code              CodeableConcept `bson:"code" json:"code"`
	Subject           *Reference      `bson:"subject,omitempty" json:"subject,omitempty"`
	Created           *string         `bson:"created,omitempty" json:"created,omitempty"`
}
type OtherBasic Basic

// MarshalJSON marshals the given Basic as JSON into a byte slice
func (r Basic) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBasic
		ResourceType string `json:"resourceType"`
	}{
		OtherBasic:   OtherBasic(r),
		ResourceType: "Basic",
	})
}

// UnmarshalBasic unmarshals a Basic.
func UnmarshalBasic(b []byte) (Basic, error) {
	var basic Basic
	if err := json.Unmarshal(b, &basic); err != nil {
		return basic, err
	}
	return basic, nil
}
