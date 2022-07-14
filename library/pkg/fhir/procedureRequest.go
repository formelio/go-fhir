package fhir

import "encoding/json"

// ProcedureRequest is documented here http://hl7.org/fhir/StructureDefinition/ProcedureRequest
type ProcedureRequest struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                    `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                 `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn           []Reference                `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces          []Reference                `bson:"replaces,omitempty" json:"replaces,omitempty"`
	Requisition       *Identifier                `bson:"requisition,omitempty" json:"requisition,omitempty"`
	Status            string                     `bson:"status" json:"status"`
	Intent            string                     `bson:"intent" json:"intent"`
	Priority          *string                    `bson:"priority,omitempty" json:"priority,omitempty"`
	DoNotPerform      *bool                      `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	Category          []CodeableConcept          `bson:"category,omitempty" json:"category,omitempty"`
	Code              CodeableConcept            `bson:"code" json:"code"`
	AuthoredOn        *string                    `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester         *ProcedureRequestRequester `bson:"requester,omitempty" json:"requester,omitempty"`
	PerformerType     *CodeableConcept           `bson:"performerType,omitempty" json:"performerType,omitempty"`
	ReasonCode        []CodeableConcept          `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	SupportingInfo    []Reference                `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Specimen          []Reference                `bson:"specimen,omitempty" json:"specimen,omitempty"`
	BodySite          []CodeableConcept          `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Note              []Annotation               `bson:"note,omitempty" json:"note,omitempty"`
	RelevantHistory   []Reference                `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
}
type ProcedureRequestRequester struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}
type OtherProcedureRequest ProcedureRequest

// MarshalJSON marshals the given ProcedureRequest as JSON into a byte slice
func (r ProcedureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherProcedureRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherProcedureRequest: OtherProcedureRequest(r),
		ResourceType:          "ProcedureRequest",
	})
}

// UnmarshalProcedureRequest unmarshals a ProcedureRequest.
func UnmarshalProcedureRequest(b []byte) (ProcedureRequest, error) {
	var procedureRequest ProcedureRequest
	if err := json.Unmarshal(b, &procedureRequest); err != nil {
		return procedureRequest, err
	}
	return procedureRequest, nil
}
