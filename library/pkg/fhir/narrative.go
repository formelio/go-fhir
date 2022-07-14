package fhir

// Narrative is documented here http://hl7.org/fhir/StructureDefinition/Narrative
type Narrative struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Status    string      `bson:"status" json:"status"`
	Div       string      `bson:"div" json:"div"`
}
