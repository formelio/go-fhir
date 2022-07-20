package fhir

import "encoding/json"

// ResearchStudy is documented here http://hl7.org/fhir/StructureDefinition/ResearchStudy
type ResearchStudy struct {
	Id                    *string            `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string            `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Title                 *string            `bson:"title,omitempty" json:"title,omitempty"`
	Protocol              []Reference        `bson:"protocol,omitempty" json:"protocol,omitempty"`
	PartOf                []Reference        `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                string             `bson:"status" json:"status"`
	Category              []CodeableConcept  `bson:"category,omitempty" json:"category,omitempty"`
	Focus                 []CodeableConcept  `bson:"focus,omitempty" json:"focus,omitempty"`
	Contact               []ContactDetail    `bson:"contact,omitempty" json:"contact,omitempty"`
	RelatedArtifact       []RelatedArtifact  `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Keyword               []CodeableConcept  `bson:"keyword,omitempty" json:"keyword,omitempty"`
	Jurisdiction          []CodeableConcept  `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Description           *string            `bson:"description,omitempty" json:"description,omitempty"`
	Enrollment            []Reference        `bson:"enrollment,omitempty" json:"enrollment,omitempty"`
	Period                *Period            `bson:"period,omitempty" json:"period,omitempty"`
	Sponsor               *Reference         `bson:"sponsor,omitempty" json:"sponsor,omitempty"`
	PrincipalInvestigator *Reference         `bson:"principalInvestigator,omitempty" json:"principalInvestigator,omitempty"`
	Site                  []Reference        `bson:"site,omitempty" json:"site,omitempty"`
	ReasonStopped         *CodeableConcept   `bson:"reasonStopped,omitempty" json:"reasonStopped,omitempty"`
	Note                  []Annotation       `bson:"note,omitempty" json:"note,omitempty"`
	Arm                   []ResearchStudyArm `bson:"arm,omitempty" json:"arm,omitempty"`
}
type ResearchStudyArm struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string           `bson:"name" json:"name"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Description       *string          `bson:"description,omitempty" json:"description,omitempty"`
}
type OtherResearchStudy ResearchStudy

// MarshalJSON marshals the given ResearchStudy as JSON into a byte slice
func (r ResearchStudy) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResearchStudy
		ResourceType string `json:"resourceType"`
	}{
		OtherResearchStudy: OtherResearchStudy(r),
		ResourceType:       "ResearchStudy",
	})
}

// UnmarshalResearchStudy unmarshals a ResearchStudy.
func UnmarshalResearchStudy(b []byte) (ResearchStudy, error) {
	var researchStudy ResearchStudy
	if err := json.Unmarshal(b, &researchStudy); err != nil {
		return researchStudy, err
	}
	return researchStudy, nil
}
