package fhir

import "encoding/json"

// TestReport is documented here http://hl7.org/fhir/StructureDefinition/TestReport
type TestReport struct {
	Id                *string                 `bson:"id" json:"id"`
	Meta              *Meta                   `bson:"meta" json:"meta"`
	ImplicitRules     *string                 `bson:"implicitRules" json:"implicitRules"`
	Language          *string                 `bson:"language" json:"language"`
	Text              *Narrative              `bson:"text" json:"text"`
	Contained         []json.RawMessage       `bson:"contained" json:"contained"`
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
type OtherTestReport TestReport

// MarshalJSON marshals the given TestReport as JSON into a byte slice
func (r TestReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTestReport
		ResourceType string `json:"resourceType"`
	}{
		OtherTestReport: OtherTestReport(r),
		ResourceType:    "TestReport",
	})
}

// UnmarshalTestReport unmarshalls a TestReport.
func UnmarshalTestReport(b []byte) (TestReport, error) {
	var testReport TestReport
	if err := json.Unmarshal(b, &testReport); err != nil {
		return testReport, err
	}
	return testReport, nil
}
