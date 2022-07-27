package fhir

import "encoding/json"

// EligibilityResponse is documented here http://hl7.org/fhir/StructureDefinition/EligibilityResponse
type EligibilityResponse struct {
	Id                  *string                        `bson:"id" json:"id"`
	Meta                *Meta                          `bson:"meta" json:"meta"`
	ImplicitRules       *string                        `bson:"implicitRules" json:"implicitRules"`
	Language            *string                        `bson:"language" json:"language"`
	Text                *Narrative                     `bson:"text" json:"text"`
	RawContained        []json.RawMessage              `bson:"contained" json:"contained"`
	Contained           []IResource                    `bson:"-" json:"-"`
	Extension           []Extension                    `bson:"extension" json:"extension"`
	ModifierExtension   []Extension                    `bson:"modifierExtension" json:"modifierExtension"`
	Identifier          []Identifier                   `bson:"identifier" json:"identifier"`
	Status              *string                        `bson:"status" json:"status"`
	Created             *string                        `bson:"created" json:"created"`
	RequestProvider     *Reference                     `bson:"requestProvider" json:"requestProvider"`
	RequestOrganization *Reference                     `bson:"requestOrganization" json:"requestOrganization"`
	Request             *Reference                     `bson:"request" json:"request"`
	Outcome             *CodeableConcept               `bson:"outcome" json:"outcome"`
	Disposition         *string                        `bson:"disposition" json:"disposition"`
	Insurer             *Reference                     `bson:"insurer" json:"insurer"`
	Inforce             *bool                          `bson:"inforce" json:"inforce"`
	Insurance           []EligibilityResponseInsurance `bson:"insurance" json:"insurance"`
	Form                *CodeableConcept               `bson:"form" json:"form"`
	Error               []EligibilityResponseError     `bson:"error" json:"error"`
}
type EligibilityResponseInsurance struct {
	Id                *string                                      `bson:"id" json:"id"`
	Extension         []Extension                                  `bson:"extension" json:"extension"`
	ModifierExtension []Extension                                  `bson:"modifierExtension" json:"modifierExtension"`
	Coverage          *Reference                                   `bson:"coverage" json:"coverage"`
	Contract          *Reference                                   `bson:"contract" json:"contract"`
	BenefitBalance    []EligibilityResponseInsuranceBenefitBalance `bson:"benefitBalance" json:"benefitBalance"`
}
type EligibilityResponseInsuranceBenefitBalance struct {
	Id                *string                                               `bson:"id" json:"id"`
	Extension         []Extension                                           `bson:"extension" json:"extension"`
	ModifierExtension []Extension                                           `bson:"modifierExtension" json:"modifierExtension"`
	Category          CodeableConcept                                       `bson:"category,omitempty" json:"category,omitempty"`
	SubCategory       *CodeableConcept                                      `bson:"subCategory" json:"subCategory"`
	Excluded          *bool                                                 `bson:"excluded" json:"excluded"`
	Name              *string                                               `bson:"name" json:"name"`
	Description       *string                                               `bson:"description" json:"description"`
	Network           *CodeableConcept                                      `bson:"network" json:"network"`
	Unit              *CodeableConcept                                      `bson:"unit" json:"unit"`
	Term              *CodeableConcept                                      `bson:"term" json:"term"`
	Financial         []EligibilityResponseInsuranceBenefitBalanceFinancial `bson:"financial" json:"financial"`
}
type EligibilityResponseInsuranceBenefitBalanceFinancial struct {
	Id                 *string         `bson:"id" json:"id"`
	Extension          []Extension     `bson:"extension" json:"extension"`
	ModifierExtension  []Extension     `bson:"modifierExtension" json:"modifierExtension"`
	Type               CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	AllowedUnsignedInt *int            `bson:"allowedUnsignedInt,omitempty" json:"allowedUnsignedInt,omitempty"`
	AllowedString      *string         `bson:"allowedString,omitempty" json:"allowedString,omitempty"`
	AllowedMoney       *Money          `bson:"allowedMoney,omitempty" json:"allowedMoney,omitempty"`
	UsedUnsignedInt    *int            `bson:"usedUnsignedInt,omitempty" json:"usedUnsignedInt,omitempty"`
	UsedMoney          *Money          `bson:"usedMoney,omitempty" json:"usedMoney,omitempty"`
}
type EligibilityResponseError struct {
	Id                *string         `bson:"id" json:"id"`
	Extension         []Extension     `bson:"extension" json:"extension"`
	ModifierExtension []Extension     `bson:"modifierExtension" json:"modifierExtension"`
	Code              CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
}

// OtherEligibilityResponse is a helper type to use the default implementations of Marshall and Unmarshal
type OtherEligibilityResponse EligibilityResponse

// MarshalJSON marshals the given EligibilityResponse as JSON into a byte slice
func (r EligibilityResponse) MarshalJSON() ([]byte, error) {
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
		OtherEligibilityResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherEligibilityResponse: OtherEligibilityResponse(r),
		ResourceType:             "EligibilityResponse",
	})
}

// UnmarshalJSON unmarshals the given byte slice into EligibilityResponse
func (r *EligibilityResponse) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherEligibilityResponse)(r)); err != nil {
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
func (r EligibilityResponse) GetResourceType() ResourceType {
	return ResourceTypeEligibilityResponse
}
