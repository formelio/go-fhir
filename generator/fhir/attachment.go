package fhir

// Attachment is documented here http://hl7.org/fhir/StructureDefinition/Attachment
type Attachment struct {
	Id          *string     `bson:"id" json:"id"`
	Extension   []Extension `bson:"extension" json:"extension"`
	ContentType *string     `bson:"contentType" json:"contentType"`
	Language    *string     `bson:"language" json:"language"`
	Data        *string     `bson:"data" json:"data"`
	Url         *string     `bson:"url" json:"url"`
	Size        *int        `bson:"size" json:"size"`
	Hash        *string     `bson:"hash" json:"hash"`
	Title       *string     `bson:"title" json:"title"`
	Creation    *string     `bson:"creation" json:"creation"`
}
