package fhir

import "encoding/json"

// CommunicationRequest is documented here http://hl7.org/fhir/StructureDefinition/CommunicationRequest
type CommunicationRequest struct {
	Id                 *string                        `bson:"id" json:"id"`
	Meta               *Meta                          `bson:"meta" json:"meta"`
	ImplicitRules      *string                        `bson:"implicitRules" json:"implicitRules"`
	Language           *string                        `bson:"language" json:"language"`
	Text               *Narrative                     `bson:"text" json:"text"`
	Contained          []json.RawMessage              `bson:"contained" json:"contained"`
	Extension          []Extension                    `bson:"extension" json:"extension"`
	ModifierExtension  []Extension                    `bson:"modifierExtension" json:"modifierExtension"`
	Identifier         []Identifier                   `bson:"identifier" json:"identifier"`
	BasedOn            []Reference                    `bson:"basedOn" json:"basedOn"`
	Replaces           []Reference                    `bson:"replaces" json:"replaces"`
	GroupIdentifier    *Identifier                    `bson:"groupIdentifier" json:"groupIdentifier"`
	Status             RequestStatus                  `bson:"status,omitempty" json:"status,omitempty"`
	Category           []CodeableConcept              `bson:"category" json:"category"`
	Priority           *RequestPriority               `bson:"priority" json:"priority"`
	Medium             []CodeableConcept              `bson:"medium" json:"medium"`
	Subject            *Reference                     `bson:"subject" json:"subject"`
	Recipient          []Reference                    `bson:"recipient" json:"recipient"`
	Topic              []Reference                    `bson:"topic" json:"topic"`
	Context            *Reference                     `bson:"context" json:"context"`
	Payload            []CommunicationRequestPayload  `bson:"payload" json:"payload"`
	OccurrenceDateTime *string                        `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period                        `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	AuthoredOn         *string                        `bson:"authoredOn" json:"authoredOn"`
	Sender             *Reference                     `bson:"sender" json:"sender"`
	Requester          *CommunicationRequestRequester `bson:"requester" json:"requester"`
	ReasonCode         []CodeableConcept              `bson:"reasonCode" json:"reasonCode"`
	ReasonReference    []Reference                    `bson:"reasonReference" json:"reasonReference"`
	Note               []Annotation                   `bson:"note" json:"note"`
}
type CommunicationRequestPayload struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	ContentString     *string     `bson:"contentString,omitempty" json:"contentString,omitempty"`
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}
type CommunicationRequestRequester struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Agent             Reference   `bson:"agent,omitempty" json:"agent,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf" json:"onBehalfOf"`
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

// UnmarshalCommunicationRequest unmarshalls a CommunicationRequest.
func UnmarshalCommunicationRequest(b []byte) (CommunicationRequest, error) {
	var communicationRequest CommunicationRequest
	if err := json.Unmarshal(b, &communicationRequest); err != nil {
		return communicationRequest, err
	}
	return communicationRequest, nil
}
