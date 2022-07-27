package fhir

import "encoding/json"

// CommunicationRequest is documented here http://hl7.org/fhir/StructureDefinition/CommunicationRequest
type CommunicationRequest struct {
	Id                 *string                        `bson:"id" json:"id"`
	Meta               *Meta                          `bson:"meta" json:"meta"`
	ImplicitRules      *string                        `bson:"implicitRules" json:"implicitRules"`
	Language           *string                        `bson:"language" json:"language"`
	Text               *Narrative                     `bson:"text" json:"text"`
	RawContained       []json.RawMessage              `bson:"contained" json:"contained"`
	Contained          []IResource                    `bson:"-" json:"-"`
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

// OtherCommunicationRequest is a helper type to use the default implementations of Marshall and Unmarshal
type OtherCommunicationRequest CommunicationRequest

// MarshalJSON marshals the given CommunicationRequest as JSON into a byte slice
func (r CommunicationRequest) MarshalJSON() ([]byte, error) {
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
		OtherCommunicationRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherCommunicationRequest: OtherCommunicationRequest(r),
		ResourceType:              "CommunicationRequest",
	})
}

// UnmarshalJSON unmarshals the given byte slice into CommunicationRequest
func (r *CommunicationRequest) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherCommunicationRequest)(r)); err != nil {
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
func (r CommunicationRequest) GetResourceType() ResourceType {
	return ResourceTypeCommunicationRequest
}
