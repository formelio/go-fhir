package fhir

import "encoding/json"

// CodeSystem is documented here http://hl7.org/fhir/StructureDefinition/CodeSystem
type CodeSystem struct {
	Id                *string                     `bson:"id" json:"id"`
	Meta              *Meta                       `bson:"meta" json:"meta"`
	ImplicitRules     *string                     `bson:"implicitRules" json:"implicitRules"`
	Language          *string                     `bson:"language" json:"language"`
	Text              *Narrative                  `bson:"text" json:"text"`
	RawContained      []json.RawMessage           `bson:"contained" json:"contained"`
	Contained         []IResource                 `bson:"-" json:"-"`
	Extension         []Extension                 `bson:"extension" json:"extension"`
	ModifierExtension []Extension                 `bson:"modifierExtension" json:"modifierExtension"`
	Url               *string                     `bson:"url" json:"url"`
	Identifier        *Identifier                 `bson:"identifier" json:"identifier"`
	Version           *string                     `bson:"version" json:"version"`
	Name              *string                     `bson:"name" json:"name"`
	Title             *string                     `bson:"title" json:"title"`
	Status            PublicationStatus           `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                       `bson:"experimental" json:"experimental"`
	Date              *string                     `bson:"date" json:"date"`
	Publisher         *string                     `bson:"publisher" json:"publisher"`
	Contact           []ContactDetail             `bson:"contact" json:"contact"`
	Description       *string                     `bson:"description" json:"description"`
	UseContext        []UsageContext              `bson:"useContext" json:"useContext"`
	Jurisdiction      []CodeableConcept           `bson:"jurisdiction" json:"jurisdiction"`
	Purpose           *string                     `bson:"purpose" json:"purpose"`
	Copyright         *string                     `bson:"copyright" json:"copyright"`
	CaseSensitive     *bool                       `bson:"caseSensitive" json:"caseSensitive"`
	ValueSet          *string                     `bson:"valueSet" json:"valueSet"`
	HierarchyMeaning  *CodeSystemHierarchyMeaning `bson:"hierarchyMeaning" json:"hierarchyMeaning"`
	Compositional     *bool                       `bson:"compositional" json:"compositional"`
	VersionNeeded     *bool                       `bson:"versionNeeded" json:"versionNeeded"`
	Content           CodeSystemContentMode       `bson:"content,omitempty" json:"content,omitempty"`
	Count             *int                        `bson:"count" json:"count"`
	Filter            []CodeSystemFilter          `bson:"filter" json:"filter"`
	Property          []CodeSystemProperty        `bson:"property" json:"property"`
	Concept           []CodeSystemConcept         `bson:"concept" json:"concept"`
}
type CodeSystemFilter struct {
	Id                *string          `bson:"id" json:"id"`
	Extension         []Extension      `bson:"extension" json:"extension"`
	ModifierExtension []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	Code              string           `bson:"code,omitempty" json:"code,omitempty"`
	Description       *string          `bson:"description" json:"description"`
	Operator          []FilterOperator `bson:"operator,omitempty" json:"operator,omitempty"`
	Value             string           `bson:"value,omitempty" json:"value,omitempty"`
}
type CodeSystemProperty struct {
	Id                *string      `bson:"id" json:"id"`
	Extension         []Extension  `bson:"extension" json:"extension"`
	ModifierExtension []Extension  `bson:"modifierExtension" json:"modifierExtension"`
	Code              string       `bson:"code,omitempty" json:"code,omitempty"`
	Uri               *string      `bson:"uri" json:"uri"`
	Description       *string      `bson:"description" json:"description"`
	Type              PropertyType `bson:"type,omitempty" json:"type,omitempty"`
}
type CodeSystemConcept struct {
	Id                *string                        `bson:"id" json:"id"`
	Extension         []Extension                    `bson:"extension" json:"extension"`
	ModifierExtension []Extension                    `bson:"modifierExtension" json:"modifierExtension"`
	Code              string                         `bson:"code,omitempty" json:"code,omitempty"`
	Display           *string                        `bson:"display" json:"display"`
	Definition        *string                        `bson:"definition" json:"definition"`
	Designation       []CodeSystemConceptDesignation `bson:"designation" json:"designation"`
	Property          []CodeSystemConceptProperty    `bson:"property" json:"property"`
	Concept           []CodeSystemConcept            `bson:"concept" json:"concept"`
}
type CodeSystemConceptDesignation struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Language          *string     `bson:"language" json:"language"`
	Use               *Coding     `bson:"use" json:"use"`
	Value             string      `bson:"value,omitempty" json:"value,omitempty"`
}
type CodeSystemConceptProperty struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Code              string      `bson:"code,omitempty" json:"code,omitempty"`
	ValueCode         *string     `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueCoding       *Coding     `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueString       *string     `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueInteger      *int        `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueBoolean      *bool       `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueDateTime     *string     `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
}

// OtherCodeSystem is a helper type to use the default implementations of Marshall and Unmarshal
type OtherCodeSystem CodeSystem

// MarshalJSON marshals the given CodeSystem as JSON into a byte slice
func (r CodeSystem) MarshalJSON() ([]byte, error) {
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
		OtherCodeSystem
		ResourceType string `json:"resourceType"`
	}{
		OtherCodeSystem: OtherCodeSystem(r),
		ResourceType:    "CodeSystem",
	})
}

// UnmarshalJSON unmarshals the given byte slice into CodeSystem
func (r *CodeSystem) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherCodeSystem)(r)); err != nil {
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
func (r CodeSystem) GetResourceType() ResourceType {
	return ResourceTypeCodeSystem
}
