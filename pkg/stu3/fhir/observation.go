package fhir

import "encoding/json"

// Observation is documented here http://hl7.org/fhir/StructureDefinition/Observation
type Observation struct {
	Id                   *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	RawContained         []json.RawMessage           `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained            []IResource                 `bson:"-,omitempty" json:"-,omitempty"`
	Extension            []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn              []Reference                 `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Status               ObservationStatus           `bson:"status,omitempty" json:"status,omitempty"`
	Category             []CodeableConcept           `bson:"category,omitempty" json:"category,omitempty"`
	Code                 CodeableConcept             `bson:"code,omitempty" json:"code,omitempty"`
	Subject              *Reference                  `bson:"subject,omitempty" json:"subject,omitempty"`
	Context              *Reference                  `bson:"context,omitempty" json:"context,omitempty"`
	EffectiveDateTime    *string                     `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod      *Period                     `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Issued               *string                     `bson:"issued,omitempty" json:"issued,omitempty"`
	Performer            []Reference                 `bson:"performer,omitempty" json:"performer,omitempty"`
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
	DataAbsentReason     *CodeableConcept            `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	Interpretation       *CodeableConcept            `bson:"interpretation,omitempty" json:"interpretation,omitempty"`
	Comment              *string                     `bson:"comment,omitempty" json:"comment,omitempty"`
	BodySite             *CodeableConcept            `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Method               *CodeableConcept            `bson:"method,omitempty" json:"method,omitempty"`
	Specimen             *Reference                  `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Device               *Reference                  `bson:"device,omitempty" json:"device,omitempty"`
	ReferenceRange       []ObservationReferenceRange `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
	Related              []ObservationRelated        `bson:"related,omitempty" json:"related,omitempty"`
	Component            []ObservationComponent      `bson:"component,omitempty" json:"component,omitempty"`
}
type ObservationReferenceRange struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Low               *Quantity         `bson:"low,omitempty" json:"low,omitempty"`
	High              *Quantity         `bson:"high,omitempty" json:"high,omitempty"`
	Type              *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	AppliesTo         []CodeableConcept `bson:"appliesTo,omitempty" json:"appliesTo,omitempty"`
	Age               *Range            `bson:"age,omitempty" json:"age,omitempty"`
	Text              *string           `bson:"text,omitempty" json:"text,omitempty"`
}
type ObservationRelated struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *ObservationRelationshipType `bson:"type,omitempty" json:"type,omitempty"`
	Target            Reference                    `bson:"target,omitempty" json:"target,omitempty"`
}
type ObservationComponent struct {
	Id                   *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
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
	DataAbsentReason     *CodeableConcept            `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	Interpretation       *CodeableConcept            `bson:"interpretation,omitempty" json:"interpretation,omitempty"`
	ReferenceRange       []ObservationReferenceRange `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
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
