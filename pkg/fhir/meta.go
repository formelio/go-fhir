package fhir

// Meta is documented here http://hl7.org/fhir/StructureDefinition/Meta
type Meta struct {
	Id          *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension   []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	VersionId   *string     `bson:"versionId,omitempty" json:"versionId,omitempty"`
	LastUpdated *string     `bson:"lastUpdated,omitempty" json:"lastUpdated,omitempty"`
	Profile     []string    `bson:"profile,omitempty" json:"profile,omitempty"`
	Security    []Coding    `bson:"security,omitempty" json:"security,omitempty"`
	Tag         []Coding    `bson:"tag,omitempty" json:"tag,omitempty"`
}
