package fhir

import "encoding/json"

// MedicationDispense is documented here http://hl7.org/fhir/StructureDefinition/MedicationDispense
type MedicationDispense struct {
	Id                           *string                         `bson:"id" json:"id"`
	Meta                         *Meta                           `bson:"meta" json:"meta"`
	ImplicitRules                *string                         `bson:"implicitRules" json:"implicitRules"`
	Language                     *string                         `bson:"language" json:"language"`
	Text                         *Narrative                      `bson:"text" json:"text"`
	Contained                    []json.RawMessage               `bson:"contained" json:"contained"`
	Extension                    []Extension                     `bson:"extension" json:"extension"`
	ModifierExtension            []Extension                     `bson:"modifierExtension" json:"modifierExtension"`
	Identifier                   []Identifier                    `bson:"identifier" json:"identifier"`
	PartOf                       []Reference                     `bson:"partOf" json:"partOf"`
	Status                       *MedicationDispenseStatus       `bson:"status" json:"status"`
	Category                     *CodeableConcept                `bson:"category" json:"category"`
	MedicationCodeableConcept    *CodeableConcept                `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference          *Reference                      `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	Subject                      *Reference                      `bson:"subject" json:"subject"`
	Context                      *Reference                      `bson:"context" json:"context"`
	SupportingInformation        []Reference                     `bson:"supportingInformation" json:"supportingInformation"`
	Performer                    []MedicationDispensePerformer   `bson:"performer" json:"performer"`
	AuthorizingPrescription      []Reference                     `bson:"authorizingPrescription" json:"authorizingPrescription"`
	Type                         *CodeableConcept                `bson:"type" json:"type"`
	Quantity                     *Quantity                       `bson:"quantity" json:"quantity"`
	DaysSupply                   *Quantity                       `bson:"daysSupply" json:"daysSupply"`
	WhenPrepared                 *string                         `bson:"whenPrepared" json:"whenPrepared"`
	WhenHandedOver               *string                         `bson:"whenHandedOver" json:"whenHandedOver"`
	Destination                  *Reference                      `bson:"destination" json:"destination"`
	Receiver                     []Reference                     `bson:"receiver" json:"receiver"`
	Note                         []Annotation                    `bson:"note" json:"note"`
	DosageInstruction            []Dosage                        `bson:"dosageInstruction" json:"dosageInstruction"`
	Substitution                 *MedicationDispenseSubstitution `bson:"substitution" json:"substitution"`
	DetectedIssue                []Reference                     `bson:"detectedIssue" json:"detectedIssue"`
	NotDone                      *bool                           `bson:"notDone" json:"notDone"`
	NotDoneReasonCodeableConcept *CodeableConcept                `bson:"notDoneReasonCodeableConcept,omitempty" json:"notDoneReasonCodeableConcept,omitempty"`
	NotDoneReasonReference       *Reference                      `bson:"notDoneReasonReference,omitempty" json:"notDoneReasonReference,omitempty"`
	EventHistory                 []Reference                     `bson:"eventHistory" json:"eventHistory"`
}
type MedicationDispensePerformer struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Actor             Reference   `bson:"actor,omitempty" json:"actor,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf" json:"onBehalfOf"`
}
type MedicationDispenseSubstitution struct {
	Id                *string           `bson:"id" json:"id"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	WasSubstituted    bool              `bson:"wasSubstituted,omitempty" json:"wasSubstituted,omitempty"`
	Type              *CodeableConcept  `bson:"type" json:"type"`
	Reason            []CodeableConcept `bson:"reason" json:"reason"`
	ResponsibleParty  []Reference       `bson:"responsibleParty" json:"responsibleParty"`
}
type OtherMedicationDispense MedicationDispense

// MarshalJSON marshals the given MedicationDispense as JSON into a byte slice
func (r MedicationDispense) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationDispense
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationDispense: OtherMedicationDispense(r),
		ResourceType:            "MedicationDispense",
	})
}

// UnmarshalMedicationDispense unmarshalls a MedicationDispense.
func UnmarshalMedicationDispense(b []byte) (MedicationDispense, error) {
	var medicationDispense MedicationDispense
	if err := json.Unmarshal(b, &medicationDispense); err != nil {
		return medicationDispense, err
	}
	return medicationDispense, nil
}
