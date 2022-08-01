package fhir

import "encoding/json"

// HealthcareService is documented here http://hl7.org/fhir/StructureDefinition/HealthcareService
type HealthcareService struct {
	Id                     *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                          `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative                       `bson:"text,omitempty" json:"text,omitempty"`
	RawContained           []json.RawMessage                `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained              []IResource                      `bson:"-,omitempty" json:"-,omitempty"`
	Extension              []Extension                      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier             []Identifier                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active                 *bool                            `bson:"active,omitempty" json:"active,omitempty"`
	ProvidedBy             *Reference                       `bson:"providedBy,omitempty" json:"providedBy,omitempty"`
	Category               *CodeableConcept                 `bson:"category,omitempty" json:"category,omitempty"`
	Type                   []CodeableConcept                `bson:"type,omitempty" json:"type,omitempty"`
	Specialty              []CodeableConcept                `bson:"specialty,omitempty" json:"specialty,omitempty"`
	Location               []Reference                      `bson:"location,omitempty" json:"location,omitempty"`
	Name                   *string                          `bson:"name,omitempty" json:"name,omitempty"`
	Comment                *string                          `bson:"comment,omitempty" json:"comment,omitempty"`
	ExtraDetails           *string                          `bson:"extraDetails,omitempty" json:"extraDetails,omitempty"`
	Photo                  *Attachment                      `bson:"photo,omitempty" json:"photo,omitempty"`
	Telecom                []ContactPoint                   `bson:"telecom,omitempty" json:"telecom,omitempty"`
	CoverageArea           []Reference                      `bson:"coverageArea,omitempty" json:"coverageArea,omitempty"`
	ServiceProvisionCode   []CodeableConcept                `bson:"serviceProvisionCode,omitempty" json:"serviceProvisionCode,omitempty"`
	Eligibility            *CodeableConcept                 `bson:"eligibility,omitempty" json:"eligibility,omitempty"`
	EligibilityNote        *string                          `bson:"eligibilityNote,omitempty" json:"eligibilityNote,omitempty"`
	ProgramName            []string                         `bson:"programName,omitempty" json:"programName,omitempty"`
	Characteristic         []CodeableConcept                `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	ReferralMethod         []CodeableConcept                `bson:"referralMethod,omitempty" json:"referralMethod,omitempty"`
	AppointmentRequired    *bool                            `bson:"appointmentRequired,omitempty" json:"appointmentRequired,omitempty"`
	AvailableTime          []HealthcareServiceAvailableTime `bson:"availableTime,omitempty" json:"availableTime,omitempty"`
	NotAvailable           []HealthcareServiceNotAvailable  `bson:"notAvailable,omitempty" json:"notAvailable,omitempty"`
	AvailabilityExceptions *string                          `bson:"availabilityExceptions,omitempty" json:"availabilityExceptions,omitempty"`
	Endpoint               []Reference                      `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}
type HealthcareServiceAvailableTime struct {
	Id                 *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	DaysOfWeek         []DaysOfWeek `bson:"daysOfWeek,omitempty" json:"daysOfWeek,omitempty"`
	AllDay             *bool        `bson:"allDay,omitempty" json:"allDay,omitempty"`
	AvailableStartTime *string      `bson:"availableStartTime,omitempty" json:"availableStartTime,omitempty"`
	AvailableEndTime   *string      `bson:"availableEndTime,omitempty" json:"availableEndTime,omitempty"`
}
type HealthcareServiceNotAvailable struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       string      `bson:"description,omitempty" json:"description,omitempty"`
	During            *Period     `bson:"during,omitempty" json:"during,omitempty"`
}

// OtherHealthcareService is a helper type to use the default implementations of Marshall and Unmarshal
type OtherHealthcareService HealthcareService

// MarshalJSON marshals the given HealthcareService as JSON into a byte slice
func (r HealthcareService) MarshalJSON() ([]byte, error) {
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
		OtherHealthcareService
		ResourceType string `json:"resourceType"`
	}{
		OtherHealthcareService: OtherHealthcareService(r),
		ResourceType:           "HealthcareService",
	})
}

// UnmarshalJSON unmarshals the given byte slice into HealthcareService
func (r *HealthcareService) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherHealthcareService)(r)); err != nil {
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
func (r HealthcareService) GetResourceType() ResourceType {
	return ResourceTypeHealthcareService
}
