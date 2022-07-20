package fhir

// DataRequirement is documented here http://hl7.org/fhir/StructureDefinition/DataRequirement
type DataRequirement struct {
	Id          *string                     `bson:"id" json:"id"`
	Extension   []Extension                 `bson:"extension" json:"extension"`
	Type        string                      `bson:"type,omitempty" json:"type,omitempty"`
	Profile     []string                    `bson:"profile" json:"profile"`
	MustSupport []string                    `bson:"mustSupport" json:"mustSupport"`
	CodeFilter  []DataRequirementCodeFilter `bson:"codeFilter" json:"codeFilter"`
	DateFilter  []DataRequirementDateFilter `bson:"dateFilter" json:"dateFilter"`
}
type DataRequirementCodeFilter struct {
	Id                   *string           `bson:"id" json:"id"`
	Extension            []Extension       `bson:"extension" json:"extension"`
	Path                 string            `bson:"path,omitempty" json:"path,omitempty"`
	ValueSetString       *string           `bson:"valueSetString,omitempty" json:"valueSetString,omitempty"`
	ValueSetReference    *Reference        `bson:"valueSetReference,omitempty" json:"valueSetReference,omitempty"`
	ValueCode            []string          `bson:"valueCode" json:"valueCode"`
	ValueCoding          []Coding          `bson:"valueCoding" json:"valueCoding"`
	ValueCodeableConcept []CodeableConcept `bson:"valueCodeableConcept" json:"valueCodeableConcept"`
}
type DataRequirementDateFilter struct {
	Id            *string     `bson:"id" json:"id"`
	Extension     []Extension `bson:"extension" json:"extension"`
	Path          string      `bson:"path,omitempty" json:"path,omitempty"`
	ValueDateTime *string     `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod   *Period     `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	ValueDuration *Duration   `bson:"valueDuration,omitempty" json:"valueDuration,omitempty"`
}
