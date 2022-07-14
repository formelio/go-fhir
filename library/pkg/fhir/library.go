package fhir

import "encoding/json"

// Library is documented here http://hl7.org/fhir/StructureDefinition/Library
type Library struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string               `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string               `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string               `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string               `bson:"title,omitempty" json:"title,omitempty"`
	Status            string                `bson:"status" json:"status"`
	Experimental      *bool                 `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Type              CodeableConcept       `bson:"type" json:"type"`
	Date              *string               `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string               `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Description       *string               `bson:"description,omitempty" json:"description,omitempty"`
	Purpose           *string               `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Usage             *string               `bson:"usage,omitempty" json:"usage,omitempty"`
	ApprovalDate      *string               `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate    *string               `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period               `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	UseContext        []UsageContext        `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept     `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Topic             []CodeableConcept     `bson:"topic,omitempty" json:"topic,omitempty"`
	Contributor       []Contributor         `bson:"contributor,omitempty" json:"contributor,omitempty"`
	Contact           []ContactDetail       `bson:"contact,omitempty" json:"contact,omitempty"`
	Copyright         *string               `bson:"copyright,omitempty" json:"copyright,omitempty"`
	RelatedArtifact   []RelatedArtifact     `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Parameter         []ParameterDefinition `bson:"parameter,omitempty" json:"parameter,omitempty"`
	DataRequirement   []DataRequirement     `bson:"dataRequirement,omitempty" json:"dataRequirement,omitempty"`
	Content           []Attachment          `bson:"content,omitempty" json:"content,omitempty"`
}
type OtherLibrary Library

// MarshalJSON marshals the given Library as JSON into a byte slice
func (r Library) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherLibrary
		ResourceType string `json:"resourceType"`
	}{
		OtherLibrary: OtherLibrary(r),
		ResourceType: "Library",
	})
}

// UnmarshalLibrary unmarshals a Library.
func UnmarshalLibrary(b []byte) (Library, error) {
	var library Library
	if err := json.Unmarshal(b, &library); err != nil {
		return library, err
	}
	return library, nil
}
