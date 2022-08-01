package fhir

import (
	"bytes"
	"encoding/json"
)

// MessageHeader is documented here http://hl7.org/fhir/StructureDefinition/MessageHeader
type MessageHeader struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage           `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource                 `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Event             Coding                      `bson:"event,omitempty" json:"event,omitempty"`
	Destination       []*MessageHeaderDestination `bson:"destination,omitempty" json:"destination,omitempty"`
	Receiver          *Reference                  `bson:"receiver,omitempty" json:"receiver,omitempty"`
	Sender            *Reference                  `bson:"sender,omitempty" json:"sender,omitempty"`
	Timestamp         string                      `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Enterer           *Reference                  `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Author            *Reference                  `bson:"author,omitempty" json:"author,omitempty"`
	Source            MessageHeaderSource         `bson:"source,omitempty" json:"source,omitempty"`
	Responsible       *Reference                  `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Reason            *CodeableConcept            `bson:"reason,omitempty" json:"reason,omitempty"`
	Response          *MessageHeaderResponse      `bson:"response,omitempty" json:"response,omitempty"`
	Focus             []*Reference                `bson:"focus,omitempty" json:"focus,omitempty"`
}
type MessageHeaderDestination struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              *string      `bson:"name,omitempty" json:"name,omitempty"`
	Target            *Reference   `bson:"target,omitempty" json:"target,omitempty"`
	Endpoint          string       `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}
type MessageHeaderSource struct {
	Id                *string       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              *string       `bson:"name,omitempty" json:"name,omitempty"`
	Software          *string       `bson:"software,omitempty" json:"software,omitempty"`
	Version           *string       `bson:"version,omitempty" json:"version,omitempty"`
	Contact           *ContactPoint `bson:"contact,omitempty" json:"contact,omitempty"`
	Endpoint          string        `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}
type MessageHeaderResponse struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        string       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code              ResponseType `bson:"code,omitempty" json:"code,omitempty"`
	Details           *Reference   `bson:"details,omitempty" json:"details,omitempty"`
}

// OtherMessageHeader is a helper type to use the default implementations of Marshall and Unmarshal
type OtherMessageHeader MessageHeader

// MarshalJSON marshals the given MessageHeader as JSON into a byte slice
func (r MessageHeader) MarshalJSON() ([]byte, error) {
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
		OtherMessageHeader
	}{
		OtherMessageHeader: OtherMessageHeader(r),
		ResourceType:       "MessageHeader",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into MessageHeader
func (r *MessageHeader) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherMessageHeader)(r)); err != nil {
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
func (r MessageHeader) GetResourceType() ResourceType {
	return ResourceTypeMessageHeader
}
