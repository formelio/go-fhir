package fhir

import "encoding/json"

// ResearchStudy is documented here http://hl7.org/fhir/StructureDefinition/ResearchStudy
type ResearchStudy struct {
	Id                    *string             `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta               `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string             `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string             `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative          `bson:"text,omitempty" json:"text,omitempty"`
	RawContained          []json.RawMessage   `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained             []IResource         `bson:"-,omitempty" json:"-,omitempty"`
	Extension             []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier        `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Title                 *string             `bson:"title,omitempty" json:"title,omitempty"`
	Protocol              []Reference         `bson:"protocol,omitempty" json:"protocol,omitempty"`
	PartOf                []Reference         `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                ResearchStudyStatus `bson:"status,omitempty" json:"status,omitempty"`
	Category              []CodeableConcept   `bson:"category,omitempty" json:"category,omitempty"`
	Focus                 []CodeableConcept   `bson:"focus,omitempty" json:"focus,omitempty"`
	Contact               []ContactDetail     `bson:"contact,omitempty" json:"contact,omitempty"`
	RelatedArtifact       []RelatedArtifact   `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Keyword               []CodeableConcept   `bson:"keyword,omitempty" json:"keyword,omitempty"`
	Jurisdiction          []CodeableConcept   `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Description           *string             `bson:"description,omitempty" json:"description,omitempty"`
	Enrollment            []Reference         `bson:"enrollment,omitempty" json:"enrollment,omitempty"`
	Period                *Period             `bson:"period,omitempty" json:"period,omitempty"`
	Sponsor               *Reference          `bson:"sponsor,omitempty" json:"sponsor,omitempty"`
	PrincipalInvestigator *Reference          `bson:"principalInvestigator,omitempty" json:"principalInvestigator,omitempty"`
	Site                  []Reference         `bson:"site,omitempty" json:"site,omitempty"`
	ReasonStopped         *CodeableConcept    `bson:"reasonStopped,omitempty" json:"reasonStopped,omitempty"`
	Note                  []Annotation        `bson:"note,omitempty" json:"note,omitempty"`
	Arm                   []ResearchStudyArm  `bson:"arm,omitempty" json:"arm,omitempty"`
}
type ResearchStudyArm struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string           `bson:"name,omitempty" json:"name,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Description       *string          `bson:"description,omitempty" json:"description,omitempty"`
}

// OtherResearchStudy is a helper type to use the default implementations of Marshall and Unmarshal
type OtherResearchStudy ResearchStudy

// MarshalJSON marshals the given ResearchStudy as JSON into a byte slice
func (r ResearchStudy) MarshalJSON() ([]byte, error) {
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
		OtherResearchStudy
		ResourceType string `json:"resourceType"`
	}{
		OtherResearchStudy: OtherResearchStudy(r),
		ResourceType:       "ResearchStudy",
	})
}

// UnmarshalJSON unmarshals the given byte slice into ResearchStudy
func (r *ResearchStudy) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherResearchStudy)(r)); err != nil {
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
func (r ResearchStudy) GetResourceType() ResourceType {
	return ResourceTypeResearchStudy
}
