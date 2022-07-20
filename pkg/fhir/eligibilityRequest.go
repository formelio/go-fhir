package fhir

import "encoding/json"

// EligibilityRequest is documented here http://hl7.org/fhir/StructureDefinition/EligibilityRequest
type EligibilityRequest struct {
	Id                  *string           `bson:"id" json:"id"`
	Meta                *Meta             `bson:"meta" json:"meta"`
	ImplicitRules       *string           `bson:"implicitRules" json:"implicitRules"`
	Language            *string           `bson:"language" json:"language"`
	Text                *Narrative        `bson:"text" json:"text"`
	Contained           []json.RawMessage `bson:"contained" json:"contained"`
	Extension           []Extension       `bson:"extension" json:"extension"`
	ModifierExtension   []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier          []Identifier      `bson:"identifier" json:"identifier"`
	Status              *string           `bson:"status" json:"status"`
	Priority            *CodeableConcept  `bson:"priority" json:"priority"`
	Patient             *Reference        `bson:"patient" json:"patient"`
	ServicedDate        *string           `bson:"servicedDate,omitempty" json:"servicedDate,omitempty"`
	ServicedPeriod      *Period           `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty"`
	Created             *string           `bson:"created" json:"created"`
	Enterer             *Reference        `bson:"enterer" json:"enterer"`
	Provider            *Reference        `bson:"provider" json:"provider"`
	Organization        *Reference        `bson:"organization" json:"organization"`
	Insurer             *Reference        `bson:"insurer" json:"insurer"`
	Facility            *Reference        `bson:"facility" json:"facility"`
	Coverage            *Reference        `bson:"coverage" json:"coverage"`
	BusinessArrangement *string           `bson:"businessArrangement" json:"businessArrangement"`
	BenefitCategory     *CodeableConcept  `bson:"benefitCategory" json:"benefitCategory"`
	BenefitSubCategory  *CodeableConcept  `bson:"benefitSubCategory" json:"benefitSubCategory"`
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

// UnmarshalEligibilityRequest unmarshalls a EligibilityRequest.
func UnmarshalEligibilityRequest(b []byte) (EligibilityRequest, error) {
	var eligibilityRequest EligibilityRequest
	if err := json.Unmarshal(b, &eligibilityRequest); err != nil {
		return eligibilityRequest, err
	}
	return eligibilityRequest, nil
}
