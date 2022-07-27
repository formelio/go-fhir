package fhir

import "encoding/json"

// Coverage is documented here http://hl7.org/fhir/StructureDefinition/Coverage
type Coverage struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	RawContained      []json.RawMessage `bson:"contained" json:"contained"`
	Contained         []IResource       `bson:"-" json:"-"`
	Extension         []Extension       `bson:"extension" json:"extension"`
	ModifierExtension []Extension       `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        []Identifier      `bson:"identifier" json:"identifier"`
	Status            *string           `bson:"status" json:"status"`
	Type              *CodeableConcept  `bson:"type" json:"type"`
	PolicyHolder      *Reference        `bson:"policyHolder" json:"policyHolder"`
	Subscriber        *Reference        `bson:"subscriber" json:"subscriber"`
	SubscriberId      *string           `bson:"subscriberId" json:"subscriberId"`
	Beneficiary       *Reference        `bson:"beneficiary" json:"beneficiary"`
	Relationship      *CodeableConcept  `bson:"relationship" json:"relationship"`
	Period            *Period           `bson:"period" json:"period"`
	Payor             []Reference       `bson:"payor" json:"payor"`
	Grouping          *CoverageGrouping `bson:"grouping" json:"grouping"`
	Dependent         *string           `bson:"dependent" json:"dependent"`
	Sequence          *string           `bson:"sequence" json:"sequence"`
	Order             *int              `bson:"order" json:"order"`
	Network           *string           `bson:"network" json:"network"`
	Contract          []Reference       `bson:"contract" json:"contract"`
}
type CoverageGrouping struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Group             *string     `bson:"group" json:"group"`
	GroupDisplay      *string     `bson:"groupDisplay" json:"groupDisplay"`
	SubGroup          *string     `bson:"subGroup" json:"subGroup"`
	SubGroupDisplay   *string     `bson:"subGroupDisplay" json:"subGroupDisplay"`
	Plan              *string     `bson:"plan" json:"plan"`
	PlanDisplay       *string     `bson:"planDisplay" json:"planDisplay"`
	SubPlan           *string     `bson:"subPlan" json:"subPlan"`
	SubPlanDisplay    *string     `bson:"subPlanDisplay" json:"subPlanDisplay"`
	Class             *string     `bson:"class" json:"class"`
	ClassDisplay      *string     `bson:"classDisplay" json:"classDisplay"`
	SubClass          *string     `bson:"subClass" json:"subClass"`
	SubClassDisplay   *string     `bson:"subClassDisplay" json:"subClassDisplay"`
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
	return json.Marshal(struct {
		OtherCoverage
		ResourceType string `json:"resourceType"`
	}{
		OtherCoverage: OtherCoverage(r),
		ResourceType:  "Coverage",
	})
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
