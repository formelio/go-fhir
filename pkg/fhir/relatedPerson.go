package fhir

import "encoding/json"

// RelatedPerson is documented here http://hl7.org/fhir/StructureDefinition/RelatedPerson
type RelatedPerson struct {
	Id                *string               `bson:"id" json:"id"`
	Meta              *Meta                 `bson:"meta" json:"meta"`
	ImplicitRules     *string               `bson:"implicitRules" json:"implicitRules"`
	Language          *string               `bson:"language" json:"language"`
	Text              *Narrative            `bson:"text" json:"text"`
	Contained         []json.RawMessage     `bson:"contained" json:"contained"`
	Extension         []Extension           `bson:"extension" json:"extension"`
	ModifierExtension []Extension           `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier          `bson:"identifier" json:"identifier"`
	Active            *bool                 `bson:"active" json:"active"`
	Patient           Reference             `bson:"patient,omitempty" json:"patient,omitempty"`
	Relationship      *CodeableConcept      `bson:"relationship" json:"relationship"`
	Name              []HumanName           `bson:"name" json:"name"`
	Telecom           []ContactPoint        `bson:"telecom" json:"telecom"`
	Gender            *AdministrativeGender `bson:"gender" json:"gender"`
	BirthDate         *string               `bson:"birthDate" json:"birthDate"`
	Address           []Address             `bson:"address" json:"address"`
	Photo             []Attachment          `bson:"photo" json:"photo"`
	Period            *Period               `bson:"period" json:"period"`
}
type OtherRelatedPerson RelatedPerson

// MarshalJSON marshals the given RelatedPerson as JSON into a byte slice
func (r RelatedPerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRelatedPerson
		ResourceType string `json:"resourceType"`
	}{
		OtherRelatedPerson: OtherRelatedPerson(r),
		ResourceType:       "RelatedPerson",
	})
}

// UnmarshalRelatedPerson unmarshalls a RelatedPerson.
func UnmarshalRelatedPerson(b []byte) (RelatedPerson, error) {
	var relatedPerson RelatedPerson
	if err := json.Unmarshal(b, &relatedPerson); err != nil {
		return relatedPerson, err
	}
	return relatedPerson, nil
}
