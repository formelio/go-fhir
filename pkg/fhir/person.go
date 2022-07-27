package fhir

import "encoding/json"

// Person is documented here http://hl7.org/fhir/StructureDefinition/Person
type Person struct {
	Id                   *string               `bson:"id" json:"id"`
	Meta                 *Meta                 `bson:"meta" json:"meta"`
	ImplicitRules        *string               `bson:"implicitRules" json:"implicitRules"`
	Language             *string               `bson:"language" json:"language"`
	Text                 *Narrative            `bson:"text" json:"text"`
	RawContained         []json.RawMessage     `bson:"contained" json:"contained"`
	Contained            []IResource           `bson:"-" json:"-"`
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

// OtherPerson is a helper type to use the default implementations of Marshall and Unmarshal
type OtherPerson Person

// MarshalJSON marshals the given Person as JSON into a byte slice
func (r Person) MarshalJSON() ([]byte, error) {
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
		OtherPerson
		ResourceType string `json:"resourceType"`
	}{
		OtherPerson:  OtherPerson(r),
		ResourceType: "Person",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Person
func (r *Person) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherPerson)(r)); err != nil {
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
func (r Person) GetResourceType() ResourceType {
	return ResourceTypePerson
}
