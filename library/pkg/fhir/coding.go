package fhir

// Coding is documented here http://hl7.org/fhir/StructureDefinition/Coding
type Coding struct {
	Id           *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension    []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	System       *string     `bson:"system,omitempty" json:"system,omitempty"`
	Version      *string     `bson:"version,omitempty" json:"version,omitempty"`
	Code         *string     `bson:"code,omitempty" json:"code,omitempty"`
	Display      *string     `bson:"display,omitempty" json:"display,omitempty"`
	UserSelected *bool       `bson:"userSelected,omitempty" json:"userSelected,omitempty"`
}
