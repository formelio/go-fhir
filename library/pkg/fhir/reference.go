package fhir

// Reference is documented here http://hl7.org/fhir/StructureDefinition/Reference
type Reference struct {
	Id         *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension  []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Reference  *string     `bson:"reference,omitempty" json:"reference,omitempty"`
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Display    *string     `bson:"display,omitempty" json:"display,omitempty"`
}
