package fhir

import (
	"bytes"
	"encoding/json"
)

// Flag is documented here http://hl7.org/fhir/StructureDefinition/Flag
type Flag struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource       `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []*Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            FlagStatus        `bson:"status,omitempty" json:"status,omitempty"`
	Category          *CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	Code              CodeableConcept   `bson:"code,omitempty" json:"code,omitempty"`
	Subject           Reference         `bson:"subject,omitempty" json:"subject,omitempty"`
	Period            *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Encounter         *Reference        `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Author            *Reference        `bson:"author,omitempty" json:"author,omitempty"`
}

// OtherFlag is a helper type to use the default implementations of Marshall and Unmarshal
type OtherFlag Flag

// MarshalJSON marshals the given Flag as JSON into a byte slice
func (r Flag) MarshalJSON() ([]byte, error) {
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
		OtherFlag
	}{
		OtherFlag:    OtherFlag(r),
		ResourceType: "Flag",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into Flag
func (r *Flag) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherFlag)(r)); err != nil {
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
func (r Flag) GetResourceType() ResourceType {
	return ResourceTypeFlag
}
