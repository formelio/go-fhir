package fhir

import (
	"bytes"
	"encoding/json"
)

// SearchParameter is documented here http://hl7.org/fhir/StructureDefinition/SearchParameter
type SearchParameter struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage           `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource                 `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                      `bson:"url,omitempty" json:"url,omitempty"`
	Version           *string                     `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                      `bson:"name,omitempty" json:"name,omitempty"`
	Status            PublicationStatus           `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                       `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                     `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                     `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []*ContactDetail            `bson:"contact,omitempty" json:"contact,omitempty"`
	UseContext        []*UsageContext             `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []*CodeableConcept          `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                     `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Code              string                      `bson:"code,omitempty" json:"code,omitempty"`
	Base              []ResourceType              `bson:"base,omitempty" json:"base,omitempty"`
	Type              SearchParamType             `bson:"type,omitempty" json:"type,omitempty"`
	DerivedFrom       *string                     `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	Description       string                      `bson:"description,omitempty" json:"description,omitempty"`
	Expression        *string                     `bson:"expression,omitempty" json:"expression,omitempty"`
	Xpath             *string                     `bson:"xpath,omitempty" json:"xpath,omitempty"`
	XpathUsage        *XPathUsageType             `bson:"xpathUsage,omitempty" json:"xpathUsage,omitempty"`
	Target            []*ResourceType             `bson:"target,omitempty" json:"target,omitempty"`
	Comparator        []*SearchComparator         `bson:"comparator,omitempty" json:"comparator,omitempty"`
	Modifier          []*SearchModifierCode       `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Chain             []*string                   `bson:"chain,omitempty" json:"chain,omitempty"`
	Component         []*SearchParameterComponent `bson:"component,omitempty" json:"component,omitempty"`
}
type SearchParameterComponent struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Definition        Reference    `bson:"definition,omitempty" json:"definition,omitempty"`
	Expression        string       `bson:"expression,omitempty" json:"expression,omitempty"`
}

// OtherSearchParameter is a helper type to use the default implementations of Marshall and Unmarshal
type OtherSearchParameter SearchParameter

// MarshalJSON marshals the given SearchParameter as JSON into a byte slice
func (r SearchParameter) MarshalJSON() ([]byte, error) {
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherSearchParameter
	}{
		OtherSearchParameter: OtherSearchParameter(r),
		ResourceType:         "SearchParameter",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into SearchParameter
func (r *SearchParameter) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherSearchParameter)(r)); err != nil {
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
func (r SearchParameter) GetResourceType() ResourceType {
	return ResourceTypeSearchParameter
}
