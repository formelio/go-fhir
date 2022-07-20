package fhir

import "encoding/json"

// DeviceRequest is documented here http://hl7.org/fhir/StructureDefinition/DeviceRequest
type DeviceRequest struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn           []Reference             `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PriorRequest      []Reference             `bson:"priorRequest,omitempty" json:"priorRequest,omitempty"`
	GroupIdentifier   *Identifier             `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	Status            *string                 `bson:"status,omitempty" json:"status,omitempty"`
	Intent            CodeableConcept         `bson:"intent" json:"intent"`
	Priority          *string                 `bson:"priority,omitempty" json:"priority,omitempty"`
	AuthoredOn        *string                 `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester         *DeviceRequestRequester `bson:"requester,omitempty" json:"requester,omitempty"`
	PerformerType     *CodeableConcept        `bson:"performerType,omitempty" json:"performerType,omitempty"`
	ReasonCode        []CodeableConcept       `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference   []Reference             `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	SupportingInfo    []Reference             `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Note              []Annotation            `bson:"note,omitempty" json:"note,omitempty"`
	RelevantHistory   []Reference             `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
}
type DeviceRequestRequester struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}
type OtherDeviceRequest DeviceRequest

// MarshalJSON marshals the given DeviceRequest as JSON into a byte slice
func (r DeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceRequest: OtherDeviceRequest(r),
		ResourceType:       "DeviceRequest",
	})
}

// UnmarshalDeviceRequest unmarshals a DeviceRequest.
func UnmarshalDeviceRequest(b []byte) (DeviceRequest, error) {
	var deviceRequest DeviceRequest
	if err := json.Unmarshal(b, &deviceRequest); err != nil {
		return deviceRequest, err
	}
	return deviceRequest, nil
}
