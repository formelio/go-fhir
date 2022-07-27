package fhir

import "encoding/json"

// Immunization is documented here http://hl7.org/fhir/StructureDefinition/Immunization
type Immunization struct {
	Id                  *string                           `bson:"id" json:"id"`
	Meta                *Meta                             `bson:"meta" json:"meta"`
	ImplicitRules       *string                           `bson:"implicitRules" json:"implicitRules"`
	Language            *string                           `bson:"language" json:"language"`
	Text                *Narrative                        `bson:"text" json:"text"`
	RawContained        []json.RawMessage                 `bson:"contained" json:"contained"`
	Contained           []IResource                       `bson:"-" json:"-"`
	Extension           []Extension                       `bson:"extension" json:"extension"`
	ModifierExtension   []Extension                       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier          []Identifier                      `bson:"identifier" json:"identifier"`
	Status              string                            `bson:"status,omitempty" json:"status,omitempty"`
	NotGiven            bool                              `bson:"notGiven,omitempty" json:"notGiven,omitempty"`
	VaccineCode         CodeableConcept                   `bson:"vaccineCode,omitempty" json:"vaccineCode,omitempty"`
	Patient             Reference                         `bson:"patient,omitempty" json:"patient,omitempty"`
	Encounter           *Reference                        `bson:"encounter" json:"encounter"`
	Date                *string                           `bson:"date" json:"date"`
	PrimarySource       bool                              `bson:"primarySource,omitempty" json:"primarySource,omitempty"`
	ReportOrigin        *CodeableConcept                  `bson:"reportOrigin" json:"reportOrigin"`
	Location            *Reference                        `bson:"location" json:"location"`
	Manufacturer        *Reference                        `bson:"manufacturer" json:"manufacturer"`
	LotNumber           *string                           `bson:"lotNumber" json:"lotNumber"`
	ExpirationDate      *string                           `bson:"expirationDate" json:"expirationDate"`
	Site                *CodeableConcept                  `bson:"site" json:"site"`
	Route               *CodeableConcept                  `bson:"route" json:"route"`
	DoseQuantity        *Quantity                         `bson:"doseQuantity" json:"doseQuantity"`
	Practitioner        []ImmunizationPractitioner        `bson:"practitioner" json:"practitioner"`
	Note                []Annotation                      `bson:"note" json:"note"`
	Explanation         *ImmunizationExplanation          `bson:"explanation" json:"explanation"`
	Reaction            []ImmunizationReaction            `bson:"reaction" json:"reaction"`
	VaccinationProtocol []ImmunizationVaccinationProtocol `bson:"vaccinationProtocol" json:"vaccinationProtocol"`
}
type ImmunizationPractitioner struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Role              *CodeableConcept `bson:"role" json:"role"`
	Actor             Reference        `bson:"actor,omitempty" json:"actor,omitempty"`
}
type ImmunizationExplanation struct {
	Id                *string           `bson:"id" json:"id"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Reason            []CodeableConcept `bson:"reason" json:"reason"`
	ReasonNotGiven    []CodeableConcept `bson:"reasonNotGiven" json:"reasonNotGiven"`
}
type ImmunizationReaction struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Date              *string     `bson:"date" json:"date"`
	Detail            *Reference  `bson:"detail" json:"detail"`
	Reported          *bool       `bson:"reported" json:"reported"`
}
type ImmunizationVaccinationProtocol struct {
	Id                *string           `bson:"id" json:"id"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	DoseSequence      *int              `bson:"doseSequence" json:"doseSequence"`
	Description       *string           `bson:"description" json:"description"`
	Authority         *Reference        `bson:"authority" json:"authority"`
	Series            *string           `bson:"series" json:"series"`
	SeriesDoses       *int              `bson:"seriesDoses" json:"seriesDoses"`
	TargetDisease     []CodeableConcept `bson:"targetDisease,omitempty" json:"targetDisease,omitempty"`
	DoseStatus        CodeableConcept   `bson:"doseStatus,omitempty" json:"doseStatus,omitempty"`
	DoseStatusReason  *CodeableConcept  `bson:"doseStatusReason" json:"doseStatusReason"`
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
	return json.Marshal(struct {
		OtherImmunization
		ResourceType string `json:"resourceType"`
	}{
		OtherImmunization: OtherImmunization(r),
		ResourceType:      "Immunization",
	})
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
