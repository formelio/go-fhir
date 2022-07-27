package fhir

import "encoding/json"

// BodySite is documented here http://hl7.org/fhir/StructureDefinition/BodySite
type BodySite struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	RawContained      []json.RawMessage `bson:"contained" json:"contained"`
	Contained         []IResource       `bson:"-" json:"-"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier      `bson:"identifier" json:"identifier"`
	Active            *bool             `bson:"active" json:"active"`
	Code              *CodeableConcept  `bson:"code" json:"code"`
	Qualifier         []CodeableConcept `bson:"qualifier" json:"qualifier"`
	Description       *string           `bson:"description" json:"description"`
	Image             []Attachment      `bson:"image" json:"image"`
	Patient           Reference         `bson:"patient,omitempty" json:"patient,omitempty"`
}

// OtherBodySite is a helper type to use the default implementations of Marshall and Unmarshal
type OtherBodySite BodySite

// MarshalJSON marshals the given BodySite as JSON into a byte slice
func (r BodySite) MarshalJSON() ([]byte, error) {
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
		OtherBodySite
		ResourceType string `json:"resourceType"`
	}{
		OtherBodySite: OtherBodySite(r),
		ResourceType:  "BodySite",
	})
}

// UnmarshalJSON unmarshals the given byte slice into BodySite
func (r *BodySite) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherBodySite)(r)); err != nil {
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
func (r BodySite) GetResourceType() ResourceType {
	return ResourceTypeBodySite
}
