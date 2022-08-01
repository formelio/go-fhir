package fhir

// Period is documented here http://hl7.org/fhir/StructureDefinition/Period
type Period struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Start     *string     `bson:"start,omitempty" json:"start,omitempty"`
	End       *string     `bson:"end,omitempty" json:"end,omitempty"`
}
