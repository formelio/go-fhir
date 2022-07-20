package fhir

import "encoding/json"

// ProcedureRequest is documented here http://hl7.org/fhir/StructureDefinition/ProcedureRequest
type ProcedureRequest struct {
	Id                      *string                    `bson:"id" json:"id"`
	Meta                    *Meta                      `bson:"meta" json:"meta"`
	ImplicitRules           *string                    `bson:"implicitRules" json:"implicitRules"`
	Language                *string                    `bson:"language" json:"language"`
	Text                    *Narrative                 `bson:"text" json:"text"`
	Contained               []json.RawMessage          `bson:"contained" json:"contained"`
	Extension               []Extension                `bson:"extension" json:"extension"`
	ModifierExtension       []Extension                `bson:"modifierExtension" json:"modifierExtension"`
	Identifier              []Identifier               `bson:"identifier" json:"identifier"`
	Definition              []Reference                `bson:"definition" json:"definition"`
	BasedOn                 []Reference                `bson:"basedOn" json:"basedOn"`
	Replaces                []Reference                `bson:"replaces" json:"replaces"`
	Requisition             *Identifier                `bson:"requisition" json:"requisition"`
	Status                  RequestStatus              `bson:"status,omitempty" json:"status,omitempty"`
	Intent                  RequestIntent              `bson:"intent,omitempty" json:"intent,omitempty"`
	Priority                *RequestPriority           `bson:"priority" json:"priority"`
	DoNotPerform            *bool                      `bson:"doNotPerform" json:"doNotPerform"`
	Category                []CodeableConcept          `bson:"category" json:"category"`
	Code                    CodeableConcept            `bson:"code,omitempty" json:"code,omitempty"`
	Subject                 Reference                  `bson:"subject,omitempty" json:"subject,omitempty"`
	Context                 *Reference                 `bson:"context" json:"context"`
	OccurrenceDateTime      *string                    `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod        *Period                    `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceTiming        *Timing                    `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	AsNeededBoolean         *bool                      `bson:"asNeededBoolean,omitempty" json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept *CodeableConcept           `bson:"asNeededCodeableConcept,omitempty" json:"asNeededCodeableConcept,omitempty"`
	AuthoredOn              *string                    `bson:"authoredOn" json:"authoredOn"`
	Requester               *ProcedureRequestRequester `bson:"requester" json:"requester"`
	PerformerType           *CodeableConcept           `bson:"performerType" json:"performerType"`
	Performer               *Reference                 `bson:"performer" json:"performer"`
	ReasonCode              []CodeableConcept          `bson:"reasonCode" json:"reasonCode"`
	ReasonReference         []Reference                `bson:"reasonReference" json:"reasonReference"`
	SupportingInfo          []Reference                `bson:"supportingInfo" json:"supportingInfo"`
	Specimen                []Reference                `bson:"specimen" json:"specimen"`
	BodySite                []CodeableConcept          `bson:"bodySite" json:"bodySite"`
	Note                    []Annotation               `bson:"note" json:"note"`
	RelevantHistory         []Reference                `bson:"relevantHistory" json:"relevantHistory"`
}
type ProcedureRequestRequester struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Agent             Reference   `bson:"agent,omitempty" json:"agent,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf" json:"onBehalfOf"`
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

// UnmarshalProcedureRequest unmarshalls a ProcedureRequest.
func UnmarshalProcedureRequest(b []byte) (ProcedureRequest, error) {
	var procedureRequest ProcedureRequest
	if err := json.Unmarshal(b, &procedureRequest); err != nil {
		return procedureRequest, err
	}
	return procedureRequest, nil
}