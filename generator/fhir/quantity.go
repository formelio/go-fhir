package fhir

// Quantity is documented here http://hl7.org/fhir/StructureDefinition/Quantity
type Quantity struct {
	Id         *string             `bson:"id" json:"id"`
	Extension  []Extension         `bson:"extension" json:"extension"`
	Value      *float64            `bson:"value" json:"value"`
	Comparator *QuantityComparator `bson:"comparator" json:"comparator"`
	Unit       *string             `bson:"unit" json:"unit"`
	System     *string             `bson:"system" json:"system"`
	Code       *string             `bson:"code" json:"code"`
}
