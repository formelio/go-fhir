package fhir

// Annotation is documented here http://hl7.org/fhir/StructureDefinition/Annotation
type Annotation struct {
	Id              *string     `bson:"id" json:"id"`
	Extension       []Extension `bson:"extension" json:"extension"`
	AuthorReference *Reference  `bson:"authorReference,omitempty" json:"authorReference,omitempty"`
	AuthorString    *string     `bson:"authorString,omitempty" json:"authorString,omitempty"`
	Time            *string     `bson:"time" json:"time"`
	Text            string      `bson:"text,omitempty" json:"text,omitempty"`
}
