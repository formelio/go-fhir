package fhir

import "encoding/json"

// MessageHeader is documented here http://hl7.org/fhir/StructureDefinition/MessageHeader
type MessageHeader struct {
	Id                *string                    `bson:"id" json:"id"`
	Meta              *Meta                      `bson:"meta" json:"meta"`
	ImplicitRules     *string                    `bson:"implicitRules" json:"implicitRules"`
	Language          *string                    `bson:"language" json:"language"`
	Text              *Narrative                 `bson:"text" json:"text"`
	Contained         []json.RawMessage          `bson:"contained" json:"contained"`
	Extension         []Extension                `bson:"extension" json:"extension"`
	ModifierExtension []Extension                `bson:"modifierExtension" json:"modifierExtension"`
	Event             Coding                     `bson:"event,omitempty" json:"event,omitempty"`
	Destination       []MessageHeaderDestination `bson:"destination" json:"destination"`
	Receiver          *Reference                 `bson:"receiver" json:"receiver"`
	Sender            *Reference                 `bson:"sender" json:"sender"`
	Timestamp         string                     `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Enterer           *Reference                 `bson:"enterer" json:"enterer"`
	Author            *Reference                 `bson:"author" json:"author"`
	Source            MessageHeaderSource        `bson:"source,omitempty" json:"source,omitempty"`
	Responsible       *Reference                 `bson:"responsible" json:"responsible"`
	Reason            *CodeableConcept           `bson:"reason" json:"reason"`
	Response          *MessageHeaderResponse     `bson:"response" json:"response"`
	Focus             []Reference                `bson:"focus" json:"focus"`
}
type MessageHeaderDestination struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Name              *string     `bson:"name" json:"name"`
	Target            *Reference  `bson:"target" json:"target"`
	Endpoint          string      `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}
type MessageHeaderSource struct {
	Id                *string       `bson:"id" json:"id"`
	Extension         []Extension   `bson:"extension" json:"extension"`
	ModifierExtension []Extension   `bson:"modifierExtension" json:"modifierExtension"`
	Name              *string       `bson:"name" json:"name"`
	Software          *string       `bson:"software" json:"software"`
	Version           *string       `bson:"version" json:"version"`
	Contact           *ContactPoint `bson:"contact" json:"contact"`
	Endpoint          string        `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}
type MessageHeaderResponse struct {
	Id                *string      `bson:"id" json:"id"`
	Extension         []Extension  `bson:"extension" json:"extension"`
	ModifierExtension []Extension  `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        string       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code              ResponseType `bson:"code,omitempty" json:"code,omitempty"`
	Details           *Reference   `bson:"details" json:"details"`
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

// UnmarshalMessageHeader unmarshalls a MessageHeader.
func UnmarshalMessageHeader(b []byte) (MessageHeader, error) {
	var messageHeader MessageHeader
	if err := json.Unmarshal(b, &messageHeader); err != nil {
		return messageHeader, err
	}
	return messageHeader, nil
}
