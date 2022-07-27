package fhir

import "encoding/json"

// Communication is documented here http://hl7.org/fhir/StructureDefinition/Communication
type Communication struct {
	Id                *string                `bson:"id" json:"id"`
	Meta              *Meta                  `bson:"meta" json:"meta"`
	ImplicitRules     *string                `bson:"implicitRules" json:"implicitRules"`
	Language          *string                `bson:"language" json:"language"`
	Text              *Narrative             `bson:"text" json:"text"`
	RawContained      []json.RawMessage      `bson:"contained" json:"contained"`
	Contained         []IResource            `bson:"-" json:"-"`
	Extension         []Extension            `bson:"extension" json:"extension"`
	ModifierExtension []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier           `bson:"identifier" json:"identifier"`
	Definition        []Reference            `bson:"definition" json:"definition"`
	BasedOn           []Reference            `bson:"basedOn" json:"basedOn"`
	PartOf            []Reference            `bson:"partOf" json:"partOf"`
	Status            EventStatus            `bson:"status,omitempty" json:"status,omitempty"`
	NotDone           *bool                  `bson:"notDone" json:"notDone"`
	NotDoneReason     *CodeableConcept       `bson:"notDoneReason" json:"notDoneReason"`
	Category          []CodeableConcept      `bson:"category" json:"category"`
	Medium            []CodeableConcept      `bson:"medium" json:"medium"`
	Subject           *Reference             `bson:"subject" json:"subject"`
	Recipient         []Reference            `bson:"recipient" json:"recipient"`
	Topic             []Reference            `bson:"topic" json:"topic"`
	Context           *Reference             `bson:"context" json:"context"`
	Sent              *string                `bson:"sent" json:"sent"`
	Received          *string                `bson:"received" json:"received"`
	Sender            *Reference             `bson:"sender" json:"sender"`
	ReasonCode        []CodeableConcept      `bson:"reasonCode" json:"reasonCode"`
	ReasonReference   []Reference            `bson:"reasonReference" json:"reasonReference"`
	Payload           []CommunicationPayload `bson:"payload" json:"payload"`
	Note              []Annotation           `bson:"note" json:"note"`
}
type CommunicationPayload struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	ContentString     *string     `bson:"contentString,omitempty" json:"contentString,omitempty"`
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
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
	return json.Marshal(struct {
		OtherCommunication
		ResourceType string `json:"resourceType"`
	}{
		OtherCommunication: OtherCommunication(r),
		ResourceType:       "Communication",
	})
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
