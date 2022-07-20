package fhir

import "encoding/json"

// ReferralRequest is documented here http://hl7.org/fhir/StructureDefinition/ReferralRequest
type ReferralRequest struct {
	Id                 *string                   `bson:"id" json:"id"`
	Meta               *Meta                     `bson:"meta" json:"meta"`
	ImplicitRules      *string                   `bson:"implicitRules" json:"implicitRules"`
	Language           *string                   `bson:"language" json:"language"`
	Text               *Narrative                `bson:"text" json:"text"`
	Contained          []json.RawMessage         `bson:"contained" json:"contained"`
	Extension          []Extension               `bson:"extension" json:"extension"`
	ModifierExtension  []Extension               `bson:"modifierExtension" json:"modifierExtension"`
	Identifier         []Identifier              `bson:"identifier" json:"identifier"`
	Definition         []Reference               `bson:"definition" json:"definition"`
	BasedOn            []Reference               `bson:"basedOn" json:"basedOn"`
	Replaces           []Reference               `bson:"replaces" json:"replaces"`
	GroupIdentifier    *Identifier               `bson:"groupIdentifier" json:"groupIdentifier"`
	Status             RequestStatus             `bson:"status,omitempty" json:"status,omitempty"`
	Intent             RequestIntent             `bson:"intent,omitempty" json:"intent,omitempty"`
	Type               *CodeableConcept          `bson:"type" json:"type"`
	Priority           *RequestPriority          `bson:"priority" json:"priority"`
	ServiceRequested   []CodeableConcept         `bson:"serviceRequested" json:"serviceRequested"`
	Subject            Reference                 `bson:"subject,omitempty" json:"subject,omitempty"`
	Context            *Reference                `bson:"context" json:"context"`
	OccurrenceDateTime *string                   `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period                   `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	AuthoredOn         *string                   `bson:"authoredOn" json:"authoredOn"`
	Requester          *ReferralRequestRequester `bson:"requester" json:"requester"`
	Specialty          *CodeableConcept          `bson:"specialty" json:"specialty"`
	Recipient          []Reference               `bson:"recipient" json:"recipient"`
	ReasonCode         []CodeableConcept         `bson:"reasonCode" json:"reasonCode"`
	ReasonReference    []Reference               `bson:"reasonReference" json:"reasonReference"`
	Description        *string                   `bson:"description" json:"description"`
	SupportingInfo     []Reference               `bson:"supportingInfo" json:"supportingInfo"`
	Note               []Annotation              `bson:"note" json:"note"`
	RelevantHistory    []Reference               `bson:"relevantHistory" json:"relevantHistory"`
}
type ReferralRequestRequester struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Agent             Reference   `bson:"agent,omitempty" json:"agent,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf" json:"onBehalfOf"`
}
type OtherReferralRequest ReferralRequest

// MarshalJSON marshals the given ReferralRequest as JSON into a byte slice
func (r ReferralRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherReferralRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherReferralRequest: OtherReferralRequest(r),
		ResourceType:         "ReferralRequest",
	})
}

// UnmarshalReferralRequest unmarshalls a ReferralRequest.
func UnmarshalReferralRequest(b []byte) (ReferralRequest, error) {
	var referralRequest ReferralRequest
	if err := json.Unmarshal(b, &referralRequest); err != nil {
		return referralRequest, err
	}
	return referralRequest, nil
}
