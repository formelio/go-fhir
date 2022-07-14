package fhir

// Annotation is documented here http://hl7.org/fhir/StructureDefinition/Annotation
type Annotation struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Time      *string     `bson:"time,omitempty" json:"time,omitempty"`
	Text      string      `bson:"text" json:"text"`
}
