package fhir

import "encoding/json"

// ImagingStudy is documented here http://hl7.org/fhir/StructureDefinition/ImagingStudy
type ImagingStudy struct {
	Id                 *string              `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string              `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	Extension          []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Uid                string               `bson:"uid" json:"uid"`
	Accession          *Identifier          `bson:"accession,omitempty" json:"accession,omitempty"`
	Identifier         []Identifier         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Availability       *string              `bson:"availability,omitempty" json:"availability,omitempty"`
	ModalityList       []Coding             `bson:"modalityList,omitempty" json:"modalityList,omitempty"`
	Patient            Reference            `bson:"patient" json:"patient"`
	Started            *string              `bson:"started,omitempty" json:"started,omitempty"`
	Referrer           *Reference           `bson:"referrer,omitempty" json:"referrer,omitempty"`
	Interpreter        []Reference          `bson:"interpreter,omitempty" json:"interpreter,omitempty"`
	Endpoint           []Reference          `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	NumberOfSeries     *int                 `bson:"numberOfSeries,omitempty" json:"numberOfSeries,omitempty"`
	NumberOfInstances  *int                 `bson:"numberOfInstances,omitempty" json:"numberOfInstances,omitempty"`
	ProcedureReference []Reference          `bson:"procedureReference,omitempty" json:"procedureReference,omitempty"`
	ProcedureCode      []CodeableConcept    `bson:"procedureCode,omitempty" json:"procedureCode,omitempty"`
	Reason             *CodeableConcept     `bson:"reason,omitempty" json:"reason,omitempty"`
	Description        *string              `bson:"description,omitempty" json:"description,omitempty"`
	Series             []ImagingStudySeries `bson:"series,omitempty" json:"series,omitempty"`
}
type ImagingStudySeries struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Uid               string                       `bson:"uid" json:"uid"`
	Number            *int                         `bson:"number,omitempty" json:"number,omitempty"`
	Modality          Coding                       `bson:"modality" json:"modality"`
	Description       *string                      `bson:"description,omitempty" json:"description,omitempty"`
	NumberOfInstances *int                         `bson:"numberOfInstances,omitempty" json:"numberOfInstances,omitempty"`
	Availability      *string                      `bson:"availability,omitempty" json:"availability,omitempty"`
	Endpoint          []Reference                  `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	BodySite          *Coding                      `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Laterality        *Coding                      `bson:"laterality,omitempty" json:"laterality,omitempty"`
	Started           *string                      `bson:"started,omitempty" json:"started,omitempty"`
	Performer         []Reference                  `bson:"performer,omitempty" json:"performer,omitempty"`
	Instance          []ImagingStudySeriesInstance `bson:"instance,omitempty" json:"instance,omitempty"`
}
type ImagingStudySeriesInstance struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Uid               string      `bson:"uid" json:"uid"`
	Number            *int        `bson:"number,omitempty" json:"number,omitempty"`
	SopClass          string      `bson:"sopClass" json:"sopClass"`
	Title             *string     `bson:"title,omitempty" json:"title,omitempty"`
}
type OtherImagingStudy ImagingStudy

// MarshalJSON marshals the given ImagingStudy as JSON into a byte slice
func (r ImagingStudy) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImagingStudy
		ResourceType string `json:"resourceType"`
	}{
		OtherImagingStudy: OtherImagingStudy(r),
		ResourceType:      "ImagingStudy",
	})
}

// UnmarshalImagingStudy unmarshals a ImagingStudy.
func UnmarshalImagingStudy(b []byte) (ImagingStudy, error) {
	var imagingStudy ImagingStudy
	if err := json.Unmarshal(b, &imagingStudy); err != nil {
		return imagingStudy, err
	}
	return imagingStudy, nil
}
