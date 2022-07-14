package fhir

import "encoding/json"

// MedicationRequest is documented here http://hl7.org/fhir/StructureDefinition/MedicationRequest
type MedicationRequest struct {
	Id                    *string                           `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                           `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative                        `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	GroupIdentifier       *Identifier                       `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	Status                *string                           `bson:"status,omitempty" json:"status,omitempty"`
	Intent                string                            `bson:"intent" json:"intent"`
	Category              *CodeableConcept                  `bson:"category,omitempty" json:"category,omitempty"`
	Priority              *string                           `bson:"priority,omitempty" json:"priority,omitempty"`
	SupportingInformation []Reference                       `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	AuthoredOn            *string                           `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester             *MedicationRequestRequester       `bson:"requester,omitempty" json:"requester,omitempty"`
	Recorder              *Reference                        `bson:"recorder,omitempty" json:"recorder,omitempty"`
	ReasonCode            []CodeableConcept                 `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	Note                  []Annotation                      `bson:"note,omitempty" json:"note,omitempty"`
	DosageInstruction     []Dosage                          `bson:"dosageInstruction,omitempty" json:"dosageInstruction,omitempty"`
	DispenseRequest       *MedicationRequestDispenseRequest `bson:"dispenseRequest,omitempty" json:"dispenseRequest,omitempty"`
	Substitution          *MedicationRequestSubstitution    `bson:"substitution,omitempty" json:"substitution,omitempty"`
	PriorPrescription     *Reference                        `bson:"priorPrescription,omitempty" json:"priorPrescription,omitempty"`
	DetectedIssue         []Reference                       `bson:"detectedIssue,omitempty" json:"detectedIssue,omitempty"`
	EventHistory          []Reference                       `bson:"eventHistory,omitempty" json:"eventHistory,omitempty"`
}
type MedicationRequestRequester struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}
type MedicationRequestDispenseRequest struct {
	Id                     *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension              []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ValidityPeriod         *Period     `bson:"validityPeriod,omitempty" json:"validityPeriod,omitempty"`
	NumberOfRepeatsAllowed *int        `bson:"numberOfRepeatsAllowed,omitempty" json:"numberOfRepeatsAllowed,omitempty"`
	Quantity               *Quantity   `bson:"quantity,omitempty" json:"quantity,omitempty"`
	ExpectedSupplyDuration *Duration   `bson:"expectedSupplyDuration,omitempty" json:"expectedSupplyDuration,omitempty"`
	Performer              *Reference  `bson:"performer,omitempty" json:"performer,omitempty"`
}
type MedicationRequestSubstitution struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Allowed           bool             `bson:"allowed" json:"allowed"`
	Reason            *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
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

// UnmarshalMedicationRequest unmarshals a MedicationRequest.
func UnmarshalMedicationRequest(b []byte) (MedicationRequest, error) {
	var medicationRequest MedicationRequest
	if err := json.Unmarshal(b, &medicationRequest); err != nil {
		return medicationRequest, err
	}
	return medicationRequest, nil
}
