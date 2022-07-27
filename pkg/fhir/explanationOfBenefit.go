package fhir

import "encoding/json"

// ExplanationOfBenefit is documented here http://hl7.org/fhir/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefit struct {
	Id                   *string                              `bson:"id" json:"id"`
	Meta                 *Meta                                `bson:"meta" json:"meta"`
	ImplicitRules        *string                              `bson:"implicitRules" json:"implicitRules"`
	Language             *string                              `bson:"language" json:"language"`
	Text                 *Narrative                           `bson:"text" json:"text"`
	RawContained         []json.RawMessage                    `bson:"contained" json:"contained"`
	Contained            []IResource                          `bson:"-" json:"-"`
	Extension            []Extension                          `bson:"extension" json:"extension"`
	ModifierExtension    []Extension                          `bson:"modifierExtension" json:"modifierExtension"`
	Identifier           []Identifier                         `bson:"identifier" json:"identifier"`
	Status               *ExplanationOfBenefitStatus          `bson:"status" json:"status"`
	Type                 *CodeableConcept                     `bson:"type" json:"type"`
	SubType              []CodeableConcept                    `bson:"subType" json:"subType"`
	Patient              *Reference                           `bson:"patient" json:"patient"`
	BillablePeriod       *Period                              `bson:"billablePeriod" json:"billablePeriod"`
	Created              *string                              `bson:"created" json:"created"`
	Enterer              *Reference                           `bson:"enterer" json:"enterer"`
	Insurer              *Reference                           `bson:"insurer" json:"insurer"`
	Provider             *Reference                           `bson:"provider" json:"provider"`
	Organization         *Reference                           `bson:"organization" json:"organization"`
	Referral             *Reference                           `bson:"referral" json:"referral"`
	Facility             *Reference                           `bson:"facility" json:"facility"`
	Claim                *Reference                           `bson:"claim" json:"claim"`
	ClaimResponse        *Reference                           `bson:"claimResponse" json:"claimResponse"`
	Outcome              *CodeableConcept                     `bson:"outcome" json:"outcome"`
	Disposition          *string                              `bson:"disposition" json:"disposition"`
	Related              []ExplanationOfBenefitRelated        `bson:"related" json:"related"`
	Prescription         *Reference                           `bson:"prescription" json:"prescription"`
	OriginalPrescription *Reference                           `bson:"originalPrescription" json:"originalPrescription"`
	Payee                *ExplanationOfBenefitPayee           `bson:"payee" json:"payee"`
	Information          []ExplanationOfBenefitInformation    `bson:"information" json:"information"`
	CareTeam             []ExplanationOfBenefitCareTeam       `bson:"careTeam" json:"careTeam"`
	Diagnosis            []ExplanationOfBenefitDiagnosis      `bson:"diagnosis" json:"diagnosis"`
	Procedure            []ExplanationOfBenefitProcedure      `bson:"procedure" json:"procedure"`
	Precedence           *int                                 `bson:"precedence" json:"precedence"`
	Insurance            *ExplanationOfBenefitInsurance       `bson:"insurance" json:"insurance"`
	Accident             *ExplanationOfBenefitAccident        `bson:"accident" json:"accident"`
	EmploymentImpacted   *Period                              `bson:"employmentImpacted" json:"employmentImpacted"`
	Hospitalization      *Period                              `bson:"hospitalization" json:"hospitalization"`
	Item                 []ExplanationOfBenefitItem           `bson:"item" json:"item"`
	AddItem              []ExplanationOfBenefitAddItem        `bson:"addItem" json:"addItem"`
	TotalCost            *Money                               `bson:"totalCost" json:"totalCost"`
	UnallocDeductable    *Money                               `bson:"unallocDeductable" json:"unallocDeductable"`
	TotalBenefit         *Money                               `bson:"totalBenefit" json:"totalBenefit"`
	Payment              *ExplanationOfBenefitPayment         `bson:"payment" json:"payment"`
	Form                 *CodeableConcept                     `bson:"form" json:"form"`
	ProcessNote          []ExplanationOfBenefitProcessNote    `bson:"processNote" json:"processNote"`
	BenefitBalance       []ExplanationOfBenefitBenefitBalance `bson:"benefitBalance" json:"benefitBalance"`
}
type ExplanationOfBenefitRelated struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Claim             *Reference       `bson:"claim" json:"claim"`
	Relationship      *CodeableConcept `bson:"relationship" json:"relationship"`
	Reference         *Identifier      `bson:"reference" json:"reference"`
}
type ExplanationOfBenefitPayee struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Type              *CodeableConcept `bson:"type" json:"type"`
	ResourceType      *CodeableConcept `bson:"resourceType" json:"resourceType"`
	Party             *Reference       `bson:"party" json:"party"`
}
type ExplanationOfBenefitInformation struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Sequence          int              `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Category          CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	Code              *CodeableConcept `bson:"code" json:"code"`
	TimingDate        *string          `bson:"timingDate,omitempty" json:"timingDate,omitempty"`
	TimingPeriod      *Period          `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	ValueString       *string          `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueQuantity     *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueAttachment   *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueReference    *Reference       `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	Reason            *Coding          `bson:"reason" json:"reason"`
}
type ExplanationOfBenefitCareTeam struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Sequence          int              `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Provider          Reference        `bson:"provider,omitempty" json:"provider,omitempty"`
	Responsible       *bool            `bson:"responsible" json:"responsible"`
	Role              *CodeableConcept `bson:"role" json:"role"`
	Qualification     *CodeableConcept `bson:"qualification" json:"qualification"`
}
type ExplanationOfBenefitDiagnosis struct {
	Id                       *string           `bson:"id" json:"id"`
	Extension                []Extension       `bson:"extension" json:"extension"`
	ModifierExtension        []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Sequence                 int               `bson:"sequence,omitempty" json:"sequence,omitempty"`
	DiagnosisCodeableConcept *CodeableConcept  `bson:"diagnosisCodeableConcept,omitempty" json:"diagnosisCodeableConcept,omitempty"`
	DiagnosisReference       *Reference        `bson:"diagnosisReference,omitempty" json:"diagnosisReference,omitempty"`
	Type                     []CodeableConcept `bson:"type" json:"type"`
	PackageCode              *CodeableConcept  `bson:"packageCode" json:"packageCode"`
}
type ExplanationOfBenefitProcedure struct {
	Id                       *string          `bson:"id" json:"id"`
	Extension                []Extension      `bson:"extension" json:"extension"`
	ModifierExtension        []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Sequence                 int              `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Date                     *string          `bson:"date" json:"date"`
	ProcedureCodeableConcept *CodeableConcept `bson:"procedureCodeableConcept,omitempty" json:"procedureCodeableConcept,omitempty"`
	ProcedureReference       *Reference       `bson:"procedureReference,omitempty" json:"procedureReference,omitempty"`
}
type ExplanationOfBenefitInsurance struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Coverage          *Reference  `bson:"coverage" json:"coverage"`
	PreAuthRef        []string    `bson:"preAuthRef" json:"preAuthRef"`
}
type ExplanationOfBenefitAccident struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Date              *string          `bson:"date" json:"date"`
	Type              *CodeableConcept `bson:"type" json:"type"`
	LocationAddress   *Address         `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference *Reference       `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
}
type ExplanationOfBenefitItem struct {
	Id                      *string                                `bson:"id" json:"id"`
	Extension               []Extension                            `bson:"extension" json:"extension"`
	ModifierExtension       []Extension                            `bson:"modifierExtension" json:"modifierExtension"`
	Sequence                int                                    `bson:"sequence,omitempty" json:"sequence,omitempty"`
	CareTeamLinkId          []int                                  `bson:"careTeamLinkId" json:"careTeamLinkId"`
	DiagnosisLinkId         []int                                  `bson:"diagnosisLinkId" json:"diagnosisLinkId"`
	ProcedureLinkId         []int                                  `bson:"procedureLinkId" json:"procedureLinkId"`
	InformationLinkId       []int                                  `bson:"informationLinkId" json:"informationLinkId"`
	Revenue                 *CodeableConcept                       `bson:"revenue" json:"revenue"`
	Category                *CodeableConcept                       `bson:"category" json:"category"`
	Service                 *CodeableConcept                       `bson:"service" json:"service"`
	Modifier                []CodeableConcept                      `bson:"modifier" json:"modifier"`
	ProgramCode             []CodeableConcept                      `bson:"programCode" json:"programCode"`
	ServicedDate            *string                                `bson:"servicedDate,omitempty" json:"servicedDate,omitempty"`
	ServicedPeriod          *Period                                `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept                       `bson:"locationCodeableConcept,omitempty" json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address                               `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference       *Reference                             `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
	Quantity                *Quantity                              `bson:"quantity" json:"quantity"`
	UnitPrice               *Money                                 `bson:"unitPrice" json:"unitPrice"`
	Factor                  *float64                               `bson:"factor" json:"factor"`
	Net                     *Money                                 `bson:"net" json:"net"`
	Udi                     []Reference                            `bson:"udi" json:"udi"`
	BodySite                *CodeableConcept                       `bson:"bodySite" json:"bodySite"`
	SubSite                 []CodeableConcept                      `bson:"subSite" json:"subSite"`
	Encounter               []Reference                            `bson:"encounter" json:"encounter"`
	NoteNumber              []int                                  `bson:"noteNumber" json:"noteNumber"`
	Adjudication            []ExplanationOfBenefitItemAdjudication `bson:"adjudication" json:"adjudication"`
	Detail                  []ExplanationOfBenefitItemDetail       `bson:"detail" json:"detail"`
}
type ExplanationOfBenefitItemAdjudication struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Category          CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	Reason            *CodeableConcept `bson:"reason" json:"reason"`
	Amount            *Money           `bson:"amount" json:"amount"`
	Value             *float64         `bson:"value" json:"value"`
}
type ExplanationOfBenefitItemDetail struct {
	Id                *string                                   `bson:"id" json:"id"`
	Extension         []Extension                               `bson:"extension" json:"extension"`
	ModifierExtension []Extension                               `bson:"modifierExtension" json:"modifierExtension"`
	Sequence          int                                       `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type              CodeableConcept                           `bson:"type,omitempty" json:"type,omitempty"`
	Revenue           *CodeableConcept                          `bson:"revenue" json:"revenue"`
	Category          *CodeableConcept                          `bson:"category" json:"category"`
	Service           *CodeableConcept                          `bson:"service" json:"service"`
	Modifier          []CodeableConcept                         `bson:"modifier" json:"modifier"`
	ProgramCode       []CodeableConcept                         `bson:"programCode" json:"programCode"`
	Quantity          *Quantity                                 `bson:"quantity" json:"quantity"`
	UnitPrice         *Money                                    `bson:"unitPrice" json:"unitPrice"`
	Factor            *float64                                  `bson:"factor" json:"factor"`
	Net               *Money                                    `bson:"net" json:"net"`
	Udi               []Reference                               `bson:"udi" json:"udi"`
	NoteNumber        []int                                     `bson:"noteNumber" json:"noteNumber"`
	Adjudication      []ExplanationOfBenefitItemAdjudication    `bson:"adjudication" json:"adjudication"`
	SubDetail         []ExplanationOfBenefitItemDetailSubDetail `bson:"subDetail" json:"subDetail"`
}
type ExplanationOfBenefitItemDetailSubDetail struct {
	Id                *string                                `bson:"id" json:"id"`
	Extension         []Extension                            `bson:"extension" json:"extension"`
	ModifierExtension []Extension                            `bson:"modifierExtension" json:"modifierExtension"`
	Sequence          int                                    `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type              CodeableConcept                        `bson:"type,omitempty" json:"type,omitempty"`
	Revenue           *CodeableConcept                       `bson:"revenue" json:"revenue"`
	Category          *CodeableConcept                       `bson:"category" json:"category"`
	Service           *CodeableConcept                       `bson:"service" json:"service"`
	Modifier          []CodeableConcept                      `bson:"modifier" json:"modifier"`
	ProgramCode       []CodeableConcept                      `bson:"programCode" json:"programCode"`
	Quantity          *Quantity                              `bson:"quantity" json:"quantity"`
	UnitPrice         *Money                                 `bson:"unitPrice" json:"unitPrice"`
	Factor            *float64                               `bson:"factor" json:"factor"`
	Net               *Money                                 `bson:"net" json:"net"`
	Udi               []Reference                            `bson:"udi" json:"udi"`
	NoteNumber        []int                                  `bson:"noteNumber" json:"noteNumber"`
	Adjudication      []ExplanationOfBenefitItemAdjudication `bson:"adjudication" json:"adjudication"`
}
type ExplanationOfBenefitAddItem struct {
	Id                *string                                `bson:"id" json:"id"`
	Extension         []Extension                            `bson:"extension" json:"extension"`
	ModifierExtension []Extension                            `bson:"modifierExtension" json:"modifierExtension"`
	SequenceLinkId    []int                                  `bson:"sequenceLinkId" json:"sequenceLinkId"`
	Revenue           *CodeableConcept                       `bson:"revenue" json:"revenue"`
	Category          *CodeableConcept                       `bson:"category" json:"category"`
	Service           *CodeableConcept                       `bson:"service" json:"service"`
	Modifier          []CodeableConcept                      `bson:"modifier" json:"modifier"`
	Fee               *Money                                 `bson:"fee" json:"fee"`
	NoteNumber        []int                                  `bson:"noteNumber" json:"noteNumber"`
	Adjudication      []ExplanationOfBenefitItemAdjudication `bson:"adjudication" json:"adjudication"`
	Detail            []ExplanationOfBenefitAddItemDetail    `bson:"detail" json:"detail"`
}
type ExplanationOfBenefitAddItemDetail struct {
	Id                *string                                `bson:"id" json:"id"`
	Extension         []Extension                            `bson:"extension" json:"extension"`
	ModifierExtension []Extension                            `bson:"modifierExtension" json:"modifierExtension"`
	Revenue           *CodeableConcept                       `bson:"revenue" json:"revenue"`
	Category          *CodeableConcept                       `bson:"category" json:"category"`
	Service           *CodeableConcept                       `bson:"service" json:"service"`
	Modifier          []CodeableConcept                      `bson:"modifier" json:"modifier"`
	Fee               *Money                                 `bson:"fee" json:"fee"`
	NoteNumber        []int                                  `bson:"noteNumber" json:"noteNumber"`
	Adjudication      []ExplanationOfBenefitItemAdjudication `bson:"adjudication" json:"adjudication"`
}
type ExplanationOfBenefitPayment struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Type              *CodeableConcept `bson:"type" json:"type"`
	Adjustment        *Money           `bson:"adjustment" json:"adjustment"`
	AdjustmentReason  *CodeableConcept `bson:"adjustmentReason" json:"adjustmentReason"`
	Date              *string          `bson:"date" json:"date"`
	Amount            *Money           `bson:"amount" json:"amount"`
	Identifier        *Identifier      `bson:"identifier" json:"identifier"`
}
type ExplanationOfBenefitProcessNote struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Number            *int             `bson:"number" json:"number"`
	Type              *CodeableConcept `bson:"type" json:"type"`
	Text              *string          `bson:"text" json:"text"`
	Language          *CodeableConcept `bson:"language" json:"language"`
}
type ExplanationOfBenefitBenefitBalance struct {
	Id                *string                                       `bson:"id" json:"id"`
	Extension         []Extension                                   `bson:"extension" json:"extension"`
	ModifierExtension []Extension                                   `bson:"modifierExtension" json:"modifierExtension"`
	Category          CodeableConcept                               `bson:"category,omitempty" json:"category,omitempty"`
	SubCategory       *CodeableConcept                              `bson:"subCategory" json:"subCategory"`
	Excluded          *bool                                         `bson:"excluded" json:"excluded"`
	Name              *string                                       `bson:"name" json:"name"`
	Description       *string                                       `bson:"description" json:"description"`
	Network           *CodeableConcept                              `bson:"network" json:"network"`
	Unit              *CodeableConcept                              `bson:"unit" json:"unit"`
	Term              *CodeableConcept                              `bson:"term" json:"term"`
	Financial         []ExplanationOfBenefitBenefitBalanceFinancial `bson:"financial" json:"financial"`
}
type ExplanationOfBenefitBenefitBalanceFinancial struct {
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
	return json.Marshal(struct {
		OtherExplanationOfBenefit
		ResourceType string `json:"resourceType"`
	}{
		OtherExplanationOfBenefit: OtherExplanationOfBenefit(r),
		ResourceType:              "ExplanationOfBenefit",
	})
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
