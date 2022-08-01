package fhir

import (
	"bytes"
	"encoding/json"
)

// MedicationStatement is documented here http://hl7.org/fhir/StructureDefinition/MedicationStatement
type MedicationStatement struct {
	Id                        *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta                      *Meta                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules             *string                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                  *string                   `bson:"language,omitempty" json:"language,omitempty"`
	Text                      *Narrative                `bson:"text,omitempty" json:"text,omitempty"`
	RawContained              []json.RawMessage         `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained                 []IResource               `bson:"-,omitempty" json:"-,omitempty"`
	Extension                 []*Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []*Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier                []*Identifier             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn                   []*Reference              `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf                    []*Reference              `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Context                   *Reference                `bson:"context,omitempty" json:"context,omitempty"`
	Status                    MedicationStatementStatus `bson:"status,omitempty" json:"status,omitempty"`
	Category                  *CodeableConcept          `bson:"category,omitempty" json:"category,omitempty"`
	MedicationCodeableConcept *CodeableConcept          `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference                `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	EffectiveDateTime         *string                   `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod           *Period                   `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	DateAsserted              *string                   `bson:"dateAsserted,omitempty" json:"dateAsserted,omitempty"`
	InformationSource         *Reference                `bson:"informationSource,omitempty" json:"informationSource,omitempty"`
	Subject                   Reference                 `bson:"subject,omitempty" json:"subject,omitempty"`
	DerivedFrom               []*Reference              `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	Taken                     MedicationStatementTaken  `bson:"taken,omitempty" json:"taken,omitempty"`
	ReasonNotTaken            []*CodeableConcept        `bson:"reasonNotTaken,omitempty" json:"reasonNotTaken,omitempty"`
	ReasonCode                []*CodeableConcept        `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference           []*Reference              `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note                      []*Annotation             `bson:"note,omitempty" json:"note,omitempty"`
	Dosage                    []*Dosage                 `bson:"dosage,omitempty" json:"dosage,omitempty"`
}

// OtherMedicationStatement is a helper type to use the default implementations of Marshall and Unmarshal
type OtherMedicationStatement MedicationStatement

// MarshalJSON marshals the given MedicationStatement as JSON into a byte slice
func (r MedicationStatement) MarshalJSON() ([]byte, error) {
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
		OtherMedicationStatement
	}{
		OtherMedicationStatement: OtherMedicationStatement(r),
		ResourceType:             "MedicationStatement",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into MedicationStatement
func (r *MedicationStatement) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherMedicationStatement)(r)); err != nil {
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
func (r MedicationStatement) GetResourceType() ResourceType {
	return ResourceTypeMedicationStatement
}
