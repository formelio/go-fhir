package fhir

import "encoding/json"

// DataElement is documented here http://hl7.org/fhir/StructureDefinition/DataElement
type DataElement struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string              `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string              `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative           `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string              `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []Identifier         `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string              `bson:"version,omitempty" json:"version,omitempty"`
	Status            string               `bson:"status" json:"status"`
	Experimental      *bool                `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string              `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string              `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Name              *string              `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string              `bson:"title,omitempty" json:"title,omitempty"`
	Contact           []ContactDetail      `bson:"contact,omitempty" json:"contact,omitempty"`
	UseContext        []UsageContext       `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept    `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Copyright         *string              `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Stringency        *string              `bson:"stringency,omitempty" json:"stringency,omitempty"`
	Mapping           []DataElementMapping `bson:"mapping,omitempty" json:"mapping,omitempty"`
	Element           []ElementDefinition  `bson:"element" json:"element"`
}
type DataElementMapping struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identity          string      `bson:"identity" json:"identity"`
	Uri               *string     `bson:"uri,omitempty" json:"uri,omitempty"`
	Name              *string     `bson:"name,omitempty" json:"name,omitempty"`
	Comment           *string     `bson:"comment,omitempty" json:"comment,omitempty"`
}
type OtherDataElement DataElement

// MarshalJSON marshals the given DataElement as JSON into a byte slice
func (r DataElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDataElement
		ResourceType string `json:"resourceType"`
	}{
		OtherDataElement: OtherDataElement(r),
		ResourceType:     "DataElement",
	})
}

// UnmarshalDataElement unmarshals a DataElement.
func UnmarshalDataElement(b []byte) (DataElement, error) {
	var dataElement DataElement
	if err := json.Unmarshal(b, &dataElement); err != nil {
		return dataElement, err
	}
	return dataElement, nil
}
