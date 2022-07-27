package fhir

import "encoding/json"

// Basic is documented here http://hl7.org/fhir/StructureDefinition/Basic
type Basic struct {
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
	Code              CodeableConcept   `bson:"code,omitempty" json:"code,omitempty"`
	Subject           *Reference        `bson:"subject" json:"subject"`
	Created           *string           `bson:"created" json:"created"`
	Author            *Reference        `bson:"author" json:"author"`
}

// OtherBasic is a helper type to use the default implementations of Marshall and Unmarshal
type OtherBasic Basic

// MarshalJSON marshals the given Basic as JSON into a byte slice
func (r Basic) MarshalJSON() ([]byte, error) {
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
		OtherBasic
		ResourceType string `json:"resourceType"`
	}{
		OtherBasic:   OtherBasic(r),
		ResourceType: "Basic",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Basic
func (r *Basic) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherBasic)(r)); err != nil {
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
func (r Basic) GetResourceType() ResourceType {
	return ResourceTypeBasic
}
