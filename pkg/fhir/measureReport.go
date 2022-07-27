package fhir

import "encoding/json"

// MeasureReport is documented here http://hl7.org/fhir/StructureDefinition/MeasureReport
type MeasureReport struct {
	Id                    *string              `bson:"id" json:"id"`
	Meta                  *Meta                `bson:"meta" json:"meta"`
	ImplicitRules         *string              `bson:"implicitRules" json:"implicitRules"`
	Language              *string              `bson:"language" json:"language"`
	Text                  *Narrative           `bson:"text" json:"text"`
	RawContained          []json.RawMessage    `bson:"contained" json:"contained"`
	Contained             []IResource          `bson:"-" json:"-"`
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

// OtherMeasureReport is a helper type to use the default implementations of Marshall and Unmarshal
type OtherMeasureReport MeasureReport

// MarshalJSON marshals the given MeasureReport as JSON into a byte slice
func (r MeasureReport) MarshalJSON() ([]byte, error) {
	// If the field has contained resources, we need to marshal them individually and store them in .RawContained
	if len(r.Contained) > 0 {
		var err error
		r.RawContained = make([]json.RawMessage, len(r.Contained))
		for i, contained := range r.Contained {
			r.RawContained[i], err = json.Marshal(contained)
			if err != nil {
				return nil, err
			}
		}
	}
	return json.Marshal(struct {
		OtherMeasureReport
		ResourceType string `json:"resourceType"`
	}{
		OtherMeasureReport: OtherMeasureReport(r),
		ResourceType:       "MeasureReport",
	})
}

// UnmarshalJSON unmarshals the given byte slice into MeasureReport
func (r *MeasureReport) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherMeasureReport)(r)); err != nil {
		return err
	}
	// If the field has contained resources, we need to unmarshal them individually and store them in .Contained
	if len(r.RawContained) > 0 {
		var err error
		r.Contained = make([]IResource, len(r.RawContained))
		for i, rawContained := range r.RawContained {
			r.Contained[i], err = UnmarshalResource(rawContained)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Returns the resourceType of the resource, makes this resource an instance of IResource
func (r MeasureReport) GetResourceType() ResourceType {
	return ResourceTypeMeasureReport
}
