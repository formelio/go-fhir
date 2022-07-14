package fhir

import "encoding/json"

// OperationDefinition is documented here http://hl7.org/fhir/StructureDefinition/OperationDefinition
type OperationDefinition struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                          `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                        `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                        `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                     `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string                        `bson:"url,omitempty" json:"url,omitempty"`
	Version           *string                        `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                         `bson:"name" json:"name"`
	Status            string                         `bson:"status" json:"status"`
	Kind              string                         `bson:"kind" json:"kind"`
	Experimental      *bool                          `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                        `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                        `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail                `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                        `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext                 `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept              `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                        `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Idempotent        *bool                          `bson:"idempotent,omitempty" json:"idempotent,omitempty"`
	Code              string                         `bson:"code" json:"code"`
	Comment           *string                        `bson:"comment,omitempty" json:"comment,omitempty"`
	Base              *Reference                     `bson:"base,omitempty" json:"base,omitempty"`
	Resource          []string                       `bson:"resource,omitempty" json:"resource,omitempty"`
	System            bool                           `bson:"system" json:"system"`
	Type              bool                           `bson:"type" json:"type"`
	Instance          bool                           `bson:"instance" json:"instance"`
	Parameter         []OperationDefinitionParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	Overload          []OperationDefinitionOverload  `bson:"overload,omitempty" json:"overload,omitempty"`
}
type OperationDefinitionParameter struct {
	Id                *string                              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                               `bson:"name" json:"name"`
	Use               string                               `bson:"use" json:"use"`
	Min               int                                  `bson:"min" json:"min"`
	Max               string                               `bson:"max" json:"max"`
	Documentation     *string                              `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Type              *string                              `bson:"type,omitempty" json:"type,omitempty"`
	SearchType        *string                              `bson:"searchType,omitempty" json:"searchType,omitempty"`
	Profile           *Reference                           `bson:"profile,omitempty" json:"profile,omitempty"`
	Binding           *OperationDefinitionParameterBinding `bson:"binding,omitempty" json:"binding,omitempty"`
	Part              []OperationDefinitionParameter       `bson:"part,omitempty" json:"part,omitempty"`
}
type OperationDefinitionParameterBinding struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Strength          string      `bson:"strength" json:"strength"`
}
type OperationDefinitionOverload struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ParameterName     []string    `bson:"parameterName,omitempty" json:"parameterName,omitempty"`
	Comment           *string     `bson:"comment,omitempty" json:"comment,omitempty"`
}
type OtherOperationDefinition OperationDefinition

// MarshalJSON marshals the given OperationDefinition as JSON into a byte slice
func (r OperationDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherOperationDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherOperationDefinition: OtherOperationDefinition(r),
		ResourceType:             "OperationDefinition",
	})
}

// UnmarshalOperationDefinition unmarshals a OperationDefinition.
func UnmarshalOperationDefinition(b []byte) (OperationDefinition, error) {
	var operationDefinition OperationDefinition
	if err := json.Unmarshal(b, &operationDefinition); err != nil {
		return operationDefinition, err
	}
	return operationDefinition, nil
}
