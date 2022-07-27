package fhir

import "encoding/json"

// Condition is documented here http://hl7.org/fhir/StructureDefinition/Condition
type Condition struct {
	Id                 *string                      `bson:"id" json:"id"`
	Meta               *Meta                        `bson:"meta" json:"meta"`
	ImplicitRules      *string                      `bson:"implicitRules" json:"implicitRules"`
	Language           *string                      `bson:"language" json:"language"`
	Text               *Narrative                   `bson:"text" json:"text"`
	RawContained       []json.RawMessage            `bson:"contained" json:"contained"`
	Contained          []IResource                  `bson:"-" json:"-"`
	Extension          []Extension                  `bson:"extension" json:"extension"`
	ModifierExtension  []Extension                  `bson:"modifierExtension" json:"modifierExtension"`
	Identifier         []Identifier                 `bson:"identifier" json:"identifier"`
	ClinicalStatus     *string                      `bson:"clinicalStatus" json:"clinicalStatus"`
	VerificationStatus *ConditionVerificationStatus `bson:"verificationStatus" json:"verificationStatus"`
	Category           []CodeableConcept            `bson:"category" json:"category"`
	Severity           *CodeableConcept             `bson:"severity" json:"severity"`
	Code               *CodeableConcept             `bson:"code" json:"code"`
	BodySite           []CodeableConcept            `bson:"bodySite" json:"bodySite"`
	Subject            Reference                    `bson:"subject,omitempty" json:"subject,omitempty"`
	Context            *Reference                   `bson:"context" json:"context"`
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
	AssertedDate       *string                      `bson:"assertedDate" json:"assertedDate"`
	Asserter           *Reference                   `bson:"asserter" json:"asserter"`
	Stage              *ConditionStage              `bson:"stage" json:"stage"`
	Evidence           []ConditionEvidence          `bson:"evidence" json:"evidence"`
	Note               []Annotation                 `bson:"note" json:"note"`
}
type ConditionStage struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Summary           *CodeableConcept `bson:"summary" json:"summary"`
	Assessment        []Reference      `bson:"assessment" json:"assessment"`
}
type ConditionEvidence struct {
	Id                *string           `bson:"id" json:"id"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Code              []CodeableConcept `bson:"code" json:"code"`
	Detail            []Reference       `bson:"detail" json:"detail"`
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
	return json.Marshal(struct {
		OtherCondition
		ResourceType string `json:"resourceType"`
	}{
		OtherCondition: OtherCondition(r),
		ResourceType:   "Condition",
	})
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
