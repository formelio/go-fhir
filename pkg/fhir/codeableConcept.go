package fhir

// CodeableConcept is documented here http://hl7.org/fhir/StructureDefinition/CodeableConcept
type CodeableConcept struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Coding    []Coding    `bson:"coding,omitempty" json:"coding,omitempty"`
	Text      *string     `bson:"text,omitempty" json:"text,omitempty"`
}
