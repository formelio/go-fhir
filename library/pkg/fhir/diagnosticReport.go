package fhir

import "encoding/json"

// DiagnosticReport is documented here http://hl7.org/fhir/StructureDefinition/DiagnosticReport
type DiagnosticReport struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            string                      `bson:"status" json:"status"`
	Category          *CodeableConcept            `bson:"category,omitempty" json:"category,omitempty"`
	Code              CodeableConcept             `bson:"code" json:"code"`
	Issued            *string                     `bson:"issued,omitempty" json:"issued,omitempty"`
	Performer         []DiagnosticReportPerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	Specimen          []Reference                 `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Result            []Reference                 `bson:"result,omitempty" json:"result,omitempty"`
	Image             []DiagnosticReportImage     `bson:"image,omitempty" json:"image,omitempty"`
	Conclusion        *string                     `bson:"conclusion,omitempty" json:"conclusion,omitempty"`
	CodedDiagnosis    []CodeableConcept           `bson:"codedDiagnosis,omitempty" json:"codedDiagnosis,omitempty"`
	PresentedForm     []Attachment                `bson:"presentedForm,omitempty" json:"presentedForm,omitempty"`
}
type DiagnosticReportPerformer struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}
type DiagnosticReportImage struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Comment           *string     `bson:"comment,omitempty" json:"comment,omitempty"`
	Link              Reference   `bson:"link" json:"link"`
}
type OtherDiagnosticReport DiagnosticReport

// MarshalJSON marshals the given DiagnosticReport as JSON into a byte slice
func (r DiagnosticReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDiagnosticReport
		ResourceType string `json:"resourceType"`
	}{
		OtherDiagnosticReport: OtherDiagnosticReport(r),
		ResourceType:          "DiagnosticReport",
	})
}

// UnmarshalDiagnosticReport unmarshals a DiagnosticReport.
func UnmarshalDiagnosticReport(b []byte) (DiagnosticReport, error) {
	var diagnosticReport DiagnosticReport
	if err := json.Unmarshal(b, &diagnosticReport); err != nil {
		return diagnosticReport, err
	}
	return diagnosticReport, nil
}
