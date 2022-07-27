package fhir

import "encoding/json"

// TestScript is documented here http://hl7.org/fhir/StructureDefinition/TestScript
type TestScript struct {
	Id                *string                 `bson:"id" json:"id"`
	Meta              *Meta                   `bson:"meta" json:"meta"`
	ImplicitRules     *string                 `bson:"implicitRules" json:"implicitRules"`
	Language          *string                 `bson:"language" json:"language"`
	Text              *Narrative              `bson:"text" json:"text"`
	RawContained      []json.RawMessage       `bson:"contained" json:"contained"`
	Contained         []IResource             `bson:"-" json:"-"`
	Extension         []Extension             `bson:"extension" json:"extension"`
	ModifierExtension []Extension             `bson:"modifierExtension" json:"modifierExtension"`
	Url               string                  `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        *Identifier             `bson:"identifier" json:"identifier"`
	Version           *string                 `bson:"version" json:"version"`
	Name              string                  `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                 `bson:"title" json:"title"`
	Status            PublicationStatus       `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                   `bson:"experimental" json:"experimental"`
	Date              *string                 `bson:"date" json:"date"`
	Publisher         *string                 `bson:"publisher" json:"publisher"`
	Contact           []ContactDetail         `bson:"contact" json:"contact"`
	Description       *string                 `bson:"description" json:"description"`
	UseContext        []UsageContext          `bson:"useContext" json:"useContext"`
	Jurisdiction      []CodeableConcept       `bson:"jurisdiction" json:"jurisdiction"`
	Purpose           *string                 `bson:"purpose" json:"purpose"`
	Copyright         *string                 `bson:"copyright" json:"copyright"`
	Origin            []TestScriptOrigin      `bson:"origin" json:"origin"`
	Destination       []TestScriptDestination `bson:"destination" json:"destination"`
	Metadata          *TestScriptMetadata     `bson:"metadata" json:"metadata"`
	Fixture           []TestScriptFixture     `bson:"fixture" json:"fixture"`
	Profile           []Reference             `bson:"profile" json:"profile"`
	Variable          []TestScriptVariable    `bson:"variable" json:"variable"`
	Rule              []TestScriptRule        `bson:"rule" json:"rule"`
	Ruleset           []TestScriptRuleset     `bson:"ruleset" json:"ruleset"`
	Setup             *TestScriptSetup        `bson:"setup" json:"setup"`
	Test              []TestScriptTest        `bson:"test" json:"test"`
	Teardown          *TestScriptTeardown     `bson:"teardown" json:"teardown"`
}
type TestScriptOrigin struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Index             int         `bson:"index,omitempty" json:"index,omitempty"`
	Profile           Coding      `bson:"profile,omitempty" json:"profile,omitempty"`
}
type TestScriptDestination struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Index             int         `bson:"index,omitempty" json:"index,omitempty"`
	Profile           Coding      `bson:"profile,omitempty" json:"profile,omitempty"`
}
type TestScriptMetadata struct {
	Id                *string                        `bson:"id" json:"id"`
	Extension         []Extension                    `bson:"extension" json:"extension"`
	ModifierExtension []Extension                    `bson:"modifierExtension" json:"modifierExtension"`
	Link              []TestScriptMetadataLink       `bson:"link" json:"link"`
	Capability        []TestScriptMetadataCapability `bson:"capability,omitempty" json:"capability,omitempty"`
}
type TestScriptMetadataLink struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Url               string      `bson:"url,omitempty" json:"url,omitempty"`
	Description       *string     `bson:"description" json:"description"`
}
type TestScriptMetadataCapability struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Required          *bool       `bson:"required" json:"required"`
	Validated         *bool       `bson:"validated" json:"validated"`
	Description       *string     `bson:"description" json:"description"`
	Origin            []int       `bson:"origin" json:"origin"`
	Destination       *int        `bson:"destination" json:"destination"`
	Link              []string    `bson:"link" json:"link"`
	Capabilities      Reference   `bson:"capabilities,omitempty" json:"capabilities,omitempty"`
}
type TestScriptFixture struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Autocreate        *bool       `bson:"autocreate" json:"autocreate"`
	Autodelete        *bool       `bson:"autodelete" json:"autodelete"`
	Resource          *Reference  `bson:"resource" json:"resource"`
}
type TestScriptVariable struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Name              string      `bson:"name,omitempty" json:"name,omitempty"`
	DefaultValue      *string     `bson:"defaultValue" json:"defaultValue"`
	Description       *string     `bson:"description" json:"description"`
	Expression        *string     `bson:"expression" json:"expression"`
	HeaderField       *string     `bson:"headerField" json:"headerField"`
	Hint              *string     `bson:"hint" json:"hint"`
	Path              *string     `bson:"path" json:"path"`
	SourceId          *string     `bson:"sourceId" json:"sourceId"`
}
type TestScriptRule struct {
	Id                *string               `bson:"id" json:"id"`
	Extension         []Extension           `bson:"extension" json:"extension"`
	ModifierExtension []Extension           `bson:"modifierExtension" json:"modifierExtension"`
	Resource          Reference             `bson:"resource,omitempty" json:"resource,omitempty"`
	Param             []TestScriptRuleParam `bson:"param" json:"param"`
}
type TestScriptRuleParam struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Name              string      `bson:"name,omitempty" json:"name,omitempty"`
	Value             *string     `bson:"value" json:"value"`
}
type TestScriptRuleset struct {
	Id                *string                 `bson:"id" json:"id"`
	Extension         []Extension             `bson:"extension" json:"extension"`
	ModifierExtension []Extension             `bson:"modifierExtension" json:"modifierExtension"`
	Resource          Reference               `bson:"resource,omitempty" json:"resource,omitempty"`
	Rule              []TestScriptRulesetRule `bson:"rule,omitempty" json:"rule,omitempty"`
}
type TestScriptRulesetRule struct {
	Id                *string                      `bson:"id" json:"id"`
	Extension         []Extension                  `bson:"extension" json:"extension"`
	ModifierExtension []Extension                  `bson:"modifierExtension" json:"modifierExtension"`
	RuleId            string                       `bson:"ruleId,omitempty" json:"ruleId,omitempty"`
	Param             []TestScriptRulesetRuleParam `bson:"param" json:"param"`
}
type TestScriptRulesetRuleParam struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Name              string      `bson:"name,omitempty" json:"name,omitempty"`
	Value             *string     `bson:"value" json:"value"`
}
type TestScriptSetup struct {
	Id                *string                 `bson:"id" json:"id"`
	Extension         []Extension             `bson:"extension" json:"extension"`
	ModifierExtension []Extension             `bson:"modifierExtension" json:"modifierExtension"`
	Action            []TestScriptSetupAction `bson:"action,omitempty" json:"action,omitempty"`
}
type TestScriptSetupAction struct {
	Id                *string                         `bson:"id" json:"id"`
	Extension         []Extension                     `bson:"extension" json:"extension"`
	ModifierExtension []Extension                     `bson:"modifierExtension" json:"modifierExtension"`
	Operation         *TestScriptSetupActionOperation `bson:"operation" json:"operation"`
	Assert            *TestScriptSetupActionAssert    `bson:"assert" json:"assert"`
}
type TestScriptSetupActionOperation struct {
	Id                *string                                       `bson:"id" json:"id"`
	Extension         []Extension                                   `bson:"extension" json:"extension"`
	ModifierExtension []Extension                                   `bson:"modifierExtension" json:"modifierExtension"`
	Type              *Coding                                       `bson:"type" json:"type"`
	Resource          *string                                       `bson:"resource" json:"resource"`
	Label             *string                                       `bson:"label" json:"label"`
	Description       *string                                       `bson:"description" json:"description"`
	Accept            *ContentType                                  `bson:"accept" json:"accept"`
	ContentType       *ContentType                                  `bson:"contentType" json:"contentType"`
	Destination       *int                                          `bson:"destination" json:"destination"`
	EncodeRequestUrl  *bool                                         `bson:"encodeRequestUrl" json:"encodeRequestUrl"`
	Origin            *int                                          `bson:"origin" json:"origin"`
	Params            *string                                       `bson:"params" json:"params"`
	RequestHeader     []TestScriptSetupActionOperationRequestHeader `bson:"requestHeader" json:"requestHeader"`
	RequestId         *string                                       `bson:"requestId" json:"requestId"`
	ResponseId        *string                                       `bson:"responseId" json:"responseId"`
	SourceId          *string                                       `bson:"sourceId" json:"sourceId"`
	TargetId          *string                                       `bson:"targetId" json:"targetId"`
	Url               *string                                       `bson:"url" json:"url"`
}
type TestScriptSetupActionOperationRequestHeader struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Field             string      `bson:"field,omitempty" json:"field,omitempty"`
	Value             string      `bson:"value,omitempty" json:"value,omitempty"`
}
type TestScriptSetupActionAssert struct {
	Id                        *string                             `bson:"id" json:"id"`
	Extension                 []Extension                         `bson:"extension" json:"extension"`
	ModifierExtension         []Extension                         `bson:"modifierExtension" json:"modifierExtension"`
	Label                     *string                             `bson:"label" json:"label"`
	Description               *string                             `bson:"description" json:"description"`
	Direction                 *AssertionDirectionType             `bson:"direction" json:"direction"`
	CompareToSourceId         *string                             `bson:"compareToSourceId" json:"compareToSourceId"`
	CompareToSourceExpression *string                             `bson:"compareToSourceExpression" json:"compareToSourceExpression"`
	CompareToSourcePath       *string                             `bson:"compareToSourcePath" json:"compareToSourcePath"`
	ContentType               *ContentType                        `bson:"contentType" json:"contentType"`
	Expression                *string                             `bson:"expression" json:"expression"`
	HeaderField               *string                             `bson:"headerField" json:"headerField"`
	MinimumId                 *string                             `bson:"minimumId" json:"minimumId"`
	NavigationLinks           *bool                               `bson:"navigationLinks" json:"navigationLinks"`
	Operator                  *AssertionOperatorType              `bson:"operator" json:"operator"`
	Path                      *string                             `bson:"path" json:"path"`
	RequestMethod             *TestScriptRequestMethodCode        `bson:"requestMethod" json:"requestMethod"`
	RequestURL                *string                             `bson:"requestURL" json:"requestURL"`
	Resource                  *string                             `bson:"resource" json:"resource"`
	Response                  *AssertionResponseTypes             `bson:"response" json:"response"`
	ResponseCode              *string                             `bson:"responseCode" json:"responseCode"`
	Rule                      *TestScriptSetupActionAssertRule    `bson:"rule" json:"rule"`
	Ruleset                   *TestScriptSetupActionAssertRuleset `bson:"ruleset" json:"ruleset"`
	SourceId                  *string                             `bson:"sourceId" json:"sourceId"`
	ValidateProfileId         *string                             `bson:"validateProfileId" json:"validateProfileId"`
	Value                     *string                             `bson:"value" json:"value"`
	WarningOnly               *bool                               `bson:"warningOnly" json:"warningOnly"`
}
type TestScriptSetupActionAssertRule struct {
	Id                *string                                `bson:"id" json:"id"`
	Extension         []Extension                            `bson:"extension" json:"extension"`
	ModifierExtension []Extension                            `bson:"modifierExtension" json:"modifierExtension"`
	RuleId            string                                 `bson:"ruleId,omitempty" json:"ruleId,omitempty"`
	Param             []TestScriptSetupActionAssertRuleParam `bson:"param" json:"param"`
}
type TestScriptSetupActionAssertRuleParam struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Name              string      `bson:"name,omitempty" json:"name,omitempty"`
	Value             string      `bson:"value,omitempty" json:"value,omitempty"`
}
type TestScriptSetupActionAssertRuleset struct {
	Id                *string                                  `bson:"id" json:"id"`
	Extension         []Extension                              `bson:"extension" json:"extension"`
	ModifierExtension []Extension                              `bson:"modifierExtension" json:"modifierExtension"`
	RulesetId         string                                   `bson:"rulesetId,omitempty" json:"rulesetId,omitempty"`
	Rule              []TestScriptSetupActionAssertRulesetRule `bson:"rule" json:"rule"`
}
type TestScriptSetupActionAssertRulesetRule struct {
	Id                *string                                       `bson:"id" json:"id"`
	Extension         []Extension                                   `bson:"extension" json:"extension"`
	ModifierExtension []Extension                                   `bson:"modifierExtension" json:"modifierExtension"`
	RuleId            string                                        `bson:"ruleId,omitempty" json:"ruleId,omitempty"`
	Param             []TestScriptSetupActionAssertRulesetRuleParam `bson:"param" json:"param"`
}
type TestScriptSetupActionAssertRulesetRuleParam struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	Name              string      `bson:"name,omitempty" json:"name,omitempty"`
	Value             string      `bson:"value,omitempty" json:"value,omitempty"`
}
type TestScriptTest struct {
	Id                *string                `bson:"id" json:"id"`
	Extension         []Extension            `bson:"extension" json:"extension"`
	ModifierExtension []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	Name              *string                `bson:"name" json:"name"`
	Description       *string                `bson:"description" json:"description"`
	Action            []TestScriptTestAction `bson:"action,omitempty" json:"action,omitempty"`
}
type TestScriptTestAction struct {
	Id                *string                         `bson:"id" json:"id"`
	Extension         []Extension                     `bson:"extension" json:"extension"`
	ModifierExtension []Extension                     `bson:"modifierExtension" json:"modifierExtension"`
	Operation         *TestScriptSetupActionOperation `bson:"operation" json:"operation"`
	Assert            *TestScriptSetupActionAssert    `bson:"assert" json:"assert"`
}
type TestScriptTeardown struct {
	Id                *string                    `bson:"id" json:"id"`
	Extension         []Extension                `bson:"extension" json:"extension"`
	ModifierExtension []Extension                `bson:"modifierExtension" json:"modifierExtension"`
	Action            []TestScriptTeardownAction `bson:"action,omitempty" json:"action,omitempty"`
}
type TestScriptTeardownAction struct {
	Id                *string                        `bson:"id" json:"id"`
	Extension         []Extension                    `bson:"extension" json:"extension"`
	ModifierExtension []Extension                    `bson:"modifierExtension" json:"modifierExtension"`
	Operation         TestScriptSetupActionOperation `bson:"operation,omitempty" json:"operation,omitempty"`
}

// OtherTestScript is a helper type to use the default implementations of Marshall and Unmarshal
type OtherTestScript TestScript

// MarshalJSON marshals the given TestScript as JSON into a byte slice
func (r TestScript) MarshalJSON() ([]byte, error) {
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
		OtherTestScript
		ResourceType string `json:"resourceType"`
	}{
		OtherTestScript: OtherTestScript(r),
		ResourceType:    "TestScript",
	})
}

// UnmarshalJSON unmarshals the given byte slice into TestScript
func (r *TestScript) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherTestScript)(r)); err != nil {
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
func (r TestScript) GetResourceType() ResourceType {
	return ResourceTypeTestScript
}
