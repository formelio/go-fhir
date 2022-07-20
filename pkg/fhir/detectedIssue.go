package fhir

import "encoding/json"

// DetectedIssue is documented here http://hl7.org/fhir/StructureDefinition/DetectedIssue
type DetectedIssue struct {
	Id                *string                   `bson:"id" json:"id"`
	Meta              *Meta                     `bson:"meta" json:"meta"`
	ImplicitRules     *string                   `bson:"implicitRules" json:"implicitRules"`
	Language          *string                   `bson:"language" json:"language"`
	Text              *Narrative                `bson:"text" json:"text"`
	Contained         []json.RawMessage         `bson:"contained" json:"contained"`
	Extension         []Extension               `bson:"extension" json:"extension"`
	ModifierExtension []Extension               `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier               `bson:"identifier" json:"identifier"`
	Status            ObservationStatus         `bson:"status,omitempty" json:"status,omitempty"`
	Category          *CodeableConcept          `bson:"category" json:"category"`
	Severity          *DetectedIssueSeverity    `bson:"severity" json:"severity"`
	Patient           *Reference                `bson:"patient" json:"patient"`
	Date              *string                   `bson:"date" json:"date"`
	Author            *Reference                `bson:"author" json:"author"`
	Implicated        []Reference               `bson:"implicated" json:"implicated"`
	Detail            *string                   `bson:"detail" json:"detail"`
	Reference         *string                   `bson:"reference" json:"reference"`
	Mitigation        []DetectedIssueMitigation `bson:"mitigation" json:"mitigation"`
}
type DetectedIssueMitigation struct {
	Id                *string         `bson:"id" json:"id"`
	Extension         []Extension     `bson:"extension" json:"extension"`
	ModifierExtension []Extension     `bson:"modifierExtension" json:"modifierExtension"`
	Action            CodeableConcept `bson:"action,omitempty" json:"action,omitempty"`
	Date              *string         `bson:"date" json:"date"`
	Author            *Reference      `bson:"author" json:"author"`
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

// UnmarshalDetectedIssue unmarshalls a DetectedIssue.
func UnmarshalDetectedIssue(b []byte) (DetectedIssue, error) {
	var detectedIssue DetectedIssue
	if err := json.Unmarshal(b, &detectedIssue); err != nil {
		return detectedIssue, err
	}
	return detectedIssue, nil
}
