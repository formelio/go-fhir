package fhir

import "encoding/json"

// MedicationStatement is documented here http://hl7.org/fhir/StructureDefinition/MedicationStatement
type MedicationStatement struct {
	Id                        *string                   `bson:"id" json:"id"`
	Meta                      *Meta                     `bson:"meta" json:"meta"`
	ImplicitRules             *string                   `bson:"implicitRules" json:"implicitRules"`
	Language                  *string                   `bson:"language" json:"language"`
	Text                      *Narrative                `bson:"text" json:"text"`
	RawContained              []json.RawMessage         `bson:"contained" json:"contained"`
	Contained                 []IResource               `bson:"-" json:"-"`
	Extension                 []Extension               `bson:"extension" json:"extension"`
	ModifierExtension         []Extension               `bson:"modifierExtension" json:"modifierExtension"`
	Identifier                []Identifier              `bson:"identifier" json:"identifier"`
	BasedOn                   []Reference               `bson:"basedOn" json:"basedOn"`
	PartOf                    []Reference               `bson:"partOf" json:"partOf"`
	Context                   *Reference                `bson:"context" json:"context"`
	Status                    MedicationStatementStatus `bson:"status,omitempty" json:"status,omitempty"`
	Category                  *CodeableConcept          `bson:"category" json:"category"`
	MedicationCodeableConcept *CodeableConcept          `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference                `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	EffectiveDateTime         *string                   `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod           *Period                   `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	DateAsserted              *string                   `bson:"dateAsserted" json:"dateAsserted"`
	InformationSource         *Reference                `bson:"informationSource" json:"informationSource"`
	Subject                   Reference                 `bson:"subject,omitempty" json:"subject,omitempty"`
	DerivedFrom               []Reference               `bson:"derivedFrom" json:"derivedFrom"`
	Taken                     MedicationStatementTaken  `bson:"taken,omitempty" json:"taken,omitempty"`
	ReasonNotTaken            []CodeableConcept         `bson:"reasonNotTaken" json:"reasonNotTaken"`
	ReasonCode                []CodeableConcept         `bson:"reasonCode" json:"reasonCode"`
	ReasonReference           []Reference               `bson:"reasonReference" json:"reasonReference"`
	Note                      []Annotation              `bson:"note" json:"note"`
	Dosage                    []Dosage                  `bson:"dosage" json:"dosage"`
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
	return json.Marshal(struct {
		OtherMedicationStatement
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationStatement: OtherMedicationStatement(r),
		ResourceType:             "MedicationStatement",
	})
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
