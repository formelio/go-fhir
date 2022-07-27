package fhir

import "encoding/json"

// StructureDefinition is documented here http://hl7.org/fhir/StructureDefinition/StructureDefinition
type StructureDefinition struct {
	Id                *string                          `bson:"id" json:"id"`
	Meta              *Meta                            `bson:"meta" json:"meta"`
	ImplicitRules     *string                          `bson:"implicitRules" json:"implicitRules"`
	Language          *string                          `bson:"language" json:"language"`
	Text              *Narrative                       `bson:"text" json:"text"`
	RawContained      []json.RawMessage                `bson:"contained" json:"contained"`
	Contained         []IResource                      `bson:"-" json:"-"`
	Extension         []Extension                      `bson:"extension" json:"extension"`
	ModifierExtension []Extension                      `bson:"modifierExtension" json:"modifierExtension"`
	Url               string                           `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        []Identifier                     `bson:"identifier" json:"identifier"`
	Version           *string                          `bson:"version" json:"version"`
	Name              string                           `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                          `bson:"title" json:"title"`
	Status            PublicationStatus                `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                            `bson:"experimental" json:"experimental"`
	Date              *string                          `bson:"date" json:"date"`
	Publisher         *string                          `bson:"publisher" json:"publisher"`
	Contact           []ContactDetail                  `bson:"contact" json:"contact"`
	Description       *string                          `bson:"description" json:"description"`
	UseContext        []UsageContext                   `bson:"useContext" json:"useContext"`
	Jurisdiction      []CodeableConcept                `bson:"jurisdiction" json:"jurisdiction"`
	Purpose           *string                          `bson:"purpose" json:"purpose"`
	Copyright         *string                          `bson:"copyright" json:"copyright"`
	Keyword           []Coding                         `bson:"keyword" json:"keyword"`
	FhirVersion       *string                          `bson:"fhirVersion" json:"fhirVersion"`
	Mapping           []StructureDefinitionMapping     `bson:"mapping" json:"mapping"`
	Kind              StructureDefinitionKind          `bson:"kind,omitempty" json:"kind,omitempty"`
	Abstract          bool                             `bson:"abstract,omitempty" json:"abstract,omitempty"`
	ContextType       *ExtensionContext                `bson:"contextType" json:"contextType"`
	Context           []string                         `bson:"context" json:"context"`
	ContextInvariant  []string                         `bson:"contextInvariant" json:"contextInvariant"`
	Type              string                           `bson:"type,omitempty" json:"type,omitempty"`
	BaseDefinition    *string                          `bson:"baseDefinition" json:"baseDefinition"`
	Derivation        *TypeDerivationRule              `bson:"derivation" json:"derivation"`
	Snapshot          *StructureDefinitionSnapshot     `bson:"snapshot" json:"snapshot"`
	Differential      *StructureDefinitionDifferential `bson:"differential" json:"differential"`
}
type StructureDefinitionMapping struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Identity          string      `bson:"identity,omitempty" json:"identity,omitempty"`
	Uri               *string     `bson:"uri" json:"uri"`
	Name              *string     `bson:"name" json:"name"`
	Comment           *string     `bson:"comment" json:"comment"`
}
type StructureDefinitionSnapshot struct {
	Id                *string             `bson:"id" json:"id"`
	Extension         []Extension         `bson:"extension" json:"extension"`
	ModifierExtension []Extension         `bson:"modifierExtension" json:"modifierExtension"`
	Element           []ElementDefinition `bson:"element,omitempty" json:"element,omitempty"`
}
type StructureDefinitionDifferential struct {
	Id                *string             `bson:"id" json:"id"`
	Extension         []Extension         `bson:"extension" json:"extension"`
	ModifierExtension []Extension         `bson:"modifierExtension" json:"modifierExtension"`
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
