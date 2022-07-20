package fhir

import "encoding/json"

// Patient is documented here http://hl7.org/fhir/StructureDefinition/Patient
type Patient struct {
	Id                   *string                `bson:"id" json:"id"`
	Meta                 *Meta                  `bson:"meta" json:"meta"`
	ImplicitRules        *string                `bson:"implicitRules" json:"implicitRules"`
	Language             *string                `bson:"language" json:"language"`
	Text                 *Narrative             `bson:"text" json:"text"`
	Contained            []json.RawMessage      `bson:"contained" json:"contained"`
	Extension            []Extension            `bson:"extension" json:"extension"`
	ModifierExtension    []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	Identifier           []Identifier           `bson:"identifier" json:"identifier"`
	Active               *bool                  `bson:"active" json:"active"`
	Name                 []HumanName            `bson:"name" json:"name"`
	Telecom              []ContactPoint         `bson:"telecom" json:"telecom"`
	Gender               *AdministrativeGender  `bson:"gender" json:"gender"`
	BirthDate            *string                `bson:"birthDate" json:"birthDate"`
	DeceasedBoolean      *bool                  `bson:"deceasedBoolean,omitempty" json:"deceasedBoolean,omitempty"`
	DeceasedDateTime     *string                `bson:"deceasedDateTime,omitempty" json:"deceasedDateTime,omitempty"`
	Address              []Address              `bson:"address" json:"address"`
	MaritalStatus        *CodeableConcept       `bson:"maritalStatus" json:"maritalStatus"`
	MultipleBirthBoolean *bool                  `bson:"multipleBirthBoolean,omitempty" json:"multipleBirthBoolean,omitempty"`
	MultipleBirthInteger *int                   `bson:"multipleBirthInteger,omitempty" json:"multipleBirthInteger,omitempty"`
	Photo                []Attachment           `bson:"photo" json:"photo"`
	Contact              []PatientContact       `bson:"contact" json:"contact"`
	Animal               *PatientAnimal         `bson:"animal" json:"animal"`
	Communication        []PatientCommunication `bson:"communication" json:"communication"`
	GeneralPractitioner  []Reference            `bson:"generalPractitioner" json:"generalPractitioner"`
	ManagingOrganization *Reference             `bson:"managingOrganization" json:"managingOrganization"`
	Link                 []PatientLink          `bson:"link" json:"link"`
}
type PatientContact struct {
	Id                *string               `bson:"id" json:"id"`
	Extension         []Extension           `bson:"extension" json:"extension"`
	ModifierExtension []Extension           `bson:"modifierExtension" json:"modifierExtension"`
	Relationship      []CodeableConcept     `bson:"relationship" json:"relationship"`
	Name              *HumanName            `bson:"name" json:"name"`
	Telecom           []ContactPoint        `bson:"telecom" json:"telecom"`
	Address           *Address              `bson:"address" json:"address"`
	Gender            *AdministrativeGender `bson:"gender" json:"gender"`
	Organization      *Reference            `bson:"organization" json:"organization"`
	Period            *Period               `bson:"period" json:"period"`
}
type PatientAnimal struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Species           CodeableConcept  `bson:"species,omitempty" json:"species,omitempty"`
	Breed             *CodeableConcept `bson:"breed" json:"breed"`
	GenderStatus      *CodeableConcept `bson:"genderStatus" json:"genderStatus"`
}
type PatientCommunication struct {
	Id                *string         `bson:"id" json:"id"`
	Extension         []Extension     `bson:"extension" json:"extension"`
	ModifierExtension []Extension     `bson:"modifierExtension" json:"modifierExtension"`
	Language          CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
	Preferred         *bool           `bson:"preferred" json:"preferred"`
}
type PatientLink struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Other             Reference   `bson:"other,omitempty" json:"other,omitempty"`
	Type              LinkType    `bson:"type,omitempty" json:"type,omitempty"`
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

// UnmarshalPatient unmarshalls a Patient.
func UnmarshalPatient(b []byte) (Patient, error) {
	var patient Patient
	if err := json.Unmarshal(b, &patient); err != nil {
		return patient, err
	}
	return patient, nil
}
