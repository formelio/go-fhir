package fhir

import (
	"bytes"
	"encoding/json"
)

// Claim is documented here http://hl7.org/fhir/StructureDefinition/Claim
type Claim struct {
	Id                   *string             `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string             `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative          `bson:"text,omitempty" json:"text,omitempty"`
	RawContained         []json.RawMessage   `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained            []IResource         `bson:"-,omitempty" json:"-,omitempty"`
	Extension            []*Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []*Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []*Identifier       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               *string             `bson:"status,omitempty" json:"status,omitempty"`
	Type                 *CodeableConcept    `bson:"type,omitempty" json:"type,omitempty"`
	SubType              []*CodeableConcept  `bson:"subType,omitempty" json:"subType,omitempty"`
	Use                  *Use                `bson:"use,omitempty" json:"use,omitempty"`
	Patient              *Reference          `bson:"patient,omitempty" json:"patient,omitempty"`
	BillablePeriod       *Period             `bson:"billablePeriod,omitempty" json:"billablePeriod,omitempty"`
	Created              *string             `bson:"created,omitempty" json:"created,omitempty"`
	Enterer              *Reference          `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Insurer              *Reference          `bson:"insurer,omitempty" json:"insurer,omitempty"`
	Provider             *Reference          `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization         *Reference          `bson:"organization,omitempty" json:"organization,omitempty"`
	Priority             *CodeableConcept    `bson:"priority,omitempty" json:"priority,omitempty"`
	FundsReserve         *CodeableConcept    `bson:"fundsReserve,omitempty" json:"fundsReserve,omitempty"`
	Related              []*ClaimRelated     `bson:"related,omitempty" json:"related,omitempty"`
	Prescription         *Reference          `bson:"prescription,omitempty" json:"prescription,omitempty"`
	OriginalPrescription *Reference          `bson:"originalPrescription,omitempty" json:"originalPrescription,omitempty"`
	Payee                *ClaimPayee         `bson:"payee,omitempty" json:"payee,omitempty"`
	Referral             *Reference          `bson:"referral,omitempty" json:"referral,omitempty"`
	Facility             *Reference          `bson:"facility,omitempty" json:"facility,omitempty"`
	CareTeam             []*ClaimCareTeam    `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	Information          []*ClaimInformation `bson:"information,omitempty" json:"information,omitempty"`
	Diagnosis            []*ClaimDiagnosis   `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Procedure            []*ClaimProcedure   `bson:"procedure,omitempty" json:"procedure,omitempty"`
	Insurance            []*ClaimInsurance   `bson:"insurance,omitempty" json:"insurance,omitempty"`
	Accident             *ClaimAccident      `bson:"accident,omitempty" json:"accident,omitempty"`
	EmploymentImpacted   *Period             `bson:"employmentImpacted,omitempty" json:"employmentImpacted,omitempty"`
	Hospitalization      *Period             `bson:"hospitalization,omitempty" json:"hospitalization,omitempty"`
	Item                 []*ClaimItem        `bson:"item,omitempty" json:"item,omitempty"`
	Total                *Money              `bson:"total,omitempty" json:"total,omitempty"`
}
type ClaimRelated struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Claim             *Reference       `bson:"claim,omitempty" json:"claim,omitempty"`
	Relationship      *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Reference         *Identifier      `bson:"reference,omitempty" json:"reference,omitempty"`
}
type ClaimPayee struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	ResourceType      *Coding         `bson:"resourceType,omitempty" json:"resourceType,omitempty"`
	Party             *Reference      `bson:"party,omitempty" json:"party,omitempty"`
}
type ClaimCareTeam struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int              `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Provider          Reference        `bson:"provider,omitempty" json:"provider,omitempty"`
	Responsible       *bool            `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Qualification     *CodeableConcept `bson:"qualification,omitempty" json:"qualification,omitempty"`
}
type ClaimInformation struct {
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
	Reason            *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
}
type ClaimDiagnosis struct {
	Id                       *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension                []*Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension        []*Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence                 int                `bson:"sequence,omitempty" json:"sequence,omitempty"`
	DiagnosisCodeableConcept *CodeableConcept   `bson:"diagnosisCodeableConcept,omitempty" json:"diagnosisCodeableConcept,omitempty"`
	DiagnosisReference       *Reference         `bson:"diagnosisReference,omitempty" json:"diagnosisReference,omitempty"`
	Type                     []*CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	PackageCode              *CodeableConcept   `bson:"packageCode,omitempty" json:"packageCode,omitempty"`
}
type ClaimProcedure struct {
	Id                       *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension                []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension        []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence                 int              `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Date                     *string          `bson:"date,omitempty" json:"date,omitempty"`
	ProcedureCodeableConcept *CodeableConcept `bson:"procedureCodeableConcept,omitempty" json:"procedureCodeableConcept,omitempty"`
	ProcedureReference       *Reference       `bson:"procedureReference,omitempty" json:"procedureReference,omitempty"`
}
type ClaimInsurance struct {
	Id                  *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence            int          `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Focal               bool         `bson:"focal,omitempty" json:"focal,omitempty"`
	Coverage            Reference    `bson:"coverage,omitempty" json:"coverage,omitempty"`
	BusinessArrangement *string      `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
	PreAuthRef          []*string    `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	ClaimResponse       *Reference   `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
}
type ClaimAccident struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Date              string           `bson:"date,omitempty" json:"date,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	LocationAddress   *Address         `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference *Reference       `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
}
type ClaimItem struct {
	Id                      *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension               []*Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []*Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence                int                `bson:"sequence,omitempty" json:"sequence,omitempty"`
	CareTeamLinkId          []*int             `bson:"careTeamLinkId,omitempty" json:"careTeamLinkId,omitempty"`
	DiagnosisLinkId         []*int             `bson:"diagnosisLinkId,omitempty" json:"diagnosisLinkId,omitempty"`
	ProcedureLinkId         []*int             `bson:"procedureLinkId,omitempty" json:"procedureLinkId,omitempty"`
	InformationLinkId       []*int             `bson:"informationLinkId,omitempty" json:"informationLinkId,omitempty"`
	Revenue                 *CodeableConcept   `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category                *CodeableConcept   `bson:"category,omitempty" json:"category,omitempty"`
	Service                 *CodeableConcept   `bson:"service,omitempty" json:"service,omitempty"`
	Modifier                []*CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode             []*CodeableConcept `bson:"programCode,omitempty" json:"programCode,omitempty"`
	ServicedDate            *string            `bson:"servicedDate,omitempty" json:"servicedDate,omitempty"`
	ServicedPeriod          *Period            `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty"`
	LocationCodeableConcept *CodeableConcept   `bson:"locationCodeableConcept,omitempty" json:"locationCodeableConcept,omitempty"`
	LocationAddress         *Address           `bson:"locationAddress,omitempty" json:"locationAddress,omitempty"`
	LocationReference       *Reference         `bson:"locationReference,omitempty" json:"locationReference,omitempty"`
	Quantity                *Quantity          `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice               *Money             `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor                  *float64           `bson:"factor,omitempty" json:"factor,omitempty"`
	Net                     *Money             `bson:"net,omitempty" json:"net,omitempty"`
	Udi                     []*Reference       `bson:"udi,omitempty" json:"udi,omitempty"`
	BodySite                *CodeableConcept   `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SubSite                 []*CodeableConcept `bson:"subSite,omitempty" json:"subSite,omitempty"`
	Encounter               []*Reference       `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Detail                  []*ClaimItemDetail `bson:"detail,omitempty" json:"detail,omitempty"`
}
type ClaimItemDetail struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int                         `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Revenue           *CodeableConcept            `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept            `bson:"category,omitempty" json:"category,omitempty"`
	Service           *CodeableConcept            `bson:"service,omitempty" json:"service,omitempty"`
	Modifier          []*CodeableConcept          `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode       []*CodeableConcept          `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity          *Quantity                   `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money                      `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *float64                    `bson:"factor,omitempty" json:"factor,omitempty"`
	Net               *Money                      `bson:"net,omitempty" json:"net,omitempty"`
	Udi               []*Reference                `bson:"udi,omitempty" json:"udi,omitempty"`
	SubDetail         []*ClaimItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}
type ClaimItemDetailSubDetail struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Sequence          int                `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Revenue           *CodeableConcept   `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept   `bson:"category,omitempty" json:"category,omitempty"`
	Service           *CodeableConcept   `bson:"service,omitempty" json:"service,omitempty"`
	Modifier          []*CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	ProgramCode       []*CodeableConcept `bson:"programCode,omitempty" json:"programCode,omitempty"`
	Quantity          *Quantity          `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice         *Money             `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor            *float64           `bson:"factor,omitempty" json:"factor,omitempty"`
	Net               *Money             `bson:"net,omitempty" json:"net,omitempty"`
	Udi               []*Reference       `bson:"udi,omitempty" json:"udi,omitempty"`
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherClaim
	}{
		OtherClaim:   OtherClaim(r),
		ResourceType: "Claim",
	})
	return buffer.Bytes(), err
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
