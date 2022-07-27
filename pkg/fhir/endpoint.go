package fhir

import "encoding/json"

// Endpoint is documented here http://hl7.org/fhir/StructureDefinition/Endpoint
type Endpoint struct {
	Id                   *string           `bson:"id" json:"id"`
	Meta                 *Meta             `bson:"meta" json:"meta"`
	ImplicitRules        *string           `bson:"implicitRules" json:"implicitRules"`
	Language             *string           `bson:"language" json:"language"`
	Text                 *Narrative        `bson:"text" json:"text"`
	RawContained         []json.RawMessage `bson:"contained" json:"contained"`
	Contained            []IResource       `bson:"-" json:"-"`
	Extension            []Extension       `bson:"extension" json:"extension"`
	ModifierExtension    []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier           []Identifier      `bson:"identifier" json:"identifier"`
	Status               EndpointStatus    `bson:"status,omitempty" json:"status,omitempty"`
	ConnectionType       Coding            `bson:"connectionType,omitempty" json:"connectionType,omitempty"`
	Name                 *string           `bson:"name" json:"name"`
	ManagingOrganization *Reference        `bson:"managingOrganization" json:"managingOrganization"`
	Contact              []ContactPoint    `bson:"contact" json:"contact"`
	Period               *Period           `bson:"period" json:"period"`
	PayloadType          []CodeableConcept `bson:"payloadType,omitempty" json:"payloadType,omitempty"`
	PayloadMimeType      []string          `bson:"payloadMimeType" json:"payloadMimeType"`
	Address              string            `bson:"address,omitempty" json:"address,omitempty"`
	Header               []string          `bson:"header" json:"header"`
}

// OtherEndpoint is a helper type to use the default implementations of Marshall and Unmarshal
type OtherEndpoint Endpoint

// MarshalJSON marshals the given Endpoint as JSON into a byte slice
func (r Endpoint) MarshalJSON() ([]byte, error) {
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
		OtherEndpoint
		ResourceType string `json:"resourceType"`
	}{
		OtherEndpoint: OtherEndpoint(r),
		ResourceType:  "Endpoint",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Endpoint
func (r *Endpoint) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherEndpoint)(r)); err != nil {
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
func (r Endpoint) GetResourceType() ResourceType {
	return ResourceTypeEndpoint
}
