package fhir

// Contributor is documented here http://hl7.org/fhir/StructureDefinition/Contributor
type Contributor struct {
	Id        *string         `bson:"id" json:"id"`
	Extension []Extension     `bson:"extension" json:"extension"`
	Type      ContributorType `bson:"type,omitempty" json:"type,omitempty"`
	Name      string          `bson:"name,omitempty" json:"name,omitempty"`
	Contact   []ContactDetail `bson:"contact" json:"contact"`
}
