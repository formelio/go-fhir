package fhir

// DataRequirement is documented here http://hl7.org/fhir/StructureDefinition/DataRequirement
type DataRequirement struct {
	Id          *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension   []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	Type        string                      `bson:"type" json:"type"`
	Profile     []string                    `bson:"profile,omitempty" json:"profile,omitempty"`
	MustSupport []string                    `bson:"mustSupport,omitempty" json:"mustSupport,omitempty"`
	CodeFilter  []DataRequirementCodeFilter `bson:"codeFilter,omitempty" json:"codeFilter,omitempty"`
	DateFilter  []DataRequirementDateFilter `bson:"dateFilter,omitempty" json:"dateFilter,omitempty"`
}
type DataRequirementCodeFilter struct {
	Id                   *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	Path                 string            `bson:"path" json:"path"`
	ValueCode            []string          `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueCoding          []Coding          `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueCodeableConcept []CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
}
type DataRequirementDateFilter struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Path      string      `bson:"path" json:"path"`
}
