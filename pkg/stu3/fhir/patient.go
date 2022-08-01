package fhir

import (
	"bytes"
	"encoding/json"
)

// Patient is documented here http://hl7.org/fhir/StructureDefinition/Patient
type Patient struct {
	Id                   *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	RawContained         []json.RawMessage       `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained            []IResource             `bson:"-,omitempty" json:"-,omitempty"`
	Extension            []*Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []*Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []*Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active               *bool                   `bson:"active,omitempty" json:"active,omitempty"`
	Name                 []*HumanName            `bson:"name,omitempty" json:"name,omitempty"`
	Telecom              []*ContactPoint         `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Gender               *AdministrativeGender   `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate            *string                 `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	DeceasedBoolean      *bool                   `bson:"deceasedBoolean,omitempty" json:"deceasedBoolean,omitempty"`
	DeceasedDateTime     *string                 `bson:"deceasedDateTime,omitempty" json:"deceasedDateTime,omitempty"`
	Address              []*Address              `bson:"address,omitempty" json:"address,omitempty"`
	MaritalStatus        *CodeableConcept        `bson:"maritalStatus,omitempty" json:"maritalStatus,omitempty"`
	MultipleBirthBoolean *bool                   `bson:"multipleBirthBoolean,omitempty" json:"multipleBirthBoolean,omitempty"`
	MultipleBirthInteger *int                    `bson:"multipleBirthInteger,omitempty" json:"multipleBirthInteger,omitempty"`
	Photo                []*Attachment           `bson:"photo,omitempty" json:"photo,omitempty"`
	Contact              []*PatientContact       `bson:"contact,omitempty" json:"contact,omitempty"`
	Animal               *PatientAnimal          `bson:"animal,omitempty" json:"animal,omitempty"`
	Communication        []*PatientCommunication `bson:"communication,omitempty" json:"communication,omitempty"`
	GeneralPractitioner  []*Reference            `bson:"generalPractitioner,omitempty" json:"generalPractitioner,omitempty"`
	ManagingOrganization *Reference              `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Link                 []*PatientLink          `bson:"link,omitempty" json:"link,omitempty"`
}
type PatientContact struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Relationship      []*CodeableConcept    `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Name              *HumanName            `bson:"name,omitempty" json:"name,omitempty"`
	Telecom           []*ContactPoint       `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address           *Address              `bson:"address,omitempty" json:"address,omitempty"`
	Gender            *AdministrativeGender `bson:"gender,omitempty" json:"gender,omitempty"`
	Organization      *Reference            `bson:"organization,omitempty" json:"organization,omitempty"`
	Period            *Period               `bson:"period,omitempty" json:"period,omitempty"`
}
type PatientAnimal struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Species           CodeableConcept  `bson:"species,omitempty" json:"species,omitempty"`
	Breed             *CodeableConcept `bson:"breed,omitempty" json:"breed,omitempty"`
	GenderStatus      *CodeableConcept `bson:"genderStatus,omitempty" json:"genderStatus,omitempty"`
}
type PatientCommunication struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Language          CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
	Preferred         *bool           `bson:"preferred,omitempty" json:"preferred,omitempty"`
}
type PatientLink struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Other             Reference    `bson:"other,omitempty" json:"other,omitempty"`
	Type              LinkType     `bson:"type,omitempty" json:"type,omitempty"`
}

// OtherPatient is a helper type to use the default implementations of Marshall and Unmarshal
type OtherPatient Patient

// MarshalJSON marshals the given Patient as JSON into a byte slice
func (r Patient) MarshalJSON() ([]byte, error) {
	// If the field has contained resources, we need to marshal them individually and store them in .RawContained
	if len(r.Contained) > 0 {
		var err error
		r.RawContained = make([]json.RawMessage, len(r.Contained))
		for i, contained := range r.Contained {
			r.RawContained[i], err = json.Marshal(contained)
			if err != nil {
				return nil, err
			}
		}
	}
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherPatient
	}{
		OtherPatient: OtherPatient(r),
		ResourceType: "Patient",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into Patient
func (r *Patient) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherPatient)(r)); err != nil {
		return err
	}
	// If the field has contained resources, we need to unmarshal them individually and store them in .Contained
	if len(r.RawContained) > 0 {
		var err error
		r.Contained = make([]IResource, len(r.RawContained))
		for i, rawContained := range r.RawContained {
			r.Contained[i], err = UnmarshalResource(rawContained)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Returns the resourceType of the resource, makes this resource an instance of IResource
func (r Patient) GetResourceType() ResourceType {
	return ResourceTypePatient
}
