package fhir

import "encoding/json"

// Media is documented here http://hl7.org/fhir/StructureDefinition/Media
type Media struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn           []Reference       `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Type              string            `bson:"type" json:"type"`
	Subtype           *CodeableConcept  `bson:"subtype,omitempty" json:"subtype,omitempty"`
	View              *CodeableConcept  `bson:"view,omitempty" json:"view,omitempty"`
	Operator          *Reference        `bson:"operator,omitempty" json:"operator,omitempty"`
	ReasonCode        []CodeableConcept `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	BodySite          *CodeableConcept  `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Height            *int              `bson:"height,omitempty" json:"height,omitempty"`
	Width             *int              `bson:"width,omitempty" json:"width,omitempty"`
	Frames            *int              `bson:"frames,omitempty" json:"frames,omitempty"`
	Duration          *int              `bson:"duration,omitempty" json:"duration,omitempty"`
	Content           Attachment        `bson:"content" json:"content"`
	Note              []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
}
type OtherMedia Media

// MarshalJSON marshals the given Media as JSON into a byte slice
func (r Media) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedia
		ResourceType string `json:"resourceType"`
	}{
		OtherMedia:   OtherMedia(r),
		ResourceType: "Media",
	})
}

// UnmarshalMedia unmarshals a Media.
func UnmarshalMedia(b []byte) (Media, error) {
	var media Media
	if err := json.Unmarshal(b, &media); err != nil {
		return media, err
	}
	return media, nil
}
