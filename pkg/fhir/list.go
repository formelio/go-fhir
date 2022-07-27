package fhir

import "encoding/json"

// List is documented here http://hl7.org/fhir/StructureDefinition/List
type List struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	RawContained      []json.RawMessage `bson:"contained" json:"contained"`
	Contained         []IResource       `bson:"-" json:"-"`
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

// OtherList is a helper type to use the default implementations of Marshall and Unmarshal
type OtherList List

// MarshalJSON marshals the given List as JSON into a byte slice
func (r List) MarshalJSON() ([]byte, error) {
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
		OtherList
		ResourceType string `json:"resourceType"`
	}{
		OtherList:    OtherList(r),
		ResourceType: "List",
	})
}

// UnmarshalJSON unmarshals the given byte slice into List
func (r *List) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherList)(r)); err != nil {
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
func (r List) GetResourceType() ResourceType {
	return ResourceTypeList
}
