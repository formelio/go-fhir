package fhir

import "encoding/json"

// Claim is documented here http://hl7.org/fhir/StructureDefinition/Claim
type Claim struct {
	Id                   *string            `bson:"id" json:"id"`
	Meta                 *Meta              `bson:"meta" json:"meta"`
	ImplicitRules        *string            `bson:"implicitRules" json:"implicitRules"`
	Language             *string            `bson:"language" json:"language"`
	Text                 *Narrative         `bson:"text" json:"text"`
	RawContained         []json.RawMessage  `bson:"contained" json:"contained"`
	Contained            []IResource        `bson:"-" json:"-"`
	Extension            []Extension        `bson:"extension" json:"extension"`
	ModifierExtension    []Extension        `bson:"modifierExtension" json:"modifierExtension"`
	Identifier           []Identifier       `bson:"identifier" json:"identifier"`
	Status               *string            `bson:"status" json:"status"`
	Type                 *CodeableConcept   `bson:"type" json:"type"`
	SubType              []CodeableConcept  `bson:"subType" json:"subType"`
	Use                  *Use               `bson:"use" json:"use"`
	Patient              *Reference         `bson:"patient" json:"patient"`
	BillablePeriod       *Period            `bson:"billablePeriod" json:"billablePeriod"`
	Created              *string            `bson:"created" json:"created"`
	Enterer              *Reference         `bson:"enterer" json:"enterer"`
	Insurer              *Reference         `bson:"insurer" json:"insurer"`
	Provider             *Reference         `bson:"provider" json:"provider"`
	Organization         *Reference         `bson:"organization" json:"organization"`
	Priority             *CodeableConcept   `bson:"priority" json:"priority"`
	FundsReserve         *CodeableConcept   `bson:"fundsReserve" json:"fundsReserve"`
	Related              []ClaimRelated     `bson:"related" json:"related"`
	Prescription         *Reference         `bson:"prescription" json:"prescription"`
	OriginalPrescription *Reference         `bson:"originalPrescription" json:"originalPrescription"`
	Payee                *ClaimPayee        `bson:"payee" json:"payee"`
	Referral             *Reference         `bson:"referral" json:"referral"`
	Facility             *Reference         `bson:"facility" json:"facility"`
	CareTeam             []ClaimCareTeam    `bson:"careTeam" json:"careTeam"`
	Information          []ClaimInformation `bson:"information" json:"information"`
	Diagnosis            []ClaimDiagnosis   `bson:"diagnosis" json:"diagnosis"`
	Procedure            []ClaimProcedure   `bson:"procedure" json:"procedure"`
	Insurance            []ClaimInsurance   `bson:"insurance" json:"insurance"`
	Accident             *ClaimAccident     `bson:"accident" json:"accident"`
	EmploymentImpacted   *Period            `bson:"employmentImpacted" json:"employmentImpacted"`
	Hospitalization      *Period            `bson:"hospitalization" json:"hospitalization"`
	Item                 []ClaimItem        `bson:"item" json:"item"`
	Total                *Money             `bson:"total" json:"total"`
}
type ClaimRelated struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Claim             *Reference       `bson:"claim" json:"claim"`
	Relationship      *CodeableConcept `bson:"relationship" json:"relationship"`
	Reference         *Identifier      `bson:"reference" json:"reference"`
}
type ClaimPayee struct {
	Id                *string         `bson:"id" json:"id"`
	Extension         []Extension     `bson:"extension" json:"extension"`
	ModifierExtension []Extension     `bson:"modifierExtension" json:"modifierExtension"`
	Type              CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	ResourceType      *Coding         `bson:"resourceType" json:"resourceType"`
	Party             *Reference      `bson:"party" json:"party"`
}
type ClaimCareTeam struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Sequence          int              `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Provider          Reference        `bson:"provider,omitempty" json:"provider,omitempty"`
	Responsible       *bool            `bson:"responsible" json:"responsible"`
	Role              *CodeableConcept `bson:"role" json:"role"`
	Qualification     *CodeableConcept `bson:"qualification" json:"qualification"`
}
type ClaimInformation struct {
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
	Reason            *CodeableConcept `bson:"reason" json:"reason"`
}
type ClaimDiagnosis struct {
	Id                       *string           `bson:"id" json:"id"`
	Extension                []Extension       `bson:"extension" json:"extension"`
	ModifierExtension        []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Sequence                 int               `bson:"sequence,omitempty" json:"sequence,omitempty"`
	DiagnosisCodeableConcept *CodeableConcept  `bson:"diagnosisCodeableConcept,omitempty" json:"diagnosisCodeableConcept,omitempty"`
	DiagnosisReference       *Reference        `bson:"diagnosisReference,omitempty" json:"diagnosisReference,omitempty"`
	Type                     []CodeableConcept `bson:"type" json:"type"`
	PackageCode              *CodeableConcept  `bson:"packageCode" json:"packageCode"`
}
type ClaimProcedure struct {
	Id                       *string          `bson:"id" json:"id"`
	Extension                []Extension      `bson:"extension" json:"extension"`
	ModifierExtension        []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Sequence                 int              `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Date                     *string          `bson:"date" json:"date"`
	ProcedureCodeableConcept *CodeableConcept `bson:"procedureCodeableConcept,omitempty" json:"procedureCodeableConcept,omitempty"`
	ProcedureReference       *Reference       `bson:"procedureReference,omitempty" json:"procedureReference,omitempty"`
}
type ClaimInsurance struct {
	Id                  *string     `bson:"id" json:"id"`
	Extension           []Extension `bson:"extension" json:"extension"`
	ModifierExtension   []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Sequence            int         `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Focal               bool        `bson:"focal,omitempty" json:"focal,omitempty"`
	Coverage            Reference   `bson:"coverage,omitempty" json:"coverage,omitempty"`
	BusinessArrangement *string     `bson:"businessArrangement" json:"businessArrangement"`
	PreAuthRef          []string    `bson:"preAuthRef" json:"preAuthRef"`
	ClaimResponse       *Reference  `bson:"claimResponse" json:"claimResponse"`
}
type ClaimAccident struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Date              string           `bson:"date,omitempty" json:"date,omitempty"`
	Type              *CodeableConcept `bson:"type" json:"type"`
	LocationAddress   *Address         `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference *Reference       `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
}
type ClaimItem struct {
	Id                      *string           `bson:"id" json:"id"`
	Extension               []Extension       `bson:"extension" json:"extension"`
	ModifierExtension       []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Sequence                int               `bson:"sequence,omitempty" json:"sequence,omitempty"`
	CareTeamLinkId          []int             `bson:"careTeamLinkId" json:"careTeamLinkId"`
	DiagnosisLinkId         []int             `bson:"diagnosisLinkId" json:"diagnosisLinkId"`
	ProcedureLinkId         []int             `bson:"procedureLinkId" json:"procedureLinkId"`
	InformationLinkId       []int             `bson:"informationLinkId" json:"informationLinkId"`
	Revenue                 *CodeableConcept  `bson:"revenue" json:"revenue"`
	Category                *CodeableConcept  `bson:"category" json:"category"`
	Service                 *CodeableConcept  `bson:"service" json:"service"`
	Modifier                []CodeableConcept `bson:"modifier" json:"modifier"`
	ProgramCode             []CodeableConcept `bson:"programCode" json:"programCode"`
	ServicedDate            *string           `bson:"servicedDate,omitempty" json:"servicedDate,omitempty"`
	ServicedPeriod          *Period           `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept  `bson:"locationCodeableConcept,omitempty" json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address          `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference       *Reference        `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
	Quantity                *Quantity         `bson:"quantity" json:"quantity"`
	UnitPrice               *Money            `bson:"unitPrice" json:"unitPrice"`
	Factor                  *float64          `bson:"factor" json:"factor"`
	Net                     *Money            `bson:"net" json:"net"`
	Udi                     []Reference       `bson:"udi" json:"udi"`
	BodySite                *CodeableConcept  `bson:"bodySite" json:"bodySite"`
	SubSite                 []CodeableConcept `bson:"subSite" json:"subSite"`
	Encounter               []Reference       `bson:"encounter" json:"encounter"`
	Detail                  []ClaimItemDetail `bson:"detail" json:"detail"`
}
type ClaimItemDetail struct {
	Id                *string                    `bson:"id" json:"id"`
	Extension         []Extension                `bson:"extension" json:"extension"`
	ModifierExtension []Extension                `bson:"modifierExtension" json:"modifierExtension"`
	Sequence          int                        `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Revenue           *CodeableConcept           `bson:"revenue" json:"revenue"`
	Category          *CodeableConcept           `bson:"category" json:"category"`
	Service           *CodeableConcept           `bson:"service" json:"service"`
	Modifier          []CodeableConcept          `bson:"modifier" json:"modifier"`
	ProgramCode       []CodeableConcept          `bson:"programCode" json:"programCode"`
	Quantity          *Quantity                  `bson:"quantity" json:"quantity"`
	UnitPrice         *Money                     `bson:"unitPrice" json:"unitPrice"`
	Factor            *float64                   `bson:"factor" json:"factor"`
	Net               *Money                     `bson:"net" json:"net"`
	Udi               []Reference                `bson:"udi" json:"udi"`
	SubDetail         []ClaimItemDetailSubDetail `bson:"subDetail" json:"subDetail"`
}
type ClaimItemDetailSubDetail struct {
	Id                *string           `bson:"id" json:"id"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Sequence          int               `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Revenue           *CodeableConcept  `bson:"revenue" json:"revenue"`
	Category          *CodeableConcept  `bson:"category" json:"category"`
	Service           *CodeableConcept  `bson:"service" json:"service"`
	Modifier          []CodeableConcept `bson:"modifier" json:"modifier"`
	ProgramCode       []CodeableConcept `bson:"programCode" json:"programCode"`
	Quantity          *Quantity         `bson:"quantity" json:"quantity"`
	UnitPrice         *Money            `bson:"unitPrice" json:"unitPrice"`
	Factor            *float64          `bson:"factor" json:"factor"`
	Net               *Money            `bson:"net" json:"net"`
	Udi               []Reference       `bson:"udi" json:"udi"`
}

// OtherClaim is a helper type to use the default implementations of Marshall and Unmarshal
type OtherClaim Claim

// MarshalJSON marshals the given Claim as JSON into a byte slice
func (r Claim) MarshalJSON() ([]byte, error) {
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
		OtherClaim
		ResourceType string `json:"resourceType"`
	}{
		OtherClaim:   OtherClaim(r),
		ResourceType: "Claim",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Claim
func (r *Claim) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherClaim)(r)); err != nil {
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
func (r Claim) GetResourceType() ResourceType {
	return ResourceTypeClaim
}
