package fhir

// ParameterDefinition is documented here http://hl7.org/fhir/StructureDefinition/ParameterDefinition
type ParameterDefinition struct {
	Id            *string               `bson:"id" json:"id"`
	Extension     []Extension           `bson:"extension" json:"extension"`
	Name          *string               `bson:"name" json:"name"`
	Use           OperationParameterUse `bson:"use,omitempty" json:"use,omitempty"`
	Min           *int                  `bson:"min" json:"min"`
	Max           *string               `bson:"max" json:"max"`
	Documentation *string               `bson:"documentation" json:"documentation"`
	Type          string                `bson:"type,omitempty" json:"type,omitempty"`
	Profile       *Reference            `bson:"profile" json:"profile"`
}
