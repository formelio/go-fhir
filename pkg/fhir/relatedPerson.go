package fhir

import "encoding/json"

// RelatedPerson is documented here http://hl7.org/fhir/StructureDefinition/RelatedPerson
type RelatedPerson struct {
	Id                *string               `bson:"id" json:"id"`
	Meta              *Meta                 `bson:"meta" json:"meta"`
	ImplicitRules     *string               `bson:"implicitRules" json:"implicitRules"`
	Language          *string               `bson:"language" json:"language"`
	Text              *Narrative            `bson:"text" json:"text"`
	RawContained      []json.RawMessage     `bson:"contained" json:"contained"`
	Contained         []IResource           `bson:"-" json:"-"`
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

// OtherRelatedPerson is a helper type to use the default implementations of Marshall and Unmarshal
type OtherRelatedPerson RelatedPerson

// MarshalJSON marshals the given RelatedPerson as JSON into a byte slice
func (r RelatedPerson) MarshalJSON() ([]byte, error) {
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
		OtherRelatedPerson
		ResourceType string `json:"resourceType"`
	}{
		OtherRelatedPerson: OtherRelatedPerson(r),
		ResourceType:       "RelatedPerson",
	})
}

// UnmarshalJSON unmarshals the given byte slice into RelatedPerson
func (r *RelatedPerson) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherRelatedPerson)(r)); err != nil {
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
func (r RelatedPerson) GetResourceType() ResourceType {
	return ResourceTypeRelatedPerson
}
