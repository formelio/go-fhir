package fhir

import "encoding/json"

// Practitioner is documented here http://hl7.org/fhir/StructureDefinition/Practitioner
type Practitioner struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active            *bool                       `bson:"active,omitempty" json:"active,omitempty"`
	Name              []HumanName                 `bson:"name,omitempty" json:"name,omitempty"`
	Telecom           []ContactPoint              `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address           []Address                   `bson:"address,omitempty" json:"address,omitempty"`
	Gender            *string                     `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate         *string                     `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Photo             []Attachment                `bson:"photo,omitempty" json:"photo,omitempty"`
	Qualification     []PractitionerQualification `bson:"qualification,omitempty" json:"qualification,omitempty"`
	Communication     []CodeableConcept           `bson:"communication,omitempty" json:"communication,omitempty"`
}
type PractitionerQualification struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code              CodeableConcept `bson:"code" json:"code"`
	Period            *Period         `bson:"period,omitempty" json:"period,omitempty"`
	Issuer            *Reference      `bson:"issuer,omitempty" json:"issuer,omitempty"`
}
type OtherPractitioner Practitioner

// MarshalJSON marshals the given Practitioner as JSON into a byte slice
func (r Practitioner) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPractitioner
		ResourceType string `json:"resourceType"`
	}{
		OtherPractitioner: OtherPractitioner(r),
		ResourceType:      "Practitioner",
	})
}

// UnmarshalPractitioner unmarshals a Practitioner.
func UnmarshalPractitioner(b []byte) (Practitioner, error) {
	var practitioner Practitioner
	if err := json.Unmarshal(b, &practitioner); err != nil {
		return practitioner, err
	}
	return practitioner, nil
}
