package fhir

import (
	"bytes"
	"encoding/json"
)

// MessageDefinition is documented here http://hl7.org/fhir/StructureDefinition/MessageDefinition
type MessageDefinition struct {
	Id                *string                             `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                             `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                          `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage                   `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource                         `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension                        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string                             `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        *Identifier                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                             `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string                             `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                             `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus                   `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                               `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              string                              `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                             `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []*ContactDetail                    `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                             `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []*UsageContext                     `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []*CodeableConcept                  `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                             `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string                             `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Base              *Reference                          `bson:"base,omitempty" json:"base,omitempty"`
	Parent            []*Reference                        `bson:"parent,omitempty" json:"parent,omitempty"`
	Replaces          []*Reference                        `bson:"replaces,omitempty" json:"replaces,omitempty"`
	Event             Coding                              `bson:"event,omitempty" json:"event,omitempty"`
	Category          *MessageSignificanceCategory        `bson:"category,omitempty" json:"category,omitempty"`
	Focus             []*MessageDefinitionFocus           `bson:"focus,omitempty" json:"focus,omitempty"`
	ResponseRequired  *bool                               `bson:"responseRequired,omitempty" json:"responseRequired,omitempty"`
	AllowedResponse   []*MessageDefinitionAllowedResponse `bson:"allowedResponse,omitempty" json:"allowedResponse,omitempty"`
}
type MessageDefinitionFocus struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              ResourceType `bson:"code,omitempty" json:"code,omitempty"`
	Profile           *Reference   `bson:"profile,omitempty" json:"profile,omitempty"`
	Min               *int         `bson:"min,omitempty" json:"min,omitempty"`
	Max               *string      `bson:"max,omitempty" json:"max,omitempty"`
}
type MessageDefinitionAllowedResponse struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Message           Reference    `bson:"message,omitempty" json:"message,omitempty"`
	Situation         *string      `bson:"situation,omitempty" json:"situation,omitempty"`
}

// OtherMessageDefinition is a helper type to use the default implementations of Marshall and Unmarshal
type OtherMessageDefinition MessageDefinition

// MarshalJSON marshals the given MessageDefinition as JSON into a byte slice
func (r MessageDefinition) MarshalJSON() ([]byte, error) {
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
		OtherMessageDefinition
	}{
		OtherMessageDefinition: OtherMessageDefinition(r),
		ResourceType:           "MessageDefinition",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into MessageDefinition
func (r *MessageDefinition) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherMessageDefinition)(r)); err != nil {
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
func (r MessageDefinition) GetResourceType() ResourceType {
	return ResourceTypeMessageDefinition
}
