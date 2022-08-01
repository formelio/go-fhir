package fhir

import "encoding/json"

// ActivityDefinition is documented here http://hl7.org/fhir/StructureDefinition/ActivityDefinition
type ActivityDefinition struct {
	Id                     *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                          `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative                       `bson:"text,omitempty" json:"text,omitempty"`
	RawContained           []json.RawMessage                `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained              []IResource                      `bson:"-,omitempty" json:"-,omitempty"`
	Extension              []Extension                      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                    *string                          `bson:"url,omitempty" json:"url,omitempty"`
	Identifier             []Identifier                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version                *string                          `bson:"version,omitempty" json:"version,omitempty"`
	Name                   *string                          `bson:"name,omitempty" json:"name,omitempty"`
	Title                  *string                          `bson:"title,omitempty" json:"title,omitempty"`
	Status                 PublicationStatus                `bson:"status,omitempty" json:"status,omitempty"`
	Experimental           *bool                            `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date                   *string                          `bson:"date,omitempty" json:"date,omitempty"`
	Publisher              *string                          `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Description            *string                          `bson:"description,omitempty" json:"description,omitempty"`
	Purpose                *string                          `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Usage                  *string                          `bson:"usage,omitempty" json:"usage,omitempty"`
	ApprovalDate           *string                          `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate         *string                          `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period                          `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	UseContext             []UsageContext                   `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Topic                  []CodeableConcept                `bson:"topic,omitempty" json:"topic,omitempty"`
	Contributor            []Contributor                    `bson:"contributor,omitempty" json:"contributor,omitempty"`
	Contact                []ContactDetail                  `bson:"contact,omitempty" json:"contact,omitempty"`
	Copyright              *string                          `bson:"copyright,omitempty" json:"copyright,omitempty"`
	RelatedArtifact        []RelatedArtifact                `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Library                []Reference                      `bson:"library,omitempty" json:"library,omitempty"`
	Kind                   *ResourceType                    `bson:"kind,omitempty" json:"kind,omitempty"`
	Code                   *CodeableConcept                 `bson:"code,omitempty" json:"code,omitempty"`
	TimingTiming           *Timing                          `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	TimingDateTime         *string                          `bson:"timingDateTime,omitempty" json:"timingDateTime,omitempty"`
	TimingPeriod           *Period                          `bson:"timingPeriod,omitempty" json:"timingPeriod,omitempty"`
	TimingRange            *Range                           `bson:"timingRange,omitempty" json:"timingRange,omitempty"`
	Location               *Reference                       `bson:"location,omitempty" json:"location,omitempty"`
	Participant            []ActivityDefinitionParticipant  `bson:"participant,omitempty" json:"participant,omitempty"`
	ProductReference       *Reference                       `bson:"productReference,omitempty" json:"productReference,omitempty"`
	ProductCodeableConcept *CodeableConcept                 `bson:"productCodeableConcept,omitempty" json:"productCodeableConcept,omitempty"`
	Quantity               *Quantity                        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Dosage                 []Dosage                         `bson:"dosage,omitempty" json:"dosage,omitempty"`
	BodySite               []CodeableConcept                `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Transform              *Reference                       `bson:"transform,omitempty" json:"transform,omitempty"`
	DynamicValue           []ActivityDefinitionDynamicValue `bson:"dynamicValue,omitempty" json:"dynamicValue,omitempty"`
}
type ActivityDefinitionParticipant struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              ActionParticipantType `bson:"type,omitempty" json:"type,omitempty"`
	Role              *CodeableConcept      `bson:"role,omitempty" json:"role,omitempty"`
}
type ActivityDefinitionDynamicValue struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
	Path              *string     `bson:"path,omitempty" json:"path,omitempty"`
	Language          *string     `bson:"language,omitempty" json:"language,omitempty"`
	Expression        *string     `bson:"expression,omitempty" json:"expression,omitempty"`
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
