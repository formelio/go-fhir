package fhir

import "encoding/json"

// MeasureReport is documented here http://hl7.org/fhir/StructureDefinition/MeasureReport
type MeasureReport struct {
	Id                    *string              `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string              `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            *Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                string               `bson:"status" json:"status"`
	Type                  string               `bson:"type" json:"type"`
	Measure               Reference            `bson:"measure" json:"measure"`
	Patient               *Reference           `bson:"patient,omitempty" json:"patient,omitempty"`
	Date                  *string              `bson:"date,omitempty" json:"date,omitempty"`
	ReportingOrganization *Reference           `bson:"reportingOrganization,omitempty" json:"reportingOrganization,omitempty"`
	Period                Period               `bson:"period" json:"period"`
	Group                 []MeasureReportGroup `bson:"group,omitempty" json:"group,omitempty"`
	EvaluatedResources    *Reference           `bson:"evaluatedResources,omitempty" json:"evaluatedResources,omitempty"`
}
type MeasureReportGroup struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        Identifier                     `bson:"identifier" json:"identifier"`
	Population        []MeasureReportGroupPopulation `bson:"population,omitempty" json:"population,omitempty"`
	MeasureScore      *string                        `bson:"measureScore,omitempty" json:"measureScore,omitempty"`
	Stratifier        []MeasureReportGroupStratifier `bson:"stratifier,omitempty" json:"stratifier,omitempty"`
}
type MeasureReportGroupPopulation struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Count             *int             `bson:"count,omitempty" json:"count,omitempty"`
	Patients          *Reference       `bson:"patients,omitempty" json:"patients,omitempty"`
}
type MeasureReportGroupStratifier struct {
	Id                *string                               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier                           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Stratum           []MeasureReportGroupStratifierStratum `bson:"stratum,omitempty" json:"stratum,omitempty"`
}
type MeasureReportGroupStratifierStratum struct {
	Id                *string                                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Value             string                                          `bson:"value" json:"value"`
	Population        []MeasureReportGroupStratifierStratumPopulation `bson:"population,omitempty" json:"population,omitempty"`
	MeasureScore      *string                                         `bson:"measureScore,omitempty" json:"measureScore,omitempty"`
}
type MeasureReportGroupStratifierStratumPopulation struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Count             *int             `bson:"count,omitempty" json:"count,omitempty"`
	Patients          *Reference       `bson:"patients,omitempty" json:"patients,omitempty"`
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

// UnmarshalMeasureReport unmarshals a MeasureReport.
func UnmarshalMeasureReport(b []byte) (MeasureReport, error) {
	var measureReport MeasureReport
	if err := json.Unmarshal(b, &measureReport); err != nil {
		return measureReport, err
	}
	return measureReport, nil
}
