package fhir

import "encoding/json"

// ResearchStudy is documented here http://hl7.org/fhir/StructureDefinition/ResearchStudy
type ResearchStudy struct {
	Id                    *string             `bson:"id" json:"id"`
	Meta                  *Meta               `bson:"meta" json:"meta"`
	ImplicitRules         *string             `bson:"implicitRules" json:"implicitRules"`
	Language              *string             `bson:"language" json:"language"`
	Text                  *Narrative          `bson:"text" json:"text"`
	RawContained          []json.RawMessage   `bson:"contained" json:"contained"`
	Contained             []IResource         `bson:"-" json:"-"`
	Extension             []Extension         `bson:"extension" json:"extension"`
	ModifierExtension     []Extension         `bson:"modifierExtension" json:"modifierExtension"`
	Identifier            []Identifier        `bson:"identifier" json:"identifier"`
	Title                 *string             `bson:"title" json:"title"`
	Protocol              []Reference         `bson:"protocol" json:"protocol"`
	PartOf                []Reference         `bson:"partOf" json:"partOf"`
	Status                ResearchStudyStatus `bson:"status,omitempty" json:"status,omitempty"`
	Category              []CodeableConcept   `bson:"category" json:"category"`
	Focus                 []CodeableConcept   `bson:"focus" json:"focus"`
	Contact               []ContactDetail     `bson:"contact" json:"contact"`
	RelatedArtifact       []RelatedArtifact   `bson:"relatedArtifact" json:"relatedArtifact"`
	Keyword               []CodeableConcept   `bson:"keyword" json:"keyword"`
	Jurisdiction          []CodeableConcept   `bson:"jurisdiction" json:"jurisdiction"`
	Description           *string             `bson:"description" json:"description"`
	Enrollment            []Reference         `bson:"enrollment" json:"enrollment"`
	Period                *Period             `bson:"period" json:"period"`
	Sponsor               *Reference          `bson:"sponsor" json:"sponsor"`
	PrincipalInvestigator *Reference          `bson:"principalInvestigator" json:"principalInvestigator"`
	Site                  []Reference         `bson:"site" json:"site"`
	ReasonStopped         *CodeableConcept    `bson:"reasonStopped" json:"reasonStopped"`
	Note                  []Annotation        `bson:"note" json:"note"`
	Arm                   []ResearchStudyArm  `bson:"arm" json:"arm"`
}
type ResearchStudyArm struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Name              string           `bson:"name,omitempty" json:"name,omitempty"`
	Code              *CodeableConcept `bson:"code" json:"code"`
	Description       *string          `bson:"description" json:"description"`
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
