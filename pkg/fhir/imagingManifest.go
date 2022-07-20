package fhir

import "encoding/json"

// ImagingManifest is documented here http://hl7.org/fhir/StructureDefinition/ImagingManifest
type ImagingManifest struct {
	Id                *string                `bson:"id" json:"id"`
	Meta              *Meta                  `bson:"meta" json:"meta"`
	ImplicitRules     *string                `bson:"implicitRules" json:"implicitRules"`
	Language          *string                `bson:"language" json:"language"`
	Text              *Narrative             `bson:"text" json:"text"`
	Contained         []json.RawMessage      `bson:"contained" json:"contained"`
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

// UnmarshalImagingManifest unmarshalls a ImagingManifest.
func UnmarshalImagingManifest(b []byte) (ImagingManifest, error) {
	var imagingManifest ImagingManifest
	if err := json.Unmarshal(b, &imagingManifest); err != nil {
		return imagingManifest, err
	}
	return imagingManifest, nil
}
