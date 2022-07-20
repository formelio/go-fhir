package fhir

import "encoding/json"

// Coverage is documented here http://hl7.org/fhir/StructureDefinition/Coverage
type Coverage struct {
	Id                *string           `bson:"id" json:"id"`
	Meta              *Meta             `bson:"meta" json:"meta"`
	ImplicitRules     *string           `bson:"implicitRules" json:"implicitRules"`
	Language          *string           `bson:"language" json:"language"`
	Text              *Narrative        `bson:"text" json:"text"`
	Contained         []json.RawMessage `bson:"contained" json:"contained"`
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

// UnmarshalCoverage unmarshalls a Coverage.
func UnmarshalCoverage(b []byte) (Coverage, error) {
	var coverage Coverage
	if err := json.Unmarshal(b, &coverage); err != nil {
		return coverage, err
	}
	return coverage, nil
}
