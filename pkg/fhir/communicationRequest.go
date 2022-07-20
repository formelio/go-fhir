package fhir

import "encoding/json"

// CommunicationRequest is documented here http://hl7.org/fhir/StructureDefinition/CommunicationRequest
type CommunicationRequest struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                          `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                        `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                        `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                     `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn           []Reference                    `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces          []Reference                    `bson:"replaces,omitempty" json:"replaces,omitempty"`
	GroupIdentifier   *Identifier                    `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	Status            string                         `bson:"status" json:"status"`
	Category          []CodeableConcept              `bson:"category,omitempty" json:"category,omitempty"`
	Priority          *string                        `bson:"priority,omitempty" json:"priority,omitempty"`
	Medium            []CodeableConcept              `bson:"medium,omitempty" json:"medium,omitempty"`
	Topic             []Reference                    `bson:"topic,omitempty" json:"topic,omitempty"`
	Payload           []CommunicationRequestPayload  `bson:"payload,omitempty" json:"payload,omitempty"`
	AuthoredOn        *string                        `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester         *CommunicationRequestRequester `bson:"requester,omitempty" json:"requester,omitempty"`
	ReasonCode        []CodeableConcept              `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	Note              []Annotation                   `bson:"note,omitempty" json:"note,omitempty"`
}
type CommunicationRequestPayload struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}
type CommunicationRequestRequester struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}
type OtherCommunicationRequest CommunicationRequest

// MarshalJSON marshals the given CommunicationRequest as JSON into a byte slice
func (r CommunicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCommunicationRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherCommunicationRequest: OtherCommunicationRequest(r),
		ResourceType:              "CommunicationRequest",
	})
}

// UnmarshalCommunicationRequest unmarshals a CommunicationRequest.
func UnmarshalCommunicationRequest(b []byte) (CommunicationRequest, error) {
	var communicationRequest CommunicationRequest
	if err := json.Unmarshal(b, &communicationRequest); err != nil {
		return communicationRequest, err
	}
	return communicationRequest, nil
}
