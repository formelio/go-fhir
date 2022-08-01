package fhir

import "encoding/json"

// MedicationRequest is documented here http://hl7.org/fhir/StructureDefinition/MedicationRequest
type MedicationRequest struct {
	Id                        *string                           `bson:"id,omitempty" json:"id,omitempty"`
	Meta                      *Meta                             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules             *string                           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                  *string                           `bson:"language,omitempty" json:"language,omitempty"`
	Text                      *Narrative                        `bson:"text,omitempty" json:"text,omitempty"`
	RawContained              []json.RawMessage                 `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained                 []IResource                       `bson:"-,omitempty" json:"-,omitempty"`
	Extension                 []Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier                []Identifier                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Definition                []Reference                       `bson:"definition,omitempty" json:"definition,omitempty"`
	BasedOn                   []Reference                       `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	GroupIdentifier           *Identifier                       `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	Status                    *MedicationRequestStatus          `bson:"status,omitempty" json:"status,omitempty"`
	Intent                    MedicationRequestIntent           `bson:"intent,omitempty" json:"intent,omitempty"`
	Category                  *CodeableConcept                  `bson:"category,omitempty" json:"category,omitempty"`
	Priority                  *MedicationRequestPriority        `bson:"priority,omitempty" json:"priority,omitempty"`
	MedicationCodeableConcept *CodeableConcept                  `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference                        `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	Subject                   Reference                         `bson:"subject,omitempty" json:"subject,omitempty"`
	Context                   *Reference                        `bson:"context,omitempty" json:"context,omitempty"`
	SupportingInformation     []Reference                       `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	AuthoredOn                *string                           `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester                 *MedicationRequestRequester       `bson:"requester,omitempty" json:"requester,omitempty"`
	Recorder                  *Reference                        `bson:"recorder,omitempty" json:"recorder,omitempty"`
	ReasonCode                []CodeableConcept                 `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference           []Reference                       `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note                      []Annotation                      `bson:"note,omitempty" json:"note,omitempty"`
	DosageInstruction         []Dosage                          `bson:"dosageInstruction,omitempty" json:"dosageInstruction,omitempty"`
	DispenseRequest           *MedicationRequestDispenseRequest `bson:"dispenseRequest,omitempty" json:"dispenseRequest,omitempty"`
	Substitution              *MedicationRequestSubstitution    `bson:"substitution,omitempty" json:"substitution,omitempty"`
	PriorPrescription         *Reference                        `bson:"priorPrescription,omitempty" json:"priorPrescription,omitempty"`
	DetectedIssue             []Reference                       `bson:"detectedIssue,omitempty" json:"detectedIssue,omitempty"`
	EventHistory              []Reference                       `bson:"eventHistory,omitempty" json:"eventHistory,omitempty"`
}
type MedicationRequestRequester struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Agent             Reference   `bson:"agent,omitempty" json:"agent,omitempty"`
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
	Allowed           bool             `bson:"allowed,omitempty" json:"allowed,omitempty"`
	Reason            *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
}

// OtherMedicationRequest is a helper type to use the default implementations of Marshall and Unmarshal
type OtherMedicationRequest MedicationRequest

// MarshalJSON marshals the given MedicationRequest as JSON into a byte slice
func (r MedicationRequest) MarshalJSON() ([]byte, error) {
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
		OtherMedicationRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationRequest: OtherMedicationRequest(r),
		ResourceType:           "MedicationRequest",
	})
}

// UnmarshalJSON unmarshals the given byte slice into MedicationRequest
func (r *MedicationRequest) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherMedicationRequest)(r)); err != nil {
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
func (r MedicationRequest) GetResourceType() ResourceType {
	return ResourceTypeMedicationRequest
}
