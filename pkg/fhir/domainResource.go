package fhir

import "encoding/json"

// DomainResource is documented here http://hl7.org/fhir/StructureDefinition/DomainResource
type DomainResource struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	RawContained      []json.RawMessage `bson:"contained" json:"contained"`
	Contained         []IResource       `bson:"-" json:"-"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
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
