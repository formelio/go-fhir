package fhir

import "encoding/json"

// Organization is documented here http://hl7.org/fhir/StructureDefinition/Organization
type Organization struct {
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
	Type              []CodeableConcept     `bson:"type" json:"type"`
	Name              *string               `bson:"name" json:"name"`
	Alias             []string              `bson:"alias" json:"alias"`
	Telecom           []ContactPoint        `bson:"telecom" json:"telecom"`
	Address           []Address             `bson:"address" json:"address"`
	PartOf            *Reference            `bson:"partOf" json:"partOf"`
	Contact           []OrganizationContact `bson:"contact" json:"contact"`
	Endpoint          []Reference           `bson:"endpoint" json:"endpoint"`
}
type OrganizationContact struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Purpose           *CodeableConcept `bson:"purpose" json:"purpose"`
	Name              *HumanName       `bson:"name" json:"name"`
	Telecom           []ContactPoint   `bson:"telecom" json:"telecom"`
	Address           *Address         `bson:"address" json:"address"`
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
