package fhir

import (
	"bytes"
	"encoding/json"
)

// Condition is documented here http://hl7.org/fhir/StructureDefinition/Condition
type Condition struct {
	Id                 *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	RawContained       []json.RawMessage            `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained          []IResource                  `bson:"-,omitempty" json:"-,omitempty"`
	Extension          []*Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []*Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         []*Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	ClinicalStatus     *string                      `bson:"clinicalStatus,omitempty" json:"clinicalStatus,omitempty"`
	VerificationStatus *ConditionVerificationStatus `bson:"verificationStatus,omitempty" json:"verificationStatus,omitempty"`
	Category           []*CodeableConcept           `bson:"category,omitempty" json:"category,omitempty"`
	Severity           *CodeableConcept             `bson:"severity,omitempty" json:"severity,omitempty"`
	Code               *CodeableConcept             `bson:"code,omitempty" json:"code,omitempty"`
	BodySite           []*CodeableConcept           `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Subject            Reference                    `bson:"subject,omitempty" json:"subject,omitempty"`
	Context            *Reference                   `bson:"context,omitempty" json:"context,omitempty"`
	OnsetDateTime      *string                      `bson:"onsetDateTime,omitempty" json:"onsetDateTime,omitempty"`
	OnsetAge           *Age                         `bson:"onsetAge,omitempty" json:"onsetAge,omitempty"`
	OnsetPeriod        *Period                      `bson:"onsetPeriod,omitempty" json:"onsetPeriod,omitempty"`
	OnsetRange         *Range                       `bson:"onsetRange,omitempty" json:"onsetRange,omitempty"`
	OnsetString        *string                      `bson:"onsetString,omitempty" json:"onsetString,omitempty"`
	AbatementDateTime  *string                      `bson:"abatementDateTime,omitempty" json:"abatementDateTime,omitempty"`
	AbatementAge       *Age                         `bson:"abatementAge,omitempty" json:"abatementAge,omitempty"`
	AbatementBoolean   *bool                        `bson:"abatementBoolean,omitempty" json:"abatementBoolean,omitempty"`
	AbatementPeriod    *Period                      `bson:"abatementPeriod,omitempty" json:"abatementPeriod,omitempty"`
	AbatementRange     *Range                       `bson:"abatementRange,omitempty" json:"abatementRange,omitempty"`
	AbatementString    *string                      `bson:"abatementString,omitempty" json:"abatementString,omitempty"`
	AssertedDate       *string                      `bson:"assertedDate,omitempty" json:"assertedDate,omitempty"`
	Asserter           *Reference                   `bson:"asserter,omitempty" json:"asserter,omitempty"`
	Stage              *ConditionStage              `bson:"stage,omitempty" json:"stage,omitempty"`
	Evidence           []*ConditionEvidence         `bson:"evidence,omitempty" json:"evidence,omitempty"`
	Note               []*Annotation                `bson:"note,omitempty" json:"note,omitempty"`
}
type ConditionStage struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Summary           *CodeableConcept `bson:"summary,omitempty" json:"summary,omitempty"`
	Assessment        []*Reference     `bson:"assessment,omitempty" json:"assessment,omitempty"`
}
type ConditionEvidence struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              []*CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Detail            []*Reference       `bson:"detail,omitempty" json:"detail,omitempty"`
}

// OtherCondition is a helper type to use the default implementations of Marshall and Unmarshal
type OtherCondition Condition

// MarshalJSON marshals the given Condition as JSON into a byte slice
func (r Condition) MarshalJSON() ([]byte, error) {
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
		OtherCondition
	}{
		OtherCondition: OtherCondition(r),
		ResourceType:   "Condition",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into Condition
func (r *Condition) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherCondition)(r)); err != nil {
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
func (r Condition) GetResourceType() ResourceType {
	return ResourceTypeCondition
}
