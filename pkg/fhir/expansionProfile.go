package fhir

import "encoding/json"

// ExpansionProfile is documented here http://hl7.org/fhir/StructureDefinition/ExpansionProfile
type ExpansionProfile struct {
	Id                     *string                         `bson:"id" json:"id"`
	Meta                   *Meta                           `bson:"meta" json:"meta"`
	ImplicitRules          *string                         `bson:"implicitRules" json:"implicitRules"`
	Language               *string                         `bson:"language" json:"language"`
	Text                   *Narrative                      `bson:"text" json:"text"`
	Contained              []json.RawMessage               `bson:"contained" json:"contained"`
	Extension              []Extension                     `bson:"extension" json:"extension"`
	ModifierExtension      []Extension                     `bson:"modifierExtension" json:"modifierExtension"`
	Url                    *string                         `bson:"url" json:"url"`
	Identifier             *Identifier                     `bson:"identifier" json:"identifier"`
	Version                *string                         `bson:"version" json:"version"`
	Name                   *string                         `bson:"name" json:"name"`
	Status                 PublicationStatus               `bson:"status,omitempty" json:"status,omitempty"`
	Experimental           *bool                           `bson:"experimental" json:"experimental"`
	Date                   *string                         `bson:"date" json:"date"`
	Publisher              *string                         `bson:"publisher" json:"publisher"`
	Contact                []ContactDetail                 `bson:"contact" json:"contact"`
	Description            *string                         `bson:"description" json:"description"`
	UseContext             []UsageContext                  `bson:"useContext" json:"useContext"`
	Jurisdiction           []CodeableConcept               `bson:"jurisdiction" json:"jurisdiction"`
	FixedVersion           []ExpansionProfileFixedVersion  `bson:"fixedVersion" json:"fixedVersion"`
	ExcludedSystem         *ExpansionProfileExcludedSystem `bson:"excludedSystem" json:"excludedSystem"`
	IncludeDesignations    *bool                           `bson:"includeDesignations" json:"includeDesignations"`
	Designation            *ExpansionProfileDesignation    `bson:"designation" json:"designation"`
	IncludeDefinition      *bool                           `bson:"includeDefinition" json:"includeDefinition"`
	ActiveOnly             *bool                           `bson:"activeOnly" json:"activeOnly"`
	ExcludeNested          *bool                           `bson:"excludeNested" json:"excludeNested"`
	ExcludeNotForUI        *bool                           `bson:"excludeNotForUI" json:"excludeNotForUI"`
	ExcludePostCoordinated *bool                           `bson:"excludePostCoordinated" json:"excludePostCoordinated"`
	DisplayLanguage        *string                         `bson:"displayLanguage" json:"displayLanguage"`
	LimitedExpansion       *bool                           `bson:"limitedExpansion" json:"limitedExpansion"`
}
type ExpansionProfileFixedVersion struct {
	Id                *string                     `bson:"id" json:"id"`
	Extension         []Extension                 `bson:"extension" json:"extension"`
	ModifierExtension []Extension                 `bson:"modifierExtension" json:"modifierExtension"`
	System            string                      `bson:"system,omitempty" json:"system,omitempty"`
	Version           string                      `bson:"version,omitempty" json:"version,omitempty"`
	Mode              SystemVersionProcessingMode `bson:"mode,omitempty" json:"mode,omitempty"`
}
type ExpansionProfileExcludedSystem struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	System            string      `bson:"system,omitempty" json:"system,omitempty"`
	Version           *string     `bson:"version" json:"version"`
}
type ExpansionProfileDesignation struct {
	Id                *string                             `bson:"id" json:"id"`
	Extension         []Extension                         `bson:"extension" json:"extension"`
	ModifierExtension []Extension                         `bson:"modifierExtension" json:"modifierExtension"`
	Include           *ExpansionProfileDesignationInclude `bson:"include" json:"include"`
	Exclude           *ExpansionProfileDesignationExclude `bson:"exclude" json:"exclude"`
}
type ExpansionProfileDesignationInclude struct {
	Id                *string                                         `bson:"id" json:"id"`
	Extension         []Extension                                     `bson:"extension" json:"extension"`
	ModifierExtension []Extension                                     `bson:"modifierExtension" json:"modifierExtension"`
	Designation       []ExpansionProfileDesignationIncludeDesignation `bson:"designation" json:"designation"`
}
type ExpansionProfileDesignationIncludeDesignation struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Language          *string     `bson:"language" json:"language"`
	Use               *Coding     `bson:"use" json:"use"`
}
type ExpansionProfileDesignationExclude struct {
	Id                *string                                         `bson:"id" json:"id"`
	Extension         []Extension                                     `bson:"extension" json:"extension"`
	ModifierExtension []Extension                                     `bson:"modifierExtension" json:"modifierExtension"`
	Designation       []ExpansionProfileDesignationExcludeDesignation `bson:"designation" json:"designation"`
}
type ExpansionProfileDesignationExcludeDesignation struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Language          *string     `bson:"language" json:"language"`
	Use               *Coding     `bson:"use" json:"use"`
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

// UnmarshalExpansionProfile unmarshalls a ExpansionProfile.
func UnmarshalExpansionProfile(b []byte) (ExpansionProfile, error) {
	var expansionProfile ExpansionProfile
	if err := json.Unmarshal(b, &expansionProfile); err != nil {
		return expansionProfile, err
	}
	return expansionProfile, nil
}
