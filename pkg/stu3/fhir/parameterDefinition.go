package fhir

// ParameterDefinition is documented here http://hl7.org/fhir/StructureDefinition/ParameterDefinition
type ParameterDefinition struct {
	Id            *string               `bson:"id,omitempty" json:"id,omitempty"`
	Extension     []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	Name          *string               `bson:"name,omitempty" json:"name,omitempty"`
	Use           OperationParameterUse `bson:"use,omitempty" json:"use,omitempty"`
	Min           *int                  `bson:"min,omitempty" json:"min,omitempty"`
	Max           *string               `bson:"max,omitempty" json:"max,omitempty"`
	Documentation *string               `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Type          string                `bson:"type,omitempty" json:"type,omitempty"`
	Profile       *Reference            `bson:"profile,omitempty" json:"profile,omitempty"`
}
