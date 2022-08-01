package fhir

import (
	"bytes"
	"encoding/json"
)

// Communication is documented here http://hl7.org/fhir/StructureDefinition/Communication
type Communication struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage       `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource             `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []*Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Definition        []*Reference            `bson:"definition,omitempty" json:"definition,omitempty"`
	BasedOn           []*Reference            `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf            []*Reference            `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status            EventStatus             `bson:"status,omitempty" json:"status,omitempty"`
	NotDone           *bool                   `bson:"notDone,omitempty" json:"notDone,omitempty"`
	NotDoneReason     *CodeableConcept        `bson:"notDoneReason,omitempty" json:"notDoneReason,omitempty"`
	Category          []*CodeableConcept      `bson:"category,omitempty" json:"category,omitempty"`
	Medium            []*CodeableConcept      `bson:"medium,omitempty" json:"medium,omitempty"`
	Subject           *Reference              `bson:"subject,omitempty" json:"subject,omitempty"`
	Recipient         []*Reference            `bson:"recipient,omitempty" json:"recipient,omitempty"`
	Topic             []*Reference            `bson:"topic,omitempty" json:"topic,omitempty"`
	Context           *Reference              `bson:"context,omitempty" json:"context,omitempty"`
	Sent              *string                 `bson:"sent,omitempty" json:"sent,omitempty"`
	Received          *string                 `bson:"received,omitempty" json:"received,omitempty"`
	Sender            *Reference              `bson:"sender,omitempty" json:"sender,omitempty"`
	ReasonCode        []*CodeableConcept      `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference   []*Reference            `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Payload           []*CommunicationPayload `bson:"payload,omitempty" json:"payload,omitempty"`
	Note              []*Annotation           `bson:"note,omitempty" json:"note,omitempty"`
}
type CommunicationPayload struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ContentString     *string      `bson:"contentString,omitempty" json:"contentString,omitempty"`
	ContentAttachment *Attachment  `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference   `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}

// OtherCommunication is a helper type to use the default implementations of Marshall and Unmarshal
type OtherCommunication Communication

// MarshalJSON marshals the given Communication as JSON into a byte slice
func (r Communication) MarshalJSON() ([]byte, error) {
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherCommunication
	}{
		OtherCommunication: OtherCommunication(r),
		ResourceType:       "Communication",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into Communication
func (r *Communication) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherCommunication)(r)); err != nil {
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
func (r Communication) GetResourceType() ResourceType {
	return ResourceTypeCommunication
}
