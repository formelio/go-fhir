package fhir

import "encoding/json"

// Communication is documented here http://hl7.org/fhir/StructureDefinition/Communication
type Communication struct {
	Id                *string                `bson:"id" json:"id"`
	Meta              *Meta                  `bson:"meta" json:"meta"`
	ImplicitRules     *string                `bson:"implicitRules" json:"implicitRules"`
	Language          *string                `bson:"language" json:"language"`
	Text              *Narrative             `bson:"text" json:"text"`
	Contained         []json.RawMessage      `bson:"contained" json:"contained"`
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
type OtherCommunication Communication

// MarshalJSON marshals the given Communication as JSON into a byte slice
func (r Communication) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCommunication
		ResourceType string `json:"resourceType"`
	}{
		OtherCommunication: OtherCommunication(r),
		ResourceType:       "Communication",
	})
}

// UnmarshalCommunication unmarshalls a Communication.
func UnmarshalCommunication(b []byte) (Communication, error) {
	var communication Communication
	if err := json.Unmarshal(b, &communication); err != nil {
		return communication, err
	}
	return communication, nil
}
