package fhir

import "encoding/json"

// Communication is documented here http://hl7.org/fhir/StructureDefinition/Communication
type Communication struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn           []Reference            `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf            []Reference            `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status            string                 `bson:"status" json:"status"`
	NotDone           *bool                  `bson:"notDone,omitempty" json:"notDone,omitempty"`
	NotDoneReason     *CodeableConcept       `bson:"notDoneReason,omitempty" json:"notDoneReason,omitempty"`
	Category          []CodeableConcept      `bson:"category,omitempty" json:"category,omitempty"`
	Medium            []CodeableConcept      `bson:"medium,omitempty" json:"medium,omitempty"`
	Topic             []Reference            `bson:"topic,omitempty" json:"topic,omitempty"`
	Sent              *string                `bson:"sent,omitempty" json:"sent,omitempty"`
	Received          *string                `bson:"received,omitempty" json:"received,omitempty"`
	ReasonCode        []CodeableConcept      `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	Payload           []CommunicationPayload `bson:"payload,omitempty" json:"payload,omitempty"`
	Note              []Annotation           `bson:"note,omitempty" json:"note,omitempty"`
}
type CommunicationPayload struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
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

// UnmarshalCommunication unmarshals a Communication.
func UnmarshalCommunication(b []byte) (Communication, error) {
	var communication Communication
	if err := json.Unmarshal(b, &communication); err != nil {
		return communication, err
	}
	return communication, nil
}
