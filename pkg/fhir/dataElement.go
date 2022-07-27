package fhir

import "encoding/json"

// DataElement is documented here http://hl7.org/fhir/StructureDefinition/DataElement
type DataElement struct {
	Id                *string                `bson:"id" json:"id"`
	Meta              *Meta                  `bson:"meta" json:"meta"`
	ImplicitRules     *string                `bson:"implicitRules" json:"implicitRules"`
	Language          *string                `bson:"language" json:"language"`
	Text              *Narrative             `bson:"text" json:"text"`
	RawContained      []json.RawMessage      `bson:"contained" json:"contained"`
	Contained         []IResource            `bson:"-" json:"-"`
	Extension         []Extension            `bson:"extension" json:"extension"`
	ModifierExtension []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	Url               *string                `bson:"url" json:"url"`
	Identifier        []Identifier           `bson:"identifier" json:"identifier"`
	Version           *string                `bson:"version" json:"version"`
	Status            PublicationStatus      `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                  `bson:"experimental" json:"experimental"`
	Date              *string                `bson:"date" json:"date"`
	Publisher         *string                `bson:"publisher" json:"publisher"`
	Name              *string                `bson:"name" json:"name"`
	Title             *string                `bson:"title" json:"title"`
	Contact           []ContactDetail        `bson:"contact" json:"contact"`
	UseContext        []UsageContext         `bson:"useContext" json:"useContext"`
	Jurisdiction      []CodeableConcept      `bson:"jurisdiction" json:"jurisdiction"`
	Copyright         *string                `bson:"copyright" json:"copyright"`
	Stringency        *DataElementStringency `bson:"stringency" json:"stringency"`
	Mapping           []DataElementMapping   `bson:"mapping" json:"mapping"`
	Element           []ElementDefinition    `bson:"element,omitempty" json:"element,omitempty"`
}
type DataElementMapping struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Identity          string      `bson:"identity,omitempty" json:"identity,omitempty"`
	Uri               *string     `bson:"uri" json:"uri"`
	Name              *string     `bson:"name" json:"name"`
	Comment           *string     `bson:"comment" json:"comment"`
}

// OtherDataElement is a helper type to use the default implementations of Marshall and Unmarshal
type OtherDataElement DataElement

// MarshalJSON marshals the given DataElement as JSON into a byte slice
func (r DataElement) MarshalJSON() ([]byte, error) {
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
		OtherDataElement
		ResourceType string `json:"resourceType"`
	}{
		OtherDataElement: OtherDataElement(r),
		ResourceType:     "DataElement",
	})
}

// UnmarshalJSON unmarshals the given byte slice into DataElement
func (r *DataElement) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherDataElement)(r)); err != nil {
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
func (r DataElement) GetResourceType() ResourceType {
	return ResourceTypeDataElement
}
