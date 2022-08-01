package fhir

import "encoding/json"

// ExpansionProfile is documented here http://hl7.org/fhir/StructureDefinition/ExpansionProfile
type ExpansionProfile struct {
	Id                     *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                         `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative                      `bson:"text,omitempty" json:"text,omitempty"`
	RawContained           []json.RawMessage               `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained              []IResource                     `bson:"-,omitempty" json:"-,omitempty"`
	Extension              []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                    *string                         `bson:"url,omitempty" json:"url,omitempty"`
	Identifier             *Identifier                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version                *string                         `bson:"version,omitempty" json:"version,omitempty"`
	Name                   *string                         `bson:"name,omitempty" json:"name,omitempty"`
	Status                 PublicationStatus               `bson:"status,omitempty" json:"status,omitempty"`
	Experimental           *bool                           `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date                   *string                         `bson:"date,omitempty" json:"date,omitempty"`
	Publisher              *string                         `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact                []ContactDetail                 `bson:"contact,omitempty" json:"contact,omitempty"`
	Description            *string                         `bson:"description,omitempty" json:"description,omitempty"`
	UseContext             []UsageContext                  `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept               `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	FixedVersion           []ExpansionProfileFixedVersion  `bson:"fixedVersion,omitempty" json:"fixedVersion,omitempty"`
	ExcludedSystem         *ExpansionProfileExcludedSystem `bson:"excludedSystem,omitempty" json:"excludedSystem,omitempty"`
	IncludeDesignations    *bool                           `bson:"includeDesignations,omitempty" json:"includeDesignations,omitempty"`
	Designation            *ExpansionProfileDesignation    `bson:"designation,omitempty" json:"designation,omitempty"`
	IncludeDefinition      *bool                           `bson:"includeDefinition,omitempty" json:"includeDefinition,omitempty"`
	ActiveOnly             *bool                           `bson:"activeOnly,omitempty" json:"activeOnly,omitempty"`
	ExcludeNested          *bool                           `bson:"excludeNested,omitempty" json:"excludeNested,omitempty"`
	ExcludeNotForUI        *bool                           `bson:"excludeNotForUI,omitempty" json:"excludeNotForUI,omitempty"`
	ExcludePostCoordinated *bool                           `bson:"excludePostCoordinated,omitempty" json:"excludePostCoordinated,omitempty"`
	DisplayLanguage        *string                         `bson:"displayLanguage,omitempty" json:"displayLanguage,omitempty"`
	LimitedExpansion       *bool                           `bson:"limitedExpansion,omitempty" json:"limitedExpansion,omitempty"`
}
type ExpansionProfileFixedVersion struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	System            string                      `bson:"system,omitempty" json:"system,omitempty"`
	Version           string                      `bson:"version,omitempty" json:"version,omitempty"`
	Mode              SystemVersionProcessingMode `bson:"mode,omitempty" json:"mode,omitempty"`
}
type ExpansionProfileExcludedSystem struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	System            string      `bson:"system,omitempty" json:"system,omitempty"`
	Version           *string     `bson:"version,omitempty" json:"version,omitempty"`
}
type ExpansionProfileDesignation struct {
	Id                *string                             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Include           *ExpansionProfileDesignationInclude `bson:"include,omitempty" json:"include,omitempty"`
	Exclude           *ExpansionProfileDesignationExclude `bson:"exclude,omitempty" json:"exclude,omitempty"`
}
type ExpansionProfileDesignationInclude struct {
	Id                *string                                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Designation       []ExpansionProfileDesignationIncludeDesignation `bson:"designation,omitempty" json:"designation,omitempty"`
}
type ExpansionProfileDesignationIncludeDesignation struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Language          *string     `bson:"language,omitempty" json:"language,omitempty"`
	Use               *Coding     `bson:"use,omitempty" json:"use,omitempty"`
}
type ExpansionProfileDesignationExclude struct {
	Id                *string                                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Designation       []ExpansionProfileDesignationExcludeDesignation `bson:"designation,omitempty" json:"designation,omitempty"`
}
type ExpansionProfileDesignationExcludeDesignation struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Language          *string     `bson:"language,omitempty" json:"language,omitempty"`
	Use               *Coding     `bson:"use,omitempty" json:"use,omitempty"`
}

// OtherExpansionProfile is a helper type to use the default implementations of Marshall and Unmarshal
type OtherExpansionProfile ExpansionProfile

// MarshalJSON marshals the given ExpansionProfile as JSON into a byte slice
func (r ExpansionProfile) MarshalJSON() ([]byte, error) {
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
		OtherExpansionProfile
		ResourceType string `json:"resourceType"`
	}{
		OtherExpansionProfile: OtherExpansionProfile(r),
		ResourceType:          "ExpansionProfile",
	})
}

// UnmarshalJSON unmarshals the given byte slice into ExpansionProfile
func (r *ExpansionProfile) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherExpansionProfile)(r)); err != nil {
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
func (r ExpansionProfile) GetResourceType() ResourceType {
	return ResourceTypeExpansionProfile
}
