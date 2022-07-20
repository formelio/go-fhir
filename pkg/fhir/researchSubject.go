package fhir

import "encoding/json"

// ResearchSubject is documented here http://hl7.org/fhir/StructureDefinition/ResearchSubject
type ResearchSubject struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string     `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative  `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            string      `bson:"status" json:"status"`
	Period            *Period     `bson:"period,omitempty" json:"period,omitempty"`
	Study             Reference   `bson:"study" json:"study"`
	Individual        Reference   `bson:"individual" json:"individual"`
	AssignedArm       *string     `bson:"assignedArm,omitempty" json:"assignedArm,omitempty"`
	ActualArm         *string     `bson:"actualArm,omitempty" json:"actualArm,omitempty"`
	Consent           *Reference  `bson:"consent,omitempty" json:"consent,omitempty"`
}
type OtherResearchSubject ResearchSubject

// MarshalJSON marshals the given ResearchSubject as JSON into a byte slice
func (r ResearchSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResearchSubject
		ResourceType string `json:"resourceType"`
	}{
		OtherResearchSubject: OtherResearchSubject(r),
		ResourceType:         "ResearchSubject",
	})
}

// UnmarshalResearchSubject unmarshals a ResearchSubject.
func UnmarshalResearchSubject(b []byte) (ResearchSubject, error) {
	var researchSubject ResearchSubject
	if err := json.Unmarshal(b, &researchSubject); err != nil {
		return researchSubject, err
	}
	return researchSubject, nil
}
