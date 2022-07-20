package fhir

import "encoding/json"

// Person is documented here http://hl7.org/fhir/StructureDefinition/Person
type Person struct {
	Id                   *string               `bson:"id" json:"id"`
	Meta                 *Meta                 `bson:"meta" json:"meta"`
	ImplicitRules        *string               `bson:"implicitRules" json:"implicitRules"`
	Language             *string               `bson:"language" json:"language"`
	Text                 *Narrative            `bson:"text" json:"text"`
	Contained            []json.RawMessage     `bson:"contained" json:"contained"`
	Extension            []Extension           `bson:"extension" json:"extension"`
	ModifierExtension    []Extension           `bson:"modifierExtension" json:"modifierExtension"`
	Identifier           []Identifier          `bson:"identifier" json:"identifier"`
	Name                 []HumanName           `bson:"name" json:"name"`
	Telecom              []ContactPoint        `bson:"telecom" json:"telecom"`
	Gender               *AdministrativeGender `bson:"gender" json:"gender"`
	BirthDate            *string               `bson:"birthDate" json:"birthDate"`
	Address              []Address             `bson:"address" json:"address"`
	Photo                *Attachment           `bson:"photo" json:"photo"`
	ManagingOrganization *Reference            `bson:"managingOrganization" json:"managingOrganization"`
	Active               *bool                 `bson:"active" json:"active"`
	Link                 []PersonLink          `bson:"link" json:"link"`
}
type PersonLink struct {
	Id                *string                 `bson:"id" json:"id"`
	Extension         []Extension             `bson:"extension" json:"extension"`
	ModifierExtension []Extension             `bson:"modifierExtension" json:"modifierExtension"`
	Target            Reference               `bson:"target,omitempty" json:"target,omitempty"`
	Assurance         *IdentityAssuranceLevel `bson:"assurance" json:"assurance"`
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

// UnmarshalPerson unmarshalls a Person.
func UnmarshalPerson(b []byte) (Person, error) {
	var person Person
	if err := json.Unmarshal(b, &person); err != nil {
		return person, err
	}
	return person, nil
}
