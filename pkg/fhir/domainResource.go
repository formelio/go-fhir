package fhir

import "encoding/json"

// DomainResource is documented here http://hl7.org/fhir/StructureDefinition/DomainResource
type DomainResource struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource       `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}

// OtherDomainResource is a helper type to use the default implementations of Marshall and Unmarshal
type OtherDomainResource DomainResource

// MarshalJSON marshals the given DomainResource as JSON into a byte slice
func (r DomainResource) MarshalJSON() ([]byte, error) {
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
		OtherDomainResource
		ResourceType string `json:"resourceType"`
	}{
		OtherDomainResource: OtherDomainResource(r),
		ResourceType:        "DomainResource",
	})
}

// UnmarshalJSON unmarshals the given byte slice into DomainResource
func (r *DomainResource) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherDomainResource)(r)); err != nil {
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
func (r DomainResource) GetResourceType() ResourceType {
	return ResourceTypeDomainResource
}
