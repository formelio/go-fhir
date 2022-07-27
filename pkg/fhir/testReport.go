package fhir

import "encoding/json"

// TestReport is documented here http://hl7.org/fhir/StructureDefinition/TestReport
type TestReport struct {
	Id                *string                 `bson:"id" json:"id"`
	Meta              *Meta                   `bson:"meta" json:"meta"`
	ImplicitRules     *string                 `bson:"implicitRules" json:"implicitRules"`
	Language          *string                 `bson:"language" json:"language"`
	Text              *Narrative              `bson:"text" json:"text"`
	RawContained      []json.RawMessage       `bson:"contained" json:"contained"`
	Contained         []IResource             `bson:"-" json:"-"`
	Extension         []Extension             `bson:"extension" json:"extension"`
	ModifierExtension []Extension             `bson:"modifierExtension" json:"modifierExtension"`
	Identifier        *Identifier             `bson:"identifier" json:"identifier"`
	Name              *string                 `bson:"name" json:"name"`
	Status            TestReportStatus        `bson:"status,omitempty" json:"status,omitempty"`
	TestScript        Reference               `bson:"testScript,omitempty" json:"testScript,omitempty"`
	Result            TestReportResult        `bson:"result,omitempty" json:"result,omitempty"`
	Score             *float64                `bson:"score" json:"score"`
	Tester            *string                 `bson:"tester" json:"tester"`
	Issued            *string                 `bson:"issued" json:"issued"`
	Participant       []TestReportParticipant `bson:"participant" json:"participant"`
	Setup             *TestReportSetup        `bson:"setup" json:"setup"`
	Test              []TestReportTest        `bson:"test" json:"test"`
	Teardown          *TestReportTeardown     `bson:"teardown" json:"teardown"`
}
type TestReportParticipant struct {
	Id                *string                   `bson:"id" json:"id"`
	Extension         []Extension               `bson:"extension" json:"extension"`
	ModifierExtension []Extension               `bson:"modifierExtension" json:"modifierExtension"`
	Type              TestReportParticipantType `bson:"type,omitempty" json:"type,omitempty"`
	Uri               string                    `bson:"uri,omitempty" json:"uri,omitempty"`
	Display           *string                   `bson:"display" json:"display"`
}
type TestReportSetup struct {
	Id                *string                 `bson:"id" json:"id"`
	Extension         []Extension             `bson:"extension" json:"extension"`
	ModifierExtension []Extension             `bson:"modifierExtension" json:"modifierExtension"`
	Action            []TestReportSetupAction `bson:"action,omitempty" json:"action,omitempty"`
}
type TestReportSetupAction struct {
	Id                *string                         `bson:"id" json:"id"`
	Extension         []Extension                     `bson:"extension" json:"extension"`
	ModifierExtension []Extension                     `bson:"modifierExtension" json:"modifierExtension"`
	Operation         *TestReportSetupActionOperation `bson:"operation" json:"operation"`
	Assert            *TestReportSetupActionAssert    `bson:"assert" json:"assert"`
}
type TestReportSetupActionOperation struct {
	Id                *string                `bson:"id" json:"id"`
	Extension         []Extension            `bson:"extension" json:"extension"`
	ModifierExtension []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	Result            TestReportActionResult `bson:"result,omitempty" json:"result,omitempty"`
	Message           *string                `bson:"message" json:"message"`
	Detail            *string                `bson:"detail" json:"detail"`
}
type TestReportSetupActionAssert struct {
	Id                *string                `bson:"id" json:"id"`
	Extension         []Extension            `bson:"extension" json:"extension"`
	ModifierExtension []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	Result            TestReportActionResult `bson:"result,omitempty" json:"result,omitempty"`
	Message           *string                `bson:"message" json:"message"`
	Detail            *string                `bson:"detail" json:"detail"`
}
type TestReportTest struct {
	Id                *string                `bson:"id" json:"id"`
	Extension         []Extension            `bson:"extension" json:"extension"`
	ModifierExtension []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	Name              *string                `bson:"name" json:"name"`
	Description       *string                `bson:"description" json:"description"`
	Action            []TestReportTestAction `bson:"action,omitempty" json:"action,omitempty"`
}
type TestReportTestAction struct {
	Id                *string                         `bson:"id" json:"id"`
	Extension         []Extension                     `bson:"extension" json:"extension"`
	ModifierExtension []Extension                     `bson:"modifierExtension" json:"modifierExtension"`
	Operation         *TestReportSetupActionOperation `bson:"operation" json:"operation"`
	Assert            *TestReportSetupActionAssert    `bson:"assert" json:"assert"`
}
type TestReportTeardown struct {
	Id                *string                    `bson:"id" json:"id"`
	Extension         []Extension                `bson:"extension" json:"extension"`
	ModifierExtension []Extension                `bson:"modifierExtension" json:"modifierExtension"`
	Action            []TestReportTeardownAction `bson:"action,omitempty" json:"action,omitempty"`
}
type TestReportTeardownAction struct {
	Id                *string                        `bson:"id" json:"id"`
	Extension         []Extension                    `bson:"extension" json:"extension"`
	ModifierExtension []Extension                    `bson:"modifierExtension" json:"modifierExtension"`
	Operation         TestReportSetupActionOperation `bson:"operation,omitempty" json:"operation,omitempty"`
}

// OtherTestReport is a helper type to use the default implementations of Marshall and Unmarshal
type OtherTestReport TestReport

// MarshalJSON marshals the given TestReport as JSON into a byte slice
func (r TestReport) MarshalJSON() ([]byte, error) {
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
		OtherTestReport
		ResourceType string `json:"resourceType"`
	}{
		OtherTestReport: OtherTestReport(r),
		ResourceType:    "TestReport",
	})
}

// UnmarshalJSON unmarshals the given byte slice into TestReport
func (r *TestReport) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherTestReport)(r)); err != nil {
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
func (r TestReport) GetResourceType() ResourceType {
	return ResourceTypeTestReport
}
