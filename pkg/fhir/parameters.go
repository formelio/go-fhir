package fhir

import "encoding/json"

// Parameters is documented here http://hl7.org/fhir/StructureDefinition/Parameters
type Parameters struct {
	Id            *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta          *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language      *string               `bson:"language,omitempty" json:"language,omitempty"`
	Parameter     []ParametersParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
}
type ParametersParameter struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                `bson:"name" json:"name"`
	Resource          json.RawMessage       `bson:"resource,omitempty" json:"resource,omitempty"`
	Part              []ParametersParameter `bson:"part,omitempty" json:"part,omitempty"`
}
type OtherParameters Parameters

// MarshalJSON marshals the given Parameters as JSON into a byte slice
func (r Parameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherParameters
		ResourceType string `json:"resourceType"`
	}{
		OtherParameters: OtherParameters(r),
		ResourceType:    "Parameters",
	})
}

// UnmarshalParameters unmarshals a Parameters.
func UnmarshalParameters(b []byte) (Parameters, error) {
	var parameters Parameters
	if err := json.Unmarshal(b, &parameters); err != nil {
		return parameters, err
	}
	return parameters, nil
}
