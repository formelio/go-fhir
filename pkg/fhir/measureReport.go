package fhir

import "encoding/json"

// MeasureReport is documented here http://hl7.org/fhir/StructureDefinition/MeasureReport
type MeasureReport struct {
	Id                    *string              `bson:"id" json:"id"`
	Meta                  *Meta                `bson:"meta" json:"meta"`
	ImplicitRules         *string              `bson:"implicitRules" json:"implicitRules"`
	Language              *string              `bson:"language" json:"language"`
	Text                  *Narrative           `bson:"text" json:"text"`
	Contained             []json.RawMessage    `bson:"contained" json:"contained"`
	Extension             []Extension          `bson:"extension" json:"extension"`
	ModifierExtension     []Extension          `bson:"modifierExtension" json:"modifierExtension"`
	Identifier            *Identifier          `bson:"identifier" json:"identifier"`
	Status                MeasureReportStatus  `bson:"status,omitempty" json:"status,omitempty"`
	Type                  MeasureReportType    `bson:"type,omitempty" json:"type,omitempty"`
	Measure               Reference            `bson:"measure,omitempty" json:"measure,omitempty"`
	Patient               *Reference           `bson:"patient" json:"patient"`
	Date                  *string              `bson:"date" json:"date"`
	ReportingOrganization *Reference           `bson:"reportingOrganization" json:"reportingOrganization"`
	Period                Period               `bson:"period,omitempty" json:"period,omitempty"`
	Group                 []MeasureReportGroup `bson:"group" json:"group"`
	EvaluatedResources    *Reference           `bson:"evaluatedResources" json:"evaluatedResources"`
}
type MeasureReportGroup struct {
	Id                *string                        `bson:"id" json:"id"`
	Extension         []Extension                    `bson:"extension" json:"extension"`
	ModifierExtension []Extension                    `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        Identifier                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Population        []MeasureReportGroupPopulation `bson:"population" json:"population"`
	MeasureScore      *float64                       `bson:"measureScore" json:"measureScore"`
	Stratifier        []MeasureReportGroupStratifier `bson:"stratifier" json:"stratifier"`
}
type MeasureReportGroupPopulation struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier      `bson:"identifier" json:"identifier"`
	Code              *CodeableConcept `bson:"code" json:"code"`
	Count             *int             `bson:"count" json:"count"`
	Patients          *Reference       `bson:"patients" json:"patients"`
}
type MeasureReportGroupStratifier struct {
	Id                *string                               `bson:"id" json:"id"`
	Extension         []Extension                           `bson:"extension" json:"extension"`
	ModifierExtension []Extension                           `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier                           `bson:"identifier" json:"identifier"`
	Stratum           []MeasureReportGroupStratifierStratum `bson:"stratum" json:"stratum"`
}
type MeasureReportGroupStratifierStratum struct {
	Id                *string                                         `bson:"id" json:"id"`
	Extension         []Extension                                     `bson:"extension" json:"extension"`
	ModifierExtension []Extension                                     `bson:"modifierExtension" json:"modifierExtension"`
	Value             string                                          `bson:"value,omitempty" json:"value,omitempty"`
	Population        []MeasureReportGroupStratifierStratumPopulation `bson:"population" json:"population"`
	MeasureScore      *float64                                        `bson:"measureScore" json:"measureScore"`
}
type MeasureReportGroupStratifierStratumPopulation struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier      `bson:"identifier" json:"identifier"`
	Code              *CodeableConcept `bson:"code" json:"code"`
	Count             *int             `bson:"count" json:"count"`
	Patients          *Reference       `bson:"patients" json:"patients"`
}
type OtherMeasureReport MeasureReport

// MarshalJSON marshals the given MeasureReport as JSON into a byte slice
func (r MeasureReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMeasureReport
		ResourceType string `json:"resourceType"`
	}{
		OtherMeasureReport: OtherMeasureReport(r),
		ResourceType:       "MeasureReport",
	})
}

// UnmarshalMeasureReport unmarshalls a MeasureReport.
func UnmarshalMeasureReport(b []byte) (MeasureReport, error) {
	var measureReport MeasureReport
	if err := json.Unmarshal(b, &measureReport); err != nil {
		return measureReport, err
	}
	return measureReport, nil
}
