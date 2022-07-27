package fhir

import "encoding/json"

// DiagnosticReport is documented here http://hl7.org/fhir/StructureDefinition/DiagnosticReport
type DiagnosticReport struct {
	Id                *string                     `bson:"id" json:"id"`
	Meta              *Meta                       `bson:"meta" json:"meta"`
	ImplicitRules     *string                     `bson:"implicitRules" json:"implicitRules"`
	Language          *string                     `bson:"language" json:"language"`
	Text              *Narrative                  `bson:"text" json:"text"`
	RawContained      []json.RawMessage           `bson:"contained" json:"contained"`
	Contained         []IResource                 `bson:"-" json:"-"`
	Extension         []Extension                 `bson:"extension" json:"extension"`
	ModifierExtension []Extension                 `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier                `bson:"identifier" json:"identifier"`
	BasedOn           []Reference                 `bson:"basedOn" json:"basedOn"`
	Status            DiagnosticReportStatus      `bson:"status,omitempty" json:"status,omitempty"`
	Category          *CodeableConcept            `bson:"category" json:"category"`
	Code              CodeableConcept             `bson:"code,omitempty" json:"code,omitempty"`
	Subject           *Reference                  `bson:"subject" json:"subject"`
	Context           *Reference                  `bson:"context" json:"context"`
	EffectiveDateTime *string                     `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *Period                     `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Issued            *string                     `bson:"issued" json:"issued"`
	Performer         []DiagnosticReportPerformer `bson:"performer" json:"performer"`
	Specimen          []Reference                 `bson:"specimen" json:"specimen"`
	Result            []Reference                 `bson:"result" json:"result"`
	ImagingStudy      []Reference                 `bson:"imagingStudy" json:"imagingStudy"`
	Image             []DiagnosticReportImage     `bson:"image" json:"image"`
	Conclusion        *string                     `bson:"conclusion" json:"conclusion"`
	CodedDiagnosis    []CodeableConcept           `bson:"codedDiagnosis" json:"codedDiagnosis"`
	PresentedForm     []Attachment                `bson:"presentedForm" json:"presentedForm"`
}
type DiagnosticReportPerformer struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Role              *CodeableConcept `bson:"role" json:"role"`
	Actor             Reference        `bson:"actor,omitempty" json:"actor,omitempty"`
}
type DiagnosticReportImage struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Comment           *string     `bson:"comment" json:"comment"`
	Link              Reference   `bson:"link,omitempty" json:"link,omitempty"`
}

// OtherDiagnosticReport is a helper type to use the default implementations of Marshall and Unmarshal
type OtherDiagnosticReport DiagnosticReport

// MarshalJSON marshals the given DiagnosticReport as JSON into a byte slice
func (r DiagnosticReport) MarshalJSON() ([]byte, error) {
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
		OtherDiagnosticReport
		ResourceType string `json:"resourceType"`
	}{
		OtherDiagnosticReport: OtherDiagnosticReport(r),
		ResourceType:          "DiagnosticReport",
	})
}

// UnmarshalJSON unmarshals the given byte slice into DiagnosticReport
func (r *DiagnosticReport) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherDiagnosticReport)(r)); err != nil {
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
func (r DiagnosticReport) GetResourceType() ResourceType {
	return ResourceTypeDiagnosticReport
}
