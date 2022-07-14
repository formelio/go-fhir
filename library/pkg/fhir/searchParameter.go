package fhir

import "encoding/json"

// SearchParameter is documented here http://hl7.org/fhir/StructureDefinition/SearchParameter
type SearchParameter struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                    `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                 `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                     `bson:"url" json:"url"`
	Version           *string                    `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                     `bson:"name" json:"name"`
	Status            string                     `bson:"status" json:"status"`
	Experimental      *bool                      `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                    `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                    `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail            `bson:"contact,omitempty" json:"contact,omitempty"`
	UseContext        []UsageContext             `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept          `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                    `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Code              string                     `bson:"code" json:"code"`
	Base              []string                   `bson:"base" json:"base"`
	Type              string                     `bson:"type" json:"type"`
	DerivedFrom       *string                    `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	Description       string                     `bson:"description" json:"description"`
	Expression        *string                    `bson:"expression,omitempty" json:"expression,omitempty"`
	Xpath             *string                    `bson:"xpath,omitempty" json:"xpath,omitempty"`
	XpathUsage        *string                    `bson:"xpathUsage,omitempty" json:"xpathUsage,omitempty"`
	Target            []string                   `bson:"target,omitempty" json:"target,omitempty"`
	Comparator        []string                   `bson:"comparator,omitempty" json:"comparator,omitempty"`
	Modifier          []string                   `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Chain             []string                   `bson:"chain,omitempty" json:"chain,omitempty"`
	Component         []SearchParameterComponent `bson:"component,omitempty" json:"component,omitempty"`
}
type SearchParameterComponent struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Definition        Reference   `bson:"definition" json:"definition"`
	Expression        string      `bson:"expression" json:"expression"`
}
type OtherSearchParameter SearchParameter

// MarshalJSON marshals the given SearchParameter as JSON into a byte slice
func (r SearchParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSearchParameter
		ResourceType string `json:"resourceType"`
	}{
		OtherSearchParameter: OtherSearchParameter(r),
		ResourceType:         "SearchParameter",
	})
}

// UnmarshalSearchParameter unmarshals a SearchParameter.
func UnmarshalSearchParameter(b []byte) (SearchParameter, error) {
	var searchParameter SearchParameter
	if err := json.Unmarshal(b, &searchParameter); err != nil {
		return searchParameter, err
	}
	return searchParameter, nil
}
