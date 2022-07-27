package fhir

import "encoding/json"

// ActivityDefinition is documented here http://hl7.org/fhir/StructureDefinition/ActivityDefinition
type ActivityDefinition struct {
	Id                     *string                          `bson:"id" json:"id"`
	Meta                   *Meta                            `bson:"meta" json:"meta"`
	ImplicitRules          *string                          `bson:"implicitRules" json:"implicitRules"`
	Language               *string                          `bson:"language" json:"language"`
	Text                   *Narrative                       `bson:"text" json:"text"`
	RawContained           []json.RawMessage                `bson:"contained" json:"contained"`
	Contained              []IResource                      `bson:"-" json:"-"`
	Extension              []Extension                      `bson:"extension" json:"extension"`
	ModifierExtension      []Extension                      `bson:"modifierExtension" json:"modifierExtension"`
	Url                    *string                          `bson:"url" json:"url"`
	Identifier             []Identifier                     `bson:"identifier" json:"identifier"`
	Version                *string                          `bson:"version" json:"version"`
	Name                   *string                          `bson:"name" json:"name"`
	Title                  *string                          `bson:"title" json:"title"`
	Status                 PublicationStatus                `bson:"status,omitempty" json:"status,omitempty"`
	Experimental           *bool                            `bson:"experimental" json:"experimental"`
	Date                   *string                          `bson:"date" json:"date"`
	Publisher              *string                          `bson:"publisher" json:"publisher"`
	Description            *string                          `bson:"description" json:"description"`
	Purpose                *string                          `bson:"purpose" json:"purpose"`
	Usage                  *string                          `bson:"usage" json:"usage"`
	ApprovalDate           *string                          `bson:"approvalDate" json:"approvalDate"`
	LastReviewDate         *string                          `bson:"lastReviewDate" json:"lastReviewDate"`
	EffectivePeriod        *Period                          `bson:"effectivePeriod" json:"effectivePeriod"`
	UseContext             []UsageContext                   `bson:"useContext" json:"useContext"`
	Jurisdiction           []CodeableConcept                `bson:"jurisdiction" json:"jurisdiction"`
	Topic                  []CodeableConcept                `bson:"topic" json:"topic"`
	Contributor            []Contributor                    `bson:"contributor" json:"contributor"`
	Contact                []ContactDetail                  `bson:"contact" json:"contact"`
	Copyright              *string                          `bson:"copyright" json:"copyright"`
	RelatedArtifact        []RelatedArtifact                `bson:"relatedArtifact" json:"relatedArtifact"`
	Library                []Reference                      `bson:"library" json:"library"`
	Kind                   *ResourceType                    `bson:"kind" json:"kind"`
	Code                   *CodeableConcept                 `bson:"code" json:"code"`
	TimingTiming           *Timing                          `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	TimingDateTime         *string                          `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty"`
	TimingPeriod           *Period                          `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	TimingRange            *Range                           `bson:"timingRange,omitempty" json:"timingRange,omitempty"`
	Location               *Reference                       `bson:"location" json:"location"`
	Participant            []ActivityDefinitionParticipant  `bson:"participant" json:"participant"`
	ProductReference       *Reference                       `bson:"productReference,omitempty" json:"productReference,omitempty"`
	ProductCodeableConcept *CodeableConcept                 `bson:"productCodeableConcept,omitempty" json:"productCodeableConcept,omitempty"`
	Quantity               *Quantity                        `bson:"quantity" json:"quantity"`
	Dosage                 []Dosage                         `bson:"dosage" json:"dosage"`
	BodySite               []CodeableConcept                `bson:"bodySite" json:"bodySite"`
	Transform              *Reference                       `bson:"transform" json:"transform"`
	DynamicValue           []ActivityDefinitionDynamicValue `bson:"dynamicValue" json:"dynamicValue"`
}
type ActivityDefinitionParticipant struct {
	Id                *string               `bson:"id" json:"id"`
	Extension         []Extension           `bson:"extension" json:"extension"`
	ModifierExtension []Extension           `bson:"modifierExtension" json:"modifierExtension"`
	Type              ActionParticipantType `bson:"type,omitempty" json:"type,omitempty"`
	Role              *CodeableConcept      `bson:"role" json:"role"`
}
type ActivityDefinitionDynamicValue struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Description       *string     `bson:"description" json:"description"`
	Path              *string     `bson:"path" json:"path"`
	Language          *string     `bson:"language" json:"language"`
	Expression        *string     `bson:"expression" json:"expression"`
}

// OtherActivityDefinition is a helper type to use the default implementations of Marshall and Unmarshal
type OtherActivityDefinition ActivityDefinition

// MarshalJSON marshals the given ActivityDefinition as JSON into a byte slice
func (r ActivityDefinition) MarshalJSON() ([]byte, error) {
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
		OtherActivityDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherActivityDefinition: OtherActivityDefinition(r),
		ResourceType:            "ActivityDefinition",
	})
}

// UnmarshalJSON unmarshals the given byte slice into ActivityDefinition
func (r *ActivityDefinition) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherActivityDefinition)(r)); err != nil {
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
func (r ActivityDefinition) GetResourceType() ResourceType {
	return ResourceTypeActivityDefinition
}
