package fhir

// ContactDetail is documented here http://hl7.org/fhir/StructureDefinition/ContactDetail
type ContactDetail struct {
	Id        *string        `bson:"id" json:"id"`
	Extension []Extension    `bson:"extension" json:"extension"`
	Name      *string        `bson:"name" json:"name"`
	Telecom   []ContactPoint `bson:"telecom" json:"telecom"`
}
