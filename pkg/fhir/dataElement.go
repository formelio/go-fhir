package fhir

import "encoding/json"

// DataElement is documented here http://hl7.org/fhir/StructureDefinition/DataElement
type DataElement struct {
	Id                *string                `bson:"id" json:"id"`
	Meta              *Meta                  `bson:"meta" json:"meta"`
	ImplicitRules     *string                `bson:"implicitRules" json:"implicitRules"`
	Language          *string                `bson:"language" json:"language"`
	Text              *Narrative             `bson:"text" json:"text"`
	Contained         []json.RawMessage      `bson:"contained" json:"contained"`
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

// UnmarshalDataElement unmarshalls a DataElement.
func UnmarshalDataElement(b []byte) (DataElement, error) {
	var dataElement DataElement
	if err := json.Unmarshal(b, &dataElement); err != nil {
		return dataElement, err
	}
	return dataElement, nil
}
