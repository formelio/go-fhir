package fhir

import "encoding/json"

// MeasureReport is documented here http://hl7.org/fhir/StructureDefinition/MeasureReport
type MeasureReport struct {
	Id                    *string              `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string              `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	RawContained          []json.RawMessage    `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained             []IResource          `bson:"-,omitempty" json:"-,omitempty"`
	Extension             []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            *Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                MeasureReportStatus  `bson:"status,omitempty" json:"status,omitempty"`
	Type                  MeasureReportType    `bson:"type,omitempty" json:"type,omitempty"`
	Measure               Reference            `bson:"measure,omitempty" json:"measure,omitempty"`
	Patient               *Reference           `bson:"patient,omitempty" json:"patient,omitempty"`
	Date                  *string              `bson:"date,omitempty" json:"date,omitempty"`
	ReportingOrganization *Reference           `bson:"reportingOrganization,omitempty" json:"reportingOrganization,omitempty"`
	Period                Period               `bson:"period,omitempty" json:"period,omitempty"`
	Group                 []MeasureReportGroup `bson:"group,omitempty" json:"group,omitempty"`
	EvaluatedResources    *Reference           `bson:"evaluatedResources,omitempty" json:"evaluatedResources,omitempty"`
}
type MeasureReportGroup struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        Identifier                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Population        []MeasureReportGroupPopulation `bson:"population,omitempty" json:"population,omitempty"`
	MeasureScore      *float64                       `bson:"measureScore,omitempty" json:"measureScore,omitempty"`
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
	Value             string                                          `bson:"value,omitempty" json:"value,omitempty"`
	Population        []MeasureReportGroupStratifierStratumPopulation `bson:"population,omitempty" json:"population,omitempty"`
	MeasureScore      *float64                                        `bson:"measureScore,omitempty" json:"measureScore,omitempty"`
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
