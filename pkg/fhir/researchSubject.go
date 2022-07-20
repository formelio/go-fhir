package fhir

import "encoding/json"

// ResearchSubject is documented here http://hl7.org/fhir/StructureDefinition/ResearchSubject
type ResearchSubject struct {
	Id                *string               `bson:"id" json:"id"`
	Meta              *Meta                 `bson:"meta" json:"meta"`
	ImplicitRules     *string               `bson:"implicitRules" json:"implicitRules"`
	Language          *string               `bson:"language" json:"language"`
	Text              *Narrative            `bson:"text" json:"text"`
	Contained         []json.RawMessage     `bson:"contained" json:"contained"`
	Extension         []Extension           `bson:"extension" json:"extension"`
	ModifierExtension []Extension           `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier           `bson:"identifier" json:"identifier"`
	Status            ResearchSubjectStatus `bson:"status,omitempty" json:"status,omitempty"`
	Period            *Period               `bson:"period" json:"period"`
	Study             Reference             `bson:"study,omitempty" json:"study,omitempty"`
	Individual        Reference             `bson:"individual,omitempty" json:"individual,omitempty"`
	AssignedArm       *string               `bson:"assignedArm" json:"assignedArm"`
	ActualArm         *string               `bson:"actualArm" json:"actualArm"`
	Consent           *Reference            `bson:"consent" json:"consent"`
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

// UnmarshalResearchSubject unmarshalls a ResearchSubject.
func UnmarshalResearchSubject(b []byte) (ResearchSubject, error) {
	var researchSubject ResearchSubject
	if err := json.Unmarshal(b, &researchSubject); err != nil {
		return researchSubject, err
	}
	return researchSubject, nil
}
