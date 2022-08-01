package fhir

import (
	"bytes"
	"encoding/json"
)

// Coverage is documented here http://hl7.org/fhir/StructureDefinition/Coverage
type Coverage struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource       `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []*Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            *string           `bson:"status,omitempty" json:"status,omitempty"`
	Type              *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	PolicyHolder      *Reference        `bson:"policyHolder,omitempty" json:"policyHolder,omitempty"`
	Subscriber        *Reference        `bson:"subscriber,omitempty" json:"subscriber,omitempty"`
	SubscriberId      *string           `bson:"subscriberId,omitempty" json:"subscriberId,omitempty"`
	Beneficiary       *Reference        `bson:"beneficiary,omitempty" json:"beneficiary,omitempty"`
	Relationship      *CodeableConcept  `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Period            *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Payor             []*Reference      `bson:"payor,omitempty" json:"payor,omitempty"`
	Grouping          *CoverageGrouping `bson:"grouping,omitempty" json:"grouping,omitempty"`
	Dependent         *string           `bson:"dependent,omitempty" json:"dependent,omitempty"`
	Sequence          *string           `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Order             *int              `bson:"order,omitempty" json:"order,omitempty"`
	Network           *string           `bson:"network,omitempty" json:"network,omitempty"`
	Contract          []*Reference      `bson:"contract,omitempty" json:"contract,omitempty"`
}
type CoverageGrouping struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Group             *string      `bson:"group,omitempty" json:"group,omitempty"`
	GroupDisplay      *string      `bson:"groupDisplay,omitempty" json:"groupDisplay,omitempty"`
	SubGroup          *string      `bson:"subGroup,omitempty" json:"subGroup,omitempty"`
	SubGroupDisplay   *string      `bson:"subGroupDisplay,omitempty" json:"subGroupDisplay,omitempty"`
	Plan              *string      `bson:"plan,omitempty" json:"plan,omitempty"`
	PlanDisplay       *string      `bson:"planDisplay,omitempty" json:"planDisplay,omitempty"`
	SubPlan           *string      `bson:"subPlan,omitempty" json:"subPlan,omitempty"`
	SubPlanDisplay    *string      `bson:"subPlanDisplay,omitempty" json:"subPlanDisplay,omitempty"`
	Class             *string      `bson:"class,omitempty" json:"class,omitempty"`
	ClassDisplay      *string      `bson:"classDisplay,omitempty" json:"classDisplay,omitempty"`
	SubClass          *string      `bson:"subClass,omitempty" json:"subClass,omitempty"`
	SubClassDisplay   *string      `bson:"subClassDisplay,omitempty" json:"subClassDisplay,omitempty"`
}

// OtherCoverage is a helper type to use the default implementations of Marshall and Unmarshal
type OtherCoverage Coverage

// MarshalJSON marshals the given Coverage as JSON into a byte slice
func (r Coverage) MarshalJSON() ([]byte, error) {
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
		OtherCoverage
	}{
		OtherCoverage: OtherCoverage(r),
		ResourceType:  "Coverage",
	})
	return buffer.Bytes(), err
}

// UnmarshalJSON unmarshals the given byte slice into Coverage
func (r *Coverage) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherCoverage)(r)); err != nil {
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
func (r Coverage) GetResourceType() ResourceType {
	return ResourceTypeCoverage
}
