package fhir

// Extension is documented here http://hl7.org/fhir/StructureDefinition/Extension
type Extension struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Url       string      `bson:"url" json:"url"`
}
