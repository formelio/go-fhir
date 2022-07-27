package fhir

import "encoding/json"

// ClaimResponse is documented here http://hl7.org/fhir/StructureDefinition/ClaimResponse
type ClaimResponse struct {
	Id                   *string                    `bson:"id" json:"id"`
	Meta                 *Meta                      `bson:"meta" json:"meta"`
	ImplicitRules        *string                    `bson:"implicitRules" json:"implicitRules"`
	Language             *string                    `bson:"language" json:"language"`
	Text                 *Narrative                 `bson:"text" json:"text"`
	RawContained         []json.RawMessage          `bson:"contained" json:"contained"`
	Contained            []IResource                `bson:"-" json:"-"`
	Extension            []Extension                `bson:"extension" json:"extension"`
	ModifierExtension    []Extension                `bson:"modifierExtension" json:"modifierExtension"`
	Identifier           []Identifier               `bson:"identifier" json:"identifier"`
	Status               *string                    `bson:"status" json:"status"`
	Patient              *Reference                 `bson:"patient" json:"patient"`
	Created              *string                    `bson:"created" json:"created"`
	Insurer              *Reference                 `bson:"insurer" json:"insurer"`
	RequestProvider      *Reference                 `bson:"requestProvider" json:"requestProvider"`
	RequestOrganization  *Reference                 `bson:"requestOrganization" json:"requestOrganization"`
	Request              *Reference                 `bson:"request" json:"request"`
	Outcome              *CodeableConcept           `bson:"outcome" json:"outcome"`
	Disposition          *string                    `bson:"disposition" json:"disposition"`
	PayeeType            *CodeableConcept           `bson:"payeeType" json:"payeeType"`
	Item                 []ClaimResponseItem        `bson:"item" json:"item"`
	AddItem              []ClaimResponseAddItem     `bson:"addItem" json:"addItem"`
	Error                []ClaimResponseError       `bson:"error" json:"error"`
	TotalCost            *Money                     `bson:"totalCost" json:"totalCost"`
	UnallocDeductable    *Money                     `bson:"unallocDeductable" json:"unallocDeductable"`
	TotalBenefit         *Money                     `bson:"totalBenefit" json:"totalBenefit"`
	Payment              *ClaimResponsePayment      `bson:"payment" json:"payment"`
	Reserved             *Coding                    `bson:"reserved" json:"reserved"`
	Form                 *CodeableConcept           `bson:"form" json:"form"`
	ProcessNote          []ClaimResponseProcessNote `bson:"processNote" json:"processNote"`
	CommunicationRequest []Reference                `bson:"communicationRequest" json:"communicationRequest"`
	Insurance            []ClaimResponseInsurance   `bson:"insurance" json:"insurance"`
}
type ClaimResponseItem struct {
	Id                *string                         `bson:"id" json:"id"`
	Extension         []Extension                     `bson:"extension" json:"extension"`
	ModifierExtension []Extension                     `bson:"modifierExtension" json:"modifierExtension"`
	SequenceLinkId    int                             `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	NoteNumber        []int                           `bson:"noteNumber" json:"noteNumber"`
	Adjudication      []ClaimResponseItemAdjudication `bson:"adjudication" json:"adjudication"`
	Detail            []ClaimResponseItemDetail       `bson:"detail" json:"detail"`
}
type ClaimResponseItemAdjudication struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Category          CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	Reason            *CodeableConcept `bson:"reason" json:"reason"`
	Amount            *Money           `bson:"amount" json:"amount"`
	Value             *float64         `bson:"value" json:"value"`
}
type ClaimResponseItemDetail struct {
	Id                *string                            `bson:"id" json:"id"`
	Extension         []Extension                        `bson:"extension" json:"extension"`
	ModifierExtension []Extension                        `bson:"modifierExtension" json:"modifierExtension"`
	SequenceLinkId    int                                `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	NoteNumber        []int                              `bson:"noteNumber" json:"noteNumber"`
	Adjudication      []ClaimResponseItemAdjudication    `bson:"adjudication" json:"adjudication"`
	SubDetail         []ClaimResponseItemDetailSubDetail `bson:"subDetail" json:"subDetail"`
}
type ClaimResponseItemDetailSubDetail struct {
	Id                *string                         `bson:"id" json:"id"`
	Extension         []Extension                     `bson:"extension" json:"extension"`
	ModifierExtension []Extension                     `bson:"modifierExtension" json:"modifierExtension"`
	SequenceLinkId    int                             `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	NoteNumber        []int                           `bson:"noteNumber" json:"noteNumber"`
	Adjudication      []ClaimResponseItemAdjudication `bson:"adjudication" json:"adjudication"`
}
type ClaimResponseAddItem struct {
	Id                *string                         `bson:"id" json:"id"`
	Extension         []Extension                     `bson:"extension" json:"extension"`
	ModifierExtension []Extension                     `bson:"modifierExtension" json:"modifierExtension"`
	SequenceLinkId    []int                           `bson:"sequenceLinkId" json:"sequenceLinkId"`
	Revenue           *CodeableConcept                `bson:"revenue" json:"revenue"`
	Category          *CodeableConcept                `bson:"category" json:"category"`
	Service           *CodeableConcept                `bson:"service" json:"service"`
	Modifier          []CodeableConcept               `bson:"modifier" json:"modifier"`
	Fee               *Money                          `bson:"fee" json:"fee"`
	NoteNumber        []int                           `bson:"noteNumber" json:"noteNumber"`
	Adjudication      []ClaimResponseItemAdjudication `bson:"adjudication" json:"adjudication"`
	Detail            []ClaimResponseAddItemDetail    `bson:"detail" json:"detail"`
}
type ClaimResponseAddItemDetail struct {
	Id                *string                         `bson:"id" json:"id"`
	Extension         []Extension                     `bson:"extension" json:"extension"`
	ModifierExtension []Extension                     `bson:"modifierExtension" json:"modifierExtension"`
	Revenue           *CodeableConcept                `bson:"revenue" json:"revenue"`
	Category          *CodeableConcept                `bson:"category" json:"category"`
	Service           *CodeableConcept                `bson:"service" json:"service"`
	Modifier          []CodeableConcept               `bson:"modifier" json:"modifier"`
	Fee               *Money                          `bson:"fee" json:"fee"`
	NoteNumber        []int                           `bson:"noteNumber" json:"noteNumber"`
	Adjudication      []ClaimResponseItemAdjudication `bson:"adjudication" json:"adjudication"`
}
type ClaimResponseError struct {
	Id                      *string         `bson:"id" json:"id"`
	Extension               []Extension     `bson:"extension" json:"extension"`
	ModifierExtension       []Extension     `bson:"modifierExtension" json:"modifierExtension"`
	SequenceLinkId          *int            `bson:"sequenceLinkId" json:"sequenceLinkId"`
	DetailSequenceLinkId    *int            `bson:"detailSequenceLinkId" json:"detailSequenceLinkId"`
	SubdetailSequenceLinkId *int            `bson:"subdetailSequenceLinkId" json:"subdetailSequenceLinkId"`
	Code                    CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
}
type ClaimResponsePayment struct {
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
type ClaimResponseProcessNote struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Number            *int             `bson:"number" json:"number"`
	Type              *CodeableConcept `bson:"type" json:"type"`
	Text              *string          `bson:"text" json:"text"`
	Language          *CodeableConcept `bson:"language" json:"language"`
}
type ClaimResponseInsurance struct {
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
	return json.Marshal(struct {
		OtherClaimResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherClaimResponse: OtherClaimResponse(r),
		ResourceType:       "ClaimResponse",
	})
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
