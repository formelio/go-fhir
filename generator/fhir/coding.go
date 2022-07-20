package fhir

// Coding is documented here http://hl7.org/fhir/StructureDefinition/Coding
type Coding struct {
	Id           *string     `bson:"id" json:"id"`
	Extension    []Extension `bson:"extension" json:"extension"`
	System       *string     `bson:"system" json:"system"`
	Version      *string     `bson:"version" json:"version"`
	Code         *string     `bson:"code" json:"code"`
	Display      *string     `bson:"display" json:"display"`
	UserSelected *bool       `bson:"userSelected" json:"userSelected"`
}
