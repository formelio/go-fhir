package fhir

// Money is documented here http://hl7.org/fhir/StructureDefinition/Money
type Money struct {
	Id         *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension  []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Value      *string     `bson:"value,omitempty" json:"value,omitempty"`
	Comparator *string     `bson:"comparator,omitempty" json:"comparator,omitempty"`
	Unit       *string     `bson:"unit,omitempty" json:"unit,omitempty"`
	System     *string     `bson:"system,omitempty" json:"system,omitempty"`
	Code       *string     `bson:"code,omitempty" json:"code,omitempty"`
}
