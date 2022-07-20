package fhir

import "encoding/json"

// Flag is documented here http://hl7.org/fhir/StructureDefinition/Flag
type Flag struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	Contained         []json.RawMessage `bson:"contained" json:"contained"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier      `bson:"identifier" json:"identifier"`
	Status            FlagStatus        `bson:"status,omitempty" json:"status,omitempty"`
	Category          *CodeableConcept  `bson:"category" json:"category"`
	Code              CodeableConcept   `bson:"code,omitempty" json:"code,omitempty"`
	Subject           Reference         `bson:"subject,omitempty" json:"subject,omitempty"`
	Period            *Period           `bson:"period" json:"period"`
	Encounter         *Reference        `bson:"encounter" json:"encounter"`
	Author            *Reference        `bson:"author" json:"author"`
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

// UnmarshalFlag unmarshalls a Flag.
func UnmarshalFlag(b []byte) (Flag, error) {
	var flag Flag
	if err := json.Unmarshal(b, &flag); err != nil {
		return flag, err
	}
	return flag, nil
}
