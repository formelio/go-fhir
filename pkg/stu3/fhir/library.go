package fhir

import (
	"bytes"
	"encoding/json"
)

// Library is documented here http://hl7.org/fhir/StructureDefinition/Library
type Library struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage      `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource            `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string                `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []*Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string                `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus      `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                  `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Type              CodeableConcept        `bson:"type,omitempty" json:"type,omitempty"`
	Date              *string                `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Description       *string                `bson:"description,omitempty" json:"description,omitempty"`
	Purpose           *string                `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Usage             *string                `bson:"usage,omitempty" json:"usage,omitempty"`
	ApprovalDate      *string                `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate    *string                `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period                `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	UseContext        []*UsageContext        `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []*CodeableConcept     `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Topic             []*CodeableConcept     `bson:"topic,omitempty" json:"topic,omitempty"`
	Contributor       []*Contributor         `bson:"contributor,omitempty" json:"contributor,omitempty"`
	Contact           []*ContactDetail       `bson:"contact,omitempty" json:"contact,omitempty"`
	Copyright         *string                `bson:"copyright,omitempty" json:"copyright,omitempty"`
	RelatedArtifact   []*RelatedArtifact     `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Parameter         []*ParameterDefinition `bson:"parameter,omitempty" json:"parameter,omitempty"`
	DataRequirement   []*DataRequirement     `bson:"dataRequirement,omitempty" json:"dataRequirement,omitempty"`
	Content           []*Attachment          `bson:"content,omitempty" json:"content,omitempty"`
}

// OtherLibrary is a helper type to use the default implementations of Marshall and Unmarshal
type OtherLibrary Library

// MarshalJSON marshals the given Library as JSON into a byte slice
func (r Library) MarshalJSON() ([]byte, error) {
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherLibrary
	}{
		OtherLibrary: OtherLibrary(r),
		ResourceType: "Library",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into Library
func (r *Library) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherLibrary)(r)); err != nil {
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
func (r Library) GetResourceType() ResourceType {
	return ResourceTypeLibrary
}
