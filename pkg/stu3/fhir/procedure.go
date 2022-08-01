package fhir

import (
	"bytes"
	"encoding/json"
)

// Procedure is documented here http://hl7.org/fhir/StructureDefinition/Procedure
type Procedure struct {
	Id                 *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	RawContained       []json.RawMessage       `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained          []IResource             `bson:"-,omitempty" json:"-,omitempty"`
	Extension          []*Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []*Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         []*Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Definition         []*Reference            `bson:"definition,omitempty" json:"definition,omitempty"`
	BasedOn            []*Reference            `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf             []*Reference            `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status             EventStatus             `bson:"status,omitempty" json:"status,omitempty"`
	NotDone            *bool                   `bson:"notDone,omitempty" json:"notDone,omitempty"`
	NotDoneReason      *CodeableConcept        `bson:"notDoneReason,omitempty" json:"notDoneReason,omitempty"`
	Category           *CodeableConcept        `bson:"category,omitempty" json:"category,omitempty"`
	Code               *CodeableConcept        `bson:"code,omitempty" json:"code,omitempty"`
	Subject            Reference               `bson:"subject,omitempty" json:"subject,omitempty"`
	Context            *Reference              `bson:"context,omitempty" json:"context,omitempty"`
	PerformedDateTime  *string                 `bson:"performedDateTime,omitempty" json:"performedDateTime,omitempty"`
	PerformedPeriod    *Period                 `bson:"performedPeriod,omitempty" json:"performedPeriod,omitempty"`
	Performer          []*ProcedurePerformer   `bson:"performer,omitempty" json:"performer,omitempty"`
	Location           *Reference              `bson:"location,omitempty" json:"location,omitempty"`
	ReasonCode         []*CodeableConcept      `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference    []*Reference            `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	BodySite           []*CodeableConcept      `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Outcome            *CodeableConcept        `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Report             []*Reference            `bson:"report,omitempty" json:"report,omitempty"`
	Complication       []*CodeableConcept      `bson:"complication,omitempty" json:"complication,omitempty"`
	ComplicationDetail []*Reference            `bson:"complicationDetail,omitempty" json:"complicationDetail,omitempty"`
	FollowUp           []*CodeableConcept      `bson:"followUp,omitempty" json:"followUp,omitempty"`
	Note               []*Annotation           `bson:"note,omitempty" json:"note,omitempty"`
	FocalDevice        []*ProcedureFocalDevice `bson:"focalDevice,omitempty" json:"focalDevice,omitempty"`
	UsedReference      []*Reference            `bson:"usedReference,omitempty" json:"usedReference,omitempty"`
	UsedCode           []*CodeableConcept      `bson:"usedCode,omitempty" json:"usedCode,omitempty"`
}
type ProcedurePerformer struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Actor             Reference        `bson:"actor,omitempty" json:"actor,omitempty"`
	OnBehalfOf        *Reference       `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}
type ProcedureFocalDevice struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Action            *CodeableConcept `bson:"action,omitempty" json:"action,omitempty"`
	Manipulated       Reference        `bson:"manipulated,omitempty" json:"manipulated,omitempty"`
}

// OtherProcedure is a helper type to use the default implementations of Marshall and Unmarshal
type OtherProcedure Procedure

// MarshalJSON marshals the given Procedure as JSON into a byte slice
func (r Procedure) MarshalJSON() ([]byte, error) {
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
		OtherProcedure
	}{
		OtherProcedure: OtherProcedure(r),
		ResourceType:   "Procedure",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into Procedure
func (r *Procedure) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherProcedure)(r)); err != nil {
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
func (r Procedure) GetResourceType() ResourceType {
	return ResourceTypeProcedure
}
