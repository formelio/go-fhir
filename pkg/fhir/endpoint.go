package fhir

import "encoding/json"

// Endpoint is documented here http://hl7.org/fhir/StructureDefinition/Endpoint
type Endpoint struct {
	Id                   *string           `bson:"id" json:"id"`
	Meta                 *Meta             `bson:"meta" json:"meta"`
	ImplicitRules        *string           `bson:"implicitRules" json:"implicitRules"`
	Language             *string           `bson:"language" json:"language"`
	Text                 *Narrative        `bson:"text" json:"text"`
	Contained            []json.RawMessage `bson:"contained" json:"contained"`
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
type OtherEndpoint Endpoint

// MarshalJSON marshals the given Endpoint as JSON into a byte slice
func (r Endpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEndpoint
		ResourceType string `json:"resourceType"`
	}{
		OtherEndpoint: OtherEndpoint(r),
		ResourceType:  "Endpoint",
	})
}

// UnmarshalEndpoint unmarshalls a Endpoint.
func UnmarshalEndpoint(b []byte) (Endpoint, error) {
	var endpoint Endpoint
	if err := json.Unmarshal(b, &endpoint); err != nil {
		return endpoint, err
	}
	return endpoint, nil
}
