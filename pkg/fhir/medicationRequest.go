package fhir

import "encoding/json"

// MedicationRequest is documented here http://hl7.org/fhir/StructureDefinition/MedicationRequest
type MedicationRequest struct {
	Id                        *string                           `bson:"id" json:"id"`
	Meta                      *Meta                             `bson:"meta" json:"meta"`
	ImplicitRules             *string                           `bson:"implicitRules" json:"implicitRules"`
	Language                  *string                           `bson:"language" json:"language"`
	Text                      *Narrative                        `bson:"text" json:"text"`
	Contained                 []json.RawMessage                 `bson:"contained" json:"contained"`
	Extension                 []Extension                       `bson:"extension" json:"extension"`
	ModifierExtension         []Extension                       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier                []Identifier                      `bson:"identifier" json:"identifier"`
	Definition                []Reference                       `bson:"definition" json:"definition"`
	BasedOn                   []Reference                       `bson:"basedOn" json:"basedOn"`
	GroupIdentifier           *Identifier                       `bson:"groupIdentifier" json:"groupIdentifier"`
	Status                    *MedicationRequestStatus          `bson:"status" json:"status"`
	Intent                    MedicationRequestIntent           `bson:"intent,omitempty" json:"intent,omitempty"`
	Category                  *CodeableConcept                  `bson:"category" json:"category"`
	Priority                  *MedicationRequestPriority        `bson:"priority" json:"priority"`
	MedicationCodeableConcept *CodeableConcept                  `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference                        `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	Subject                   Reference                         `bson:"subject,omitempty" json:"subject,omitempty"`
	Context                   *Reference                        `bson:"context" json:"context"`
	SupportingInformation     []Reference                       `bson:"supportingInformation" json:"supportingInformation"`
	AuthoredOn                *string                           `bson:"authoredOn" json:"authoredOn"`
	Requester                 *MedicationRequestRequester       `bson:"requester" json:"requester"`
	Recorder                  *Reference                        `bson:"recorder" json:"recorder"`
	ReasonCode                []CodeableConcept                 `bson:"reasonCode" json:"reasonCode"`
	ReasonReference           []Reference                       `bson:"reasonReference" json:"reasonReference"`
	Note                      []Annotation                      `bson:"note" json:"note"`
	DosageInstruction         []Dosage                          `bson:"dosageInstruction" json:"dosageInstruction"`
	DispenseRequest           *MedicationRequestDispenseRequest `bson:"dispenseRequest" json:"dispenseRequest"`
	Substitution              *MedicationRequestSubstitution    `bson:"substitution" json:"substitution"`
	PriorPrescription         *Reference                        `bson:"priorPrescription" json:"priorPrescription"`
	DetectedIssue             []Reference                       `bson:"detectedIssue" json:"detectedIssue"`
	EventHistory              []Reference                       `bson:"eventHistory" json:"eventHistory"`
}
type MedicationRequestRequester struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Agent             Reference   `bson:"agent,omitempty" json:"agent,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf" json:"onBehalfOf"`
}
type MedicationRequestDispenseRequest struct {
	Id                     *string     `bson:"id" json:"id"`
	Extension              []Extension `bson:"extension" json:"extension"`
	ModifierExtension      []Extension `bson:"modifierExtension" json:"modifierExtension"`
	ValidityPeriod         *Period     `bson:"validityPeriod" json:"validityPeriod"`
	NumberOfRepeatsAllowed *int        `bson:"numberOfRepeatsAllowed" json:"numberOfRepeatsAllowed"`
	Quantity               *Quantity   `bson:"quantity" json:"quantity"`
	ExpectedSupplyDuration *Duration   `bson:"expectedSupplyDuration" json:"expectedSupplyDuration"`
	Performer              *Reference  `bson:"performer" json:"performer"`
}
type MedicationRequestSubstitution struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Allowed           bool             `bson:"allowed,omitempty" json:"allowed,omitempty"`
	Reason            *CodeableConcept `bson:"reason" json:"reason"`
}
type OtherMedicationRequest MedicationRequest

// MarshalJSON marshals the given MedicationRequest as JSON into a byte slice
func (r MedicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationRequest: OtherMedicationRequest(r),
		ResourceType:           "MedicationRequest",
	})
}

// UnmarshalMedicationRequest unmarshalls a MedicationRequest.
func UnmarshalMedicationRequest(b []byte) (MedicationRequest, error) {
	var medicationRequest MedicationRequest
	if err := json.Unmarshal(b, &medicationRequest); err != nil {
		return medicationRequest, err
	}
	return medicationRequest, nil
}
