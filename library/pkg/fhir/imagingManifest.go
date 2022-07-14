package fhir

import "encoding/json"

// ImagingManifest is documented here http://hl7.org/fhir/StructureDefinition/ImagingManifest
type ImagingManifest struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Patient           Reference              `bson:"patient" json:"patient"`
	AuthoringTime     *string                `bson:"authoringTime,omitempty" json:"authoringTime,omitempty"`
	Description       *string                `bson:"description,omitempty" json:"description,omitempty"`
	Study             []ImagingManifestStudy `bson:"study" json:"study"`
}
type ImagingManifestStudy struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Uid               string                       `bson:"uid" json:"uid"`
	ImagingStudy      *Reference                   `bson:"imagingStudy,omitempty" json:"imagingStudy,omitempty"`
	Endpoint          []Reference                  `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	Series            []ImagingManifestStudySeries `bson:"series" json:"series"`
}
type ImagingManifestStudySeries struct {
	Id                *string                              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Uid               string                               `bson:"uid" json:"uid"`
	Endpoint          []Reference                          `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	Instance          []ImagingManifestStudySeriesInstance `bson:"instance" json:"instance"`
}
type ImagingManifestStudySeriesInstance struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SopClass          string      `bson:"sopClass" json:"sopClass"`
	Uid               string      `bson:"uid" json:"uid"`
}
type OtherImagingManifest ImagingManifest

// MarshalJSON marshals the given ImagingManifest as JSON into a byte slice
func (r ImagingManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImagingManifest
		ResourceType string `json:"resourceType"`
	}{
		OtherImagingManifest: OtherImagingManifest(r),
		ResourceType:         "ImagingManifest",
	})
}

// UnmarshalImagingManifest unmarshals a ImagingManifest.
func UnmarshalImagingManifest(b []byte) (ImagingManifest, error) {
	var imagingManifest ImagingManifest
	if err := json.Unmarshal(b, &imagingManifest); err != nil {
		return imagingManifest, err
	}
	return imagingManifest, nil
}
