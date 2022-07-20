package fhir

import "encoding/json"

// MessageDefinition is documented here http://hl7.org/fhir/StructureDefinition/MessageDefinition
type MessageDefinition struct {
	Id                *string                            `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                            `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                         `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string                            `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        *Identifier                        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                            `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string                            `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                            `bson:"title,omitempty" json:"title,omitempty"`
	Status            string                             `bson:"status" json:"status"`
	Experimental      *bool                              `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              string                             `bson:"date" json:"date"`
	Publisher         *string                            `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail                    `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                            `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext                     `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                  `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                            `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string                            `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Base              *Reference                         `bson:"base,omitempty" json:"base,omitempty"`
	Replaces          []Reference                        `bson:"replaces,omitempty" json:"replaces,omitempty"`
	Event             Coding                             `bson:"event" json:"event"`
	Category          *string                            `bson:"category,omitempty" json:"category,omitempty"`
	Focus             []MessageDefinitionFocus           `bson:"focus,omitempty" json:"focus,omitempty"`
	ResponseRequired  *bool                              `bson:"responseRequired,omitempty" json:"responseRequired,omitempty"`
	AllowedResponse   []MessageDefinitionAllowedResponse `bson:"allowedResponse,omitempty" json:"allowedResponse,omitempty"`
}
type MessageDefinitionFocus struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string      `bson:"code" json:"code"`
	Profile           *Reference  `bson:"profile,omitempty" json:"profile,omitempty"`
	Min               *int        `bson:"min,omitempty" json:"min,omitempty"`
	Max               *string     `bson:"max,omitempty" json:"max,omitempty"`
}
type MessageDefinitionAllowedResponse struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Message           Reference   `bson:"message" json:"message"`
	Situation         *string     `bson:"situation,omitempty" json:"situation,omitempty"`
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

// UnmarshalMessageDefinition unmarshals a MessageDefinition.
func UnmarshalMessageDefinition(b []byte) (MessageDefinition, error) {
	var messageDefinition MessageDefinition
	if err := json.Unmarshal(b, &messageDefinition); err != nil {
		return messageDefinition, err
	}
	return messageDefinition, nil
}
