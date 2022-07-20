package fhir

// Reference is documented here http://hl7.org/fhir/StructureDefinition/Reference
type Reference struct {
	Id         *string     `bson:"id" json:"id"`
	Extension  []Extension `bson:"extension" json:"extension"`
	Reference  *string     `bson:"reference" json:"reference"`
	Identifier *Identifier `bson:"identifier" json:"identifier"`
	Display    *string     `bson:"display" json:"display"`
}
