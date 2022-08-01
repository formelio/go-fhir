package fhir

import "encoding/json"

// TestReport is documented here http://hl7.org/fhir/StructureDefinition/TestReport
type TestReport struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage       `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource             `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name              *string                 `bson:"name,omitempty" json:"name,omitempty"`
	Status            TestReportStatus        `bson:"status,omitempty" json:"status,omitempty"`
	TestScript        Reference               `bson:"testScript,omitempty" json:"testScript,omitempty"`
	Result            TestReportResult        `bson:"result,omitempty" json:"result,omitempty"`
	Score             *float64                `bson:"score,omitempty" json:"score,omitempty"`
	Tester            *string                 `bson:"tester,omitempty" json:"tester,omitempty"`
	Issued            *string                 `bson:"issued,omitempty" json:"issued,omitempty"`
	Participant       []TestReportParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	Setup             *TestReportSetup        `bson:"setup,omitempty" json:"setup,omitempty"`
	Test              []TestReportTest        `bson:"test,omitempty" json:"test,omitempty"`
	Teardown          *TestReportTeardown     `bson:"teardown,omitempty" json:"teardown,omitempty"`
}
type TestReportParticipant struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              TestReportParticipantType `bson:"type,omitempty" json:"type,omitempty"`
	Uri               string                    `bson:"uri,omitempty" json:"uri,omitempty"`
	Display           *string                   `bson:"display,omitempty" json:"display,omitempty"`
}
type TestReportSetup struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Action            []TestReportSetupAction `bson:"action,omitempty" json:"action,omitempty"`
}
type TestReportSetupAction struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Operation         *TestReportSetupActionOperation `bson:"operation,omitempty" json:"operation,omitempty"`
	Assert            *TestReportSetupActionAssert    `bson:"assert,omitempty" json:"assert,omitempty"`
}
type TestReportSetupActionOperation struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Result            TestReportActionResult `bson:"result,omitempty" json:"result,omitempty"`
	Message           *string                `bson:"message,omitempty" json:"message,omitempty"`
	Detail            *string                `bson:"detail,omitempty" json:"detail,omitempty"`
}
type TestReportSetupActionAssert struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Result            TestReportActionResult `bson:"result,omitempty" json:"result,omitempty"`
	Message           *string                `bson:"message,omitempty" json:"message,omitempty"`
	Detail            *string                `bson:"detail,omitempty" json:"detail,omitempty"`
}
type TestReportTest struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              *string                `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string                `bson:"description,omitempty" json:"description,omitempty"`
	Action            []TestReportTestAction `bson:"action,omitempty" json:"action,omitempty"`
}
type TestReportTestAction struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Operation         *TestReportSetupActionOperation `bson:"operation,omitempty" json:"operation,omitempty"`
	Assert            *TestReportSetupActionAssert    `bson:"assert,omitempty" json:"assert,omitempty"`
}
type TestReportTeardown struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Action            []TestReportTeardownAction `bson:"action,omitempty" json:"action,omitempty"`
}
type TestReportTeardownAction struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
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
