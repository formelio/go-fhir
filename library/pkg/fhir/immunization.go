package fhir

import "encoding/json"

// Immunization is documented here http://hl7.org/fhir/StructureDefinition/Immunization
type Immunization struct {
	Id                  *string                           `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string                           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string                           `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative                        `bson:"text,omitempty" json:"text,omitempty"`
	Extension           []Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []Identifier                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              string                            `bson:"status" json:"status"`
	NotGiven            bool                              `bson:"notGiven" json:"notGiven"`
	VaccineCode         CodeableConcept                   `bson:"vaccineCode" json:"vaccineCode"`
	Patient             Reference                         `bson:"patient" json:"patient"`
	Encounter           *Reference                        `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Date                *string                           `bson:"date,omitempty" json:"date,omitempty"`
	PrimarySource       bool                              `bson:"primarySource" json:"primarySource"`
	ReportOrigin        *CodeableConcept                  `bson:"reportOrigin,omitempty" json:"reportOrigin,omitempty"`
	Location            *Reference                        `bson:"location,omitempty" json:"location,omitempty"`
	Manufacturer        *Reference                        `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	LotNumber           *string                           `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	ExpirationDate      *string                           `bson:"expirationDate,omitempty" json:"expirationDate,omitempty"`
	Site                *CodeableConcept                  `bson:"site,omitempty" json:"site,omitempty"`
	Route               *CodeableConcept                  `bson:"route,omitempty" json:"route,omitempty"`
	DoseQuantity        *Quantity                         `bson:"doseQuantity,omitempty" json:"doseQuantity,omitempty"`
	Practitioner        []ImmunizationPractitioner        `bson:"practitioner,omitempty" json:"practitioner,omitempty"`
	Note                []Annotation                      `bson:"note,omitempty" json:"note,omitempty"`
	Explanation         *ImmunizationExplanation          `bson:"explanation,omitempty" json:"explanation,omitempty"`
	Reaction            []ImmunizationReaction            `bson:"reaction,omitempty" json:"reaction,omitempty"`
	VaccinationProtocol []ImmunizationVaccinationProtocol `bson:"vaccinationProtocol,omitempty" json:"vaccinationProtocol,omitempty"`
}
type ImmunizationPractitioner struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Actor             Reference        `bson:"actor" json:"actor"`
}
type ImmunizationExplanation struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Reason            []CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	ReasonNotGiven    []CodeableConcept `bson:"reasonNotGiven,omitempty" json:"reasonNotGiven,omitempty"`
}
type ImmunizationReaction struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Date              *string     `bson:"date,omitempty" json:"date,omitempty"`
	Detail            *Reference  `bson:"detail,omitempty" json:"detail,omitempty"`
	Reported          *bool       `bson:"reported,omitempty" json:"reported,omitempty"`
}
type ImmunizationVaccinationProtocol struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	DoseSequence      *int              `bson:"doseSequence,omitempty" json:"doseSequence,omitempty"`
	Description       *string           `bson:"description,omitempty" json:"description,omitempty"`
	Authority         *Reference        `bson:"authority,omitempty" json:"authority,omitempty"`
	Series            *string           `bson:"series,omitempty" json:"series,omitempty"`
	SeriesDoses       *int              `bson:"seriesDoses,omitempty" json:"seriesDoses,omitempty"`
	TargetDisease     []CodeableConcept `bson:"targetDisease" json:"targetDisease"`
	DoseStatus        CodeableConcept   `bson:"doseStatus" json:"doseStatus"`
	DoseStatusReason  *CodeableConcept  `bson:"doseStatusReason,omitempty" json:"doseStatusReason,omitempty"`
}
type OtherImmunization Immunization

// MarshalJSON marshals the given Immunization as JSON into a byte slice
func (r Immunization) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImmunization
		ResourceType string `json:"resourceType"`
	}{
		OtherImmunization: OtherImmunization(r),
		ResourceType:      "Immunization",
	})
}

// UnmarshalImmunization unmarshals a Immunization.
func UnmarshalImmunization(b []byte) (Immunization, error) {
	var immunization Immunization
	if err := json.Unmarshal(b, &immunization); err != nil {
		return immunization, err
	}
	return immunization, nil
}
