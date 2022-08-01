package fhir

import (
	"bytes"
	"encoding/json"
)

// ExplanationOfBenefit is documented here http://hl7.org/fhir/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefit struct {
	Id                   *string                               `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                               `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative                            `bson:"text,omitempty" json:"text,omitempty"`
	RawContained         []json.RawMessage                     `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained            []IResource                           `bson:"-,omitempty" json:"-,omitempty"`
	Extension            []*Extension                          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []*Extension                          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []*Identifier                         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               *ExplanationOfBenefitStatus           `bson:"status,omitempty" json:"status,omitempty"`
	Type                 *CodeableConcept                      `bson:"type,omitempty" json:"type,omitempty"`
	SubType              []*CodeableConcept                    `bson:"subType,omitempty" json:"subType,omitempty"`
	Patient              *Reference                            `bson:"patient,omitempty" json:"patient,omitempty"`
	BillablePeriod       *Period                               `bson:"billablePeriod,omitempty" json:"billablePeriod,omitempty"`
	Created              *string                               `bson:"created,omitempty" json:"created,omitempty"`
	Enterer              *Reference                            `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Insurer              *Reference                            `bson:"insurer,omitempty" json:"insurer,omitempty"`
	Provider             *Reference                            `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization         *Reference                            `bson:"organization,omitempty" json:"organization,omitempty"`
	Referral             *Reference                            `bson:"referral,omitempty" json:"referral,omitempty"`
	Facility             *Reference                            `bson:"facility,omitempty" json:"facility,omitempty"`
	Claim                *Reference                            `bson:"claim,omitempty" json:"claim,omitempty"`
	ClaimResponse        *Reference                            `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
	Outcome              *CodeableConcept                      `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition          *string                               `bson:"disposition,omitempty" json:"disposition,omitempty"`
	Related              []*ExplanationOfBenefitRelated        `bson:"related,omitempty" json:"related,omitempty"`
	Prescription         *Reference                            `bson:"prescription,omitempty" json:"prescription,omitempty"`
	OriginalPrescription *Reference                            `bson:"originalPrescription,omitempty" json:"originalPrescription,omitempty"`
	Payee                *ExplanationOfBenefitPayee            `bson:"payee,omitempty" json:"payee,omitempty"`
	Information          []*ExplanationOfBenefitInformation    `bson:"information,omitempty" json:"information,omitempty"`
	CareTeam             []*ExplanationOfBenefitCareTeam       `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	Diagnosis            []*ExplanationOfBenefitDiagnosis      `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Procedure            []*ExplanationOfBenefitProcedure      `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Precedence           *int                                  `bson:"precedence,omitempty" json:"precedence,omitempty"`
	Insurance            *ExplanationOfBenefitInsurance        `bson:"insurance,omitempty" json:"insurance,omitempty"`
	Accident             *ExplanationOfBenefitAccident         `bson:"accident,omitempty" json:"accident,omitempty"`
	EmploymentImpacted   *Period                               `bson:"employmentImpacted,omitempty" json:"employmentImpacted,omitempty"`
	Hospitalization      *Period                               `bson:"hospitalization,omitempty" json:"hospitalization,omitempty"`
	Item                 []*ExplanationOfBenefitItem           `bson:"item,omitempty" json:"item,omitempty"`
	AddItem              []*ExplanationOfBenefitAddItem        `bson:"addItem,omitempty" json:"addItem,omitempty"`
	TotalCost            *Money                                `bson:"totalCost,omitempty" json:"totalCost,omitempty"`
	UnallocDeductable    *Money                                `bson:"unallocDeductable,omitempty" json:"unallocDeductable,omitempty"`
	TotalBenefit         *Money                                `bson:"totalBenefit,omitempty" json:"totalBenefit,omitempty"`
	Payment              *ExplanationOfBenefitPayment          `bson:"payment,omitempty" json:"payment,omitempty"`
	Form                 *CodeableConcept                      `bson:"form,omitempty" json:"form,omitempty"`
	ProcessNote          []*ExplanationOfBenefitProcessNote    `bson:"processNote,omitempty" json:"processNote,omitempty"`
	BenefitBalance       []*ExplanationOfBenefitBenefitBalance `bson:"benefitBalance,omitempty" json:"benefitBalance,omitempty"`
}
type ExplanationOfBenefitRelated struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Claim             *Reference       `bson:"claim,omitempty" json:"claim,omitempty"`
	Relationship      *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Reference         *Identifier      `bson:"reference,omitempty" json:"reference,omitempty"`
}
type ExplanationOfBenefitPayee struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	ResourceType      *CodeableConcept `bson:"resourceType,omitempty" json:"resourceType,omitempty"`
	Party             *Reference       `bson:"party,omitempty" json:"party,omitempty"`
}
type ExplanationOfBenefitInformation struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int              `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Category          CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	TimingDate        *string          `bson:"timingDate,omitempty" json:"timingDate,omitempty"`
	TimingPeriod      *Period          `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	ValueString       *string          `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueQuantity     *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueAttachment   *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueReference    *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	Reason            *Coding          `bson:"reason,omitempty" json:"reason,omitempty"`
}
type ExplanationOfBenefitCareTeam struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int              `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Provider          Reference        `bson:"provider,omitempty" json:"provider,omitempty"`
	Responsible       *bool            `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Qualification     *CodeableConcept `bson:"qualification,omitempty" json:"qualification,omitempty"`
}
type ExplanationOfBenefitDiagnosis struct {
	Id                       *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension                []*Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension        []*Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence                 int                `bson:"sequence,omitempty" json:"sequence,omitempty"`
	DiagnosisCodeableConcept *CodeableConcept   `bson:"diagnosisCodeableConcept,omitempty" json:"diagnosisCodeableConcept,omitempty"`
	DiagnosisReference       *Reference         `bson:"diagnosisReference,omitempty" json:"diagnosisReference,omitempty"`
	Type                     []*CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	PackageCode              *CodeableConcept   `bson:"packageCode,omitempty" json:"packageCode,omitempty"`
}
type ExplanationOfBenefitProcedure struct {
	Id                       *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension                []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension        []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence                 int              `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Date                     *string          `bson:"date,omitempty" json:"date,omitempty"`
	ProcedureCodeableConcept *CodeableConcept `bson:"procedureCodeableConcept,omitempty" json:"procedureCodeableConcept,omitempty"`
	ProcedureReference       *Reference       `bson:"procedureReference,omitempty" json:"procedureReference,omitempty"`
}
type ExplanationOfBenefitInsurance struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Coverage          *Reference   `bson:"coverage,omitempty" json:"coverage,omitempty"`
	PreAuthRef        []*string    `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
}
type ExplanationOfBenefitAccident struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Date              *string          `bson:"date,omitempty" json:"date,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	LocationAddress   *Address         `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference *Reference       `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
}
type ExplanationOfBenefitItem struct {
	Id                      *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension               []*Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []*Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence                int                                     `bson:"sequence,omitempty" json:"sequence,omitempty"`
	CareTeamLinkId          []*int                                  `bson:"careTeamLinkId,omitempty" json:"careTeamLinkId,omitempty"`
	DiagnosisLinkId         []*int                                  `bson:"diagnosisLinkId,omitempty" json:"diagnosisLinkId,omitempty"`
	ProcedureLinkId         []*int                                  `bson:"procedureLinkId,omitempty" json:"procedureLinkId,omitempty"`
	InformationLinkId       []*int                                  `bson:"informationLinkId,omitempty" json:"informationLinkId,omitempty"`
	Revenue                 *CodeableConcept                        `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category                *CodeableConcept                        `bson:"category,omitempty" json:"category,omitempty"`
	Service                 *CodeableConcept                        `bson:"service,omitempty" json:"service,omitempty"`
	Modifier                []*CodeableConcept                      `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode             []*CodeableConcept                      `bson:"programCode,omitempty" json:"programCode,omitempty"`
	ServicedDate            *string                                 `bson:"servicedDate,omitempty" json:"servicedDate,omitempty"`
	ServicedPeriod          *Period                                 `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept                        `bson:"locationCodeableConcept,omitempty" json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address                                `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference       *Reference                              `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
	Quantity                *Quantity                               `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice               *Money                                  `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor                  *float64                                `bson:"factor,omitempty" json:"factor,omitempty"`
	Net                     *Money                                  `bson:"net,omitempty" json:"net,omitempty"`
	Udi                     []*Reference                            `bson:"udi,omitempty" json:"udi,omitempty"`
	BodySite                *CodeableConcept                        `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SubSite                 []*CodeableConcept                      `bson:"subSite,omitempty" json:"subSite,omitempty"`
	Encounter               []*Reference                            `bson:"encounter,omitempty" json:"encounter,omitempty"`
	NoteNumber              []*int                                  `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication            []*ExplanationOfBenefitItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail                  []*ExplanationOfBenefitItemDetail       `bson:"detail,omitempty" json:"detail,omitempty"`
}
type ExplanationOfBenefitItemAdjudication struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	Reason            *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	Amount            *Money           `bson:"amount,omitempty" json:"amount,omitempty"`
	Value             *float64         `bson:"value,omitempty" json:"value,omitempty"`
}
type ExplanationOfBenefitItemDetail struct {
	Id                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int                                        `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type              CodeableConcept                            `bson:"type,omitempty" json:"type,omitempty"`
	Revenue           *CodeableConcept                           `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept                           `bson:"category,omitempty" json:"category,omitempty"`
	Service           *CodeableConcept                           `bson:"service,omitempty" json:"service,omitempty"`
	Modifier          []*CodeableConcept                         `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode       []*CodeableConcept                         `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity          *Quantity                                  `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money                                     `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *float64                                   `bson:"factor,omitempty" json:"factor,omitempty"`
	Net               *Money                                     `bson:"net,omitempty" json:"net,omitempty"`
	Udi               []*Reference                               `bson:"udi,omitempty" json:"udi,omitempty"`
	NoteNumber        []*int                                     `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []*ExplanationOfBenefitItemAdjudication    `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	SubDetail         []*ExplanationOfBenefitItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}
type ExplanationOfBenefitItemDetailSubDetail struct {
	Id                *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int                                     `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type              CodeableConcept                         `bson:"type,omitempty" json:"type,omitempty"`
	Revenue           *CodeableConcept                        `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept                        `bson:"category,omitempty" json:"category,omitempty"`
	Service           *CodeableConcept                        `bson:"service,omitempty" json:"service,omitempty"`
	Modifier          []*CodeableConcept                      `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode       []*CodeableConcept                      `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity          *Quantity                               `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money                                  `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *float64                                `bson:"factor,omitempty" json:"factor,omitempty"`
	Net               *Money                                  `bson:"net,omitempty" json:"net,omitempty"`
	Udi               []*Reference                            `bson:"udi,omitempty" json:"udi,omitempty"`
	NoteNumber        []*int                                  `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []*ExplanationOfBenefitItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}
type ExplanationOfBenefitAddItem struct {
	Id                *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SequenceLinkId    []*int                                  `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	Revenue           *CodeableConcept                        `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept                        `bson:"category,omitempty" json:"category,omitempty"`
	Service           *CodeableConcept                        `bson:"service,omitempty" json:"service,omitempty"`
	Modifier          []*CodeableConcept                      `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Fee               *Money                                  `bson:"fee,omitempty" json:"fee,omitempty"`
	NoteNumber        []*int                                  `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []*ExplanationOfBenefitItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail            []*ExplanationOfBenefitAddItemDetail    `bson:"detail,omitempty" json:"detail,omitempty"`
}
type ExplanationOfBenefitAddItemDetail struct {
	Id                *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Revenue           *CodeableConcept                        `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept                        `bson:"category,omitempty" json:"category,omitempty"`
	Service           *CodeableConcept                        `bson:"service,omitempty" json:"service,omitempty"`
	Modifier          []*CodeableConcept                      `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Fee               *Money                                  `bson:"fee,omitempty" json:"fee,omitempty"`
	NoteNumber        []*int                                  `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []*ExplanationOfBenefitItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}
type ExplanationOfBenefitPayment struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Adjustment        *Money           `bson:"adjustment,omitempty" json:"adjustment,omitempty"`
	AdjustmentReason  *CodeableConcept `bson:"adjustmentReason,omitempty" json:"adjustmentReason,omitempty"`
	Date              *string          `bson:"date,omitempty" json:"date,omitempty"`
	Amount            *Money           `bson:"amount,omitempty" json:"amount,omitempty"`
	Identifier        *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
}
type ExplanationOfBenefitProcessNote struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Number            *int             `bson:"number,omitempty" json:"number,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Text              *string          `bson:"text,omitempty" json:"text,omitempty"`
	Language          *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
}
type ExplanationOfBenefitBenefitBalance struct {
	Id                *string                                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          CodeableConcept                                `bson:"category,omitempty" json:"category,omitempty"`
	SubCategory       *CodeableConcept                               `bson:"subCategory,omitempty" json:"subCategory,omitempty"`
	Excluded          *bool                                          `bson:"excluded,omitempty" json:"excluded,omitempty"`
	Name              *string                                        `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string                                        `bson:"description,omitempty" json:"description,omitempty"`
	Network           *CodeableConcept                               `bson:"network,omitempty" json:"network,omitempty"`
	Unit              *CodeableConcept                               `bson:"unit,omitempty" json:"unit,omitempty"`
	Term              *CodeableConcept                               `bson:"term,omitempty" json:"term,omitempty"`
	Financial         []*ExplanationOfBenefitBenefitBalanceFinancial `bson:"financial,omitempty" json:"financial,omitempty"`
}
type ExplanationOfBenefitBenefitBalanceFinancial struct {
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

// OtherExplanationOfBenefit is a helper type to use the default implementations of Marshall and Unmarshal
type OtherExplanationOfBenefit ExplanationOfBenefit

// MarshalJSON marshals the given ExplanationOfBenefit as JSON into a byte slice
func (r ExplanationOfBenefit) MarshalJSON() ([]byte, error) {
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
		OtherExplanationOfBenefit
	}{
		OtherExplanationOfBenefit: OtherExplanationOfBenefit(r),
		ResourceType:              "ExplanationOfBenefit",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into ExplanationOfBenefit
func (r *ExplanationOfBenefit) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherExplanationOfBenefit)(r)); err != nil {
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
func (r ExplanationOfBenefit) GetResourceType() ResourceType {
	return ResourceTypeExplanationOfBenefit
}
