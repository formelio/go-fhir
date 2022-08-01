package fhir

// Narrative is documented here http://hl7.org/fhir/StructureDefinition/Narrative
type Narrative struct {
	Id        *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	Status    NarrativeStatus `bson:"status,omitempty" json:"status,omitempty"`
	Div       string          `bson:"div,omitempty" json:"div,omitempty"`
}
