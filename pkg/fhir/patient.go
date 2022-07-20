package fhir

import "encoding/json"

// Patient is documented here http://hl7.org/fhir/StructureDefinition/Patient
type Patient struct {
	Id                   *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active               *bool                  `bson:"active,omitempty" json:"active,omitempty"`
	Name                 []HumanName            `bson:"name,omitempty" json:"name,omitempty"`
	Telecom              []ContactPoint         `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Gender               *string                `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate            *string                `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Address              []Address              `bson:"address,omitempty" json:"address,omitempty"`
	MaritalStatus        *CodeableConcept       `bson:"maritalStatus,omitempty" json:"maritalStatus,omitempty"`
	Photo                []Attachment           `bson:"photo,omitempty" json:"photo,omitempty"`
	Contact              []PatientContact       `bson:"contact,omitempty" json:"contact,omitempty"`
	Animal               *PatientAnimal         `bson:"animal,omitempty" json:"animal,omitempty"`
	Communication        []PatientCommunication `bson:"communication,omitempty" json:"communication,omitempty"`
	ManagingOrganization *Reference             `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Link                 []PatientLink          `bson:"link,omitempty" json:"link,omitempty"`
}
type PatientContact struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Relationship      []CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Name              *HumanName        `bson:"name,omitempty" json:"name,omitempty"`
	Telecom           []ContactPoint    `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address           *Address          `bson:"address,omitempty" json:"address,omitempty"`
	Gender            *string           `bson:"gender,omitempty" json:"gender,omitempty"`
	Organization      *Reference        `bson:"organization,omitempty" json:"organization,omitempty"`
	Period            *Period           `bson:"period,omitempty" json:"period,omitempty"`
}
type PatientAnimal struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Species           CodeableConcept  `bson:"species" json:"species"`
	Breed             *CodeableConcept `bson:"breed,omitempty" json:"breed,omitempty"`
	GenderStatus      *CodeableConcept `bson:"genderStatus,omitempty" json:"genderStatus,omitempty"`
}
type PatientCommunication struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Language          CodeableConcept `bson:"language" json:"language"`
	Preferred         *bool           `bson:"preferred,omitempty" json:"preferred,omitempty"`
}
type PatientLink struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              string      `bson:"type" json:"type"`
}
type OtherPatient Patient

// MarshalJSON marshals the given Patient as JSON into a byte slice
func (r Patient) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPatient
		ResourceType string `json:"resourceType"`
	}{
		OtherPatient: OtherPatient(r),
		ResourceType: "Patient",
	})
}

// UnmarshalPatient unmarshals a Patient.
func UnmarshalPatient(b []byte) (Patient, error) {
	var patient Patient
	if err := json.Unmarshal(b, &patient); err != nil {
		return patient, err
	}
	return patient, nil
}
