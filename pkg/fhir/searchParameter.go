package fhir

import "encoding/json"

// SearchParameter is documented here http://hl7.org/fhir/StructureDefinition/SearchParameter
type SearchParameter struct {
	Id                *string                    `bson:"id" json:"id"`
	Meta              *Meta                      `bson:"meta" json:"meta"`
	ImplicitRules     *string                    `bson:"implicitRules" json:"implicitRules"`
	Language          *string                    `bson:"language" json:"language"`
	Text              *Narrative                 `bson:"text" json:"text"`
	RawContained      []json.RawMessage          `bson:"contained" json:"contained"`
	Contained         []IResource                `bson:"-" json:"-"`
	Extension         []Extension                `bson:"extension" json:"extension"`
	ModifierExtension []Extension                `bson:"modifierExtension" json:"modifierExtension"`
	Url               string                     `bson:"url,omitempty" json:"url,omitempty"`
	Version           *string                    `bson:"version" json:"version"`
	Name              string                     `bson:"name,omitempty" json:"name,omitempty"`
	Status            PublicationStatus          `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                      `bson:"experimental" json:"experimental"`
	Date              *string                    `bson:"date" json:"date"`
	Publisher         *string                    `bson:"publisher" json:"publisher"`
	Contact           []ContactDetail            `bson:"contact" json:"contact"`
	UseContext        []UsageContext             `bson:"useContext" json:"useContext"`
	Jurisdiction      []CodeableConcept          `bson:"jurisdiction" json:"jurisdiction"`
	Purpose           *string                    `bson:"purpose" json:"purpose"`
	Code              string                     `bson:"code,omitempty" json:"code,omitempty"`
	Base              []ResourceType             `bson:"base,omitempty" json:"base,omitempty"`
	Type              SearchParamType            `bson:"type,omitempty" json:"type,omitempty"`
	DerivedFrom       *string                    `bson:"derivedFrom" json:"derivedFrom"`
	Description       string                     `bson:"description,omitempty" json:"description,omitempty"`
	Expression        *string                    `bson:"expression" json:"expression"`
	Xpath             *string                    `bson:"xpath" json:"xpath"`
	XpathUsage        *XPathUsageType            `bson:"xpathUsage" json:"xpathUsage"`
	Target            []ResourceType             `bson:"target" json:"target"`
	Comparator        []SearchComparator         `bson:"comparator" json:"comparator"`
	Modifier          []SearchModifierCode       `bson:"modifier" json:"modifier"`
	Chain             []string                   `bson:"chain" json:"chain"`
	Component         []SearchParameterComponent `bson:"component" json:"component"`
}
type SearchParameterComponent struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Definition        Reference   `bson:"definition,omitempty" json:"definition,omitempty"`
	Expression        string      `bson:"expression,omitempty" json:"expression,omitempty"`
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
	return json.Marshal(struct {
		OtherSearchParameter
		ResourceType string `json:"resourceType"`
	}{
		OtherSearchParameter: OtherSearchParameter(r),
		ResourceType:         "SearchParameter",
	})
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
