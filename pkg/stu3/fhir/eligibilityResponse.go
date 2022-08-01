package fhir

import (
	"bytes"
	"encoding/json"
)

// EligibilityResponse is documented here http://hl7.org/fhir/StructureDefinition/EligibilityResponse
type EligibilityResponse struct {
	Id                  *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string                         `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative                      `bson:"text,omitempty" json:"text,omitempty"`
	RawContained        []json.RawMessage               `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained           []IResource                     `bson:"-,omitempty" json:"-,omitempty"`
	Extension           []*Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []*Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []*Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              *string                         `bson:"status,omitempty" json:"status,omitempty"`
	Created             *string                         `bson:"created,omitempty" json:"created,omitempty"`
	RequestProvider     *Reference                      `bson:"requestProvider,omitempty" json:"requestProvider,omitempty"`
	RequestOrganization *Reference                      `bson:"requestOrganization,omitempty" json:"requestOrganization,omitempty"`
	Request             *Reference                      `bson:"request,omitempty" json:"request,omitempty"`
	Outcome             *CodeableConcept                `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition         *string                         `bson:"disposition,omitempty" json:"disposition,omitempty"`
	Insurer             *Reference                      `bson:"insurer,omitempty" json:"insurer,omitempty"`
	Inforce             *bool                           `bson:"inforce,omitempty" json:"inforce,omitempty"`
	Insurance           []*EligibilityResponseInsurance `bson:"insurance,omitempty" json:"insurance,omitempty"`
	Form                *CodeableConcept                `bson:"form,omitempty" json:"form,omitempty"`
	Error               []*EligibilityResponseError     `bson:"error,omitempty" json:"error,omitempty"`
}
type EligibilityResponseInsurance struct {
	Id                *string                                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Coverage          *Reference                                    `bson:"coverage,omitempty" json:"coverage,omitempty"`
	Contract          *Reference                                    `bson:"contract,omitempty" json:"contract,omitempty"`
	BenefitBalance    []*EligibilityResponseInsuranceBenefitBalance `bson:"benefitBalance,omitempty" json:"benefitBalance,omitempty"`
}
type EligibilityResponseInsuranceBenefitBalance struct {
	Id                *string                                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                                           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                                           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          CodeableConcept                                        `bson:"category,omitempty" json:"category,omitempty"`
	SubCategory       *CodeableConcept                                       `bson:"subCategory,omitempty" json:"subCategory,omitempty"`
	Excluded          *bool                                                  `bson:"excluded,omitempty" json:"excluded,omitempty"`
	Name              *string                                                `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string                                                `bson:"description,omitempty" json:"description,omitempty"`
	Network           *CodeableConcept                                       `bson:"network,omitempty" json:"network,omitempty"`
	Unit              *CodeableConcept                                       `bson:"unit,omitempty" json:"unit,omitempty"`
	Term              *CodeableConcept                                       `bson:"term,omitempty" json:"term,omitempty"`
	Financial         []*EligibilityResponseInsuranceBenefitBalanceFinancial `bson:"financial,omitempty" json:"financial,omitempty"`
}
type EligibilityResponseInsuranceBenefitBalanceFinancial struct {
	Id                 *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []*Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []*Extension    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type               CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	AllowedUnsignedInt *int            `bson:"allowedUnsignedInt,omitempty" json:"allowedUnsignedInt,omitempty"`
	AllowedString      *string         `bson:"allowedString,omitempty" json:"allowedString,omitempty"`
	AllowedMoney       *Money          `bson:"allowedMoney,omitempty" json:"allowedMoney,omitempty"`
	UsedUnsignedInt    *int            `bson:"usedUnsignedInt,omitempty" json:"usedUnsignedInt,omitempty"`
	UsedMoney          *Money          `bson:"usedMoney,omitempty" json:"usedMoney,omitempty"`
}
type EligibilityResponseError struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherEligibilityResponse
	}{
		OtherEligibilityResponse: OtherEligibilityResponse(r),
		ResourceType:             "EligibilityResponse",
	})
	return buffer.Bytes(), err
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
