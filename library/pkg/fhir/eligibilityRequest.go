package fhir

import "encoding/json"

// EligibilityRequest is documented here http://hl7.org/fhir/StructureDefinition/EligibilityRequest
type EligibilityRequest struct {
	Id                  *string          `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string          `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative       `bson:"text,omitempty" json:"text,omitempty"`
	Extension           []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              *string          `bson:"status,omitempty" json:"status,omitempty"`
	Priority            *CodeableConcept `bson:"priority,omitempty" json:"priority,omitempty"`
	Patient             *Reference       `bson:"patient,omitempty" json:"patient,omitempty"`
	Created             *string          `bson:"created,omitempty" json:"created,omitempty"`
	Enterer             *Reference       `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Provider            *Reference       `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization        *Reference       `bson:"organization,omitempty" json:"organization,omitempty"`
	Insurer             *Reference       `bson:"insurer,omitempty" json:"insurer,omitempty"`
	Facility            *Reference       `bson:"facility,omitempty" json:"facility,omitempty"`
	Coverage            *Reference       `bson:"coverage,omitempty" json:"coverage,omitempty"`
	BusinessArrangement *string          `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
	BenefitCategory     *CodeableConcept `bson:"benefitCategory,omitempty" json:"benefitCategory,omitempty"`
	BenefitSubCategory  *CodeableConcept `bson:"benefitSubCategory,omitempty" json:"benefitSubCategory,omitempty"`
}
type OtherEligibilityRequest EligibilityRequest

// MarshalJSON marshals the given EligibilityRequest as JSON into a byte slice
func (r EligibilityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEligibilityRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherEligibilityRequest: OtherEligibilityRequest(r),
		ResourceType:            "EligibilityRequest",
	})
}

// UnmarshalEligibilityRequest unmarshals a EligibilityRequest.
func UnmarshalEligibilityRequest(b []byte) (EligibilityRequest, error) {
	var eligibilityRequest EligibilityRequest
	if err := json.Unmarshal(b, &eligibilityRequest); err != nil {
		return eligibilityRequest, err
	}
	return eligibilityRequest, nil
}
