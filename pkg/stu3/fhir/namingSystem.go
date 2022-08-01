package fhir

import "encoding/json"

// NamingSystem is documented here http://hl7.org/fhir/StructureDefinition/NamingSystem
type NamingSystem struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage      `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource            `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                 `bson:"name,omitempty" json:"name,omitempty"`
	Status            PublicationStatus      `bson:"status,omitempty" json:"status,omitempty"`
	Kind              NamingSystemType       `bson:"kind,omitempty" json:"kind,omitempty"`
	Date              string                 `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail        `bson:"contact,omitempty" json:"contact,omitempty"`
	Responsible       *string                `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Type              *CodeableConcept       `bson:"type,omitempty" json:"type,omitempty"`
	Description       *string                `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []UsageContext         `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept      `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Usage             *string                `bson:"usage,omitempty" json:"usage,omitempty"`
	UniqueId          []NamingSystemUniqueId `bson:"uniqueId,omitempty" json:"uniqueId,omitempty"`
	ReplacedBy        *Reference             `bson:"replacedBy,omitempty" json:"replacedBy,omitempty"`
}
type NamingSystemUniqueId struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              NamingSystemIdentifierType `bson:"type,omitempty" json:"type,omitempty"`
	Value             string                     `bson:"value,omitempty" json:"value,omitempty"`
	Preferred         *bool                      `bson:"preferred,omitempty" json:"preferred,omitempty"`
	Comment           *string                    `bson:"comment,omitempty" json:"comment,omitempty"`
	Period            *Period                    `bson:"period,omitempty" json:"period,omitempty"`
}

// OtherNamingSystem is a helper type to use the default implementations of Marshall and Unmarshal
type OtherNamingSystem NamingSystem

// MarshalJSON marshals the given NamingSystem as JSON into a byte slice
func (r NamingSystem) MarshalJSON() ([]byte, error) {
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
		OtherNamingSystem
		ResourceType string `json:"resourceType"`
	}{
		OtherNamingSystem: OtherNamingSystem(r),
		ResourceType:      "NamingSystem",
	})
}

// UnmarshalJSON unmarshals the given byte slice into NamingSystem
func (r *NamingSystem) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherNamingSystem)(r)); err != nil {
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
func (r NamingSystem) GetResourceType() ResourceType {
	return ResourceTypeNamingSystem
}
