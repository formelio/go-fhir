package fhir

import "encoding/json"

// Media is documented here http://hl7.org/fhir/StructureDefinition/Media
type Media struct {
	Id                 *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	RawContained       []json.RawMessage `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained          []IResource       `bson:"-,omitempty" json:"-,omitempty"`
	Extension          []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn            []Reference       `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Type               DigitalMediaType  `bson:"type,omitempty" json:"type,omitempty"`
	Subtype            *CodeableConcept  `bson:"subtype,omitempty" json:"subtype,omitempty"`
	View               *CodeableConcept  `bson:"view,omitempty" json:"view,omitempty"`
	Subject            *Reference        `bson:"subject,omitempty" json:"subject,omitempty"`
	Context            *Reference        `bson:"context,omitempty" json:"context,omitempty"`
	OccurrenceDateTime *string           `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period           `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	Operator           *Reference        `bson:"operator,omitempty" json:"operator,omitempty"`
	ReasonCode         []CodeableConcept `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	BodySite           *CodeableConcept  `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Device             *Reference        `bson:"device,omitempty" json:"device,omitempty"`
	Height             *int              `bson:"height,omitempty" json:"height,omitempty"`
	Width              *int              `bson:"width,omitempty" json:"width,omitempty"`
	Frames             *int              `bson:"frames,omitempty" json:"frames,omitempty"`
	Duration           *int              `bson:"duration,omitempty" json:"duration,omitempty"`
	Content            Attachment        `bson:"content,omitempty" json:"content,omitempty"`
	Note               []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
}

// OtherMedia is a helper type to use the default implementations of Marshall and Unmarshal
type OtherMedia Media

// MarshalJSON marshals the given Media as JSON into a byte slice
func (r Media) MarshalJSON() ([]byte, error) {
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
		OtherMedia
		ResourceType string `json:"resourceType"`
	}{
		OtherMedia:   OtherMedia(r),
		ResourceType: "Media",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Media
func (r *Media) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherMedia)(r)); err != nil {
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
func (r Media) GetResourceType() ResourceType {
	return ResourceTypeMedia
}
