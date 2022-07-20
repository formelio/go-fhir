package fhir

import "encoding/json"

// MessageDefinition is documented here http://hl7.org/fhir/StructureDefinition/MessageDefinition
type MessageDefinition struct {
	Id                *string                            `bson:"id" json:"id"`
	Meta              *Meta                              `bson:"meta" json:"meta"`
	ImplicitRules     *string                            `bson:"implicitRules" json:"implicitRules"`
	Language          *string                            `bson:"language" json:"language"`
	Text              *Narrative                         `bson:"text" json:"text"`
	Contained         []json.RawMessage                  `bson:"contained" json:"contained"`
	Extension         []Extension                        `bson:"extension" json:"extension"`
	ModifierExtension []Extension                        `bson:"modifierExtension" json:"modifierExtension"`
	Url               *string                            `bson:"url" json:"url"`
	Identifier        *Identifier                        `bson:"identifier" json:"identifier"`
	Version           *string                            `bson:"version" json:"version"`
	Name              *string                            `bson:"name" json:"name"`
	Title             *string                            `bson:"title" json:"title"`
	Status            PublicationStatus                  `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                              `bson:"experimental" json:"experimental"`
	Date              string                             `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                            `bson:"publisher" json:"publisher"`
	Contact           []ContactDetail                    `bson:"contact" json:"contact"`
	Description       *string                            `bson:"description" json:"description"`
	UseContext        []UsageContext                     `bson:"useContext" json:"useContext"`
	Jurisdiction      []CodeableConcept                  `bson:"jurisdiction" json:"jurisdiction"`
	Purpose           *string                            `bson:"purpose" json:"purpose"`
	Copyright         *string                            `bson:"copyright" json:"copyright"`
	Base              *Reference                         `bson:"base" json:"base"`
	Parent            []Reference                        `bson:"parent" json:"parent"`
	Replaces          []Reference                        `bson:"replaces" json:"replaces"`
	Event             Coding                             `bson:"event,omitempty" json:"event,omitempty"`
	Category          *MessageSignificanceCategory       `bson:"category" json:"category"`
	Focus             []MessageDefinitionFocus           `bson:"focus" json:"focus"`
	ResponseRequired  *bool                              `bson:"responseRequired" json:"responseRequired"`
	AllowedResponse   []MessageDefinitionAllowedResponse `bson:"allowedResponse" json:"allowedResponse"`
}
type MessageDefinitionFocus struct {
	Id                *string      `bson:"id" json:"id"`
	Extension         []Extension  `bson:"extension" json:"extension"`
	ModifierExtension []Extension  `bson:"modifierExtension" json:"modifierExtension"`
	Code              ResourceType `bson:"code,omitempty" json:"code,omitempty"`
	Profile           *Reference   `bson:"profile" json:"profile"`
	Min               *int         `bson:"min" json:"min"`
	Max               *string      `bson:"max" json:"max"`
}
type MessageDefinitionAllowedResponse struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Message           Reference   `bson:"message,omitempty" json:"message,omitempty"`
	Situation         *string     `bson:"situation" json:"situation"`
}
type OtherMessageDefinition MessageDefinition

// MarshalJSON marshals the given MessageDefinition as JSON into a byte slice
func (r MessageDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMessageDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherMessageDefinition: OtherMessageDefinition(r),
		ResourceType:           "MessageDefinition",
	})
}

// UnmarshalMessageDefinition unmarshalls a MessageDefinition.
func UnmarshalMessageDefinition(b []byte) (MessageDefinition, error) {
	var messageDefinition MessageDefinition
	if err := json.Unmarshal(b, &messageDefinition); err != nil {
		return messageDefinition, err
	}
	return messageDefinition, nil
}
