package fhir

// ContactPoint is documented here http://hl7.org/fhir/StructureDefinition/ContactPoint
type ContactPoint struct {
	Id        *string             `bson:"id" json:"id"`
	Extension []Extension         `bson:"extension" json:"extension"`
	System    *ContactPointSystem `bson:"system" json:"system"`
	Value     *string             `bson:"value" json:"value"`
	Use       *ContactPointUse    `bson:"use" json:"use"`
	Rank      *int                `bson:"rank" json:"rank"`
	Period    *Period             `bson:"period" json:"period"`
}
