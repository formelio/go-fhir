package fhir

import "encoding/json"

// RelatedPerson is documented here http://hl7.org/fhir/StructureDefinition/RelatedPerson
type RelatedPerson struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string          `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative       `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active            *bool            `bson:"active,omitempty" json:"active,omitempty"`
	Patient           Reference        `bson:"patient" json:"patient"`
	Relationship      *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Name              []HumanName      `bson:"name,omitempty" json:"name,omitempty"`
	Telecom           []ContactPoint   `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Gender            *string          `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate         *string          `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Address           []Address        `bson:"address,omitempty" json:"address,omitempty"`
	Photo             []Attachment     `bson:"photo,omitempty" json:"photo,omitempty"`
	Period            *Period          `bson:"period,omitempty" json:"period,omitempty"`
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

// UnmarshalRelatedPerson unmarshals a RelatedPerson.
func UnmarshalRelatedPerson(b []byte) (RelatedPerson, error) {
	var relatedPerson RelatedPerson
	if err := json.Unmarshal(b, &relatedPerson); err != nil {
		return relatedPerson, err
	}
	return relatedPerson, nil
}
