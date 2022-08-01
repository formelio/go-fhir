package fhir

// Identifier is documented here http://hl7.org/fhir/StructureDefinition/Identifier
type Identifier struct {
	Id        *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	Use       *IdentifierUse   `bson:"use,omitempty" json:"use,omitempty"`
	Type      *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	System    *string          `bson:"system,omitempty" json:"system,omitempty"`
	Value     *string          `bson:"value,omitempty" json:"value,omitempty"`
	Period    *Period          `bson:"period,omitempty" json:"period,omitempty"`
	Assigner  *Reference       `bson:"assigner,omitempty" json:"assigner,omitempty"`
}
