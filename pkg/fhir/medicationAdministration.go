package fhir

import "encoding/json"

// MedicationAdministration is documented here http://hl7.org/fhir/StructureDefinition/MedicationAdministration
type MedicationAdministration struct {
	Id                        *string                             `bson:"id" json:"id"`
	Meta                      *Meta                               `bson:"meta" json:"meta"`
	ImplicitRules             *string                             `bson:"implicitRules" json:"implicitRules"`
	Language                  *string                             `bson:"language" json:"language"`
	Text                      *Narrative                          `bson:"text" json:"text"`
	RawContained              []json.RawMessage                   `bson:"contained" json:"contained"`
	Contained                 []IResource                         `bson:"-" json:"-"`
	Extension                 []Extension                         `bson:"extension" json:"extension"`
	ModifierExtension         []Extension                         `bson:"modifierExtension" json:"modifierExtension"`
	Identifier                []Identifier                        `bson:"identifier" json:"identifier"`
	Definition                []Reference                         `bson:"definition" json:"definition"`
	PartOf                    []Reference                         `bson:"partOf" json:"partOf"`
	Status                    MedicationAdministrationStatus      `bson:"status,omitempty" json:"status,omitempty"`
	Category                  *CodeableConcept                    `bson:"category" json:"category"`
	MedicationCodeableConcept *CodeableConcept                    `bson:"medicationCodeableConcept,omitempty" json:"medicationCodeableConcept,omitempty"`
	MedicationReference       *Reference                          `bson:"medicationReference,omitempty" json:"medicationReference,omitempty"`
	Subject                   Reference                           `bson:"subject,omitempty" json:"subject,omitempty"`
	Context                   *Reference                          `bson:"context" json:"context"`
	SupportingInformation     []Reference                         `bson:"supportingInformation" json:"supportingInformation"`
	EffectiveDateTime         *string                             `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod           *Period                             `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Performer                 []MedicationAdministrationPerformer `bson:"performer" json:"performer"`
	NotGiven                  *bool                               `bson:"notGiven" json:"notGiven"`
	ReasonNotGiven            []CodeableConcept                   `bson:"reasonNotGiven" json:"reasonNotGiven"`
	ReasonCode                []CodeableConcept                   `bson:"reasonCode" json:"reasonCode"`
	ReasonReference           []Reference                         `bson:"reasonReference" json:"reasonReference"`
	Prescription              *Reference                          `bson:"prescription" json:"prescription"`
	Device                    []Reference                         `bson:"device" json:"device"`
	Note                      []Annotation                        `bson:"note" json:"note"`
	Dosage                    *MedicationAdministrationDosage     `bson:"dosage" json:"dosage"`
	EventHistory              []Reference                         `bson:"eventHistory" json:"eventHistory"`
}
type MedicationAdministrationPerformer struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Actor             Reference   `bson:"actor,omitempty" json:"actor,omitempty"`
	OnBehalfOf        *Reference  `bson:"onBehalfOf" json:"onBehalfOf"`
}
type MedicationAdministrationDosage struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Text              *string          `bson:"text" json:"text"`
	Site              *CodeableConcept `bson:"site" json:"site"`
	Route             *CodeableConcept `bson:"route" json:"route"`
	Method            *CodeableConcept `bson:"method" json:"method"`
	Dose              *Quantity        `bson:"dose" json:"dose"`
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
	return json.Marshal(struct {
		OtherMedicationAdministration
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationAdministration: OtherMedicationAdministration(r),
		ResourceType:                  "MedicationAdministration",
	})
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
