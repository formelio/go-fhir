package fhir

// Contributor is documented here http://hl7.org/fhir/StructureDefinition/Contributor
type Contributor struct {
	Id        *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	Type      string          `bson:"type" json:"type"`
	Name      string          `bson:"name" json:"name"`
	Contact   []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
}
