package fhir

import (
	"bytes"
	"encoding/json"
)

// EligibilityRequest is documented here http://hl7.org/fhir/StructureDefinition/EligibilityRequest
type EligibilityRequest struct {
	Id                  *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	RawContained        []json.RawMessage `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained           []IResource       `bson:"-,omitempty" json:"-,omitempty"`
	Extension           []*Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []*Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []*Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              *string           `bson:"status,omitempty" json:"status,omitempty"`
	Priority            *CodeableConcept  `bson:"priority,omitempty" json:"priority,omitempty"`
	Patient             *Reference        `bson:"patient,omitempty" json:"patient,omitempty"`
	ServicedDate        *string           `bson:"servicedDate,omitempty" json:"servicedDate,omitempty"`
	ServicedPeriod      *Period           `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty"`
	Created             *string           `bson:"created,omitempty" json:"created,omitempty"`
	Enterer             *Reference        `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Provider            *Reference        `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization        *Reference        `bson:"organization,omitempty" json:"organization,omitempty"`
	Insurer             *Reference        `bson:"insurer,omitempty" json:"insurer,omitempty"`
	Facility            *Reference        `bson:"facility,omitempty" json:"facility,omitempty"`
	Coverage            *Reference        `bson:"coverage,omitempty" json:"coverage,omitempty"`
	BusinessArrangement *string           `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
	BenefitCategory     *CodeableConcept  `bson:"benefitCategory,omitempty" json:"benefitCategory,omitempty"`
	BenefitSubCategory  *CodeableConcept  `bson:"benefitSubCategory,omitempty" json:"benefitSubCategory,omitempty"`
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherEligibilityRequest
	}{
		OtherEligibilityRequest: OtherEligibilityRequest(r),
		ResourceType:            "EligibilityRequest",
	})
	return buffer.Bytes(), err
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
