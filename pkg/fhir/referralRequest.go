package fhir

import "encoding/json"

// ReferralRequest is documented here http://hl7.org/fhir/StructureDefinition/ReferralRequest
type ReferralRequest struct {
	Id                 *string                   `bson:"id" json:"id"`
	Meta               *Meta                     `bson:"meta" json:"meta"`
	ImplicitRules      *string                   `bson:"implicitRules" json:"implicitRules"`
	Language           *string                   `bson:"language" json:"language"`
	Text               *Narrative                `bson:"text" json:"text"`
	RawContained       []json.RawMessage         `bson:"contained" json:"contained"`
	Contained          []IResource               `bson:"-" json:"-"`
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

// OtherReferralRequest is a helper type to use the default implementations of Marshall and Unmarshal
type OtherReferralRequest ReferralRequest

// MarshalJSON marshals the given ReferralRequest as JSON into a byte slice
func (r ReferralRequest) MarshalJSON() ([]byte, error) {
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
		OtherReferralRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherReferralRequest: OtherReferralRequest(r),
		ResourceType:         "ReferralRequest",
	})
}

// UnmarshalJSON unmarshals the given byte slice into ReferralRequest
func (r *ReferralRequest) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherReferralRequest)(r)); err != nil {
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
func (r ReferralRequest) GetResourceType() ResourceType {
	return ResourceTypeReferralRequest
}
