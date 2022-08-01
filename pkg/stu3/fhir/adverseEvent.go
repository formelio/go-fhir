package fhir

import "encoding/json"

// AdverseEvent is documented here http://hl7.org/fhir/StructureDefinition/AdverseEvent
type AdverseEvent struct {
	Id                    *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	RawContained          []json.RawMessage           `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained             []IResource                 `bson:"-,omitempty" json:"-,omitempty"`
	Extension             []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            *Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Category              *AdverseEventCategory       `bson:"category,omitempty" json:"category,omitempty"`
	Type                  *CodeableConcept            `bson:"type,omitempty" json:"type,omitempty"`
	Subject               *Reference                  `bson:"subject,omitempty" json:"subject,omitempty"`
	Date                  *string                     `bson:"date,omitempty" json:"date,omitempty"`
	Reaction              []Reference                 `bson:"reaction,omitempty" json:"reaction,omitempty"`
	Location              *Reference                  `bson:"location,omitempty" json:"location,omitempty"`
	Seriousness           *CodeableConcept            `bson:"seriousness,omitempty" json:"seriousness,omitempty"`
	Outcome               *CodeableConcept            `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Recorder              *Reference                  `bson:"recorder,omitempty" json:"recorder,omitempty"`
	EventParticipant      *Reference                  `bson:"eventParticipant,omitempty" json:"eventParticipant,omitempty"`
	Description           *string                     `bson:"description,omitempty" json:"description,omitempty"`
	SuspectEntity         []AdverseEventSuspectEntity `bson:"suspectEntity,omitempty" json:"suspectEntity,omitempty"`
	SubjectMedicalHistory []Reference                 `bson:"subjectMedicalHistory,omitempty" json:"subjectMedicalHistory,omitempty"`
	ReferenceDocument     []Reference                 `bson:"referenceDocument,omitempty" json:"referenceDocument,omitempty"`
	Study                 []Reference                 `bson:"study,omitempty" json:"study,omitempty"`
}
type AdverseEventSuspectEntity struct {
	Id                          *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension                   []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension           []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Instance                    Reference              `bson:"instance,omitempty" json:"instance,omitempty"`
	Causality                   *AdverseEventCausality `bson:"causality,omitempty" json:"causality,omitempty"`
	CausalityAssessment         *CodeableConcept       `bson:"causalityAssessment,omitempty" json:"causalityAssessment,omitempty"`
	CausalityProductRelatedness *string                `bson:"causalityProductRelatedness,omitempty" json:"causalityProductRelatedness,omitempty"`
	CausalityMethod             *CodeableConcept       `bson:"causalityMethod,omitempty" json:"causalityMethod,omitempty"`
	CausalityAuthor             *Reference             `bson:"causalityAuthor,omitempty" json:"causalityAuthor,omitempty"`
	CausalityResult             *CodeableConcept       `bson:"causalityResult,omitempty" json:"causalityResult,omitempty"`
}

// OtherAdverseEvent is a helper type to use the default implementations of Marshall and Unmarshal
type OtherAdverseEvent AdverseEvent

// MarshalJSON marshals the given AdverseEvent as JSON into a byte slice
func (r AdverseEvent) MarshalJSON() ([]byte, error) {
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
		OtherAdverseEvent
		ResourceType string `json:"resourceType"`
	}{
		OtherAdverseEvent: OtherAdverseEvent(r),
		ResourceType:      "AdverseEvent",
	})
}

// UnmarshalJSON unmarshals the given byte slice into AdverseEvent
func (r *AdverseEvent) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherAdverseEvent)(r)); err != nil {
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
func (r AdverseEvent) GetResourceType() ResourceType {
	return ResourceTypeAdverseEvent
}
