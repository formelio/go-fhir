package fhir

import "encoding/json"

// HealthcareService is documented here http://hl7.org/fhir/StructureDefinition/HealthcareService
type HealthcareService struct {
	Id                     *string                          `bson:"id" json:"id"`
	Meta                   *Meta                            `bson:"meta" json:"meta"`
	ImplicitRules          *string                          `bson:"implicitRules" json:"implicitRules"`
	Language               *string                          `bson:"language" json:"language"`
	Text                   *Narrative                       `bson:"text" json:"text"`
	RawContained           []json.RawMessage                `bson:"contained" json:"contained"`
	Contained              []IResource                      `bson:"-" json:"-"`
	Extension              []Extension                      `bson:"extension" json:"extension"`
	ModifierExtension      []Extension                      `bson:"modifierExtension" json:"modifierExtension"`
	Identifier             []Identifier                     `bson:"identifier" json:"identifier"`
	Active                 *bool                            `bson:"active" json:"active"`
	ProvidedBy             *Reference                       `bson:"providedBy" json:"providedBy"`
	Category               *CodeableConcept                 `bson:"category" json:"category"`
	Type                   []CodeableConcept                `bson:"type" json:"type"`
	Specialty              []CodeableConcept                `bson:"specialty" json:"specialty"`
	Location               []Reference                      `bson:"location" json:"location"`
	Name                   *string                          `bson:"name" json:"name"`
	Comment                *string                          `bson:"comment" json:"comment"`
	ExtraDetails           *string                          `bson:"extraDetails" json:"extraDetails"`
	Photo                  *Attachment                      `bson:"photo" json:"photo"`
	Telecom                []ContactPoint                   `bson:"telecom" json:"telecom"`
	CoverageArea           []Reference                      `bson:"coverageArea" json:"coverageArea"`
	ServiceProvisionCode   []CodeableConcept                `bson:"serviceProvisionCode" json:"serviceProvisionCode"`
	Eligibility            *CodeableConcept                 `bson:"eligibility" json:"eligibility"`
	EligibilityNote        *string                          `bson:"eligibilityNote" json:"eligibilityNote"`
	ProgramName            []string                         `bson:"programName" json:"programName"`
	Characteristic         []CodeableConcept                `bson:"characteristic" json:"characteristic"`
	ReferralMethod         []CodeableConcept                `bson:"referralMethod" json:"referralMethod"`
	AppointmentRequired    *bool                            `bson:"appointmentRequired" json:"appointmentRequired"`
	AvailableTime          []HealthcareServiceAvailableTime `bson:"availableTime" json:"availableTime"`
	NotAvailable           []HealthcareServiceNotAvailable  `bson:"notAvailable" json:"notAvailable"`
	AvailabilityExceptions *string                          `bson:"availabilityExceptions" json:"availabilityExceptions"`
	Endpoint               []Reference                      `bson:"endpoint" json:"endpoint"`
}
type HealthcareServiceAvailableTime struct {
	Id                 *string      `bson:"id" json:"id"`
	Extension          []Extension  `bson:"extension" json:"extension"`
	ModifierExtension  []Extension  `bson:"modifierExtension" json:"modifierExtension"`
	DaysOfWeek         []DaysOfWeek `bson:"daysOfWeek" json:"daysOfWeek"`
	AllDay             *bool        `bson:"allDay" json:"allDay"`
	AvailableStartTime *string      `bson:"availableStartTime" json:"availableStartTime"`
	AvailableEndTime   *string      `bson:"availableEndTime" json:"availableEndTime"`
}
type HealthcareServiceNotAvailable struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Description       string      `bson:"description,omitempty" json:"description,omitempty"`
	During            *Period     `bson:"during" json:"during"`
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
