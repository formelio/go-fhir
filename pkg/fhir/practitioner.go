package fhir

import "encoding/json"

// Practitioner is documented here http://hl7.org/fhir/StructureDefinition/Practitioner
type Practitioner struct {
	Id                *string                     `bson:"id" json:"id"`
	Meta              *Meta                       `bson:"meta" json:"meta"`
	ImplicitRules     *string                     `bson:"implicitRules" json:"implicitRules"`
	Language          *string                     `bson:"language" json:"language"`
	Text              *Narrative                  `bson:"text" json:"text"`
	Contained         []json.RawMessage           `bson:"contained" json:"contained"`
	Extension         []Extension                 `bson:"extension" json:"extension"`
	ModifierExtension []Extension                 `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier                `bson:"identifier" json:"identifier"`
	Active            *bool                       `bson:"active" json:"active"`
	Name              []HumanName                 `bson:"name" json:"name"`
	Telecom           []ContactPoint              `bson:"telecom" json:"telecom"`
	Address           []Address                   `bson:"address" json:"address"`
	Gender            *AdministrativeGender       `bson:"gender" json:"gender"`
	BirthDate         *string                     `bson:"birthDate" json:"birthDate"`
	Photo             []Attachment                `bson:"photo" json:"photo"`
	Qualification     []PractitionerQualification `bson:"qualification" json:"qualification"`
	Communication     []CodeableConcept           `bson:"communication" json:"communication"`
}
type PractitionerQualification struct {
	Id                *string         `bson:"id" json:"id"`
	Extension         []Extension     `bson:"extension" json:"extension"`
	ModifierExtension []Extension     `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier    `bson:"identifier" json:"identifier"`
	Code              CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Period            *Period         `bson:"period" json:"period"`
	Issuer            *Reference      `bson:"issuer" json:"issuer"`
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

// UnmarshalPractitioner unmarshalls a Practitioner.
func UnmarshalPractitioner(b []byte) (Practitioner, error) {
	var practitioner Practitioner
	if err := json.Unmarshal(b, &practitioner); err != nil {
		return practitioner, err
	}
	return practitioner, nil
}
