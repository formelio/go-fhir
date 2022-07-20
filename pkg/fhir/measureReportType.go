package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MeasureReportType is documented here http://hl7.org/fhir/ValueSet/measure-report-type
type MeasureReportType int

const (
	MeasureReportTypeIndividual MeasureReportType = iota
	MeasureReportTypePatientList
	MeasureReportTypeSummary
)

func (code MeasureReportType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *MeasureReportType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "individual":
		*code = MeasureReportTypeIndividual
	case "patient-list":
		*code = MeasureReportTypePatientList
	case "summary":
		*code = MeasureReportTypeSummary
	default:
		return fmt.Errorf("unknown MeasureReportType code `%s`", s)
	}
	return nil
}
func (code MeasureReportType) String() string {
	return code.Code()
}
func (code MeasureReportType) Code() string {
	switch code {
	case MeasureReportTypeIndividual:
		return "individual"
	case MeasureReportTypePatientList:
		return "patient-list"
	case MeasureReportTypeSummary:
		return "summary"
	}
	return "<unknown>"
}
func (code MeasureReportType) Display() string {
	switch code {
	case MeasureReportTypeIndividual:
		return "Individual"
	case MeasureReportTypePatientList:
		return "Patient List"
	case MeasureReportTypeSummary:
		return "Summary"
	}
	return "<unknown>"
}
func (code MeasureReportType) Definition() string {
	switch code {
	case MeasureReportTypeIndividual:
		return "An individual report that provides information on the performance for a given measure with respect to a single patient"
	case MeasureReportTypePatientList:
		return "A patient list report that includes a listing of patients that satisfied each population criteria in the measure"
	case MeasureReportTypeSummary:
		return "A summary report that returns the number of patients in each population criteria for the measure"
	}
	return "<unknown>"
}
