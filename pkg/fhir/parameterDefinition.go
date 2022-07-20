package fhir

// ParameterDefinition is documented here http://hl7.org/fhir/StructureDefinition/ParameterDefinition
type ParameterDefinition struct {
	Id            *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension     []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Name          *string     `bson:"name,omitempty" json:"name,omitempty"`
	Use           string      `bson:"use" json:"use"`
	Min           *int        `bson:"min,omitempty" json:"min,omitempty"`
	Max           *string     `bson:"max,omitempty" json:"max,omitempty"`
	Documentation *string     `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Type          string      `bson:"type" json:"type"`
	Profile       *Reference  `bson:"profile,omitempty" json:"profile,omitempty"`
}
