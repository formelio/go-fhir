package fhir

// Narrative is documented here http://hl7.org/fhir/StructureDefinition/Narrative
type Narrative struct {
	Id        *string         `bson:"id" json:"id"`
	Extension []Extension     `bson:"extension" json:"extension"`
	Status    NarrativeStatus `bson:"status,omitempty" json:"status,omitempty"`
	Div       string          `bson:"div,omitempty" json:"div,omitempty"`
}
