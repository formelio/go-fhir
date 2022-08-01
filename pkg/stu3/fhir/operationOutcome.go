package fhir

import (
	"bytes"
	"encoding/json"
)

// OperationOutcome is documented here http://hl7.org/fhir/StructureDefinition/OperationOutcome
type OperationOutcome struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage       `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource             `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Issue             []OperationOutcomeIssue `bson:"issue,omitempty" json:"issue,omitempty"`
}
type OperationOutcomeIssue struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Severity          IssueSeverity    `bson:"severity,omitempty" json:"severity,omitempty"`
	Code              IssueType        `bson:"code,omitempty" json:"code,omitempty"`
	Details           *CodeableConcept `bson:"details,omitempty" json:"details,omitempty"`
	Diagnostics       *string          `bson:"diagnostics,omitempty" json:"diagnostics,omitempty"`
	Location          []*string        `bson:"location,omitempty" json:"location,omitempty"`
	Expression        []*string        `bson:"expression,omitempty" json:"expression,omitempty"`
}

// OtherOperationOutcome is a helper type to use the default implementations of Marshall and Unmarshal
type OtherOperationOutcome OperationOutcome

// MarshalJSON marshals the given OperationOutcome as JSON into a byte slice
func (r OperationOutcome) MarshalJSON() ([]byte, error) {
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
		OtherOperationOutcome
	}{
		OtherOperationOutcome: OtherOperationOutcome(r),
		ResourceType:          "OperationOutcome",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into OperationOutcome
func (r *OperationOutcome) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherOperationOutcome)(r)); err != nil {
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
func (r OperationOutcome) GetResourceType() ResourceType {
	return ResourceTypeOperationOutcome
}
