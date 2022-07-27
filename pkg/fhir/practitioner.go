package fhir

import "encoding/json"

// Practitioner is documented here http://hl7.org/fhir/StructureDefinition/Practitioner
type Practitioner struct {
	Id                *string                     `bson:"id" json:"id"`
	Meta              *Meta                       `bson:"meta" json:"meta"`
	ImplicitRules     *string                     `bson:"implicitRules" json:"implicitRules"`
	Language          *string                     `bson:"language" json:"language"`
	Text              *Narrative                  `bson:"text" json:"text"`
	RawContained      []json.RawMessage           `bson:"contained" json:"contained"`
	Contained         []IResource                 `bson:"-" json:"-"`
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

// OtherPractitioner is a helper type to use the default implementations of Marshall and Unmarshal
type OtherPractitioner Practitioner

// MarshalJSON marshals the given Practitioner as JSON into a byte slice
func (r Practitioner) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(struct {
		OtherPractitioner
		ResourceType string `json:"resourceType"`
	}{
		OtherPractitioner: OtherPractitioner(r),
		ResourceType:      "Practitioner",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Practitioner
func (r *Practitioner) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherPractitioner)(r)); err != nil {
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
func (r Practitioner) GetResourceType() ResourceType {
	return ResourceTypePractitioner
}
