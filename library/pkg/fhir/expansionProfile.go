package fhir

import "encoding/json"

// ExpansionProfile is documented here http://hl7.org/fhir/StructureDefinition/ExpansionProfile
type ExpansionProfile struct {
	Id                     *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                         `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative                      `bson:"text,omitempty" json:"text,omitempty"`
	Extension              []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                    *string                         `bson:"url,omitempty" json:"url,omitempty"`
	Identifier             *Identifier                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version                *string                         `bson:"version,omitempty" json:"version,omitempty"`
	Name                   *string                         `bson:"name,omitempty" json:"name,omitempty"`
	Status                 string                          `bson:"status" json:"status"`
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
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	System            string      `bson:"system" json:"system"`
	Version           string      `bson:"version" json:"version"`
	Mode              string      `bson:"mode" json:"mode"`
}
type ExpansionProfileExcludedSystem struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	System            string      `bson:"system" json:"system"`
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
type OtherExpansionProfile ExpansionProfile

// MarshalJSON marshals the given ExpansionProfile as JSON into a byte slice
func (r ExpansionProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherExpansionProfile
		ResourceType string `json:"resourceType"`
	}{
		OtherExpansionProfile: OtherExpansionProfile(r),
		ResourceType:          "ExpansionProfile",
	})
}

// UnmarshalExpansionProfile unmarshals a ExpansionProfile.
func UnmarshalExpansionProfile(b []byte) (ExpansionProfile, error) {
	var expansionProfile ExpansionProfile
	if err := json.Unmarshal(b, &expansionProfile); err != nil {
		return expansionProfile, err
	}
	return expansionProfile, nil
}
