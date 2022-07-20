package fhir

import "encoding/json"

// AdverseEvent is documented here http://hl7.org/fhir/StructureDefinition/AdverseEvent
type AdverseEvent struct {
	Id                    *string                     `bson:"id" json:"id"`
	Meta                  *Meta                       `bson:"meta" json:"meta"`
	ImplicitRules         *string                     `bson:"implicitRules" json:"implicitRules"`
	Language              *string                     `bson:"language" json:"language"`
	Text                  *Narrative                  `bson:"text" json:"text"`
	Contained             []json.RawMessage           `bson:"contained" json:"contained"`
	Extension             []Extension                 `bson:"extension" json:"extension"`
	ModifierExtension     []Extension                 `bson:"modifierExtension" json:"modifierExtension"`
	Identifier            *Identifier                 `bson:"identifier" json:"identifier"`
	Category              *AdverseEventCategory       `bson:"category" json:"category"`
	Type                  *CodeableConcept            `bson:"type" json:"type"`
	Subject               *Reference                  `bson:"subject" json:"subject"`
	Date                  *string                     `bson:"date" json:"date"`
	Reaction              []Reference                 `bson:"reaction" json:"reaction"`
	Location              *Reference                  `bson:"location" json:"location"`
	Seriousness           *CodeableConcept            `bson:"seriousness" json:"seriousness"`
	Outcome               *CodeableConcept            `bson:"outcome" json:"outcome"`
	Recorder              *Reference                  `bson:"recorder" json:"recorder"`
	EventParticipant      *Reference                  `bson:"eventParticipant" json:"eventParticipant"`
	Description           *string                     `bson:"description" json:"description"`
	SuspectEntity         []AdverseEventSuspectEntity `bson:"suspectEntity" json:"suspectEntity"`
	SubjectMedicalHistory []Reference                 `bson:"subjectMedicalHistory" json:"subjectMedicalHistory"`
	ReferenceDocument     []Reference                 `bson:"referenceDocument" json:"referenceDocument"`
	Study                 []Reference                 `bson:"study" json:"study"`
}
type AdverseEventSuspectEntity struct {
	Id                          *string                `bson:"id" json:"id"`
	Extension                   []Extension            `bson:"extension" json:"extension"`
	ModifierExtension           []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	Instance                    Reference              `bson:"instance,omitempty" json:"instance,omitempty"`
	Causality                   *AdverseEventCausality `bson:"causality" json:"causality"`
	CausalityAssessment         *CodeableConcept       `bson:"causalityAssessment" json:"causalityAssessment"`
	CausalityProductRelatedness *string                `bson:"causalityProductRelatedness" json:"causalityProductRelatedness"`
	CausalityMethod             *CodeableConcept       `bson:"causalityMethod" json:"causalityMethod"`
	CausalityAuthor             *Reference             `bson:"causalityAuthor" json:"causalityAuthor"`
	CausalityResult             *CodeableConcept       `bson:"causalityResult" json:"causalityResult"`
}
type OtherAdverseEvent AdverseEvent

// MarshalJSON marshals the given AdverseEvent as JSON into a byte slice
func (r AdverseEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAdverseEvent
		ResourceType string `json:"resourceType"`
	}{
		OtherAdverseEvent: OtherAdverseEvent(r),
		ResourceType:      "AdverseEvent",
	})
}

// UnmarshalAdverseEvent unmarshalls a AdverseEvent.
func UnmarshalAdverseEvent(b []byte) (AdverseEvent, error) {
	var adverseEvent AdverseEvent
	if err := json.Unmarshal(b, &adverseEvent); err != nil {
		return adverseEvent, err
	}
	return adverseEvent, nil
}
