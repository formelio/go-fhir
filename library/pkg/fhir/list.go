package fhir

import "encoding/json"

// List is documented here http://hl7.org/fhir/StructureDefinition/List
type List struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string          `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative       `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            string           `bson:"status" json:"status"`
	Mode              string           `bson:"mode" json:"mode"`
	Title             *string          `bson:"title,omitempty" json:"title,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Encounter         *Reference       `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Date              *string          `bson:"date,omitempty" json:"date,omitempty"`
	OrderedBy         *CodeableConcept `bson:"orderedBy,omitempty" json:"orderedBy,omitempty"`
	Note              []Annotation     `bson:"note,omitempty" json:"note,omitempty"`
	Entry             []ListEntry      `bson:"entry,omitempty" json:"entry,omitempty"`
	EmptyReason       *CodeableConcept `bson:"emptyReason,omitempty" json:"emptyReason,omitempty"`
}
type ListEntry struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Flag              *CodeableConcept `bson:"flag,omitempty" json:"flag,omitempty"`
	Deleted           *bool            `bson:"deleted,omitempty" json:"deleted,omitempty"`
	Date              *string          `bson:"date,omitempty" json:"date,omitempty"`
	Item              Reference        `bson:"item" json:"item"`
}
type OtherList List

// MarshalJSON marshals the given List as JSON into a byte slice
func (r List) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherList
		ResourceType string `json:"resourceType"`
	}{
		OtherList:    OtherList(r),
		ResourceType: "List",
	})
}

// UnmarshalList unmarshals a List.
func UnmarshalList(b []byte) (List, error) {
	var list List
	if err := json.Unmarshal(b, &list); err != nil {
		return list, err
	}
	return list, nil
}
