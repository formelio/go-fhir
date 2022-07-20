package fhir

// RelatedArtifact is documented here http://hl7.org/fhir/StructureDefinition/RelatedArtifact
type RelatedArtifact struct {
	Id        *string             `bson:"id" json:"id"`
	Extension []Extension         `bson:"extension" json:"extension"`
	Type      RelatedArtifactType `bson:"type,omitempty" json:"type,omitempty"`
	Display   *string             `bson:"display" json:"display"`
	Citation  *string             `bson:"citation" json:"citation"`
	Url       *string             `bson:"url" json:"url"`
	Document  *Attachment         `bson:"document" json:"document"`
	Resource  *Reference          `bson:"resource" json:"resource"`
}
