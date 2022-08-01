package fhir

import (
	"bytes"
	"encoding/json"
)

// MedicationAdministration is documented here http://hl7.org/fhir/StructureDefinition/MedicationAdministration
type MedicationAdministration struct {
	Id                        *string                              `bson:"id,omitempty" json:"id,omitempty"`
	Meta                      *Meta                                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules             *string                              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                  *string                              `bson:"language,omitempty" json:"language,omitempty"`
	Text                      *Narrative                           `bson:"text,omitempty" json:"text,omitempty"`
	RawContained              []json.RawMessage                    `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained                 []IResource                          `bson:"-,omitempty" json:"-,omitempty"`
	Extension                 []*Extension                         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []*Extension                         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier                []*Identifier                        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Definition                []*Reference                         `bson:"definition,omitempty" json:"definition,omitempty"`
	PartOf                    []*Reference                         `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                    MedicationAdministrationStatus       `bson:"status,omitempty" json:"status,omitempty"`
	Category                  *CodeableConcept                     `bson:"category,omitempty" json:"category,omitempty"`
	MedicationCodeableConcept *CodeableConcept                     `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference                           `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	Subject                   Reference                            `bson:"subject,omitempty" json:"subject,omitempty"`
	Context                   *Reference                           `bson:"context,omitempty" json:"context,omitempty"`
	SupportingInformation     []*Reference                         `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	EffectiveDateTime         *string                              `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod           *Period                              `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Performer                 []*MedicationAdministrationPerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	NotGiven                  *bool                                `bson:"notGiven,omitempty" json:"notGiven,omitempty"`
	ReasonNotGiven            []*CodeableConcept                   `bson:"reasonNotGiven,omitempty" json:"reasonNotGiven,omitempty"`
	ReasonCode                []*CodeableConcept                   `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference           []*Reference                         `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Prescription              *Reference                           `bson:"prescription,omitempty" json:"prescription,omitempty"`
	Device                    []*Reference                         `bson:"device,omitempty" json:"device,omitempty"`
	Note                      []*Annotation                        `bson:"note,omitempty" json:"note,omitempty"`
	Dosage                    *MedicationAdministrationDosage      `bson:"dosage,omitempty" json:"dosage,omitempty"`
	EventHistory              []*Reference                         `bson:"eventHistory,omitempty" json:"eventHistory,omitempty"`
}
type MedicationAdministrationPerformer struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Actor             Reference    `bson:"actor,omitempty" json:"actor,omitempty"`
	OnBehalfOf        *Reference   `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}
type MedicationAdministrationDosage struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Text              *string          `bson:"text,omitempty" json:"text,omitempty"`
	Site              *CodeableConcept `bson:"site,omitempty" json:"site,omitempty"`
	Route             *CodeableConcept `bson:"route,omitempty" json:"route,omitempty"`
	Method            *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	Dose              *Quantity        `bson:"dose,omitempty" json:"dose,omitempty"`
	RateRatio         *Ratio           `bson:"rateRatio,omitempty" json:"rateRatio,omitempty"`
	RateQuantity      *Quantity        `bson:"rateQuantity,omitempty" json:"rateQuantity,omitempty"`
}

// OtherMedicationAdministration is a helper type to use the default implementations of Marshall and Unmarshal
type OtherMedicationAdministration MedicationAdministration

// MarshalJSON marshals the given MedicationAdministration as JSON into a byte slice
func (r MedicationAdministration) MarshalJSON() ([]byte, error) {
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
		OtherMedicationAdministration
	}{
		OtherMedicationAdministration: OtherMedicationAdministration(r),
		ResourceType:                  "MedicationAdministration",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into MedicationAdministration
func (r *MedicationAdministration) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherMedicationAdministration)(r)); err != nil {
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
func (r MedicationAdministration) GetResourceType() ResourceType {
	return ResourceTypeMedicationAdministration
}
