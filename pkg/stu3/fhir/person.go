package fhir

import "encoding/json"

// Person is documented here http://hl7.org/fhir/StructureDefinition/Person
type Person struct {
	Id                   *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	RawContained         []json.RawMessage     `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained            []IResource           `bson:"-,omitempty" json:"-,omitempty"`
	Extension            []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name                 []HumanName           `bson:"name,omitempty" json:"name,omitempty"`
	Telecom              []ContactPoint        `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Gender               *AdministrativeGender `bson:"gender,omitempty" json:"gender,omitempty"`
	BirthDate            *string               `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Address              []Address             `bson:"address,omitempty" json:"address,omitempty"`
	Photo                *Attachment           `bson:"photo,omitempty" json:"photo,omitempty"`
	ManagingOrganization *Reference            `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	Active               *bool                 `bson:"active,omitempty" json:"active,omitempty"`
	Link                 []PersonLink          `bson:"link,omitempty" json:"link,omitempty"`
}
type PersonLink struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Target            Reference               `bson:"target,omitempty" json:"target,omitempty"`
	Assurance         *IdentityAssuranceLevel `bson:"assurance,omitempty" json:"assurance,omitempty"`
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
