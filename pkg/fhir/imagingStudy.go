package fhir

import "encoding/json"

// ImagingStudy is documented here http://hl7.org/fhir/StructureDefinition/ImagingStudy
type ImagingStudy struct {
	Id                 *string               `bson:"id" json:"id"`
	Meta               *Meta                 `bson:"meta" json:"meta"`
	ImplicitRules      *string               `bson:"implicitRules" json:"implicitRules"`
	Language           *string               `bson:"language" json:"language"`
	Text               *Narrative            `bson:"text" json:"text"`
	RawContained       []json.RawMessage     `bson:"contained" json:"contained"`
	Contained          []IResource           `bson:"-" json:"-"`
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

// OtherImagingStudy is a helper type to use the default implementations of Marshall and Unmarshal
type OtherImagingStudy ImagingStudy

// MarshalJSON marshals the given ImagingStudy as JSON into a byte slice
func (r ImagingStudy) MarshalJSON() ([]byte, error) {
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
		OtherImagingStudy
		ResourceType string `json:"resourceType"`
	}{
		OtherImagingStudy: OtherImagingStudy(r),
		ResourceType:      "ImagingStudy",
	})
}

// UnmarshalJSON unmarshals the given byte slice into ImagingStudy
func (r *ImagingStudy) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherImagingStudy)(r)); err != nil {
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
func (r ImagingStudy) GetResourceType() ResourceType {
	return ResourceTypeImagingStudy
}
