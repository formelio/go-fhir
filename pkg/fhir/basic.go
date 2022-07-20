package fhir

import "encoding/json"

// Basic is documented here http://hl7.org/fhir/StructureDefinition/Basic
type Basic struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	Contained         []json.RawMessage `bson:"contained" json:"contained"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier      `bson:"identifier" json:"identifier"`
	Code              CodeableConcept   `bson:"code,omitempty" json:"code,omitempty"`
	Subject           *Reference        `bson:"subject" json:"subject"`
	Created           *string           `bson:"created" json:"created"`
	Author            *Reference        `bson:"author" json:"author"`
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

// UnmarshalBasic unmarshalls a Basic.
func UnmarshalBasic(b []byte) (Basic, error) {
	var basic Basic
	if err := json.Unmarshal(b, &basic); err != nil {
		return basic, err
	}
	return basic, nil
}
