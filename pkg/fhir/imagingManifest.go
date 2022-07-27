package fhir

import "encoding/json"

// ImagingManifest is documented here http://hl7.org/fhir/StructureDefinition/ImagingManifest
type ImagingManifest struct {
	Id                *string                `bson:"id" json:"id"`
	Meta              *Meta                  `bson:"meta" json:"meta"`
	ImplicitRules     *string                `bson:"implicitRules" json:"implicitRules"`
	Language          *string                `bson:"language" json:"language"`
	Text              *Narrative             `bson:"text" json:"text"`
	RawContained      []json.RawMessage      `bson:"contained" json:"contained"`
	Contained         []IResource            `bson:"-" json:"-"`
	Extension         []Extension            `bson:"extension" json:"extension"`
	ModifierExtension []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier            `bson:"identifier" json:"identifier"`
	Patient           Reference              `bson:"patient,omitempty" json:"patient,omitempty"`
	AuthoringTime     *string                `bson:"authoringTime" json:"authoringTime"`
	Author            *Reference             `bson:"author" json:"author"`
	Description       *string                `bson:"description" json:"description"`
	Study             []ImagingManifestStudy `bson:"study,omitempty" json:"study,omitempty"`
}
type ImagingManifestStudy struct {
	Id                *string                      `bson:"id" json:"id"`
	Extension         []Extension                  `bson:"extension" json:"extension"`
	ModifierExtension []Extension                  `bson:"modifierExtension" json:"modifierExtension"`
	Uid               string                       `bson:"uid,omitempty" json:"uid,omitempty"`
	ImagingStudy      *Reference                   `bson:"imagingStudy" json:"imagingStudy"`
	Endpoint          []Reference                  `bson:"endpoint" json:"endpoint"`
	Series            []ImagingManifestStudySeries `bson:"series,omitempty" json:"series,omitempty"`
}
type ImagingManifestStudySeries struct {
	Id                *string                              `bson:"id" json:"id"`
	Extension         []Extension                          `bson:"extension" json:"extension"`
	ModifierExtension []Extension                          `bson:"modifierExtension" json:"modifierExtension"`
	Uid               string                               `bson:"uid,omitempty" json:"uid,omitempty"`
	Endpoint          []Reference                          `bson:"endpoint" json:"endpoint"`
	Instance          []ImagingManifestStudySeriesInstance `bson:"instance,omitempty" json:"instance,omitempty"`
}
type ImagingManifestStudySeriesInstance struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	SopClass          string      `bson:"sopClass,omitempty" json:"sopClass,omitempty"`
	Uid               string      `bson:"uid,omitempty" json:"uid,omitempty"`
}

// OtherImagingManifest is a helper type to use the default implementations of Marshall and Unmarshal
type OtherImagingManifest ImagingManifest

// MarshalJSON marshals the given ImagingManifest as JSON into a byte slice
func (r ImagingManifest) MarshalJSON() ([]byte, error) {
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
		OtherImagingManifest
		ResourceType string `json:"resourceType"`
	}{
		OtherImagingManifest: OtherImagingManifest(r),
		ResourceType:         "ImagingManifest",
	})
}

// UnmarshalJSON unmarshals the given byte slice into ImagingManifest
func (r *ImagingManifest) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherImagingManifest)(r)); err != nil {
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
func (r ImagingManifest) GetResourceType() ResourceType {
	return ResourceTypeImagingManifest
}
