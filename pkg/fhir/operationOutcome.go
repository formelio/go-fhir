package fhir

import "encoding/json"

// OperationOutcome is documented here http://hl7.org/fhir/StructureDefinition/OperationOutcome
type OperationOutcome struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Issue             []OperationOutcomeIssue `bson:"issue" json:"issue"`
}
type OperationOutcomeIssue struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Severity          string           `bson:"severity" json:"severity"`
	Code              string           `bson:"code" json:"code"`
	Details           *CodeableConcept `bson:"details,omitempty" json:"details,omitempty"`
	Diagnostics       *string          `bson:"diagnostics,omitempty" json:"diagnostics,omitempty"`
	Location          []string         `bson:"location,omitempty" json:"location,omitempty"`
	Expression        []string         `bson:"expression,omitempty" json:"expression,omitempty"`
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

// UnmarshalOperationOutcome unmarshals a OperationOutcome.
func UnmarshalOperationOutcome(b []byte) (OperationOutcome, error) {
	var operationOutcome OperationOutcome
	if err := json.Unmarshal(b, &operationOutcome); err != nil {
		return operationOutcome, err
	}
	return operationOutcome, nil
}
