package fhir

import (
	"bytes"
	"encoding/json"
)

// OperationDefinition is documented here http://hl7.org/fhir/StructureDefinition/OperationDefinition
type OperationDefinition struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                         `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                      `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage               `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource                     `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string                         `bson:"url,omitempty" json:"url,omitempty"`
	Version           *string                         `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                          `bson:"name,omitempty" json:"name,omitempty"`
	Status            PublicationStatus               `bson:"status,omitempty" json:"status,omitempty"`
	Kind              OperationKind                   `bson:"kind,omitempty" json:"kind,omitempty"`
	Experimental      *bool                           `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                         `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                         `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []*ContactDetail                `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                         `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []*UsageContext                 `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []*CodeableConcept              `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                         `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Idempotent        *bool                           `bson:"idempotent,omitempty" json:"idempotent,omitempty"`
	Code              string                          `bson:"code,omitempty" json:"code,omitempty"`
	Comment           *string                         `bson:"comment,omitempty" json:"comment,omitempty"`
	Base              *Reference                      `bson:"base,omitempty" json:"base,omitempty"`
	Resource          []*ResourceType                 `bson:"resource,omitempty" json:"resource,omitempty"`
	System            bool                            `bson:"system,omitempty" json:"system,omitempty"`
	Type              bool                            `bson:"type,omitempty" json:"type,omitempty"`
	Instance          bool                            `bson:"instance,omitempty" json:"instance,omitempty"`
	Parameter         []*OperationDefinitionParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	Overload          []*OperationDefinitionOverload  `bson:"overload,omitempty" json:"overload,omitempty"`
}
type OperationDefinitionParameter struct {
	Id                *string                              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                               `bson:"name,omitempty" json:"name,omitempty"`
	Use               OperationParameterUse                `bson:"use,omitempty" json:"use,omitempty"`
	Min               int                                  `bson:"min,omitempty" json:"min,omitempty"`
	Max               string                               `bson:"max,omitempty" json:"max,omitempty"`
	Documentation     *string                              `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Type              *string                              `bson:"type,omitempty" json:"type,omitempty"`
	SearchType        *SearchParamType                     `bson:"searchType,omitempty" json:"searchType,omitempty"`
	Profile           *Reference                           `bson:"profile,omitempty" json:"profile,omitempty"`
	Binding           *OperationDefinitionParameterBinding `bson:"binding,omitempty" json:"binding,omitempty"`
	Part              []*OperationDefinitionParameter      `bson:"part,omitempty" json:"part,omitempty"`
}
type OperationDefinitionParameterBinding struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Strength          BindingStrength `bson:"strength,omitempty" json:"strength,omitempty"`
	ValueSetUri       *string         `bson:"valueSetUri,omitempty" json:"valueSetUri,omitempty"`
	ValueSetReference *Reference      `bson:"valueSetReference,omitempty" json:"valueSetReference,omitempty"`
}
type OperationDefinitionOverload struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ParameterName     []*string    `bson:"parameterName,omitempty" json:"parameterName,omitempty"`
	Comment           *string      `bson:"comment,omitempty" json:"comment,omitempty"`
}

// OtherOperationDefinition is a helper type to use the default implementations of Marshall and Unmarshal
type OtherOperationDefinition OperationDefinition

// MarshalJSON marshals the given OperationDefinition as JSON into a byte slice
func (r OperationDefinition) MarshalJSON() ([]byte, error) {
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
		OtherOperationDefinition
	}{
		OtherOperationDefinition: OtherOperationDefinition(r),
		ResourceType:             "OperationDefinition",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into OperationDefinition
func (r *OperationDefinition) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherOperationDefinition)(r)); err != nil {
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
func (r OperationDefinition) GetResourceType() ResourceType {
	return ResourceTypeOperationDefinition
}
