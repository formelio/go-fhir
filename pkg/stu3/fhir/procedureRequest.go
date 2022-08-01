package fhir

import (
	"bytes"
	"encoding/json"
)

// ProcedureRequest is documented here http://hl7.org/fhir/StructureDefinition/ProcedureRequest
type ProcedureRequest struct {
	Id                      *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta                    *Meta                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules           *string                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                *string                    `bson:"language,omitempty" json:"language,omitempty"`
	Text                    *Narrative                 `bson:"text,omitempty" json:"text,omitempty"`
	RawContained            []json.RawMessage          `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained               []IResource                `bson:"-,omitempty" json:"-,omitempty"`
	Extension               []*Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []*Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier              []*Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Definition              []*Reference               `bson:"definition,omitempty" json:"definition,omitempty"`
	BasedOn                 []*Reference               `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces                []*Reference               `bson:"replaces,omitempty" json:"replaces,omitempty"`
	Requisition             *Identifier                `bson:"requisition,omitempty" json:"requisition,omitempty"`
	Status                  RequestStatus              `bson:"status,omitempty" json:"status,omitempty"`
	Intent                  RequestIntent              `bson:"intent,omitempty" json:"intent,omitempty"`
	Priority                *RequestPriority           `bson:"priority,omitempty" json:"priority,omitempty"`
	DoNotPerform            *bool                      `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	Category                []*CodeableConcept         `bson:"category,omitempty" json:"category,omitempty"`
	Code                    CodeableConcept            `bson:"code,omitempty" json:"code,omitempty"`
	Subject                 Reference                  `bson:"subject,omitempty" json:"subject,omitempty"`
	Context                 *Reference                 `bson:"context,omitempty" json:"context,omitempty"`
	OccurrenceDateTime      *string                    `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod        *Period                    `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceTiming        *Timing                    `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	AsNeededBoolean         *bool                      `bson:"asNeededBoolean,omitempty" json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept           `bson:"asNeededCodeableConcept,omitempty" json:"asNeededCodeableConcept,omitempty"`
	AuthoredOn              *string                    `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester               *ProcedureRequestRequester `bson:"requester,omitempty" json:"requester,omitempty"`
	PerformerType           *CodeableConcept           `bson:"performerType,omitempty" json:"performerType,omitempty"`
	Performer               *Reference                 `bson:"performer,omitempty" json:"performer,omitempty"`
	ReasonCode              []*CodeableConcept         `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference         []*Reference               `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	SupportingInfo          []*Reference               `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Specimen                []*Reference               `bson:"specimen,omitempty" json:"specimen,omitempty"`
	BodySite                []*CodeableConcept         `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Note                    []*Annotation              `bson:"note,omitempty" json:"note,omitempty"`
	RelevantHistory         []*Reference               `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
}
type ProcedureRequestRequester struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Agent             Reference    `bson:"agent,omitempty" json:"agent,omitempty"`
	OnBehalfOf        *Reference   `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}

// OtherProcedureRequest is a helper type to use the default implementations of Marshall and Unmarshal
type OtherProcedureRequest ProcedureRequest

// MarshalJSON marshals the given ProcedureRequest as JSON into a byte slice
func (r ProcedureRequest) MarshalJSON() ([]byte, error) {
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
		OtherProcedureRequest
	}{
		OtherProcedureRequest: OtherProcedureRequest(r),
		ResourceType:          "ProcedureRequest",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into ProcedureRequest
func (r *ProcedureRequest) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherProcedureRequest)(r)); err != nil {
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
func (r ProcedureRequest) GetResourceType() ResourceType {
	return ResourceTypeProcedureRequest
}
