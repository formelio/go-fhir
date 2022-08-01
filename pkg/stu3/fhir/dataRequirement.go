package fhir

// DataRequirement is documented here http://hl7.org/fhir/StructureDefinition/DataRequirement
type DataRequirement struct {
	Id          *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension   []*Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	Type        string                       `bson:"type,omitempty" json:"type,omitempty"`
	Profile     []*string                    `bson:"profile,omitempty" json:"profile,omitempty"`
	MustSupport []*string                    `bson:"mustSupport,omitempty" json:"mustSupport,omitempty"`
	CodeFilter  []*DataRequirementCodeFilter `bson:"codeFilter,omitempty" json:"codeFilter,omitempty"`
	DateFilter  []*DataRequirementDateFilter `bson:"dateFilter,omitempty" json:"dateFilter,omitempty"`
}
type DataRequirementCodeFilter struct {
	Id                   *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []*Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	Path                 string             `bson:"path,omitempty" json:"path,omitempty"`
	ValueSetString       *string            `bson:"valueSetString,omitempty" json:"valueSetString,omitempty"`
	ValueSetReference    *Reference         `bson:"valueSetReference,omitempty" json:"valueSetReference,omitempty"`
	ValueCode            []*string          `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueCoding          []*Coding          `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueCodeableConcept []*CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
}
type DataRequirementDateFilter struct {
	Id            *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension     []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Path          string       `bson:"path,omitempty" json:"path,omitempty"`
	ValueDateTime *string      `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod   *Period      `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	ValueDuration *Duration    `bson:"valueDuration,omitempty" json:"valueDuration,omitempty"`
}
