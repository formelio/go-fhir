package fhir

import "encoding/json"

// Coverage is documented here http://hl7.org/fhir/StructureDefinition/Coverage
type Coverage struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            *string           `bson:"status,omitempty" json:"status,omitempty"`
	Type              *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	SubscriberId      *string           `bson:"subscriberId,omitempty" json:"subscriberId,omitempty"`
	Beneficiary       *Reference        `bson:"beneficiary,omitempty" json:"beneficiary,omitempty"`
	Relationship      *CodeableConcept  `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Period            *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Grouping          *CoverageGrouping `bson:"grouping,omitempty" json:"grouping,omitempty"`
	Dependent         *string           `bson:"dependent,omitempty" json:"dependent,omitempty"`
	Sequence          *string           `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Order             *int              `bson:"order,omitempty" json:"order,omitempty"`
	Network           *string           `bson:"network,omitempty" json:"network,omitempty"`
	Contract          []Reference       `bson:"contract,omitempty" json:"contract,omitempty"`
}
type CoverageGrouping struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Group             *string     `bson:"group,omitempty" json:"group,omitempty"`
	GroupDisplay      *string     `bson:"groupDisplay,omitempty" json:"groupDisplay,omitempty"`
	SubGroup          *string     `bson:"subGroup,omitempty" json:"subGroup,omitempty"`
	SubGroupDisplay   *string     `bson:"subGroupDisplay,omitempty" json:"subGroupDisplay,omitempty"`
	Plan              *string     `bson:"plan,omitempty" json:"plan,omitempty"`
	PlanDisplay       *string     `bson:"planDisplay,omitempty" json:"planDisplay,omitempty"`
	SubPlan           *string     `bson:"subPlan,omitempty" json:"subPlan,omitempty"`
	SubPlanDisplay    *string     `bson:"subPlanDisplay,omitempty" json:"subPlanDisplay,omitempty"`
	Class             *string     `bson:"class,omitempty" json:"class,omitempty"`
	ClassDisplay      *string     `bson:"classDisplay,omitempty" json:"classDisplay,omitempty"`
	SubClass          *string     `bson:"subClass,omitempty" json:"subClass,omitempty"`
	SubClassDisplay   *string     `bson:"subClassDisplay,omitempty" json:"subClassDisplay,omitempty"`
}
type OtherCoverage Coverage

// MarshalJSON marshals the given Coverage as JSON into a byte slice
func (r Coverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCoverage
		ResourceType string `json:"resourceType"`
	}{
		OtherCoverage: OtherCoverage(r),
		ResourceType:  "Coverage",
	})
}

// UnmarshalCoverage unmarshals a Coverage.
func UnmarshalCoverage(b []byte) (Coverage, error) {
	var coverage Coverage
	if err := json.Unmarshal(b, &coverage); err != nil {
		return coverage, err
	}
	return coverage, nil
}
