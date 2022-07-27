package fhir

import "encoding/json"

// ValueSet is documented here http://hl7.org/fhir/StructureDefinition/ValueSet
type ValueSet struct {
	Id                *string            `bson:"id" json:"id"`
	Meta              *Meta              `bson:"meta" json:"meta"`
	ImplicitRules     *string            `bson:"implicitRules" json:"implicitRules"`
	Language          *string            `bson:"language" json:"language"`
	Text              *Narrative         `bson:"text" json:"text"`
	RawContained      []json.RawMessage  `bson:"contained" json:"contained"`
	Contained         []IResource        `bson:"-" json:"-"`
	Extension         []Extension        `bson:"extension" json:"extension"`
	ModifierExtension []Extension        `bson:"modifierExtension" json:"modifierExtension"`
	Url               *string            `bson:"url" json:"url"`
	Identifier        []Identifier       `bson:"identifier" json:"identifier"`
	Version           *string            `bson:"version" json:"version"`
	Name              *string            `bson:"name" json:"name"`
	Title             *string            `bson:"title" json:"title"`
	Status            PublicationStatus  `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool              `bson:"experimental" json:"experimental"`
	Date              *string            `bson:"date" json:"date"`
	Publisher         *string            `bson:"publisher" json:"publisher"`
	Contact           []ContactDetail    `bson:"contact" json:"contact"`
	Description       *string            `bson:"description" json:"description"`
	UseContext        []UsageContext     `bson:"useContext" json:"useContext"`
	Jurisdiction      []CodeableConcept  `bson:"jurisdiction" json:"jurisdiction"`
	Immutable         *bool              `bson:"immutable" json:"immutable"`
	Purpose           *string            `bson:"purpose" json:"purpose"`
	Copyright         *string            `bson:"copyright" json:"copyright"`
	Extensible        *bool              `bson:"extensible" json:"extensible"`
	Compose           *ValueSetCompose   `bson:"compose" json:"compose"`
	Expansion         *ValueSetExpansion `bson:"expansion" json:"expansion"`
}
type ValueSetCompose struct {
	Id                *string                  `bson:"id" json:"id"`
	Extension         []Extension              `bson:"extension" json:"extension"`
	ModifierExtension []Extension              `bson:"modifierExtension" json:"modifierExtension"`
	LockedDate        *string                  `bson:"lockedDate" json:"lockedDate"`
	Inactive          *bool                    `bson:"inactive" json:"inactive"`
	Include           []ValueSetComposeInclude `bson:"include,omitempty" json:"include,omitempty"`
	Exclude           []ValueSetComposeInclude `bson:"exclude" json:"exclude"`
}
type ValueSetComposeInclude struct {
	Id                *string                         `bson:"id" json:"id"`
	Extension         []Extension                     `bson:"extension" json:"extension"`
	ModifierExtension []Extension                     `bson:"modifierExtension" json:"modifierExtension"`
	System            *string                         `bson:"system" json:"system"`
	Version           *string                         `bson:"version" json:"version"`
	Concept           []ValueSetComposeIncludeConcept `bson:"concept" json:"concept"`
	Filter            []ValueSetComposeIncludeFilter  `bson:"filter" json:"filter"`
	ValueSet          []string                        `bson:"valueSet" json:"valueSet"`
}
type ValueSetComposeIncludeConcept struct {
	Id                *string                                    `bson:"id" json:"id"`
	Extension         []Extension                                `bson:"extension" json:"extension"`
	ModifierExtension []Extension                                `bson:"modifierExtension" json:"modifierExtension"`
	Code              string                                     `bson:"code,omitempty" json:"code,omitempty"`
	Display           *string                                    `bson:"display" json:"display"`
	Designation       []ValueSetComposeIncludeConceptDesignation `bson:"designation" json:"designation"`
}
type ValueSetComposeIncludeConceptDesignation struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Language          *string     `bson:"language" json:"language"`
	Use               *Coding     `bson:"use" json:"use"`
	Value             string      `bson:"value,omitempty" json:"value,omitempty"`
}
type ValueSetComposeIncludeFilter struct {
	Id                *string        `bson:"id" json:"id"`
	Extension         []Extension    `bson:"extension" json:"extension"`
	ModifierExtension []Extension    `bson:"modifierExtension" json:"modifierExtension"`
	Property          string         `bson:"property,omitempty" json:"property,omitempty"`
	Op                FilterOperator `bson:"op,omitempty" json:"op,omitempty"`
	Value             string         `bson:"value,omitempty" json:"value,omitempty"`
}
type ValueSetExpansion struct {
	Id                *string                      `bson:"id" json:"id"`
	Extension         []Extension                  `bson:"extension" json:"extension"`
	ModifierExtension []Extension                  `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        string                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Timestamp         string                       `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Total             *int                         `bson:"total" json:"total"`
	Offset            *int                         `bson:"offset" json:"offset"`
	Parameter         []ValueSetExpansionParameter `bson:"parameter" json:"parameter"`
	Contains          []ValueSetExpansionContains  `bson:"contains" json:"contains"`
}
type ValueSetExpansionParameter struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Name              string      `bson:"name,omitempty" json:"name,omitempty"`
	ValueString       *string     `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueBoolean      *bool       `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueInteger      *int        `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDecimal      *float64    `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueUri          *string     `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueCode         *string     `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
}
type ValueSetExpansionContains struct {
	Id                *string                                    `bson:"id" json:"id"`
	Extension         []Extension                                `bson:"extension" json:"extension"`
	ModifierExtension []Extension                                `bson:"modifierExtension" json:"modifierExtension"`
	System            *string                                    `bson:"system" json:"system"`
	Abstract          *bool                                      `bson:"abstract" json:"abstract"`
	Inactive          *bool                                      `bson:"inactive" json:"inactive"`
	Version           *string                                    `bson:"version" json:"version"`
	Code              *string                                    `bson:"code" json:"code"`
	Display           *string                                    `bson:"display" json:"display"`
	Designation       []ValueSetComposeIncludeConceptDesignation `bson:"designation" json:"designation"`
	Contains          []ValueSetExpansionContains                `bson:"contains" json:"contains"`
}

// OtherValueSet is a helper type to use the default implementations of Marshall and Unmarshal
type OtherValueSet ValueSet

// MarshalJSON marshals the given ValueSet as JSON into a byte slice
func (r ValueSet) MarshalJSON() ([]byte, error) {
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
		OtherValueSet
		ResourceType string `json:"resourceType"`
	}{
		OtherValueSet: OtherValueSet(r),
		ResourceType:  "ValueSet",
	})
}

// UnmarshalJSON unmarshals the given byte slice into ValueSet
func (r *ValueSet) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherValueSet)(r)); err != nil {
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
func (r ValueSet) GetResourceType() ResourceType {
	return ResourceTypeValueSet
}
