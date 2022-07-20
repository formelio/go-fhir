package fhir

import "encoding/json"

// MedicationDispense is documented here http://hl7.org/fhir/StructureDefinition/MedicationDispense
type MedicationDispense struct {
	Id                      *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta                    *Meta                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules           *string                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                *string                         `bson:"language,omitempty" json:"language,omitempty"`
	Text                    *Narrative                      `bson:"text,omitempty" json:"text,omitempty"`
	Extension               []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier              []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	PartOf                  []Reference                     `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                  *string                         `bson:"status,omitempty" json:"status,omitempty"`
	Category                *CodeableConcept                `bson:"category,omitempty" json:"category,omitempty"`
	SupportingInformation   []Reference                     `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	Performer               []MedicationDispensePerformer   `bson:"performer,omitempty" json:"performer,omitempty"`
	AuthorizingPrescription []Reference                     `bson:"authorizingPrescription,omitempty" json:"authorizingPrescription,omitempty"`
	Type                    *CodeableConcept                `bson:"type,omitempty" json:"type,omitempty"`
	Quantity                *Quantity                       `bson:"quantity,omitempty" json:"quantity,omitempty"`
	DaysSupply              *Quantity                       `bson:"daysSupply,omitempty" json:"daysSupply,omitempty"`
	WhenPrepared            *string                         `bson:"whenPrepared,omitempty" json:"whenPrepared,omitempty"`
	WhenHandedOver          *string                         `bson:"whenHandedOver,omitempty" json:"whenHandedOver,omitempty"`
	Destination             *Reference                      `bson:"destination,omitempty" json:"destination,omitempty"`
	Note                    []Annotation                    `bson:"note,omitempty" json:"note,omitempty"`
	DosageInstruction       []Dosage                        `bson:"dosageInstruction,omitempty" json:"dosageInstruction,omitempty"`
	Substitution            *MedicationDispenseSubstitution `bson:"substitution,omitempty" json:"substitution,omitempty"`
	DetectedIssue           []Reference                     `bson:"detectedIssue,omitempty" json:"detectedIssue,omitempty"`
	NotDone                 *bool                           `bson:"notDone,omitempty" json:"notDone,omitempty"`
	EventHistory            []Reference                     `bson:"eventHistory,omitempty" json:"eventHistory,omitempty"`
}
type MedicationDispensePerformer struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}
type MedicationDispenseSubstitution struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	WasSubstituted    bool              `bson:"wasSubstituted" json:"wasSubstituted"`
	Type              *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Reason            []CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	ResponsibleParty  []Reference       `bson:"responsibleParty,omitempty" json:"responsibleParty,omitempty"`
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

// UnmarshalMedicationDispense unmarshals a MedicationDispense.
func UnmarshalMedicationDispense(b []byte) (MedicationDispense, error) {
	var medicationDispense MedicationDispense
	if err := json.Unmarshal(b, &medicationDispense); err != nil {
		return medicationDispense, err
	}
	return medicationDispense, nil
}
