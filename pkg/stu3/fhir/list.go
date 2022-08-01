package fhir

import "encoding/json"

// List is documented here http://hl7.org/fhir/StructureDefinition/List
type List struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource       `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            ListStatus        `bson:"status,omitempty" json:"status,omitempty"`
	Mode              ListMode          `bson:"mode,omitempty" json:"mode,omitempty"`
	Title             *string           `bson:"title,omitempty" json:"title,omitempty"`
	Code              *CodeableConcept  `bson:"code,omitempty" json:"code,omitempty"`
	Subject           *Reference        `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter         *Reference        `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Date              *string           `bson:"date,omitempty" json:"date,omitempty"`
	Source            *Reference        `bson:"source,omitempty" json:"source,omitempty"`
	OrderedBy         *CodeableConcept  `bson:"orderedBy,omitempty" json:"orderedBy,omitempty"`
	Note              []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
	Entry             []ListEntry       `bson:"entry,omitempty" json:"entry,omitempty"`
	EmptyReason       *CodeableConcept  `bson:"emptyReason,omitempty" json:"emptyReason,omitempty"`
}
type ListEntry struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Flag              *CodeableConcept `bson:"flag,omitempty" json:"flag,omitempty"`
	Deleted           *bool            `bson:"deleted,omitempty" json:"deleted,omitempty"`
	Date              *string          `bson:"date,omitempty" json:"date,omitempty"`
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
