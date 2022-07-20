package fhir

// RelatedArtifact is documented here http://hl7.org/fhir/StructureDefinition/RelatedArtifact
type RelatedArtifact struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Type      string      `bson:"type" json:"type"`
	Display   *string     `bson:"display,omitempty" json:"display,omitempty"`
	Citation  *string     `bson:"citation,omitempty" json:"citation,omitempty"`
	Url       *string     `bson:"url,omitempty" json:"url,omitempty"`
	Document  *Attachment `bson:"document,omitempty" json:"document,omitempty"`
	Resource  *Reference  `bson:"resource,omitempty" json:"resource,omitempty"`
}
