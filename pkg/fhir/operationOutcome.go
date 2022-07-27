package fhir

import "encoding/json"

// OperationOutcome is documented here http://hl7.org/fhir/StructureDefinition/OperationOutcome
type OperationOutcome struct {
	Id                *string                 `bson:"id" json:"id"`
	Meta              *Meta                   `bson:"meta" json:"meta"`
	ImplicitRules     *string                 `bson:"implicitRules" json:"implicitRules"`
	Language          *string                 `bson:"language" json:"language"`
	Text              *Narrative              `bson:"text" json:"text"`
	RawContained      []json.RawMessage       `bson:"contained" json:"contained"`
	Contained         []IResource             `bson:"-" json:"-"`
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
	return json.Marshal(struct {
		OtherOperationOutcome
		ResourceType string `json:"resourceType"`
	}{
		OtherOperationOutcome: OtherOperationOutcome(r),
		ResourceType:          "OperationOutcome",
	})
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
