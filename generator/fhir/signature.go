package fhir

// Signature is documented here http://hl7.org/fhir/StructureDefinition/Signature
type Signature struct {
	Id                  *string     `bson:"id" json:"id"`
	Extension           []Extension `bson:"extension" json:"extension"`
	Type                []Coding    `bson:"type,omitempty" json:"type,omitempty"`
	When                string      `bson:"when,omitempty" json:"when,omitempty"`
	WhoUri              *string     `bson:"whoUri,omitempty" json:"whoUri,omitempty"`
	WhoReference        *Reference  `bson:"whoReference,omitempty" json:"whoReference,omitempty"`
	OnBehalfOfUri       *string     `bson:"onBehalfOfUri,omitempty" json:"onBehalfOfUri,omitempty"`
	OnBehalfOfReference *Reference  `bson:"onBehalfOfReference,omitempty" json:"onBehalfOfReference,omitempty"`
	ContentType         *string     `bson:"contentType" json:"contentType"`
	Blob                *string     `bson:"blob" json:"blob"`
}
