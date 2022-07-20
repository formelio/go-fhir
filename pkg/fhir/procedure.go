package fhir

import "encoding/json"

// Procedure is documented here http://hl7.org/fhir/StructureDefinition/Procedure
type Procedure struct {
	Id                 *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Extension          []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         []Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status             string                 `bson:"status" json:"status"`
	NotDone            *bool                  `bson:"notDone,omitempty" json:"notDone,omitempty"`
	NotDoneReason      *CodeableConcept       `bson:"notDoneReason,omitempty" json:"notDoneReason,omitempty"`
	Category           *CodeableConcept       `bson:"category,omitempty" json:"category,omitempty"`
	Code               *CodeableConcept       `bson:"code,omitempty" json:"code,omitempty"`
	Performer          []ProcedurePerformer   `bson:"performer,omitempty" json:"performer,omitempty"`
	Location           *Reference             `bson:"location,omitempty" json:"location,omitempty"`
	ReasonCode         []CodeableConcept      `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	BodySite           []CodeableConcept      `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Outcome            *CodeableConcept       `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Report             []Reference            `bson:"report,omitempty" json:"report,omitempty"`
	Complication       []CodeableConcept      `bson:"complication,omitempty" json:"complication,omitempty"`
	ComplicationDetail []Reference            `bson:"complicationDetail,omitempty" json:"complicationDetail,omitempty"`
	FollowUp           []CodeableConcept      `bson:"followUp,omitempty" json:"followUp,omitempty"`
	Note               []Annotation           `bson:"note,omitempty" json:"note,omitempty"`
	FocalDevice        []ProcedureFocalDevice `bson:"focalDevice,omitempty" json:"focalDevice,omitempty"`
	UsedCode           []CodeableConcept      `bson:"usedCode,omitempty" json:"usedCode,omitempty"`
}
type ProcedurePerformer struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	OnBehalfOf        *Reference       `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}
type ProcedureFocalDevice struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Action            *CodeableConcept `bson:"action,omitempty" json:"action,omitempty"`
	Manipulated       Reference        `bson:"manipulated" json:"manipulated"`
}
type OtherProcedure Procedure

// MarshalJSON marshals the given Procedure as JSON into a byte slice
func (r Procedure) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherProcedure
		ResourceType string `json:"resourceType"`
	}{
		OtherProcedure: OtherProcedure(r),
		ResourceType:   "Procedure",
	})
}

// UnmarshalProcedure unmarshals a Procedure.
func UnmarshalProcedure(b []byte) (Procedure, error) {
	var procedure Procedure
	if err := json.Unmarshal(b, &procedure); err != nil {
		return procedure, err
	}
	return procedure, nil
}
