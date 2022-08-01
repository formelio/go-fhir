package fhir

import (
	"bytes"
	"encoding/json"
)

// DiagnosticReport is documented here http://hl7.org/fhir/StructureDefinition/DiagnosticReport
type DiagnosticReport struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage            `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource                  `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []*Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn           []*Reference                 `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Status            DiagnosticReportStatus       `bson:"status,omitempty" json:"status,omitempty"`
	Category          *CodeableConcept             `bson:"category,omitempty" json:"category,omitempty"`
	Code              CodeableConcept              `bson:"code,omitempty" json:"code,omitempty"`
	Subject           *Reference                   `bson:"subject,omitempty" json:"subject,omitempty"`
	Context           *Reference                   `bson:"context,omitempty" json:"context,omitempty"`
	EffectiveDateTime *string                      `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *Period                      `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Issued            *string                      `bson:"issued,omitempty" json:"issued,omitempty"`
	Performer         []*DiagnosticReportPerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	Specimen          []*Reference                 `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Result            []*Reference                 `bson:"result,omitempty" json:"result,omitempty"`
	ImagingStudy      []*Reference                 `bson:"imagingStudy,omitempty" json:"imagingStudy,omitempty"`
	Image             []*DiagnosticReportImage     `bson:"image,omitempty" json:"image,omitempty"`
	Conclusion        *string                      `bson:"conclusion,omitempty" json:"conclusion,omitempty"`
	CodedDiagnosis    []*CodeableConcept           `bson:"codedDiagnosis,omitempty" json:"codedDiagnosis,omitempty"`
	PresentedForm     []*Attachment                `bson:"presentedForm,omitempty" json:"presentedForm,omitempty"`
}
type DiagnosticReportPerformer struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Actor             Reference        `bson:"actor,omitempty" json:"actor,omitempty"`
}
type DiagnosticReportImage struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Comment           *string      `bson:"comment,omitempty" json:"comment,omitempty"`
	Link              Reference    `bson:"link,omitempty" json:"link,omitempty"`
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherDiagnosticReport
	}{
		OtherDiagnosticReport: OtherDiagnosticReport(r),
		ResourceType:          "DiagnosticReport",
	})
	return buffer.Bytes(), err
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
