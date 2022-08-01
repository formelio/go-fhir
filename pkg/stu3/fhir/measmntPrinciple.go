package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MeasmntPrinciple is documented here http://hl7.org/fhir/ValueSet/measurement-principle
type MeasmntPrinciple int

const (
	MeasmntPrincipleOther MeasmntPrinciple = iota
	MeasmntPrincipleChemical
	MeasmntPrincipleElectrical
	MeasmntPrincipleImpedance
	MeasmntPrincipleNuclear
	MeasmntPrincipleOptical
	MeasmntPrincipleThermal
	MeasmntPrincipleBiological
	MeasmntPrincipleMechanical
	MeasmntPrincipleAcoustical
	MeasmntPrincipleManual
)

func (code MeasmntPrinciple) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *MeasmntPrinciple) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "other":
		*code = MeasmntPrincipleOther
	case "chemical":
		*code = MeasmntPrincipleChemical
	case "electrical":
		*code = MeasmntPrincipleElectrical
	case "impedance":
		*code = MeasmntPrincipleImpedance
	case "nuclear":
		*code = MeasmntPrincipleNuclear
	case "optical":
		*code = MeasmntPrincipleOptical
	case "thermal":
		*code = MeasmntPrincipleThermal
	case "biological":
		*code = MeasmntPrincipleBiological
	case "mechanical":
		*code = MeasmntPrincipleMechanical
	case "acoustical":
		*code = MeasmntPrincipleAcoustical
	case "manual":
		*code = MeasmntPrincipleManual
	default:
		return fmt.Errorf("unknown MeasmntPrinciple code `%s`", s)
	}
	return nil
}
func (code MeasmntPrinciple) String() string {
	return code.Code()
}
func (code MeasmntPrinciple) Code() string {
	switch code {
	case MeasmntPrincipleOther:
		return "other"
	case MeasmntPrincipleChemical:
		return "chemical"
	case MeasmntPrincipleElectrical:
		return "electrical"
	case MeasmntPrincipleImpedance:
		return "impedance"
	case MeasmntPrincipleNuclear:
		return "nuclear"
	case MeasmntPrincipleOptical:
		return "optical"
	case MeasmntPrincipleThermal:
		return "thermal"
	case MeasmntPrincipleBiological:
		return "biological"
	case MeasmntPrincipleMechanical:
		return "mechanical"
	case MeasmntPrincipleAcoustical:
		return "acoustical"
	case MeasmntPrincipleManual:
		return "manual"
	}
	return "<unknown>"
}
func (code MeasmntPrinciple) Display() string {
	switch code {
	case MeasmntPrincipleOther:
		return "MSP Other"
	case MeasmntPrincipleChemical:
		return "MSP Chemical"
	case MeasmntPrincipleElectrical:
		return "MSP Electrical"
	case MeasmntPrincipleImpedance:
		return "MSP Impedance"
	case MeasmntPrincipleNuclear:
		return "MSP Nuclear"
	case MeasmntPrincipleOptical:
		return "MSP Optical"
	case MeasmntPrincipleThermal:
		return "MSP Thermal"
	case MeasmntPrincipleBiological:
		return "MSP Biological"
	case MeasmntPrincipleMechanical:
		return "MSP Mechanical"
	case MeasmntPrincipleAcoustical:
		return "MSP Acoustical"
	case MeasmntPrincipleManual:
		return "MSP Manual"
	}
	return "<unknown>"
}
func (code MeasmntPrinciple) Definition() string {
	switch code {
	case MeasmntPrincipleOther:
		return "Measurement principle isn't in the list."
	case MeasmntPrincipleChemical:
		return "Measurement is done using the chemical principle."
	case MeasmntPrincipleElectrical:
		return "Measurement is done using the electrical principle."
	case MeasmntPrincipleImpedance:
		return "Measurement is done using the impedance principle."
	case MeasmntPrincipleNuclear:
		return "Measurement is done using the nuclear principle."
	case MeasmntPrincipleOptical:
		return "Measurement is done using the optical principle."
	case MeasmntPrincipleThermal:
		return "Measurement is done using the thermal principle."
	case MeasmntPrincipleBiological:
		return "Measurement is done using the biological principle."
	case MeasmntPrincipleMechanical:
		return "Measurement is done using the mechanical principle."
	case MeasmntPrincipleAcoustical:
		return "Measurement is done using the acoustical principle."
	case MeasmntPrincipleManual:
		return "Measurement is done using the manual principle."
	}
	return "<unknown>"
}
