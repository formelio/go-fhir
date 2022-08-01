package fhir

import "encoding/json"

// ImagingStudy is documented here http://hl7.org/fhir/StructureDefinition/ImagingStudy
type ImagingStudy struct {
	Id                 *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	RawContained       []json.RawMessage     `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained          []IResource           `bson:"-,omitempty" json:"-,omitempty"`
	Extension          []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Uid                string                `bson:"uid,omitempty" json:"uid,omitempty"`
	Accession          *Identifier           `bson:"accession,omitempty" json:"accession,omitempty"`
	Identifier         []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Availability       *InstanceAvailability `bson:"availability,omitempty" json:"availability,omitempty"`
	ModalityList       []Coding              `bson:"modalityList,omitempty" json:"modalityList,omitempty"`
	Patient            Reference             `bson:"patient,omitempty" json:"patient,omitempty"`
	Context            *Reference            `bson:"context,omitempty" json:"context,omitempty"`
	Started            *string               `bson:"started,omitempty" json:"started,omitempty"`
	BasedOn            []Reference           `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Referrer           *Reference            `bson:"referrer,omitempty" json:"referrer,omitempty"`
	Interpreter        []Reference           `bson:"interpreter,omitempty" json:"interpreter,omitempty"`
	Endpoint           []Reference           `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	NumberOfSeries     *int                  `bson:"numberOfSeries,omitempty" json:"numberOfSeries,omitempty"`
	NumberOfInstances  *int                  `bson:"numberOfInstances,omitempty" json:"numberOfInstances,omitempty"`
	ProcedureReference []Reference           `bson:"procedureReference,omitempty" json:"procedureReference,omitempty"`
	ProcedureCode      []CodeableConcept     `bson:"procedureCode,omitempty" json:"procedureCode,omitempty"`
	Reason             *CodeableConcept      `bson:"reason,omitempty" json:"reason,omitempty"`
	Description        *string               `bson:"description,omitempty" json:"description,omitempty"`
	Series             []ImagingStudySeries  `bson:"series,omitempty" json:"series,omitempty"`
}
type ImagingStudySeries struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Uid               string                       `bson:"uid,omitempty" json:"uid,omitempty"`
	Number            *int                         `bson:"number,omitempty" json:"number,omitempty"`
	Modality          Coding                       `bson:"modality,omitempty" json:"modality,omitempty"`
	Description       *string                      `bson:"description,omitempty" json:"description,omitempty"`
	NumberOfInstances *int                         `bson:"numberOfInstances,omitempty" json:"numberOfInstances,omitempty"`
	Availability      *InstanceAvailability        `bson:"availability,omitempty" json:"availability,omitempty"`
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
	Uid               string      `bson:"uid,omitempty" json:"uid,omitempty"`
	Number            *int        `bson:"number,omitempty" json:"number,omitempty"`
	SopClass          string      `bson:"sopClass,omitempty" json:"sopClass,omitempty"`
	Title             *string     `bson:"title,omitempty" json:"title,omitempty"`
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
