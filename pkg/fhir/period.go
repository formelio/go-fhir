package fhir

// Period is documented here http://hl7.org/fhir/StructureDefinition/Period
type Period struct {
	Id        *string     `bson:"id" json:"id"`
	Extension []Extension `bson:"extension" json:"extension"`
	Start     *string     `bson:"start" json:"start"`
	End       *string     `bson:"end" json:"end"`
}
