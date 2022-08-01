package fhir

import (
	"bytes"
	"encoding/json"
)

// MedicationDispense is documented here http://hl7.org/fhir/StructureDefinition/MedicationDispense
type MedicationDispense struct {
	Id                           *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta                         *Meta                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules                *string                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                     *string                         `bson:"language,omitempty" json:"language,omitempty"`
	Text                         *Narrative                      `bson:"text,omitempty" json:"text,omitempty"`
	RawContained                 []json.RawMessage               `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained                    []IResource                     `bson:"-,omitempty" json:"-,omitempty"`
	Extension                    []*Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension            []*Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier                   []*Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	PartOf                       []*Reference                    `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                       *MedicationDispenseStatus       `bson:"status,omitempty" json:"status,omitempty"`
	Category                     *CodeableConcept                `bson:"category,omitempty" json:"category,omitempty"`
	MedicationCodeableConcept    *CodeableConcept                `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference          *Reference                      `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	Subject                      *Reference                      `bson:"subject,omitempty" json:"subject,omitempty"`
	Context                      *Reference                      `bson:"context,omitempty" json:"context,omitempty"`
	SupportingInformation        []*Reference                    `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	Performer                    []*MedicationDispensePerformer  `bson:"performer,omitempty" json:"performer,omitempty"`
	AuthorizingPrescription      []*Reference                    `bson:"authorizingPrescription,omitempty" json:"authorizingPrescription,omitempty"`
	Type                         *CodeableConcept                `bson:"type,omitempty" json:"type,omitempty"`
	Quantity                     *Quantity                       `bson:"quantity,omitempty" json:"quantity,omitempty"`
	DaysSupply                   *Quantity                       `bson:"daysSupply,omitempty" json:"daysSupply,omitempty"`
	WhenPrepared                 *string                         `bson:"whenPrepared,omitempty" json:"whenPrepared,omitempty"`
	WhenHandedOver               *string                         `bson:"whenHandedOver,omitempty" json:"whenHandedOver,omitempty"`
	Destination                  *Reference                      `bson:"destination,omitempty" json:"destination,omitempty"`
	Receiver                     []*Reference                    `bson:"receiver,omitempty" json:"receiver,omitempty"`
	Note                         []*Annotation                   `bson:"note,omitempty" json:"note,omitempty"`
	DosageInstruction            []*Dosage                       `bson:"dosageInstruction,omitempty" json:"dosageInstruction,omitempty"`
	Substitution                 *MedicationDispenseSubstitution `bson:"substitution,omitempty" json:"substitution,omitempty"`
	DetectedIssue                []*Reference                    `bson:"detectedIssue,omitempty" json:"detectedIssue,omitempty"`
	NotDone                      *bool                           `bson:"notDone,omitempty" json:"notDone,omitempty"`
	NotDoneReasonCodeableConcept *CodeableConcept                `bson:"notDoneReasonCodeableConcept,omitempty" json:"notDoneReasonCodeableConcept,omitempty"`
	NotDoneReasonReference       *Reference                      `bson:"notDoneReasonReference,omitempty" json:"notDoneReasonReference,omitempty"`
	EventHistory                 []*Reference                    `bson:"eventHistory,omitempty" json:"eventHistory,omitempty"`
}
type MedicationDispensePerformer struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Actor             Reference    `bson:"actor,omitempty" json:"actor,omitempty"`
	OnBehalfOf        *Reference   `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}
type MedicationDispenseSubstitution struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	WasSubstituted    bool               `bson:"wasSubstituted,omitempty" json:"wasSubstituted,omitempty"`
	Type              *CodeableConcept   `bson:"type,omitempty" json:"type,omitempty"`
	Reason            []*CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	ResponsibleParty  []*Reference       `bson:"responsibleParty,omitempty" json:"responsibleParty,omitempty"`
}

// OtherMedicationDispense is a helper type to use the default implementations of Marshall and Unmarshal
type OtherMedicationDispense MedicationDispense

// MarshalJSON marshals the given MedicationDispense as JSON into a byte slice
func (r MedicationDispense) MarshalJSON() ([]byte, error) {
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
		OtherMedicationDispense
	}{
		OtherMedicationDispense: OtherMedicationDispense(r),
		ResourceType:            "MedicationDispense",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into MedicationDispense
func (r *MedicationDispense) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherMedicationDispense)(r)); err != nil {
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
func (r MedicationDispense) GetResourceType() ResourceType {
	return ResourceTypeMedicationDispense
}
