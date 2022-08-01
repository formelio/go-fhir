package fhir

import (
	"bytes"
	"encoding/json"
)

// ClaimResponse is documented here http://hl7.org/fhir/StructureDefinition/ClaimResponse
type ClaimResponse struct {
	Id                   *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	RawContained         []json.RawMessage           `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained            []IResource                 `bson:"-,omitempty" json:"-,omitempty"`
	Extension            []*Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []*Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []*Identifier               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               *string                     `bson:"status,omitempty" json:"status,omitempty"`
	Patient              *Reference                  `bson:"patient,omitempty" json:"patient,omitempty"`
	Created              *string                     `bson:"created,omitempty" json:"created,omitempty"`
	Insurer              *Reference                  `bson:"insurer,omitempty" json:"insurer,omitempty"`
	RequestProvider      *Reference                  `bson:"requestProvider,omitempty" json:"requestProvider,omitempty"`
	RequestOrganization  *Reference                  `bson:"requestOrganization,omitempty" json:"requestOrganization,omitempty"`
	Request              *Reference                  `bson:"request,omitempty" json:"request,omitempty"`
	Outcome              *CodeableConcept            `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition          *string                     `bson:"disposition,omitempty" json:"disposition,omitempty"`
	PayeeType            *CodeableConcept            `bson:"payeeType,omitempty" json:"payeeType,omitempty"`
	Item                 []*ClaimResponseItem        `bson:"item,omitempty" json:"item,omitempty"`
	AddItem              []*ClaimResponseAddItem     `bson:"addItem,omitempty" json:"addItem,omitempty"`
	Error                []*ClaimResponseError       `bson:"error,omitempty" json:"error,omitempty"`
	TotalCost            *Money                      `bson:"totalCost,omitempty" json:"totalCost,omitempty"`
	UnallocDeductable    *Money                      `bson:"unallocDeductable,omitempty" json:"unallocDeductable,omitempty"`
	TotalBenefit         *Money                      `bson:"totalBenefit,omitempty" json:"totalBenefit,omitempty"`
	Payment              *ClaimResponsePayment       `bson:"payment,omitempty" json:"payment,omitempty"`
	Reserved             *Coding                     `bson:"reserved,omitempty" json:"reserved,omitempty"`
	Form                 *CodeableConcept            `bson:"form,omitempty" json:"form,omitempty"`
	ProcessNote          []*ClaimResponseProcessNote `bson:"processNote,omitempty" json:"processNote,omitempty"`
	CommunicationRequest []*Reference                `bson:"communicationRequest,omitempty" json:"communicationRequest,omitempty"`
	Insurance            []*ClaimResponseInsurance   `bson:"insurance,omitempty" json:"insurance,omitempty"`
}
type ClaimResponseItem struct {
	Id                *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SequenceLinkId    int                              `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	NoteNumber        []*int                           `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []*ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail            []*ClaimResponseItemDetail       `bson:"detail,omitempty" json:"detail,omitempty"`
}
type ClaimResponseItemAdjudication struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	Reason            *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	Amount            *Money           `bson:"amount,omitempty" json:"amount,omitempty"`
	Value             *float64         `bson:"value,omitempty" json:"value,omitempty"`
}
type ClaimResponseItemDetail struct {
	Id                *string                             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SequenceLinkId    int                                 `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	NoteNumber        []*int                              `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []*ClaimResponseItemAdjudication    `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	SubDetail         []*ClaimResponseItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}
type ClaimResponseItemDetailSubDetail struct {
	Id                *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SequenceLinkId    int                              `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	NoteNumber        []*int                           `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []*ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}
type ClaimResponseAddItem struct {
	Id                *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SequenceLinkId    []*int                           `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	Revenue           *CodeableConcept                 `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept                 `bson:"category,omitempty" json:"category,omitempty"`
	Service           *CodeableConcept                 `bson:"service,omitempty" json:"service,omitempty"`
	Modifier          []*CodeableConcept               `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Fee               *Money                           `bson:"fee,omitempty" json:"fee,omitempty"`
	NoteNumber        []*int                           `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []*ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail            []*ClaimResponseAddItemDetail    `bson:"detail,omitempty" json:"detail,omitempty"`
}
type ClaimResponseAddItemDetail struct {
	Id                *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Revenue           *CodeableConcept                 `bson:"revenue,omitempty" json:"revenue,omitempty"`
	Category          *CodeableConcept                 `bson:"category,omitempty" json:"category,omitempty"`
	Service           *CodeableConcept                 `bson:"service,omitempty" json:"service,omitempty"`
	Modifier          []*CodeableConcept               `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Fee               *Money                           `bson:"fee,omitempty" json:"fee,omitempty"`
	NoteNumber        []*int                           `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication      []*ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}
type ClaimResponseError struct {
	Id                      *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension               []*Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []*Extension    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SequenceLinkId          *int            `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	DetailSequenceLinkId    *int            `bson:"detailSequenceLinkId,omitempty" json:"detailSequenceLinkId,omitempty"`
	SubdetailSequenceLinkId *int            `bson:"subdetailSequenceLinkId,omitempty" json:"subdetailSequenceLinkId,omitempty"`
	Code                    CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
}
type ClaimResponsePayment struct {
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
type ClaimResponseProcessNote struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Number            *int             `bson:"number,omitempty" json:"number,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Text              *string          `bson:"text,omitempty" json:"text,omitempty"`
	Language          *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
}
type ClaimResponseInsurance struct {
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

// OtherClaimResponse is a helper type to use the default implementations of Marshall and Unmarshal
type OtherClaimResponse ClaimResponse

// MarshalJSON marshals the given ClaimResponse as JSON into a byte slice
func (r ClaimResponse) MarshalJSON() ([]byte, error) {
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
		OtherClaimResponse
	}{
		OtherClaimResponse: OtherClaimResponse(r),
		ResourceType:       "ClaimResponse",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into ClaimResponse
func (r *ClaimResponse) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherClaimResponse)(r)); err != nil {
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
func (r ClaimResponse) GetResourceType() ResourceType {
	return ResourceTypeClaimResponse
}
