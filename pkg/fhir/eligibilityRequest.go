package fhir

import "encoding/json"

// EligibilityRequest is documented here http://hl7.org/fhir/StructureDefinition/EligibilityRequest
type EligibilityRequest struct {
	Id                  *string           `bson:"id" json:"id"`
	Meta                *Meta             `bson:"meta" json:"meta"`
	ImplicitRules       *string           `bson:"implicitRules" json:"implicitRules"`
	Language            *string           `bson:"language" json:"language"`
	Text                *Narrative        `bson:"text" json:"text"`
	RawContained        []json.RawMessage `bson:"contained" json:"contained"`
	Contained           []IResource       `bson:"-" json:"-"`
	Extension           []Extension       `bson:"extension" json:"extension"`
	ModifierExtension   []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier          []Identifier      `bson:"identifier" json:"identifier"`
	Status              *string           `bson:"status" json:"status"`
	Priority            *CodeableConcept  `bson:"priority" json:"priority"`
	Patient             *Reference        `bson:"patient" json:"patient"`
	ServicedDate        *string           `bson:"servicedDate,omitempty" json:"servicedDate,omitempty"`
	ServicedPeriod      *Period           `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty"`
	Created             *string           `bson:"created" json:"created"`
	Enterer             *Reference        `bson:"enterer" json:"enterer"`
	Provider            *Reference        `bson:"provider" json:"provider"`
	Organization        *Reference        `bson:"organization" json:"organization"`
	Insurer             *Reference        `bson:"insurer" json:"insurer"`
	Facility            *Reference        `bson:"facility" json:"facility"`
	Coverage            *Reference        `bson:"coverage" json:"coverage"`
	BusinessArrangement *string           `bson:"businessArrangement" json:"businessArrangement"`
	BenefitCategory     *CodeableConcept  `bson:"benefitCategory" json:"benefitCategory"`
	BenefitSubCategory  *CodeableConcept  `bson:"benefitSubCategory" json:"benefitSubCategory"`
}

// OtherEligibilityRequest is a helper type to use the default implementations of Marshall and Unmarshal
type OtherEligibilityRequest EligibilityRequest

// MarshalJSON marshals the given EligibilityRequest as JSON into a byte slice
func (r EligibilityRequest) MarshalJSON() ([]byte, error) {
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
		OtherEligibilityRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherEligibilityRequest: OtherEligibilityRequest(r),
		ResourceType:            "EligibilityRequest",
	})
}

// UnmarshalJSON unmarshals the given byte slice into EligibilityRequest
func (r *EligibilityRequest) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherEligibilityRequest)(r)); err != nil {
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
func (r EligibilityRequest) GetResourceType() ResourceType {
	return ResourceTypeEligibilityRequest
}
