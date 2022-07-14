package fhir

import "encoding/json"

// MessageHeader is documented here http://hl7.org/fhir/StructureDefinition/MessageHeader
type MessageHeader struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                    `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                 `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Event             Coding                     `bson:"event" json:"event"`
	Destination       []MessageHeaderDestination `bson:"destination,omitempty" json:"destination,omitempty"`
	Timestamp         string                     `bson:"timestamp" json:"timestamp"`
	Enterer           *Reference                 `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Author            *Reference                 `bson:"author,omitempty" json:"author,omitempty"`
	Source            MessageHeaderSource        `bson:"source" json:"source"`
	Reason            *CodeableConcept           `bson:"reason,omitempty" json:"reason,omitempty"`
	Response          *MessageHeaderResponse     `bson:"response,omitempty" json:"response,omitempty"`
	Focus             []Reference                `bson:"focus,omitempty" json:"focus,omitempty"`
}
type MessageHeaderDestination struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              *string     `bson:"name,omitempty" json:"name,omitempty"`
	Target            *Reference  `bson:"target,omitempty" json:"target,omitempty"`
	Endpoint          string      `bson:"endpoint" json:"endpoint"`
}
type MessageHeaderSource struct {
	Id                *string       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              *string       `bson:"name,omitempty" json:"name,omitempty"`
	Software          *string       `bson:"software,omitempty" json:"software,omitempty"`
	Version           *string       `bson:"version,omitempty" json:"version,omitempty"`
	Contact           *ContactPoint `bson:"contact,omitempty" json:"contact,omitempty"`
	Endpoint          string        `bson:"endpoint" json:"endpoint"`
}
type MessageHeaderResponse struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        string      `bson:"identifier" json:"identifier"`
	Code              string      `bson:"code" json:"code"`
	Details           *Reference  `bson:"details,omitempty" json:"details,omitempty"`
}
type OtherMessageHeader MessageHeader

// MarshalJSON marshals the given MessageHeader as JSON into a byte slice
func (r MessageHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMessageHeader
		ResourceType string `json:"resourceType"`
	}{
		OtherMessageHeader: OtherMessageHeader(r),
		ResourceType:       "MessageHeader",
	})
}

// UnmarshalMessageHeader unmarshals a MessageHeader.
func UnmarshalMessageHeader(b []byte) (MessageHeader, error) {
	var messageHeader MessageHeader
	if err := json.Unmarshal(b, &messageHeader); err != nil {
		return messageHeader, err
	}
	return messageHeader, nil
}
