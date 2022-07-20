package fhir

import "encoding/json"

// Person is documented here http://hl7.org/fhir/StructureDefinition/Person
type Person struct {
	Id                   *string        `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta          `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string        `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string        `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative     `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name                 []HumanName    `bson:"name,omitempty" json:"name,omitempty"`
	Telecom              []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Gender               *string        `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate            *string        `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Address              []Address      `bson:"address,omitempty" json:"address,omitempty"`
	Photo                *Attachment    `bson:"photo,omitempty" json:"photo,omitempty"`
	ManagingOrganization *Reference     `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Active               *bool          `bson:"active,omitempty" json:"active,omitempty"`
	Link                 []PersonLink   `bson:"link,omitempty" json:"link,omitempty"`
}
type PersonLink struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Assurance         *string     `bson:"assurance,omitempty" json:"assurance,omitempty"`
}
type OtherPerson Person

// MarshalJSON marshals the given Person as JSON into a byte slice
func (r Person) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPerson
		ResourceType string `json:"resourceType"`
	}{
		OtherPerson:  OtherPerson(r),
		ResourceType: "Person",
	})
}

// UnmarshalPerson unmarshals a Person.
func UnmarshalPerson(b []byte) (Person, error) {
	var person Person
	if err := json.Unmarshal(b, &person); err != nil {
		return person, err
	}
	return person, nil
}
