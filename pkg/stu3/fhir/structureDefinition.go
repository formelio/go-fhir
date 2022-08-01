package fhir

import "encoding/json"

// StructureDefinition is documented here http://hl7.org/fhir/StructureDefinition/StructureDefinition
type StructureDefinition struct {
	Id                *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                          `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                       `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage                `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource                      `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension                      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                           `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []Identifier                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                          `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                           `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                          `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus                `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                            `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                          `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                          `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail                  `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                          `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext                   `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                          `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string                          `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Keyword           []Coding                         `bson:"keyword,omitempty" json:"keyword,omitempty"`
	FhirVersion       *string                          `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	Mapping           []StructureDefinitionMapping     `bson:"mapping,omitempty" json:"mapping,omitempty"`
	Kind              StructureDefinitionKind          `bson:"kind,omitempty" json:"kind,omitempty"`
	Abstract          bool                             `bson:"abstract,omitempty" json:"abstract,omitempty"`
	ContextType       *ExtensionContext                `bson:"contextType,omitempty" json:"contextType,omitempty"`
	Context           []string                         `bson:"context,omitempty" json:"context,omitempty"`
	ContextInvariant  []string                         `bson:"contextInvariant,omitempty" json:"contextInvariant,omitempty"`
	Type              string                           `bson:"type,omitempty" json:"type,omitempty"`
	BaseDefinition    *string                          `bson:"baseDefinition,omitempty" json:"baseDefinition,omitempty"`
	Derivation        *TypeDerivationRule              `bson:"derivation,omitempty" json:"derivation,omitempty"`
	Snapshot          *StructureDefinitionSnapshot     `bson:"snapshot,omitempty" json:"snapshot,omitempty"`
	Differential      *StructureDefinitionDifferential `bson:"differential,omitempty" json:"differential,omitempty"`
}
type StructureDefinitionMapping struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identity          string      `bson:"identity,omitempty" json:"identity,omitempty"`
	Uri               *string     `bson:"uri,omitempty" json:"uri,omitempty"`
	Name              *string     `bson:"name,omitempty" json:"name,omitempty"`
	Comment           *string     `bson:"comment,omitempty" json:"comment,omitempty"`
}
type StructureDefinitionSnapshot struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Element           []ElementDefinition `bson:"element,omitempty" json:"element,omitempty"`
}
type StructureDefinitionDifferential struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Element           []ElementDefinition `bson:"element,omitempty" json:"element,omitempty"`
}

// OtherStructureDefinition is a helper type to use the default implementations of Marshall and Unmarshal
type OtherStructureDefinition StructureDefinition

// MarshalJSON marshals the given StructureDefinition as JSON into a byte slice
func (r StructureDefinition) MarshalJSON() ([]byte, error) {
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
		OtherStructureDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherStructureDefinition: OtherStructureDefinition(r),
		ResourceType:             "StructureDefinition",
	})
}

// UnmarshalJSON unmarshals the given byte slice into StructureDefinition
func (r *StructureDefinition) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherStructureDefinition)(r)); err != nil {
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
func (r StructureDefinition) GetResourceType() ResourceType {
	return ResourceTypeStructureDefinition
}
