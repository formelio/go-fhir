package fhir

import (
	"bytes"
	"encoding/json"
)

// CodeSystem is documented here http://hl7.org/fhir/StructureDefinition/CodeSystem
type CodeSystem struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage           `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource                 `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string                     `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        *Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                     `bson:"version,omitempty" json:"version,omitempty"`
	Name              *string                     `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                     `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus           `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                       `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                     `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                     `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []*ContactDetail            `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                     `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []*UsageContext             `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []*CodeableConcept          `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                     `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string                     `bson:"copyright,omitempty" json:"copyright,omitempty"`
	CaseSensitive     *bool                       `bson:"caseSensitive,omitempty" json:"caseSensitive,omitempty"`
	ValueSet          *string                     `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
	HierarchyMeaning  *CodeSystemHierarchyMeaning `bson:"hierarchyMeaning,omitempty" json:"hierarchyMeaning,omitempty"`
	Compositional     *bool                       `bson:"compositional,omitempty" json:"compositional,omitempty"`
	VersionNeeded     *bool                       `bson:"versionNeeded,omitempty" json:"versionNeeded,omitempty"`
	Content           CodeSystemContentMode       `bson:"content,omitempty" json:"content,omitempty"`
	Count             *int                        `bson:"count,omitempty" json:"count,omitempty"`
	Filter            []*CodeSystemFilter         `bson:"filter,omitempty" json:"filter,omitempty"`
	Property          []*CodeSystemProperty       `bson:"property,omitempty" json:"property,omitempty"`
	Concept           []*CodeSystemConcept        `bson:"concept,omitempty" json:"concept,omitempty"`
}
type CodeSystemFilter struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string           `bson:"code,omitempty" json:"code,omitempty"`
	Description       *string          `bson:"description,omitempty" json:"description,omitempty"`
	Operator          []FilterOperator `bson:"operator,omitempty" json:"operator,omitempty"`
	Value             string           `bson:"value,omitempty" json:"value,omitempty"`
}
type CodeSystemProperty struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string       `bson:"code,omitempty" json:"code,omitempty"`
	Uri               *string      `bson:"uri,omitempty" json:"uri,omitempty"`
	Description       *string      `bson:"description,omitempty" json:"description,omitempty"`
	Type              PropertyType `bson:"type,omitempty" json:"type,omitempty"`
}
type CodeSystemConcept struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string                          `bson:"code,omitempty" json:"code,omitempty"`
	Display           *string                         `bson:"display,omitempty" json:"display,omitempty"`
	Definition        *string                         `bson:"definition,omitempty" json:"definition,omitempty"`
	Designation       []*CodeSystemConceptDesignation `bson:"designation,omitempty" json:"designation,omitempty"`
	Property          []*CodeSystemConceptProperty    `bson:"property,omitempty" json:"property,omitempty"`
	Concept           []*CodeSystemConcept            `bson:"concept,omitempty" json:"concept,omitempty"`
}
type CodeSystemConceptDesignation struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Language          *string      `bson:"language,omitempty" json:"language,omitempty"`
	Use               *Coding      `bson:"use,omitempty" json:"use,omitempty"`
	Value             string       `bson:"value,omitempty" json:"value,omitempty"`
}
type CodeSystemConceptProperty struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              string       `bson:"code,omitempty" json:"code,omitempty"`
	ValueCode         *string      `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueCoding       *Coding      `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueString       *string      `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueInteger      *int         `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
	ValueBoolean      *bool        `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueDateTime     *string      `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherCodeSystem
	}{
		OtherCodeSystem: OtherCodeSystem(r),
		ResourceType:    "CodeSystem",
	})
	return buffer.Bytes(), err
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
