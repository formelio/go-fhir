package fhir

import "encoding/json"

// ReferralRequest is documented here http://hl7.org/fhir/StructureDefinition/ReferralRequest
type ReferralRequest struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                   `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Replaces          []Reference               `bson:"replaces,omitempty" json:"replaces,omitempty"`
	GroupIdentifier   *Identifier               `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	Status            string                    `bson:"status" json:"status"`
	Intent            string                    `bson:"intent" json:"intent"`
	Type              *CodeableConcept          `bson:"type,omitempty" json:"type,omitempty"`
	Priority          *string                   `bson:"priority,omitempty" json:"priority,omitempty"`
	ServiceRequested  []CodeableConcept         `bson:"serviceRequested,omitempty" json:"serviceRequested,omitempty"`
	AuthoredOn        *string                   `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester         *ReferralRequestRequester `bson:"requester,omitempty" json:"requester,omitempty"`
	Specialty         *CodeableConcept          `bson:"specialty,omitempty" json:"specialty,omitempty"`
	ReasonCode        []CodeableConcept         `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	Description       *string                   `bson:"description,omitempty" json:"description,omitempty"`
	SupportingInfo    []Reference               `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Note              []Annotation              `bson:"note,omitempty" json:"note,omitempty"`
	RelevantHistory   []Reference               `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
}
type ReferralRequestRequester struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
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

// UnmarshalReferralRequest unmarshals a ReferralRequest.
func UnmarshalReferralRequest(b []byte) (ReferralRequest, error) {
	var referralRequest ReferralRequest
	if err := json.Unmarshal(b, &referralRequest); err != nil {
		return referralRequest, err
	}
	return referralRequest, nil
}
