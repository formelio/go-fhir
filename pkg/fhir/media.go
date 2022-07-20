package fhir

import "encoding/json"

// Media is documented here http://hl7.org/fhir/StructureDefinition/Media
type Media struct {
	Id                 *string           `bson:"id" json:"id"`
	Meta               *Meta             `bson:"meta" json:"meta"`
	ImplicitRules      *string           `bson:"implicitRules" json:"implicitRules"`
	Language           *string           `bson:"language" json:"language"`
	Text               *Narrative        `bson:"text" json:"text"`
	Contained          []json.RawMessage `bson:"contained" json:"contained"`
	Extension          []Extension       `bson:"extension" json:"extension"`
	ModifierExtension  []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier         []Identifier      `bson:"identifier" json:"identifier"`
	BasedOn            []Reference       `bson:"basedOn" json:"basedOn"`
	Type               DigitalMediaType  `bson:"type,omitempty" json:"type,omitempty"`
	Subtype            *CodeableConcept  `bson:"subtype" json:"subtype"`
	View               *CodeableConcept  `bson:"view" json:"view"`
	Subject            *Reference        `bson:"subject" json:"subject"`
	Context            *Reference        `bson:"context" json:"context"`
	OccurrenceDateTime *string           `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period           `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	Operator           *Reference        `bson:"operator" json:"operator"`
	ReasonCode         []CodeableConcept `bson:"reasonCode" json:"reasonCode"`
	BodySite           *CodeableConcept  `bson:"bodySite" json:"bodySite"`
	Device             *Reference        `bson:"device" json:"device"`
	Height             *int              `bson:"height" json:"height"`
	Width              *int              `bson:"width" json:"width"`
	Frames             *int              `bson:"frames" json:"frames"`
	Duration           *int              `bson:"duration" json:"duration"`
	Content            Attachment        `bson:"content,omitempty" json:"content,omitempty"`
	Note               []Annotation      `bson:"note" json:"note"`
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

// UnmarshalMedia unmarshalls a Media.
func UnmarshalMedia(b []byte) (Media, error) {
	var media Media
	if err := json.Unmarshal(b, &media); err != nil {
		return media, err
	}
	return media, nil
}
