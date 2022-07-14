package fhir

import "encoding/json"

// Claim is documented here http://hl7.org/fhir/StructureDefinition/Claim
type Claim struct {
	Id                   *string            `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string            `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               *string            `bson:"status,omitempty" json:"status,omitempty"`
	Type                 *CodeableConcept   `bson:"type,omitempty" json:"type,omitempty"`
	SubType              []CodeableConcept  `bson:"subType,omitempty" json:"subType,omitempty"`
	Use                  *string            `bson:"use,omitempty" json:"use,omitempty"`
	Patient              *Reference         `bson:"patient,omitempty" json:"patient,omitempty"`
	BillablePeriod       *Period            `bson:"billablePeriod,omitempty" json:"billablePeriod,omitempty"`
	Created              *string            `bson:"created,omitempty" json:"created,omitempty"`
	Enterer              *Reference         `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Insurer              *Reference         `bson:"insurer,omitempty" json:"insurer,omitempty"`
	Provider             *Reference         `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization         *Reference         `bson:"organization,omitempty" json:"organization,omitempty"`
	Priority             *CodeableConcept   `bson:"priority,omitempty" json:"priority,omitempty"`
	FundsReserve         *CodeableConcept   `bson:"fundsReserve,omitempty" json:"fundsReserve,omitempty"`
	Related              []ClaimRelated     `bson:"related,omitempty" json:"related,omitempty"`
	OriginalPrescription *Reference         `bson:"originalPrescription,omitempty" json:"originalPrescription,omitempty"`
	Payee                *ClaimPayee        `bson:"payee,omitempty" json:"payee,omitempty"`
	Referral             *Reference         `bson:"referral,omitempty" json:"referral,omitempty"`
	Facility             *Reference         `bson:"facility,omitempty" json:"facility,omitempty"`
	CareTeam             []ClaimCareTeam    `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	Information          []ClaimInformation `bson:"information,omitempty" json:"information,omitempty"`
	Diagnosis            []ClaimDiagnosis   `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Procedure            []ClaimProcedure   `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Insurance            []ClaimInsurance   `bson:"insurance,omitempty" json:"insurance,omitempty"`
	Accident             *ClaimAccident     `bson:"accident,omitempty" json:"accident,omitempty"`
	EmploymentImpacted   *Period            `bson:"employmentImpacted,omitempty" json:"employmentImpacted,omitempty"`
	Hospitalization      *Period            `bson:"hospitalization,omitempty" json:"hospitalization,omitempty"`
	Item                 []ClaimItem        `bson:"item,omitempty" json:"item,omitempty"`
	Total                *Money             `bson:"total,omitempty" json:"total,omitempty"`
}
type ClaimRelated struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Claim             *Reference       `bson:"claim,omitempty" json:"claim,omitempty"`
	Relationship      *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Reference         *Identifier      `bson:"reference,omitempty" json:"reference,omitempty"`
}
type ClaimPayee struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type" json:"type"`
	ResourceType      *Coding         `bson:"resourceType,omitempty" json:"resourceType,omitempty"`
}
type ClaimCareTeam struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int              `bson:"sequence" json:"sequence"`
	Responsible       *bool            `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Qualification     *CodeableConcept `bson:"qualification,omitempty" json:"qualification,omitempty"`
}
type ClaimInformation struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int              `bson:"sequence" json:"sequence"`
	Category          CodeableConcept  `bson:"category" json:"category"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Reason            *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
}
type ClaimDiagnosis struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int               `bson:"sequence" json:"sequence"`
	Type              []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	PackageCode       *CodeableConcept  `bson:"packageCode,omitempty" json:"packageCode,omitempty"`
}
type ClaimProcedure struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int         `bson:"sequence" json:"sequence"`
	Date              *string     `bson:"date,omitempty" json:"date,omitempty"`
}
type ClaimInsurance struct {
	Id                  *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence            int         `bson:"sequence" json:"sequence"`
	Focal               bool        `bson:"focal" json:"focal"`
	Coverage            Reference   `bson:"coverage" json:"coverage"`
	BusinessArrangement *string     `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
	PreAuthRef          []string    `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	ClaimResponse       *Reference  `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
}
type ClaimAccident struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Date              string           `bson:"date" json:"date"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
}
type ClaimItem struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int               `bson:"sequence" json:"sequence"`
	CareTeamLinkId    []int             `bson:"careTeamLinkId,omitempty" json:"careTeamLinkId,omitempty"`
	DiagnosisLinkId   []int             `bson:"diagnosisLinkId,omitempty" json:"diagnosisLinkId,omitempty"`
	ProcedureLinkId   []int             `bson:"procedureLinkId,omitempty" json:"procedureLinkId,omitempty"`
	InformationLinkId []int             `bson:"informationLinkId,omitempty" json:"informationLinkId,omitempty"`
	Revenue           *CodeableConcept  `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	Service           *CodeableConcept  `bson:"service,omitempty" json:"service,omitempty"`
	Modifier          []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity          *Quantity         `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money            `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *string           `bson:"factor,omitempty" json:"factor,omitempty"`
	Net               *Money            `bson:"net,omitempty" json:"net,omitempty"`
	Udi               []Reference       `bson:"udi,omitempty" json:"udi,omitempty"`
	BodySite          *CodeableConcept  `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SubSite           []CodeableConcept `bson:"subSite,omitempty" json:"subSite,omitempty"`
	Encounter         []Reference       `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Detail            []ClaimItemDetail `bson:"detail,omitempty" json:"detail,omitempty"`
}
type ClaimItemDetail struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int                        `bson:"sequence" json:"sequence"`
	Revenue           *CodeableConcept           `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept           `bson:"category,omitempty" json:"category,omitempty"`
	Service           *CodeableConcept           `bson:"service,omitempty" json:"service,omitempty"`
	Modifier          []CodeableConcept          `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept          `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity          *Quantity                  `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money                     `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *string                    `bson:"factor,omitempty" json:"factor,omitempty"`
	Net               *Money                     `bson:"net,omitempty" json:"net,omitempty"`
	Udi               []Reference                `bson:"udi,omitempty" json:"udi,omitempty"`
	SubDetail         []ClaimItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}
type ClaimItemDetailSubDetail struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int               `bson:"sequence" json:"sequence"`
	Revenue           *CodeableConcept  `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	Service           *CodeableConcept  `bson:"service,omitempty" json:"service,omitempty"`
	Modifier          []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode       []CodeableConcept `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity          *Quantity         `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money            `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *string           `bson:"factor,omitempty" json:"factor,omitempty"`
	Net               *Money            `bson:"net,omitempty" json:"net,omitempty"`
	Udi               []Reference       `bson:"udi,omitempty" json:"udi,omitempty"`
}
type OtherClaim Claim

// MarshalJSON marshals the given Claim as JSON into a byte slice
func (r Claim) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClaim
		ResourceType string `json:"resourceType"`
	}{
		OtherClaim:   OtherClaim(r),
		ResourceType: "Claim",
	})
}

// UnmarshalClaim unmarshals a Claim.
func UnmarshalClaim(b []byte) (Claim, error) {
	var claim Claim
	if err := json.Unmarshal(b, &claim); err != nil {
		return claim, err
	}
	return claim, nil
}
