package fhir

// Meta is documented here http://hl7.org/fhir/StructureDefinition/Meta
type Meta struct {
	Id          *string     `bson:"id" json:"id"`
	Extension   []Extension `bson:"extension" json:"extension"`
	VersionId   *string     `bson:"versionId" json:"versionId"`
	LastUpdated *string     `bson:"lastUpdated" json:"lastUpdated"`
	Profile     []string    `bson:"profile" json:"profile"`
	Security    []Coding    `bson:"security" json:"security"`
	Tag         []Coding    `bson:"tag" json:"tag"`
}
