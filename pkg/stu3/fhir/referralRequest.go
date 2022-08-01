package fhir

import "encoding/json"

// ReferralRequest is documented here http://hl7.org/fhir/StructureDefinition/ReferralRequest
type ReferralRequest struct {
	Id                 *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string                   `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative                `bson:"text,omitempty" json:"text,omitempty"`
	RawContained       []json.RawMessage         `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained          []IResource               `bson:"-,omitempty" json:"-,omitempty"`
	Extension          []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         []Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Definition         []Reference               `bson:"definition,omitempty" json:"definition,omitempty"`
	BasedOn            []Reference               `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces           []Reference               `bson:"replaces,omitempty" json:"replaces,omitempty"`
	GroupIdentifier    *Identifier               `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	Status             RequestStatus             `bson:"status,omitempty" json:"status,omitempty"`
	Intent             RequestIntent             `bson:"intent,omitempty" json:"intent,omitempty"`
	Type               *CodeableConcept          `bson:"type,omitempty" json:"type,omitempty"`
	Priority           *RequestPriority          `bson:"priority,omitempty" json:"priority,omitempty"`
	ServiceRequested   []CodeableConcept         `bson:"serviceRequested,omitempty" json:"serviceRequested,omitempty"`
	Subject            Reference                 `bson:"subject,omitempty" json:"subject,omitempty"`
	Context            *Reference                `bson:"context,omitempty" json:"context,omitempty"`
	OccurrenceDateTime *string                   `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period                   `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	AuthoredOn         *string                   `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester          *ReferralRequestRequester `bson:"requester,omitempty" json:"requester,omitempty"`
	Specialty          *CodeableConcept          `bson:"specialty,omitempty" json:"specialty,omitempty"`
	Recipient          []Reference               `bson:"recipient,omitempty" json:"recipient,omitempty"`
	ReasonCode         []CodeableConcept         `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference    []Reference               `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Description        *string                   `bson:"description,omitempty" json:"description,omitempty"`
	SupportingInfo     []Reference               `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Note               []Annotation              `bson:"note,omitempty" json:"note,omitempty"`
	RelevantHistory    []Reference               `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
}
type ReferralRequestRequester struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Agent             Reference   `bson:"agent,omitempty" json:"agent,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
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
