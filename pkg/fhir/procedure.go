package fhir

import "encoding/json"

// Procedure is documented here http://hl7.org/fhir/StructureDefinition/Procedure
type Procedure struct {
	Id                 *string                `bson:"id" json:"id"`
	Meta               *Meta                  `bson:"meta" json:"meta"`
	ImplicitRules      *string                `bson:"implicitRules" json:"implicitRules"`
	Language           *string                `bson:"language" json:"language"`
	Text               *Narrative             `bson:"text" json:"text"`
	RawContained       []json.RawMessage      `bson:"contained" json:"contained"`
	Contained          []IResource            `bson:"-" json:"-"`
	Extension          []Extension            `bson:"extension" json:"extension"`
	ModifierExtension  []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	Identifier         []Identifier           `bson:"identifier" json:"identifier"`
	Definition         []Reference            `bson:"definition" json:"definition"`
	BasedOn            []Reference            `bson:"basedOn" json:"basedOn"`
	PartOf             []Reference            `bson:"partOf" json:"partOf"`
	Status             EventStatus            `bson:"status,omitempty" json:"status,omitempty"`
	NotDone            *bool                  `bson:"notDone" json:"notDone"`
	NotDoneReason      *CodeableConcept       `bson:"notDoneReason" json:"notDoneReason"`
	Category           *CodeableConcept       `bson:"category" json:"category"`
	Code               *CodeableConcept       `bson:"code" json:"code"`
	Subject            Reference              `bson:"subject,omitempty" json:"subject,omitempty"`
	Context            *Reference             `bson:"context" json:"context"`
	PerformedDateTime  *string                `bson:"performedDateTime,omitempty" json:"performedDateTime,omitempty"`
	PerformedPeriod    *Period                `bson:"performedPeriod,omitempty" json:"performedPeriod,omitempty"`
	Performer          []ProcedurePerformer   `bson:"performer" json:"performer"`
	Location           *Reference             `bson:"location" json:"location"`
	ReasonCode         []CodeableConcept      `bson:"reasonCode" json:"reasonCode"`
	ReasonReference    []Reference            `bson:"reasonReference" json:"reasonReference"`
	BodySite           []CodeableConcept      `bson:"bodySite" json:"bodySite"`
	Outcome            *CodeableConcept       `bson:"outcome" json:"outcome"`
	Report             []Reference            `bson:"report" json:"report"`
	Complication       []CodeableConcept      `bson:"complication" json:"complication"`
	ComplicationDetail []Reference            `bson:"complicationDetail" json:"complicationDetail"`
	FollowUp           []CodeableConcept      `bson:"followUp" json:"followUp"`
	Note               []Annotation           `bson:"note" json:"note"`
	FocalDevice        []ProcedureFocalDevice `bson:"focalDevice" json:"focalDevice"`
	UsedReference      []Reference            `bson:"usedReference" json:"usedReference"`
	UsedCode           []CodeableConcept      `bson:"usedCode" json:"usedCode"`
}
type ProcedurePerformer struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Role              *CodeableConcept `bson:"role" json:"role"`
	Actor             Reference        `bson:"actor,omitempty" json:"actor,omitempty"`
	OnBehalfOf        *Reference       `bson:"onBehalfOf" json:"onBehalfOf"`
}
type ProcedureFocalDevice struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Action            *CodeableConcept `bson:"action" json:"action"`
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
	return json.Marshal(struct {
		OtherProcedure
		ResourceType string `json:"resourceType"`
	}{
		OtherProcedure: OtherProcedure(r),
		ResourceType:   "Procedure",
	})
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
