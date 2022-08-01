package fhir

import (
	"bytes"
	"encoding/json"
)

// Immunization is documented here http://hl7.org/fhir/StructureDefinition/Immunization
type Immunization struct {
	Id                  *string                            `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string                            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string                            `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative                         `bson:"text,omitempty" json:"text,omitempty"`
	RawContained        []json.RawMessage                  `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained           []IResource                        `bson:"-,omitempty" json:"-,omitempty"`
	Extension           []*Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []*Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []*Identifier                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              string                             `bson:"status,omitempty" json:"status,omitempty"`
	NotGiven            bool                               `bson:"notGiven,omitempty" json:"notGiven,omitempty"`
	VaccineCode         CodeableConcept                    `bson:"vaccineCode,omitempty" json:"vaccineCode,omitempty"`
	Patient             Reference                          `bson:"patient,omitempty" json:"patient,omitempty"`
	Encounter           *Reference                         `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Date                *string                            `bson:"date,omitempty" json:"date,omitempty"`
	PrimarySource       bool                               `bson:"primarySource,omitempty" json:"primarySource,omitempty"`
	ReportOrigin        *CodeableConcept                   `bson:"reportOrigin,omitempty" json:"reportOrigin,omitempty"`
	Location            *Reference                         `bson:"location,omitempty" json:"location,omitempty"`
	Manufacturer        *Reference                         `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	LotNumber           *string                            `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	ExpirationDate      *string                            `bson:"expirationDate,omitempty" json:"expirationDate,omitempty"`
	Site                *CodeableConcept                   `bson:"site,omitempty" json:"site,omitempty"`
	Route               *CodeableConcept                   `bson:"route,omitempty" json:"route,omitempty"`
	DoseQuantity        *Quantity                          `bson:"doseQuantity,omitempty" json:"doseQuantity,omitempty"`
	Practitioner        []*ImmunizationPractitioner        `bson:"practitioner,omitempty" json:"practitioner,omitempty"`
	Note                []*Annotation                      `bson:"note,omitempty" json:"note,omitempty"`
	Explanation         *ImmunizationExplanation           `bson:"explanation,omitempty" json:"explanation,omitempty"`
	Reaction            []*ImmunizationReaction            `bson:"reaction,omitempty" json:"reaction,omitempty"`
	VaccinationProtocol []*ImmunizationVaccinationProtocol `bson:"vaccinationProtocol,omitempty" json:"vaccinationProtocol,omitempty"`
}
type ImmunizationPractitioner struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Actor             Reference        `bson:"actor,omitempty" json:"actor,omitempty"`
}
type ImmunizationExplanation struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Reason            []*CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	ReasonNotGiven    []*CodeableConcept `bson:"reasonNotGiven,omitempty" json:"reasonNotGiven,omitempty"`
}
type ImmunizationReaction struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Date              *string      `bson:"date,omitempty" json:"date,omitempty"`
	Detail            *Reference   `bson:"detail,omitempty" json:"detail,omitempty"`
	Reported          *bool        `bson:"reported,omitempty" json:"reported,omitempty"`
}
type ImmunizationVaccinationProtocol struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	DoseSequence      *int              `bson:"doseSequence,omitempty" json:"doseSequence,omitempty"`
	Description       *string           `bson:"description,omitempty" json:"description,omitempty"`
	Authority         *Reference        `bson:"authority,omitempty" json:"authority,omitempty"`
	Series            *string           `bson:"series,omitempty" json:"series,omitempty"`
	SeriesDoses       *int              `bson:"seriesDoses,omitempty" json:"seriesDoses,omitempty"`
	TargetDisease     []CodeableConcept `bson:"targetDisease,omitempty" json:"targetDisease,omitempty"`
	DoseStatus        CodeableConcept   `bson:"doseStatus,omitempty" json:"doseStatus,omitempty"`
	DoseStatusReason  *CodeableConcept  `bson:"doseStatusReason,omitempty" json:"doseStatusReason,omitempty"`
}

// OtherImmunization is a helper type to use the default implementations of Marshall and Unmarshal
type OtherImmunization Immunization

// MarshalJSON marshals the given Immunization as JSON into a byte slice
func (r Immunization) MarshalJSON() ([]byte, error) {
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
		OtherImmunization
	}{
		OtherImmunization: OtherImmunization(r),
		ResourceType:      "Immunization",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into Immunization
func (r *Immunization) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherImmunization)(r)); err != nil {
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
func (r Immunization) GetResourceType() ResourceType {
	return ResourceTypeImmunization
}
