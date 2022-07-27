package fhir

import "encoding/json"

// Media is documented here http://hl7.org/fhir/StructureDefinition/Media
type Media struct {
	Id                 *string           `bson:"id" json:"id"`
	Meta               *Meta             `bson:"meta" json:"meta"`
	ImplicitRules      *string           `bson:"implicitRules" json:"implicitRules"`
	Language           *string           `bson:"language" json:"language"`
	Text               *Narrative        `bson:"text" json:"text"`
	RawContained       []json.RawMessage `bson:"contained" json:"contained"`
	Contained          []IResource       `bson:"-" json:"-"`
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
