package fhir

// Signature is documented here http://hl7.org/fhir/StructureDefinition/Signature
type Signature struct {
	Id          *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension   []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Type        []Coding    `bson:"type" json:"type"`
	When        string      `bson:"when" json:"when"`
	ContentType *string     `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Blob        *string     `bson:"blob,omitempty" json:"blob,omitempty"`
}
