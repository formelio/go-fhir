package fhir

import "encoding/json"

// ClaimResponse is documented here http://hl7.org/fhir/StructureDefinition/ClaimResponse
type ClaimResponse struct {
	Id                   *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                    `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative                 `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               *string                    `bson:"status,omitempty" json:"status,omitempty"`
	Patient              *Reference                 `bson:"patient,omitempty" json:"patient,omitempty"`
	Created              *string                    `bson:"created,omitempty" json:"created,omitempty"`
	Insurer              *Reference                 `bson:"insurer,omitempty" json:"insurer,omitempty"`
	RequestProvider      *Reference                 `bson:"requestProvider,omitempty" json:"requestProvider,omitempty"`
	RequestOrganization  *Reference                 `bson:"requestOrganization,omitempty" json:"requestOrganization,omitempty"`
	Request              *Reference                 `bson:"request,omitempty" json:"request,omitempty"`
	Outcome              *CodeableConcept           `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition          *string                    `bson:"disposition,omitempty" json:"disposition,omitempty"`
	PayeeType            *CodeableConcept           `bson:"payeeType,omitempty" json:"payeeType,omitempty"`
	Item                 []ClaimResponseItem        `bson:"item,omitempty" json:"item,omitempty"`
	AddItem              []ClaimResponseAddItem     `bson:"addItem,omitempty" json:"addItem,omitempty"`
	Error                []ClaimResponseError       `bson:"error,omitempty" json:"error,omitempty"`
	TotalCost            *Money                     `bson:"totalCost,omitempty" json:"totalCost,omitempty"`
	UnallocDeductable    *Money                     `bson:"unallocDeductable,omitempty" json:"unallocDeductable,omitempty"`
	TotalBenefit         *Money                     `bson:"totalBenefit,omitempty" json:"totalBenefit,omitempty"`
	Payment              *ClaimResponsePayment      `bson:"payment,omitempty" json:"payment,omitempty"`
	Reserved             *Coding                    `bson:"reserved,omitempty" json:"reserved,omitempty"`
	Form                 *CodeableConcept           `bson:"form,omitempty" json:"form,omitempty"`
	ProcessNote          []ClaimResponseProcessNote `bson:"processNote,omitempty" json:"processNote,omitempty"`
	CommunicationRequest []Reference                `bson:"communicationRequest,omitempty" json:"communicationRequest,omitempty"`
	Insurance            []ClaimResponseInsurance   `bson:"insurance,omitempty" json:"insurance,omitempty"`
}
type ClaimResponseItem struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SequenceLinkId    int                             `bson:"sequenceLinkId" json:"sequenceLinkId"`
	NoteNumber        []int                           `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail            []ClaimResponseItemDetail       `bson:"detail,omitempty" json:"detail,omitempty"`
}
type ClaimResponseItemAdjudication struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          CodeableConcept  `bson:"category" json:"category"`
	Reason            *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	Amount            *Money           `bson:"amount,omitempty" json:"amount,omitempty"`
	Value             *string          `bson:"value,omitempty" json:"value,omitempty"`
}
type ClaimResponseItemDetail struct {
	Id                *string                            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SequenceLinkId    int                                `bson:"sequenceLinkId" json:"sequenceLinkId"`
	NoteNumber        []int                              `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication    `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	SubDetail         []ClaimResponseItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}
type ClaimResponseItemDetailSubDetail struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SequenceLinkId    int                             `bson:"sequenceLinkId" json:"sequenceLinkId"`
	NoteNumber        []int                           `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}
type ClaimResponseAddItem struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SequenceLinkId    []int                           `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	Revenue           *CodeableConcept                `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept                `bson:"category,omitempty" json:"category,omitempty"`
	Service           *CodeableConcept                `bson:"service,omitempty" json:"service,omitempty"`
	Modifier          []CodeableConcept               `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Fee               *Money                          `bson:"fee,omitempty" json:"fee,omitempty"`
	NoteNumber        []int                           `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail            []ClaimResponseAddItemDetail    `bson:"detail,omitempty" json:"detail,omitempty"`
}
type ClaimResponseAddItemDetail struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Revenue           *CodeableConcept                `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept                `bson:"category,omitempty" json:"category,omitempty"`
	Service           *CodeableConcept                `bson:"service,omitempty" json:"service,omitempty"`
	Modifier          []CodeableConcept               `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Fee               *Money                          `bson:"fee,omitempty" json:"fee,omitempty"`
	NoteNumber        []int                           `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}
type ClaimResponseError struct {
	Id                      *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension               []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SequenceLinkId          *int            `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	DetailSequenceLinkId    *int            `bson:"detailSequenceLinkId,omitempty" json:"detailSequenceLinkId,omitempty"`
	SubdetailSequenceLinkId *int            `bson:"subdetailSequenceLinkId,omitempty" json:"subdetailSequenceLinkId,omitempty"`
	Code                    CodeableConcept `bson:"code" json:"code"`
}
type ClaimResponsePayment struct {
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
type ClaimResponseProcessNote struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Number            *int             `bson:"number,omitempty" json:"number,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Text              *string          `bson:"text,omitempty" json:"text,omitempty"`
	Language          *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
}
type ClaimResponseInsurance struct {
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
type OtherClaimResponse ClaimResponse

// MarshalJSON marshals the given ClaimResponse as JSON into a byte slice
func (r ClaimResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClaimResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherClaimResponse: OtherClaimResponse(r),
		ResourceType:       "ClaimResponse",
	})
}

// UnmarshalClaimResponse unmarshals a ClaimResponse.
func UnmarshalClaimResponse(b []byte) (ClaimResponse, error) {
	var claimResponse ClaimResponse
	if err := json.Unmarshal(b, &claimResponse); err != nil {
		return claimResponse, err
	}
	return claimResponse, nil
}
