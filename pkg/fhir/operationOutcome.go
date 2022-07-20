package fhir

import "encoding/json"

// OperationOutcome is documented here http://hl7.org/fhir/StructureDefinition/OperationOutcome
type OperationOutcome struct {
	Id                *string                 `bson:"id" json:"id"`
	Meta              *Meta                   `bson:"meta" json:"meta"`
	ImplicitRules     *string                 `bson:"implicitRules" json:"implicitRules"`
	Language          *string                 `bson:"language" json:"language"`
	Text              *Narrative              `bson:"text" json:"text"`
	Contained         []json.RawMessage       `bson:"contained" json:"contained"`
	Extension         []Extension             `bson:"extension" json:"extension"`
	ModifierExtension []Extension             `bson:"modifierExtension" json:"modifierExtension"`
	Issue             []OperationOutcomeIssue `bson:"issue,omitempty" json:"issue,omitempty"`
}
type OperationOutcomeIssue struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Severity          IssueSeverity    `bson:"severity,omitempty" json:"severity,omitempty"`
	Code              IssueType        `bson:"code,omitempty" json:"code,omitempty"`
	Details           *CodeableConcept `bson:"details" json:"details"`
	Diagnostics       *string          `bson:"diagnostics" json:"diagnostics"`
	Location          []string         `bson:"location" json:"location"`
	Expression        []string         `bson:"expression" json:"expression"`
}
type OtherOperationOutcome OperationOutcome

// MarshalJSON marshals the given OperationOutcome as JSON into a byte slice
func (r OperationOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherOperationOutcome
		ResourceType string `json:"resourceType"`
	}{
		OtherOperationOutcome: OtherOperationOutcome(r),
		ResourceType:          "OperationOutcome",
	})
}

// UnmarshalOperationOutcome unmarshalls a OperationOutcome.
func UnmarshalOperationOutcome(b []byte) (OperationOutcome, error) {
	var operationOutcome OperationOutcome
	if err := json.Unmarshal(b, &operationOutcome); err != nil {
		return operationOutcome, err
	}
	return operationOutcome, nil
}
