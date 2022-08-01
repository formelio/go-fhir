package fhir

import "encoding/json"

// DeviceRequest is documented here http://hl7.org/fhir/StructureDefinition/DeviceRequest
type DeviceRequest struct {
	Id                  *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	RawContained        []json.RawMessage       `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained           []IResource             `bson:"-,omitempty" json:"-,omitempty"`
	Extension           []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Definition          []Reference             `bson:"definition,omitempty" json:"definition,omitempty"`
	BasedOn             []Reference             `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PriorRequest        []Reference             `bson:"priorRequest,omitempty" json:"priorRequest,omitempty"`
	GroupIdentifier     *Identifier             `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	Status              *RequestStatus          `bson:"status,omitempty" json:"status,omitempty"`
	Intent              CodeableConcept         `bson:"intent,omitempty" json:"intent,omitempty"`
	Priority            *RequestPriority        `bson:"priority,omitempty" json:"priority,omitempty"`
	CodeReference       *Reference              `bson:"codeReference,omitempty" json:"codeReference,omitempty"`
	CodeCodeableConcept *CodeableConcept        `bson:"codeCodeableConcept,omitempty" json:"codeCodeableConcept,omitempty"`
	Subject             Reference               `bson:"subject,omitempty" json:"subject,omitempty"`
	Context             *Reference              `bson:"context,omitempty" json:"context,omitempty"`
	OccurrenceDateTime  *string                 `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod    *Period                 `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceTiming    *Timing                 `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	AuthoredOn          *string                 `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester           *DeviceRequestRequester `bson:"requester,omitempty" json:"requester,omitempty"`
	PerformerType       *CodeableConcept        `bson:"performerType,omitempty" json:"performerType,omitempty"`
	Performer           *Reference              `bson:"performer,omitempty" json:"performer,omitempty"`
	ReasonCode          []CodeableConcept       `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference     []Reference             `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	SupportingInfo      []Reference             `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Note                []Annotation            `bson:"note,omitempty" json:"note,omitempty"`
	RelevantHistory     []Reference             `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
}
type DeviceRequestRequester struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Agent             Reference   `bson:"agent,omitempty" json:"agent,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}

// OtherDeviceRequest is a helper type to use the default implementations of Marshall and Unmarshal
type OtherDeviceRequest DeviceRequest

// MarshalJSON marshals the given DeviceRequest as JSON into a byte slice
func (r DeviceRequest) MarshalJSON() ([]byte, error) {
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
		OtherDeviceRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceRequest: OtherDeviceRequest(r),
		ResourceType:       "DeviceRequest",
	})
}

// UnmarshalJSON unmarshals the given byte slice into DeviceRequest
func (r *DeviceRequest) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherDeviceRequest)(r)); err != nil {
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
func (r DeviceRequest) GetResourceType() ResourceType {
	return ResourceTypeDeviceRequest
}
