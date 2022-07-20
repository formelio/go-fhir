package fhir

import "encoding/json"

// Observation is documented here http://hl7.org/fhir/StructureDefinition/Observation
type Observation struct {
	Id                   *string                     `bson:"id" json:"id"`
	Meta                 *Meta                       `bson:"meta" json:"meta"`
	ImplicitRules        *string                     `bson:"implicitRules" json:"implicitRules"`
	Language             *string                     `bson:"language" json:"language"`
	Text                 *Narrative                  `bson:"text" json:"text"`
	Contained            []json.RawMessage           `bson:"contained" json:"contained"`
	Extension            []Extension                 `bson:"extension" json:"extension"`
	ModifierExtension    []Extension                 `bson:"modifierExtension" json:"modifierExtension"`
	Identifier           []Identifier                `bson:"identifier" json:"identifier"`
	BasedOn              []Reference                 `bson:"basedOn" json:"basedOn"`
	Status               ObservationStatus           `bson:"status,omitempty" json:"status,omitempty"`
	Category             []CodeableConcept           `bson:"category" json:"category"`
	Code                 CodeableConcept             `bson:"code,omitempty" json:"code,omitempty"`
	Subject              *Reference                  `bson:"subject" json:"subject"`
	Context              *Reference                  `bson:"context" json:"context"`
	EffectiveDateTime    *string                     `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod      *Period                     `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Issued               *string                     `bson:"issued" json:"issued"`
	Performer            []Reference                 `bson:"performer" json:"performer"`
	ValueQuantity        *Quantity                   `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueCodeableConcept *CodeableConcept            `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueString          *string                     `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueBoolean         *bool                       `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueRange           *Range                      `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio                      `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData                `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueAttachment      *Attachment                 `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueTime            *string                     `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueDateTime        *string                     `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod          *Period                     `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	DataAbsentReason     *CodeableConcept            `bson:"dataAbsentReason" json:"dataAbsentReason"`
	Interpretation       *CodeableConcept            `bson:"interpretation" json:"interpretation"`
	Comment              *string                     `bson:"comment" json:"comment"`
	BodySite             *CodeableConcept            `bson:"bodySite" json:"bodySite"`
	Method               *CodeableConcept            `bson:"method" json:"method"`
	Specimen             *Reference                  `bson:"specimen" json:"specimen"`
	Device               *Reference                  `bson:"device" json:"device"`
	ReferenceRange       []ObservationReferenceRange `bson:"referenceRange" json:"referenceRange"`
	Related              []ObservationRelated        `bson:"related" json:"related"`
	Component            []ObservationComponent      `bson:"component" json:"component"`
}
type ObservationReferenceRange struct {
	Id                *string           `bson:"id" json:"id"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Low               *Quantity         `bson:"low" json:"low"`
	High              *Quantity         `bson:"high" json:"high"`
	Type              *CodeableConcept  `bson:"type" json:"type"`
	AppliesTo         []CodeableConcept `bson:"appliesTo" json:"appliesTo"`
	Age               *Range            `bson:"age" json:"age"`
	Text              *string           `bson:"text" json:"text"`
}
type ObservationRelated struct {
	Id                *string                      `bson:"id" json:"id"`
	Extension         []Extension                  `bson:"extension" json:"extension"`
	ModifierExtension []Extension                  `bson:"modifierExtension" json:"modifierExtension"`
	Type              *ObservationRelationshipType `bson:"type" json:"type"`
	Target            Reference                    `bson:"target,omitempty" json:"target,omitempty"`
}
type ObservationComponent struct {
	Id                   *string                     `bson:"id" json:"id"`
	Extension            []Extension                 `bson:"extension" json:"extension"`
	ModifierExtension    []Extension                 `bson:"modifierExtension" json:"modifierExtension"`
	Code                 CodeableConcept             `bson:"code,omitempty" json:"code,omitempty"`
	ValueQuantity        *Quantity                   `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueCodeableConcept *CodeableConcept            `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueString          *string                     `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueRange           *Range                      `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueRatio           *Ratio                      `bson:"valueRatio,omitempty" json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData                `bson:"valueSampledData,omitempty" json:"valueSampledData,omitempty"`
	ValueAttachment      *Attachment                 `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
	ValueTime            *string                     `bson:"valueTime,omitempty" json:"valueTime,omitempty"`
	ValueDateTime        *string                     `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod          *Period                     `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	DataAbsentReason     *CodeableConcept            `bson:"dataAbsentReason" json:"dataAbsentReason"`
	Interpretation       *CodeableConcept            `bson:"interpretation" json:"interpretation"`
	ReferenceRange       []ObservationReferenceRange `bson:"referenceRange" json:"referenceRange"`
}
type OtherObservation Observation

// MarshalJSON marshals the given Observation as JSON into a byte slice
func (r Observation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherObservation
		ResourceType string `json:"resourceType"`
	}{
		OtherObservation: OtherObservation(r),
		ResourceType:     "Observation",
	})
}

// UnmarshalObservation unmarshalls a Observation.
func UnmarshalObservation(b []byte) (Observation, error) {
	var observation Observation
	if err := json.Unmarshal(b, &observation); err != nil {
		return observation, err
	}
	return observation, nil
}
