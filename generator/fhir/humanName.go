package fhir

// HumanName is documented here http://hl7.org/fhir/StructureDefinition/HumanName
type HumanName struct {
	Id        *string     `bson:"id" json:"id"`
	Extension []Extension `bson:"extension" json:"extension"`
	Use       *NameUse    `bson:"use" json:"use"`
	Text      *string     `bson:"text" json:"text"`
	Family    *string     `bson:"family" json:"family"`
	Given     []string    `bson:"given" json:"given"`
	Prefix    []string    `bson:"prefix" json:"prefix"`
	Suffix    []string    `bson:"suffix" json:"suffix"`
	Period    *Period     `bson:"period" json:"period"`
}
