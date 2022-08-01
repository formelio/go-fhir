package fhir

// Attachment is documented here http://hl7.org/fhir/StructureDefinition/Attachment
type Attachment struct {
	Id          *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension   []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ContentType *string      `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Language    *string      `bson:"language,omitempty" json:"language,omitempty"`
	Data        *string      `bson:"data,omitempty" json:"data,omitempty"`
	Url         *string      `bson:"url,omitempty" json:"url,omitempty"`
	Size        *int         `bson:"size,omitempty" json:"size,omitempty"`
	Hash        *string      `bson:"hash,omitempty" json:"hash,omitempty"`
	Title       *string      `bson:"title,omitempty" json:"title,omitempty"`
	Creation    *string      `bson:"creation,omitempty" json:"creation,omitempty"`
}
