package fhir

import "encoding/json"

// CommunicationRequest is documented here http://hl7.org/fhir/StructureDefinition/CommunicationRequest
type CommunicationRequest struct {
	Id                 *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                          `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string                        `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string                        `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative                     `bson:"text,omitempty" json:"text,omitempty"`
	RawContained       []json.RawMessage              `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained          []IResource                    `bson:"-,omitempty" json:"-,omitempty"`
	Extension          []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn            []Reference                    `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces           []Reference                    `bson:"replaces,omitempty" json:"replaces,omitempty"`
	GroupIdentifier    *Identifier                    `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	Status             RequestStatus                  `bson:"status,omitempty" json:"status,omitempty"`
	Category           []CodeableConcept              `bson:"category,omitempty" json:"category,omitempty"`
	Priority           *RequestPriority               `bson:"priority,omitempty" json:"priority,omitempty"`
	Medium             []CodeableConcept              `bson:"medium,omitempty" json:"medium,omitempty"`
	Subject            *Reference                     `bson:"subject,omitempty" json:"subject,omitempty"`
	Recipient          []Reference                    `bson:"recipient,omitempty" json:"recipient,omitempty"`
	Topic              []Reference                    `bson:"topic,omitempty" json:"topic,omitempty"`
	Context            *Reference                     `bson:"context,omitempty" json:"context,omitempty"`
	Payload            []CommunicationRequestPayload  `bson:"payload,omitempty" json:"payload,omitempty"`
	OccurrenceDateTime *string                        `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period                        `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	AuthoredOn         *string                        `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Sender             *Reference                     `bson:"sender,omitempty" json:"sender,omitempty"`
	Requester          *CommunicationRequestRequester `bson:"requester,omitempty" json:"requester,omitempty"`
	ReasonCode         []CodeableConcept              `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference    []Reference                    `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note               []Annotation                   `bson:"note,omitempty" json:"note,omitempty"`
}
type CommunicationRequestPayload struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ContentString     *string     `bson:"contentString,omitempty" json:"contentString,omitempty"`
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}
type CommunicationRequestRequester struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Agent             Reference   `bson:"agent,omitempty" json:"agent,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
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
