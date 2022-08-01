package fhir

import (
	"bytes"
	"encoding/json"
)

// TestScript is documented here http://hl7.org/fhir/StructureDefinition/TestScript
type TestScript struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                    `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                  `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                  `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative               `bson:"text,omitempty" json:"text,omitempty"`
	RawContained      []json.RawMessage        `bson:"contained,omitempty" json:"contained,omitempty"`
	Contained         []IResource              `bson:"-,omitempty" json:"-,omitempty"`
	Extension         []*Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string                   `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        *Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version           *string                  `bson:"version,omitempty" json:"version,omitempty"`
	Name              string                   `bson:"name,omitempty" json:"name,omitempty"`
	Title             *string                  `bson:"title,omitempty" json:"title,omitempty"`
	Status            PublicationStatus        `bson:"status,omitempty" json:"status,omitempty"`
	Experimental      *bool                    `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date              *string                  `bson:"date,omitempty" json:"date,omitempty"`
	Publisher         *string                  `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []*ContactDetail         `bson:"contact,omitempty" json:"contact,omitempty"`
	Description       *string                  `bson:"description,omitempty" json:"description,omitempty"`
	UseContext        []*UsageContext          `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction      []*CodeableConcept       `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose           *string                  `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright         *string                  `bson:"copyright,omitempty" json:"copyright,omitempty"`
	Origin            []*TestScriptOrigin      `bson:"origin,omitempty" json:"origin,omitempty"`
	Destination       []*TestScriptDestination `bson:"destination,omitempty" json:"destination,omitempty"`
	Metadata          *TestScriptMetadata      `bson:"metadata,omitempty" json:"metadata,omitempty"`
	Fixture           []*TestScriptFixture     `bson:"fixture,omitempty" json:"fixture,omitempty"`
	Profile           []*Reference             `bson:"profile,omitempty" json:"profile,omitempty"`
	Variable          []*TestScriptVariable    `bson:"variable,omitempty" json:"variable,omitempty"`
	Rule              []*TestScriptRule        `bson:"rule,omitempty" json:"rule,omitempty"`
	Ruleset           []*TestScriptRuleset     `bson:"ruleset,omitempty" json:"ruleset,omitempty"`
	Setup             *TestScriptSetup         `bson:"setup,omitempty" json:"setup,omitempty"`
	Test              []*TestScriptTest        `bson:"test,omitempty" json:"test,omitempty"`
	Teardown          *TestScriptTeardown      `bson:"teardown,omitempty" json:"teardown,omitempty"`
}
type TestScriptOrigin struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Index             int          `bson:"index,omitempty" json:"index,omitempty"`
	Profile           Coding       `bson:"profile,omitempty" json:"profile,omitempty"`
}
type TestScriptDestination struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Index             int          `bson:"index,omitempty" json:"index,omitempty"`
	Profile           Coding       `bson:"profile,omitempty" json:"profile,omitempty"`
}
type TestScriptMetadata struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Link              []*TestScriptMetadataLink      `bson:"link,omitempty" json:"link,omitempty"`
	Capability        []TestScriptMetadataCapability `bson:"capability,omitempty" json:"capability,omitempty"`
}
type TestScriptMetadataLink struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               string       `bson:"url,omitempty" json:"url,omitempty"`
	Description       *string      `bson:"description,omitempty" json:"description,omitempty"`
}
type TestScriptMetadataCapability struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Required          *bool        `bson:"required,omitempty" json:"required,omitempty"`
	Validated         *bool        `bson:"validated,omitempty" json:"validated,omitempty"`
	Description       *string      `bson:"description,omitempty" json:"description,omitempty"`
	Origin            []*int       `bson:"origin,omitempty" json:"origin,omitempty"`
	Destination       *int         `bson:"destination,omitempty" json:"destination,omitempty"`
	Link              []*string    `bson:"link,omitempty" json:"link,omitempty"`
	Capabilities      Reference    `bson:"capabilities,omitempty" json:"capabilities,omitempty"`
}
type TestScriptFixture struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Autocreate        *bool        `bson:"autocreate,omitempty" json:"autocreate,omitempty"`
	Autodelete        *bool        `bson:"autodelete,omitempty" json:"autodelete,omitempty"`
	Resource          *Reference   `bson:"resource,omitempty" json:"resource,omitempty"`
}
type TestScriptVariable struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string       `bson:"name,omitempty" json:"name,omitempty"`
	DefaultValue      *string      `bson:"defaultValue,omitempty" json:"defaultValue,omitempty"`
	Description       *string      `bson:"description,omitempty" json:"description,omitempty"`
	Expression        *string      `bson:"expression,omitempty" json:"expression,omitempty"`
	HeaderField       *string      `bson:"headerField,omitempty" json:"headerField,omitempty"`
	Hint              *string      `bson:"hint,omitempty" json:"hint,omitempty"`
	Path              *string      `bson:"path,omitempty" json:"path,omitempty"`
	SourceId          *string      `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
}
type TestScriptRule struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Resource          Reference              `bson:"resource,omitempty" json:"resource,omitempty"`
	Param             []*TestScriptRuleParam `bson:"param,omitempty" json:"param,omitempty"`
}
type TestScriptRuleParam struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string       `bson:"name,omitempty" json:"name,omitempty"`
	Value             *string      `bson:"value,omitempty" json:"value,omitempty"`
}
type TestScriptRuleset struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Resource          Reference               `bson:"resource,omitempty" json:"resource,omitempty"`
	Rule              []TestScriptRulesetRule `bson:"rule,omitempty" json:"rule,omitempty"`
}
type TestScriptRulesetRule struct {
	Id                *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	RuleId            string                        `bson:"ruleId,omitempty" json:"ruleId,omitempty"`
	Param             []*TestScriptRulesetRuleParam `bson:"param,omitempty" json:"param,omitempty"`
}
type TestScriptRulesetRuleParam struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string       `bson:"name,omitempty" json:"name,omitempty"`
	Value             *string      `bson:"value,omitempty" json:"value,omitempty"`
}
type TestScriptSetup struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Action            []TestScriptSetupAction `bson:"action,omitempty" json:"action,omitempty"`
}
type TestScriptSetupAction struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Operation         *TestScriptSetupActionOperation `bson:"operation,omitempty" json:"operation,omitempty"`
	Assert            *TestScriptSetupActionAssert    `bson:"assert,omitempty" json:"assert,omitempty"`
}
type TestScriptSetupActionOperation struct {
	Id                *string                                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *Coding                                        `bson:"type,omitempty" json:"type,omitempty"`
	Resource          *string                                        `bson:"resource,omitempty" json:"resource,omitempty"`
	Label             *string                                        `bson:"label,omitempty" json:"label,omitempty"`
	Description       *string                                        `bson:"description,omitempty" json:"description,omitempty"`
	Accept            *ContentType                                   `bson:"accept,omitempty" json:"accept,omitempty"`
	ContentType       *ContentType                                   `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Destination       *int                                           `bson:"destination,omitempty" json:"destination,omitempty"`
	EncodeRequestUrl  *bool                                          `bson:"encodeRequestUrl,omitempty" json:"encodeRequestUrl,omitempty"`
	Origin            *int                                           `bson:"origin,omitempty" json:"origin,omitempty"`
	Params            *string                                        `bson:"params,omitempty" json:"params,omitempty"`
	RequestHeader     []*TestScriptSetupActionOperationRequestHeader `bson:"requestHeader,omitempty" json:"requestHeader,omitempty"`
	RequestId         *string                                        `bson:"requestId,omitempty" json:"requestId,omitempty"`
	ResponseId        *string                                        `bson:"responseId,omitempty" json:"responseId,omitempty"`
	SourceId          *string                                        `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
	TargetId          *string                                        `bson:"targetId,omitempty" json:"targetId,omitempty"`
	Url               *string                                        `bson:"url,omitempty" json:"url,omitempty"`
}
type TestScriptSetupActionOperationRequestHeader struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Field             string       `bson:"field,omitempty" json:"field,omitempty"`
	Value             string       `bson:"value,omitempty" json:"value,omitempty"`
}
type TestScriptSetupActionAssert struct {
	Id                        *string                             `bson:"id,omitempty" json:"id,omitempty"`
	Extension                 []*Extension                        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []*Extension                        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Label                     *string                             `bson:"label,omitempty" json:"label,omitempty"`
	Description               *string                             `bson:"description,omitempty" json:"description,omitempty"`
	Direction                 *AssertionDirectionType             `bson:"direction,omitempty" json:"direction,omitempty"`
	CompareToSourceId         *string                             `bson:"compareToSourceId,omitempty" json:"compareToSourceId,omitempty"`
	CompareToSourceExpression *string                             `bson:"compareToSourceExpression,omitempty" json:"compareToSourceExpression,omitempty"`
	CompareToSourcePath       *string                             `bson:"compareToSourcePath,omitempty" json:"compareToSourcePath,omitempty"`
	ContentType               *ContentType                        `bson:"contentType,omitempty" json:"contentType,omitempty"`
	Expression                *string                             `bson:"expression,omitempty" json:"expression,omitempty"`
	HeaderField               *string                             `bson:"headerField,omitempty" json:"headerField,omitempty"`
	MinimumId                 *string                             `bson:"minimumId,omitempty" json:"minimumId,omitempty"`
	NavigationLinks           *bool                               `bson:"navigationLinks,omitempty" json:"navigationLinks,omitempty"`
	Operator                  *AssertionOperatorType              `bson:"operator,omitempty" json:"operator,omitempty"`
	Path                      *string                             `bson:"path,omitempty" json:"path,omitempty"`
	RequestMethod             *TestScriptRequestMethodCode        `bson:"requestMethod,omitempty" json:"requestMethod,omitempty"`
	RequestURL                *string                             `bson:"requestURL,omitempty" json:"requestURL,omitempty"`
	Resource                  *string                             `bson:"resource,omitempty" json:"resource,omitempty"`
	Response                  *AssertionResponseTypes             `bson:"response,omitempty" json:"response,omitempty"`
	ResponseCode              *string                             `bson:"responseCode,omitempty" json:"responseCode,omitempty"`
	Rule                      *TestScriptSetupActionAssertRule    `bson:"rule,omitempty" json:"rule,omitempty"`
	Ruleset                   *TestScriptSetupActionAssertRuleset `bson:"ruleset,omitempty" json:"ruleset,omitempty"`
	SourceId                  *string                             `bson:"sourceId,omitempty" json:"sourceId,omitempty"`
	ValidateProfileId         *string                             `bson:"validateProfileId,omitempty" json:"validateProfileId,omitempty"`
	Value                     *string                             `bson:"value,omitempty" json:"value,omitempty"`
	WarningOnly               *bool                               `bson:"warningOnly,omitempty" json:"warningOnly,omitempty"`
}
type TestScriptSetupActionAssertRule struct {
	Id                *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	RuleId            string                                  `bson:"ruleId,omitempty" json:"ruleId,omitempty"`
	Param             []*TestScriptSetupActionAssertRuleParam `bson:"param,omitempty" json:"param,omitempty"`
}
type TestScriptSetupActionAssertRuleParam struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string       `bson:"name,omitempty" json:"name,omitempty"`
	Value             string       `bson:"value,omitempty" json:"value,omitempty"`
}
type TestScriptSetupActionAssertRuleset struct {
	Id                *string                                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	RulesetId         string                                    `bson:"rulesetId,omitempty" json:"rulesetId,omitempty"`
	Rule              []*TestScriptSetupActionAssertRulesetRule `bson:"rule,omitempty" json:"rule,omitempty"`
}
type TestScriptSetupActionAssertRulesetRule struct {
	Id                *string                                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	RuleId            string                                         `bson:"ruleId,omitempty" json:"ruleId,omitempty"`
	Param             []*TestScriptSetupActionAssertRulesetRuleParam `bson:"param,omitempty" json:"param,omitempty"`
}
type TestScriptSetupActionAssertRulesetRuleParam struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string       `bson:"name,omitempty" json:"name,omitempty"`
	Value             string       `bson:"value,omitempty" json:"value,omitempty"`
}
type TestScriptTest struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              *string                `bson:"name,omitempty" json:"name,omitempty"`
	Description       *string                `bson:"description,omitempty" json:"description,omitempty"`
	Action            []TestScriptTestAction `bson:"action,omitempty" json:"action,omitempty"`
}
type TestScriptTestAction struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Operation         *TestScriptSetupActionOperation `bson:"operation,omitempty" json:"operation,omitempty"`
	Assert            *TestScriptSetupActionAssert    `bson:"assert,omitempty" json:"assert,omitempty"`
}
type TestScriptTeardown struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Action            []TestScriptTeardownAction `bson:"action,omitempty" json:"action,omitempty"`
}
type TestScriptTeardownAction struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []*Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []*Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
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
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(struct {
		ResourceType string `json:"resourceType"`
		OtherTestScript
	}{
		OtherTestScript: OtherTestScript(r),
		ResourceType:    "TestScript",
	})
	return buffer.Bytes(), err
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
