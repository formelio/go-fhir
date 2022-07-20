package fhir

import "encoding/json"

// List is documented here http://hl7.org/fhir/StructureDefinition/List
type List struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	Contained         []json.RawMessage `bson:"contained" json:"contained"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier      `bson:"identifier" json:"identifier"`
	Status            ListStatus        `bson:"status,omitempty" json:"status,omitempty"`
	Mode              ListMode          `bson:"mode,omitempty" json:"mode,omitempty"`
	Title             *string           `bson:"title" json:"title"`
	Code              *CodeableConcept  `bson:"code" json:"code"`
	Subject           *Reference        `bson:"subject" json:"subject"`
	Encounter         *Reference        `bson:"encounter" json:"encounter"`
	Date              *string           `bson:"date" json:"date"`
	Source            *Reference        `bson:"source" json:"source"`
	OrderedBy         *CodeableConcept  `bson:"orderedBy" json:"orderedBy"`
	Note              []Annotation      `bson:"note" json:"note"`
	Entry             []ListEntry       `bson:"entry" json:"entry"`
	EmptyReason       *CodeableConcept  `bson:"emptyReason" json:"emptyReason"`
}
type ListEntry struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Flag              *CodeableConcept `bson:"flag" json:"flag"`
	Deleted           *bool            `bson:"deleted" json:"deleted"`
	Date              *string          `bson:"date" json:"date"`
	Item              Reference        `bson:"item,omitempty" json:"item,omitempty"`
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

// UnmarshalList unmarshalls a List.
func UnmarshalList(b []byte) (List, error) {
	var list List
	if err := json.Unmarshal(b, &list); err != nil {
		return list, err
	}
	return list, nil
}
