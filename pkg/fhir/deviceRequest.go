package fhir

import "encoding/json"

// DeviceRequest is documented here http://hl7.org/fhir/StructureDefinition/DeviceRequest
type DeviceRequest struct {
	Id                  *string                 `bson:"id" json:"id"`
	Meta                *Meta                   `bson:"meta" json:"meta"`
	ImplicitRules       *string                 `bson:"implicitRules" json:"implicitRules"`
	Language            *string                 `bson:"language" json:"language"`
	Text                *Narrative              `bson:"text" json:"text"`
	RawContained        []json.RawMessage       `bson:"contained" json:"contained"`
	Contained           []IResource             `bson:"-" json:"-"`
	Extension           []Extension             `bson:"extension" json:"extension"`
	ModifierExtension   []Extension             `bson:"modifierExtension" json:"modifierExtension"`
	Identifier          []Identifier            `bson:"identifier" json:"identifier"`
	Definition          []Reference             `bson:"definition" json:"definition"`
	BasedOn             []Reference             `bson:"basedOn" json:"basedOn"`
	PriorRequest        []Reference             `bson:"priorRequest" json:"priorRequest"`
	GroupIdentifier     *Identifier             `bson:"groupIdentifier" json:"groupIdentifier"`
	Status              *RequestStatus          `bson:"status" json:"status"`
	Intent              CodeableConcept         `bson:"intent,omitempty" json:"intent,omitempty"`
	Priority            *RequestPriority        `bson:"priority" json:"priority"`
	CodeReference       *Reference              `bson:"codeReference,omitempty" json:"codeReference,omitempty"`
	CodeCodeableConcept *CodeableConcept        `bson:"codeCodeableConcept,omitempty" json:"codeCodeableConcept,omitempty"`
	Subject             Reference               `bson:"subject,omitempty" json:"subject,omitempty"`
	Context             *Reference              `bson:"context" json:"context"`
	OccurrenceDateTime  *string                 `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod    *Period                 `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceTiming    *Timing                 `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	AuthoredOn          *string                 `bson:"authoredOn" json:"authoredOn"`
	Requester           *DeviceRequestRequester `bson:"requester" json:"requester"`
	PerformerType       *CodeableConcept        `bson:"performerType" json:"performerType"`
	Performer           *Reference              `bson:"performer" json:"performer"`
	ReasonCode          []CodeableConcept       `bson:"reasonCode" json:"reasonCode"`
	ReasonReference     []Reference             `bson:"reasonReference" json:"reasonReference"`
	SupportingInfo      []Reference             `bson:"supportingInfo" json:"supportingInfo"`
	Note                []Annotation            `bson:"note" json:"note"`
	RelevantHistory     []Reference             `bson:"relevantHistory" json:"relevantHistory"`
}
type DeviceRequestRequester struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Agent             Reference   `bson:"agent,omitempty" json:"agent,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf" json:"onBehalfOf"`
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
