package fhir

import "encoding/json"

// ExplanationOfBenefit is documented here http://hl7.org/fhir/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefit struct {
	Id                   *string                              `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                              `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative                           `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension                          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension                          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               *string                              `bson:"status,omitempty" json:"status,omitempty"`
	Type                 *CodeableConcept                     `bson:"type,omitempty" json:"type,omitempty"`
	SubType              []CodeableConcept                    `bson:"subType,omitempty" json:"subType,omitempty"`
	Patient              *Reference                           `bson:"patient,omitempty" json:"patient,omitempty"`
	BillablePeriod       *Period                              `bson:"billablePeriod,omitempty" json:"billablePeriod,omitempty"`
	Created              *string                              `bson:"created,omitempty" json:"created,omitempty"`
	Enterer              *Reference                           `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Insurer              *Reference                           `bson:"insurer,omitempty" json:"insurer,omitempty"`
	Provider             *Reference                           `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization         *Reference                           `bson:"organization,omitempty" json:"organization,omitempty"`
	Referral             *Reference                           `bson:"referral,omitempty" json:"referral,omitempty"`
	Facility             *Reference                           `bson:"facility,omitempty" json:"facility,omitempty"`
	Claim                *Reference                           `bson:"claim,omitempty" json:"claim,omitempty"`
	ClaimResponse        *Reference                           `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
	Outcome              *CodeableConcept                     `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition          *string                              `bson:"disposition,omitempty" json:"disposition,omitempty"`
	Related              []ExplanationOfBenefitRelated        `bson:"related,omitempty" json:"related,omitempty"`
	OriginalPrescription *Reference                           `bson:"originalPrescription,omitempty" json:"originalPrescription,omitempty"`
	Payee                *ExplanationOfBenefitPayee           `bson:"payee,omitempty" json:"payee,omitempty"`
	Information          []ExplanationOfBenefitInformation    `bson:"information,omitempty" json:"information,omitempty"`
	CareTeam             []ExplanationOfBenefitCareTeam       `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	Diagnosis            []ExplanationOfBenefitDiagnosis      `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Procedure            []ExplanationOfBenefitProcedure      `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Precedence           *int                                 `bson:"precedence,omitempty" json:"precedence,omitempty"`
	Insurance            *ExplanationOfBenefitInsurance       `bson:"insurance,omitempty" json:"insurance,omitempty"`
	Accident             *ExplanationOfBenefitAccident        `bson:"accident,omitempty" json:"accident,omitempty"`
	EmploymentImpacted   *Period                              `bson:"employmentImpacted,omitempty" json:"employmentImpacted,omitempty"`
	Hospitalization      *Period                              `bson:"hospitalization,omitempty" json:"hospitalization,omitempty"`
	Item                 []ExplanationOfBenefitItem           `bson:"item,omitempty" json:"item,omitempty"`
	AddItem              []ExplanationOfBenefitAddItem        `bson:"addItem,omitempty" json:"addItem,omitempty"`
	TotalCost            *Money                               `bson:"totalCost,omitempty" json:"totalCost,omitempty"`
	UnallocDeductable    *Money                               `bson:"unallocDeductable,omitempty" json:"unallocDeductable,omitempty"`
	TotalBenefit         *Money                               `bson:"totalBenefit,omitempty" json:"totalBenefit,omitempty"`
	Payment              *ExplanationOfBenefitPayment         `bson:"payment,omitempty" json:"payment,omitempty"`
	Form                 *CodeableConcept                     `bson:"form,omitempty" json:"form,omitempty"`
	ProcessNote          []ExplanationOfBenefitProcessNote    `bson:"processNote,omitempty" json:"processNote,omitempty"`
	BenefitBalance       []ExplanationOfBenefitBenefitBalance `bson:"benefitBalance,omitempty" json:"benefitBalance,omitempty"`
}
type ExplanationOfBenefitRelated struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Claim             *Reference       `bson:"claim,omitempty" json:"claim,omitempty"`
	Relationship      *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Reference         *Identifier      `bson:"reference,omitempty" json:"reference,omitempty"`
}
type ExplanationOfBenefitPayee struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	ResourceType      *CodeableConcept `bson:"resourceType,omitempty" json:"resourceType,omitempty"`
}
type ExplanationOfBenefitInformation struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int              `bson:"sequence" json:"sequence"`
	Category          CodeableConcept  `bson:"category" json:"category"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Reason            *Coding          `bson:"reason,omitempty" json:"reason,omitempty"`
}
type ExplanationOfBenefitCareTeam struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int              `bson:"sequence" json:"sequence"`
	Responsible       *bool            `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Qualification     *CodeableConcept `bson:"qualification,omitempty" json:"qualification,omitempty"`
}
type ExplanationOfBenefitDiagnosis struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int               `bson:"sequence" json:"sequence"`
	Type              []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	PackageCode       *CodeableConcept  `bson:"packageCode,omitempty" json:"packageCode,omitempty"`
}
type ExplanationOfBenefitProcedure struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int         `bson:"sequence" json:"sequence"`
	Date              *string     `bson:"date,omitempty" json:"date,omitempty"`
}
type ExplanationOfBenefitInsurance struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Coverage          *Reference  `bson:"coverage,omitempty" json:"coverage,omitempty"`
	PreAuthRef        []string    `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
}
type ExplanationOfBenefitAccident struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Date              *string          `bson:"date,omitempty" json:"date,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
}
type ExplanationOfBenefitItem struct {
	Id                *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int                                    `bson:"sequence" json:"sequence"`
	CareTeamLinkId    []int                                  `bson:"careTeamLinkId,omitempty" json:"careTeamLinkId,omitempty"`
	DiagnosisLinkId   []int                                  `bson:"diagnosisLinkId,omitempty" json:"diagnosisLinkId,omitempty"`
	ProcedureLinkId   []int                                  `bson:"procedureLinkId,omitempty" json:"procedureLinkId,omitempty"`
	InformationLinkId []int                                  `bson:"informationLinkId,omitempty" json:"informationLinkId,omitempty"`
	Revenue           *CodeableConcept                       `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept                       `bson:"category,omitempty" json:"category,omitempty"`
	Service           *CodeableConcept                       `bson:"service,omitempty" json:"service,omitempty"`
	Modifier          []CodeableConcept                      `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept                      `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity          *Quantity                              `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money                                 `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *string                                `bson:"factor,omitempty" json:"factor,omitempty"`
	Net               *Money                                 `bson:"net,omitempty" json:"net,omitempty"`
	Udi               []Reference                            `bson:"udi,omitempty" json:"udi,omitempty"`
	BodySite          *CodeableConcept                       `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SubSite           []CodeableConcept                      `bson:"subSite,omitempty" json:"subSite,omitempty"`
	Encounter         []Reference                            `bson:"encounter,omitempty" json:"encounter,omitempty"`
	NoteNumber        []int                                  `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ExplanationOfBenefitItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail            []ExplanationOfBenefitItemDetail       `bson:"detail,omitempty" json:"detail,omitempty"`
}
type ExplanationOfBenefitItemAdjudication struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          CodeableConcept  `bson:"category" json:"category"`
	Reason            *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	Amount            *Money           `bson:"amount,omitempty" json:"amount,omitempty"`
	Value             *string          `bson:"value,omitempty" json:"value,omitempty"`
}
type ExplanationOfBenefitItemDetail struct {
	Id                *string                                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int                                       `bson:"sequence" json:"sequence"`
	Type              CodeableConcept                           `bson:"type" json:"type"`
	Revenue           *CodeableConcept                          `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept                          `bson:"category,omitempty" json:"category,omitempty"`
	Service           *CodeableConcept                          `bson:"service,omitempty" json:"service,omitempty"`
	Modifier          []CodeableConcept                         `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept                         `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity          *Quantity                                 `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money                                    `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *string                                   `bson:"factor,omitempty" json:"factor,omitempty"`
	Net               *Money                                    `bson:"net,omitempty" json:"net,omitempty"`
	Udi               []Reference                               `bson:"udi,omitempty" json:"udi,omitempty"`
	NoteNumber        []int                                     `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ExplanationOfBenefitItemAdjudication    `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	SubDetail         []ExplanationOfBenefitItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}
type ExplanationOfBenefitItemDetailSubDetail struct {
	Id                *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int                                    `bson:"sequence" json:"sequence"`
	Type              CodeableConcept                        `bson:"type" json:"type"`
	Revenue           *CodeableConcept                       `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept                       `bson:"category,omitempty" json:"category,omitempty"`
	Service           *CodeableConcept                       `bson:"service,omitempty" json:"service,omitempty"`
	Modifier          []CodeableConcept                      `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept                      `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity          *Quantity                              `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money                                 `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *string                                `bson:"factor,omitempty" json:"factor,omitempty"`
	Net               *Money                                 `bson:"net,omitempty" json:"net,omitempty"`
	Udi               []Reference                            `bson:"udi,omitempty" json:"udi,omitempty"`
	NoteNumber        []int                                  `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ExplanationOfBenefitItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}
type ExplanationOfBenefitAddItem struct {
	Id                *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SequenceLinkId    []int                                  `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	Revenue           *CodeableConcept                       `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept                       `bson:"category,omitempty" json:"category,omitempty"`
	Service           *CodeableConcept                       `bson:"service,omitempty" json:"service,omitempty"`
	Modifier          []CodeableConcept                      `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Fee               *Money                                 `bson:"fee,omitempty" json:"fee,omitempty"`
	NoteNumber        []int                                  `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ExplanationOfBenefitItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail            []ExplanationOfBenefitAddItemDetail    `bson:"detail,omitempty" json:"detail,omitempty"`
}
type ExplanationOfBenefitAddItemDetail struct {
	Id                *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Revenue           *CodeableConcept                       `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept                       `bson:"category,omitempty" json:"category,omitempty"`
	Service           *CodeableConcept                       `bson:"service,omitempty" json:"service,omitempty"`
	Modifier          []CodeableConcept                      `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Fee               *Money                                 `bson:"fee,omitempty" json:"fee,omitempty"`
	NoteNumber        []int                                  `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ExplanationOfBenefitItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}
type ExplanationOfBenefitPayment struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Adjustment        *Money           `bson:"adjustment,omitempty" json:"adjustment,omitempty"`
	AdjustmentReason  *CodeableConcept `bson:"adjustmentReason,omitempty" json:"adjustmentReason,omitempty"`
	Date              *string          `bson:"date,omitempty" json:"date,omitempty"`
	Amount            *Money           `bson:"amount,omitempty" json:"amount,omitempty"`
	Identifier        *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
}
type ExplanationOfBenefitProcessNote struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Number            *int             `bson:"number,omitempty" json:"number,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Text              *string          `bson:"text,omitempty" json:"text,omitempty"`
	Language          *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
}
type ExplanationOfBenefitBenefitBalance struct {
	Id                *string                                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          CodeableConcept                               `bson:"category" json:"category"`
	SubCategory       *CodeableConcept                              `bson:"subCategory,omitempty" json:"subCategory,omitempty"`
	Excluded          *bool                                         `bson:"excluded,omitempty" json:"excluded,omitempty"`
	Name              *string                                       `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string                                       `bson:"description,omitempty" json:"description,omitempty"`
	Network           *CodeableConcept                              `bson:"network,omitempty" json:"network,omitempty"`
	Unit              *CodeableConcept                              `bson:"unit,omitempty" json:"unit,omitempty"`
	Term              *CodeableConcept                              `bson:"term,omitempty" json:"term,omitempty"`
	Financial         []ExplanationOfBenefitBenefitBalanceFinancial `bson:"financial,omitempty" json:"financial,omitempty"`
}
type ExplanationOfBenefitBenefitBalanceFinancial struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type" json:"type"`
}
type OtherExplanationOfBenefit ExplanationOfBenefit

// MarshalJSON marshals the given ExplanationOfBenefit as JSON into a byte slice
func (r ExplanationOfBenefit) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherExplanationOfBenefit
		ResourceType string `json:"resourceType"`
	}{
		OtherExplanationOfBenefit: OtherExplanationOfBenefit(r),
		ResourceType:              "ExplanationOfBenefit",
	})
}

// UnmarshalExplanationOfBenefit unmarshals a ExplanationOfBenefit.
func UnmarshalExplanationOfBenefit(b []byte) (ExplanationOfBenefit, error) {
	var explanationOfBenefit ExplanationOfBenefit
	if err := json.Unmarshal(b, &explanationOfBenefit); err != nil {
		return explanationOfBenefit, err
	}
	return explanationOfBenefit, nil
}
