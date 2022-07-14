package fhir

import "encoding/json"

// DetectedIssue is documented here http://hl7.org/fhir/StructureDefinition/DetectedIssue
type DetectedIssue struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                   `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            string                    `bson:"status" json:"status"`
	Category          *CodeableConcept          `bson:"category,omitempty" json:"category,omitempty"`
	Severity          *string                   `bson:"severity,omitempty" json:"severity,omitempty"`
	Patient           *Reference                `bson:"patient,omitempty" json:"patient,omitempty"`
	Date              *string                   `bson:"date,omitempty" json:"date,omitempty"`
	Implicated        []Reference               `bson:"implicated,omitempty" json:"implicated,omitempty"`
	Detail            *string                   `bson:"detail,omitempty" json:"detail,omitempty"`
	Reference         *string                   `bson:"reference,omitempty" json:"reference,omitempty"`
	Mitigation        []DetectedIssueMitigation `bson:"mitigation,omitempty" json:"mitigation,omitempty"`
}
type DetectedIssueMitigation struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Action            CodeableConcept `bson:"action" json:"action"`
	Date              *string         `bson:"date,omitempty" json:"date,omitempty"`
	Author            *Reference      `bson:"author,omitempty" json:"author,omitempty"`
}
type OtherDetectedIssue DetectedIssue

// MarshalJSON marshals the given DetectedIssue as JSON into a byte slice
func (r DetectedIssue) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDetectedIssue
		ResourceType string `json:"resourceType"`
	}{
		OtherDetectedIssue: OtherDetectedIssue(r),
		ResourceType:       "DetectedIssue",
	})
}

// UnmarshalDetectedIssue unmarshals a DetectedIssue.
func UnmarshalDetectedIssue(b []byte) (DetectedIssue, error) {
	var detectedIssue DetectedIssue
	if err := json.Unmarshal(b, &detectedIssue); err != nil {
		return detectedIssue, err
	}
	return detectedIssue, nil
}
