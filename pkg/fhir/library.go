package fhir

import "encoding/json"

// Library is documented here http://hl7.org/fhir/StructureDefinition/Library
type Library struct {
	Id                *string               `bson:"id" json:"id"`
	Meta              *Meta                 `bson:"meta" json:"meta"`
	ImplicitRules     *string               `bson:"implicitRules" json:"implicitRules"`
	Language          *string               `bson:"language" json:"language"`
	Text              *Narrative            `bson:"text" json:"text"`
	RawContained      []json.RawMessage     `bson:"contained" json:"contained"`
	Contained         []IResource           `bson:"-" json:"-"`
	Extension         []Extension           `bson:"extension" json:"extension"`
	ModifierExtension []Extension           `bson:"modifierExtension" json:"modifierExtension"`
	Url               *string               `bson:"url" json:"url"`
	Identifier        []Identifier          `bson:"identifier" json:"identifier"`
	Version           *string               `bson:"version" json:"version"`
	Name              *string               `bson:"name" json:"name"`
	Title             *string               `bson:"title" json:"title"`
	Status            PublicationStatus     `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                 `bson:"experimental" json:"experimental"`
	Type              CodeableConcept       `bson:"type,omitempty" json:"type,omitempty"`
	Date              *string               `bson:"date" json:"date"`
	Publisher         *string               `bson:"publisher" json:"publisher"`
	Description       *string               `bson:"description" json:"description"`
	Purpose           *string               `bson:"purpose" json:"purpose"`
	Usage             *string               `bson:"usage" json:"usage"`
	ApprovalDate      *string               `bson:"approvalDate" json:"approvalDate"`
	LastReviewDate    *string               `bson:"lastReviewDate" json:"lastReviewDate"`
	EffectivePeriod   *Period               `bson:"effectivePeriod" json:"effectivePeriod"`
	UseContext        []UsageContext        `bson:"useContext" json:"useContext"`
	Jurisdiction      []CodeableConcept     `bson:"jurisdiction" json:"jurisdiction"`
	Topic             []CodeableConcept     `bson:"topic" json:"topic"`
	Contributor       []Contributor         `bson:"contributor" json:"contributor"`
	Contact           []ContactDetail       `bson:"contact" json:"contact"`
	Copyright         *string               `bson:"copyright" json:"copyright"`
	RelatedArtifact   []RelatedArtifact     `bson:"relatedArtifact" json:"relatedArtifact"`
	Parameter         []ParameterDefinition `bson:"parameter" json:"parameter"`
	DataRequirement   []DataRequirement     `bson:"dataRequirement" json:"dataRequirement"`
	Content           []Attachment          `bson:"content" json:"content"`
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
	return json.Marshal(struct {
		OtherLibrary
		ResourceType string `json:"resourceType"`
	}{
		OtherLibrary: OtherLibrary(r),
		ResourceType: "Library",
	})
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
