package fhir

import "encoding/json"

// Organization is documented here http://hl7.org/fhir/StructureDefinition/Organization
type Organization struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage     `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource           `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active            *bool                 `bson:"active,omitempty" json:"active,omitempty"`
	Type              []CodeableConcept     `bson:"type,omitempty" json:"type,omitempty"`
	Name              *string               `bson:"name,omitempty" json:"name,omitempty"`
	Alias             []string              `bson:"alias,omitempty" json:"alias,omitempty"`
	Telecom           []ContactPoint        `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address           []Address             `bson:"address,omitempty" json:"address,omitempty"`
	PartOf            *Reference            `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Contact           []OrganizationContact `bson:"contact,omitempty" json:"contact,omitempty"`
	Endpoint          []Reference           `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}
type OrganizationContact struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Purpose           *CodeableConcept `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Name              *HumanName       `bson:"name,omitempty" json:"name,omitempty"`
	Telecom           []ContactPoint   `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address           *Address         `bson:"address,omitempty" json:"address,omitempty"`
}

// OtherOrganization is a helper type to use the default implementations of Marshall and Unmarshal
type OtherOrganization Organization

// MarshalJSON marshals the given Organization as JSON into a byte slice
func (r Organization) MarshalJSON() ([]byte, error) {
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
		OtherOrganization
		ResourceType string `json:"resourceType"`
	}{
		OtherOrganization: OtherOrganization(r),
		ResourceType:      "Organization",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Organization
func (r *Organization) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherOrganization)(r)); err != nil {
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
func (r Organization) GetResourceType() ResourceType {
	return ResourceTypeOrganization
}
