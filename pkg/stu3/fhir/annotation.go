package fhir

// Annotation is documented here http://hl7.org/fhir/StructureDefinition/Annotation
type Annotation struct {
	Id              *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension       []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	AuthorReference *Reference  `bson:"authorReference,omitempty" json:"authorReference,omitempty"`
	AuthorString    *string     `bson:"authorString,omitempty" json:"authorString,omitempty"`
	Time            *string     `bson:"time,omitempty" json:"time,omitempty"`
	Text            string      `bson:"text,omitempty" json:"text,omitempty"`
}
