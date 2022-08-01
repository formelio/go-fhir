package fhir

// ContactDetail is documented here http://hl7.org/fhir/StructureDefinition/ContactDetail
type ContactDetail struct {
	Id        *string        `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	Name      *string        `bson:"name,omitempty" json:"name,omitempty"`
	Telecom   []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}
