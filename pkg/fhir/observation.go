package fhir

import "encoding/json"

// Observation is documented here http://hl7.org/fhir/StructureDefinition/Observation
type Observation struct {
	Id                   *string                     `bson:"id" json:"id"`
	Meta                 *Meta                       `bson:"meta" json:"meta"`
	ImplicitRules        *string                     `bson:"implicitRules" json:"implicitRules"`
	Language             *string                     `bson:"language" json:"language"`
	Text                 *Narrative                  `bson:"text" json:"text"`
	RawContained         []json.RawMessage           `bson:"contained" json:"contained"`
	Contained            []IResource                 `bson:"-" json:"-"`
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

// OtherObservation is a helper type to use the default implementations of Marshall and Unmarshal
type OtherObservation Observation

// MarshalJSON marshals the given Observation as JSON into a byte slice
func (r Observation) MarshalJSON() ([]byte, error) {
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
		OtherObservation
		ResourceType string `json:"resourceType"`
	}{
		OtherObservation: OtherObservation(r),
		ResourceType:     "Observation",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Observation
func (r *Observation) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherObservation)(r)); err != nil {
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
func (r Observation) GetResourceType() ResourceType {
	return ResourceTypeObservation
}
