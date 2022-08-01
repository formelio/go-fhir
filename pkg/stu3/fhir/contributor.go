package fhir

// Contributor is documented here http://hl7.org/fhir/StructureDefinition/Contributor
type Contributor struct {
	Id        *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	Type      ContributorType  `bson:"type,omitempty" json:"type,omitempty"`
	Name      string           `bson:"name,omitempty" json:"name,omitempty"`
	Contact   []*ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
}
