package fhir

// Signature is documented here http://hl7.org/fhir/StructureDefinition/Signature
type Signature struct {
	Id                  *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Type                []Coding     `bson:"type,omitempty" json:"type,omitempty"`
	When                string       `bson:"when,omitempty" json:"when,omitempty"`
	WhoUri              *string      `bson:"whoUri,omitempty" json:"whoUri,omitempty"`
	WhoReference        *Reference   `bson:"whoReference,omitempty" json:"whoReference,omitempty"`
	OnBehalfOfUri       *string      `bson:"onBehalfOfUri,omitempty" json:"onBehalfOfUri,omitempty"`
	OnBehalfOfReference *Reference   `bson:"onBehalfOfReference,omitempty" json:"onBehalfOfReference,omitempty"`
	ContentType         *string      `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Blob                *string      `bson:"blob,omitempty" json:"blob,omitempty"`
}
