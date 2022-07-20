package fhir

import "encoding/json"

// OperationDefinition is documented here http://hl7.org/fhir/StructureDefinition/OperationDefinition
type OperationDefinition struct {
	Id                *string                        `bson:"id" json:"id"`
	Meta              *Meta                          `bson:"meta" json:"meta"`
	ImplicitRules     *string                        `bson:"implicitRules" json:"implicitRules"`
	Language          *string                        `bson:"language" json:"language"`
	Text              *Narrative                     `bson:"text" json:"text"`
	Contained         []json.RawMessage              `bson:"contained" json:"contained"`
	Extension         []Extension                    `bson:"extension" json:"extension"`
	ModifierExtension []Extension                    `bson:"modifierExtension" json:"modifierExtension"`
	Url               *string                        `bson:"url" json:"url"`
	Version           *string                        `bson:"version" json:"version"`
	Name              string                         `bson:"name,omitempty" json:"name,omitempty"`
	Status            PublicationStatus              `bson:"status,omitempty" json:"status,omitempty"`
	Kind              OperationKind                  `bson:"kind,omitempty" json:"kind,omitempty"`
	Experimental      *bool                          `bson:"experimental" json:"experimental"`
	Date              *string                        `bson:"date" json:"date"`
	Publisher         *string                        `bson:"publisher" json:"publisher"`
	Contact           []ContactDetail                `bson:"contact" json:"contact"`
	Description       *string                        `bson:"description" json:"description"`
	UseContext        []UsageContext                 `bson:"useContext" json:"useContext"`
	Jurisdiction      []CodeableConcept              `bson:"jurisdiction" json:"jurisdiction"`
	Purpose           *string                        `bson:"purpose" json:"purpose"`
	Idempotent        *bool                          `bson:"idempotent" json:"idempotent"`
	Code              string                         `bson:"code,omitempty" json:"code,omitempty"`
	Comment           *string                        `bson:"comment" json:"comment"`
	Base              *Reference                     `bson:"base" json:"base"`
	Resource          []ResourceType                 `bson:"resource" json:"resource"`
	System            bool                           `bson:"system,omitempty" json:"system,omitempty"`
	Type              bool                           `bson:"type,omitempty" json:"type,omitempty"`
	Instance          bool                           `bson:"instance,omitempty" json:"instance,omitempty"`
	Parameter         []OperationDefinitionParameter `bson:"parameter" json:"parameter"`
	Overload          []OperationDefinitionOverload  `bson:"overload" json:"overload"`
}
type OperationDefinitionParameter struct {
	Id                *string                              `bson:"id" json:"id"`
	Extension         []Extension                          `bson:"extension" json:"extension"`
	ModifierExtension []Extension                          `bson:"modifierExtension" json:"modifierExtension"`
	Name              string                               `bson:"name,omitempty" json:"name,omitempty"`
	Use               OperationParameterUse                `bson:"use,omitempty" json:"use,omitempty"`
	Min               int                                  `bson:"min,omitempty" json:"min,omitempty"`
	Max               string                               `bson:"max,omitempty" json:"max,omitempty"`
	Documentation     *string                              `bson:"documentation" json:"documentation"`
	Type              *string                              `bson:"type" json:"type"`
	SearchType        *SearchParamType                     `bson:"searchType" json:"searchType"`
	Profile           *Reference                           `bson:"profile" json:"profile"`
	Binding           *OperationDefinitionParameterBinding `bson:"binding" json:"binding"`
	Part              []OperationDefinitionParameter       `bson:"part" json:"part"`
}
type OperationDefinitionParameterBinding struct {
	Id                *string         `bson:"id" json:"id"`
	Extension         []Extension     `bson:"extension" json:"extension"`
	ModifierExtension []Extension     `bson:"modifierExtension" json:"modifierExtension"`
	Strength          BindingStrength `bson:"strength,omitempty" json:"strength,omitempty"`
	ValueSetUri       *string         `bson:"valueSetUri,omitempty" json:"valueSetUri,omitempty"`
	ValueSetReference *Reference      `bson:"valueSetReference,omitempty" json:"valueSetReference,omitempty"`
}
type OperationDefinitionOverload struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	ParameterName     []string    `bson:"parameterName" json:"parameterName"`
	Comment           *string     `bson:"comment" json:"comment"`
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

// UnmarshalOperationDefinition unmarshalls a OperationDefinition.
func UnmarshalOperationDefinition(b []byte) (OperationDefinition, error) {
	var operationDefinition OperationDefinition
	if err := json.Unmarshal(b, &operationDefinition); err != nil {
		return operationDefinition, err
	}
	return operationDefinition, nil
}
