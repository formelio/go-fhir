package fhir

import "encoding/json"

// ImagingStudy is documented here http://hl7.org/fhir/StructureDefinition/ImagingStudy
type ImagingStudy struct {
	Id                 *string               `bson:"id" json:"id"`
	Meta               *Meta                 `bson:"meta" json:"meta"`
	ImplicitRules      *string               `bson:"implicitRules" json:"implicitRules"`
	Language           *string               `bson:"language" json:"language"`
	Text               *Narrative            `bson:"text" json:"text"`
	Contained          []json.RawMessage     `bson:"contained" json:"contained"`
	Extension          []Extension           `bson:"extension" json:"extension"`
	ModifierExtension  []Extension           `bson:"modifierExtension" json:"modifierExtension"`
	Uid                string                `bson:"uid,omitempty" json:"uid,omitempty"`
	Accession          *Identifier           `bson:"accession" json:"accession"`
	Identifier         []Identifier          `bson:"identifier" json:"identifier"`
	Availability       *InstanceAvailability `bson:"availability" json:"availability"`
	ModalityList       []Coding              `bson:"modalityList" json:"modalityList"`
	Patient            Reference             `bson:"patient,omitempty" json:"patient,omitempty"`
	Context            *Reference            `bson:"context" json:"context"`
	Started            *string               `bson:"started" json:"started"`
	BasedOn            []Reference           `bson:"basedOn" json:"basedOn"`
	Referrer           *Reference            `bson:"referrer" json:"referrer"`
	Interpreter        []Reference           `bson:"interpreter" json:"interpreter"`
	Endpoint           []Reference           `bson:"endpoint" json:"endpoint"`
	NumberOfSeries     *int                  `bson:"numberOfSeries" json:"numberOfSeries"`
	NumberOfInstances  *int                  `bson:"numberOfInstances" json:"numberOfInstances"`
	ProcedureReference []Reference           `bson:"procedureReference" json:"procedureReference"`
	ProcedureCode      []CodeableConcept     `bson:"procedureCode" json:"procedureCode"`
	Reason             *CodeableConcept      `bson:"reason" json:"reason"`
	Description        *string               `bson:"description" json:"description"`
	Series             []ImagingStudySeries  `bson:"series" json:"series"`
}
type ImagingStudySeries struct {
	Id                *string                      `bson:"id" json:"id"`
	Extension         []Extension                  `bson:"extension" json:"extension"`
	ModifierExtension []Extension                  `bson:"modifierExtension" json:"modifierExtension"`
	Uid               string                       `bson:"uid,omitempty" json:"uid,omitempty"`
	Number            *int                         `bson:"number" json:"number"`
	Modality          Coding                       `bson:"modality,omitempty" json:"modality,omitempty"`
	Description       *string                      `bson:"description" json:"description"`
	NumberOfInstances *int                         `bson:"numberOfInstances" json:"numberOfInstances"`
	Availability      *InstanceAvailability        `bson:"availability" json:"availability"`
	Endpoint          []Reference                  `bson:"endpoint" json:"endpoint"`
	BodySite          *Coding                      `bson:"bodySite" json:"bodySite"`
	Laterality        *Coding                      `bson:"laterality" json:"laterality"`
	Started           *string                      `bson:"started" json:"started"`
	Performer         []Reference                  `bson:"performer" json:"performer"`
	Instance          []ImagingStudySeriesInstance `bson:"instance" json:"instance"`
}
type ImagingStudySeriesInstance struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Uid               string      `bson:"uid,omitempty" json:"uid,omitempty"`
	Number            *int        `bson:"number" json:"number"`
	SopClass          string      `bson:"sopClass,omitempty" json:"sopClass,omitempty"`
	Title             *string     `bson:"title" json:"title"`
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

// UnmarshalImagingStudy unmarshalls a ImagingStudy.
func UnmarshalImagingStudy(b []byte) (ImagingStudy, error) {
	var imagingStudy ImagingStudy
	if err := json.Unmarshal(b, &imagingStudy); err != nil {
		return imagingStudy, err
	}
	return imagingStudy, nil
}
