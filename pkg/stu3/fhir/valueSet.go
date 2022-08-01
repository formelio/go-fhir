package fhir

import "encoding/json"

// ValueSet is documented here http://hl7.org/fhir/StructureDefinition/ValueSet
type ValueSet struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string            `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage  `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource        `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string            `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []Identifier       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string            `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string            `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string            `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus  `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool              `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string            `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string            `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail    `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string            `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext     `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept  `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Immutable         *bool              `bson:"immutable,omitempty" json:"immutable,omitempty"`
	Purpose           *string            `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string            `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Extensible        *bool              `bson:"extensible,omitempty" json:"extensible,omitempty"`
	Compose           *ValueSetCompose   `bson:"compose,omitempty" json:"compose,omitempty"`
	Expansion         *ValueSetExpansion `bson:"expansion,omitempty" json:"expansion,omitempty"`
}
type ValueSetCompose struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LockedDate        *string                  `bson:"lockedDate,omitempty" json:"lockedDate,omitempty"`
	Inactive          *bool                    `bson:"inactive,omitempty" json:"inactive,omitempty"`
	Include           []ValueSetComposeInclude `bson:"include,omitempty" json:"include,omitempty"`
	Exclude           []ValueSetComposeInclude `bson:"exclude,omitempty" json:"exclude,omitempty"`
}
type ValueSetComposeInclude struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	System            *string                         `bson:"system,omitempty" json:"system,omitempty"`
	Version           *string                         `bson:"version,omitempty" json:"version,omitempty"`
	Concept           []ValueSetComposeIncludeConcept `bson:"concept,omitempty" json:"concept,omitempty"`
	Filter            []ValueSetComposeIncludeFilter  `bson:"filter,omitempty" json:"filter,omitempty"`
	ValueSet          []string                        `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
}
type ValueSetComposeIncludeConcept struct {
	Id                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string                                     `bson:"code,omitempty" json:"code,omitempty"`
	Display           *string                                    `bson:"display,omitempty" json:"display,omitempty"`
	Designation       []ValueSetComposeIncludeConceptDesignation `bson:"designation,omitempty" json:"designation,omitempty"`
}
type ValueSetComposeIncludeConceptDesignation struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Language          *string     `bson:"language,omitempty" json:"language,omitempty"`
	Use               *Coding     `bson:"use,omitempty" json:"use,omitempty"`
	Value             string      `bson:"value,omitempty" json:"value,omitempty"`
}
type ValueSetComposeIncludeFilter struct {
	Id                *string        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Property          string         `bson:"property,omitempty" json:"property,omitempty"`
	Op                FilterOperator `bson:"op,omitempty" json:"op,omitempty"`
	Value             string         `bson:"value,omitempty" json:"value,omitempty"`
}
type ValueSetExpansion struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        string                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Timestamp         string                       `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Total             *int                         `bson:"total,omitempty" json:"total,omitempty"`
	Offset            *int                         `bson:"offset,omitempty" json:"offset,omitempty"`
	Parameter         []ValueSetExpansionParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	Contains          []ValueSetExpansionContains  `bson:"contains,omitempty" json:"contains,omitempty"`
}
type ValueSetExpansionParameter struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string      `bson:"name,omitempty" json:"name,omitempty"`
	ValueString       *string     `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueBoolean      *bool       `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueInteger      *int        `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueDecimal      *float64    `bson:"valueDecimal,omitempty" json:"valueDecimal,omitempty"`
	ValueUri          *string     `bson:"valueUri,omitempty" json:"valueUri,omitempty"`
	ValueCode         *string     `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
}
type ValueSetExpansionContains struct {
	Id                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	System            *string                                    `bson:"system,omitempty" json:"system,omitempty"`
	Abstract          *bool                                      `bson:"abstract,omitempty" json:"abstract,omitempty"`
	Inactive          *bool                                      `bson:"inactive,omitempty" json:"inactive,omitempty"`
	Version           *string                                    `bson:"version,omitempty" json:"version,omitempty"`
	Code              *string                                    `bson:"code,omitempty" json:"code,omitempty"`
	Display           *string                                    `bson:"display,omitempty" json:"display,omitempty"`
	Designation       []ValueSetComposeIncludeConceptDesignation `bson:"designation,omitempty" json:"designation,omitempty"`
	Contains          []ValueSetExpansionContains                `bson:"contains,omitempty" json:"contains,omitempty"`
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
