package fhir

// Identifier is documented here http://hl7.org/fhir/StructureDefinition/Identifier
type Identifier struct {
	Id        *string          `bson:"id" json:"id"`
	Extension []Extension      `bson:"extension" json:"extension"`
	Use       *IdentifierUse   `bson:"use" json:"use"`
	Type      *CodeableConcept `bson:"type" json:"type"`
	System    *string          `bson:"system" json:"system"`
	Value     *string          `bson:"value" json:"value"`
	Period    *Period          `bson:"period" json:"period"`
	Assigner  *Reference       `bson:"assigner" json:"assigner"`
}
