package fhir

import "encoding/json"

// EligibilityResponse is documented here http://hl7.org/fhir/StructureDefinition/EligibilityResponse
type EligibilityResponse struct {
	Id                  *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                          `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string                        `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string                        `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative                     `bson:"text,omitempty" json:"text,omitempty"`
	Extension           []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              *string                        `bson:"status,omitempty" json:"status,omitempty"`
	Created             *string                        `bson:"created,omitempty" json:"created,omitempty"`
	RequestProvider     *Reference                     `bson:"requestProvider,omitempty" json:"requestProvider,omitempty"`
	RequestOrganization *Reference                     `bson:"requestOrganization,omitempty" json:"requestOrganization,omitempty"`
	Request             *Reference                     `bson:"request,omitempty" json:"request,omitempty"`
	Outcome             *CodeableConcept               `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition         *string                        `bson:"disposition,omitempty" json:"disposition,omitempty"`
	Insurer             *Reference                     `bson:"insurer,omitempty" json:"insurer,omitempty"`
	Inforce             *bool                          `bson:"inforce,omitempty" json:"inforce,omitempty"`
	Insurance           []EligibilityResponseInsurance `bson:"insurance,omitempty" json:"insurance,omitempty"`
	Form                *CodeableConcept               `bson:"form,omitempty" json:"form,omitempty"`
	Error               []EligibilityResponseError     `bson:"error,omitempty" json:"error,omitempty"`
}
type EligibilityResponseInsurance struct {
	Id                *string                                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Coverage          *Reference                                   `bson:"coverage,omitempty" json:"coverage,omitempty"`
	Contract          *Reference                                   `bson:"contract,omitempty" json:"contract,omitempty"`
	BenefitBalance    []EligibilityResponseInsuranceBenefitBalance `bson:"benefitBalance,omitempty" json:"benefitBalance,omitempty"`
}
type EligibilityResponseInsuranceBenefitBalance struct {
	Id                *string                                               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          CodeableConcept                                       `bson:"category" json:"category"`
	SubCategory       *CodeableConcept                                      `bson:"subCategory,omitempty" json:"subCategory,omitempty"`
	Excluded          *bool                                                 `bson:"excluded,omitempty" json:"excluded,omitempty"`
	Name              *string                                               `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string                                               `bson:"description,omitempty" json:"description,omitempty"`
	Network           *CodeableConcept                                      `bson:"network,omitempty" json:"network,omitempty"`
	Unit              *CodeableConcept                                      `bson:"unit,omitempty" json:"unit,omitempty"`
	Term              *CodeableConcept                                      `bson:"term,omitempty" json:"term,omitempty"`
	Financial         []EligibilityResponseInsuranceBenefitBalanceFinancial `bson:"financial,omitempty" json:"financial,omitempty"`
}
type EligibilityResponseInsuranceBenefitBalanceFinancial struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type" json:"type"`
}
type EligibilityResponseError struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept `bson:"code" json:"code"`
}
type OtherEligibilityResponse EligibilityResponse

// MarshalJSON marshals the given EligibilityResponse as JSON into a byte slice
func (r EligibilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEligibilityResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherEligibilityResponse: OtherEligibilityResponse(r),
		ResourceType:             "EligibilityResponse",
	})
}

// UnmarshalEligibilityResponse unmarshals a EligibilityResponse.
func UnmarshalEligibilityResponse(b []byte) (EligibilityResponse, error) {
	var eligibilityResponse EligibilityResponse
	if err := json.Unmarshal(b, &eligibilityResponse); err != nil {
		return eligibilityResponse, err
	}
	return eligibilityResponse, nil
}
