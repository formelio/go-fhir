package fhir

import "encoding/json"

// Flag is documented here http://hl7.org/fhir/StructureDefinition/Flag
type Flag struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string          `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative       `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            string           `bson:"status" json:"status"`
	Category          *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Code              CodeableConcept  `bson:"code" json:"code"`
	Period            *Period          `bson:"period,omitempty" json:"period,omitempty"`
	Encounter         *Reference       `bson:"encounter,omitempty" json:"encounter,omitempty"`
}
type OtherFlag Flag

// MarshalJSON marshals the given Flag as JSON into a byte slice
func (r Flag) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherFlag
		ResourceType string `json:"resourceType"`
	}{
		OtherFlag:    OtherFlag(r),
		ResourceType: "Flag",
	})
}

// UnmarshalFlag unmarshals a Flag.
func UnmarshalFlag(b []byte) (Flag, error) {
	var flag Flag
	if err := json.Unmarshal(b, &flag); err != nil {
		return flag, err
	}
	return flag, nil
}
