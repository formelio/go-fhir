package fhir

// CodeableConcept is documented here http://hl7.org/fhir/StructureDefinition/CodeableConcept
type CodeableConcept struct {
	Id        *string     `bson:"id" json:"id"`
	Extension []Extension `bson:"extension" json:"extension"`
	Coding    []Coding    `bson:"coding" json:"coding"`
	Text      *string     `bson:"text" json:"text"`
}
