package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// QuestionnaireItemType is documented here http://hl7.org/fhir/ValueSet/item-type
type QuestionnaireItemType int

const (
	QuestionnaireItemTypeGroup QuestionnaireItemType = iota
	QuestionnaireItemTypeDisplay
	QuestionnaireItemTypeQuestion
	QuestionnaireItemTypeBoolean
	QuestionnaireItemTypeDecimal
	QuestionnaireItemTypeInteger
	QuestionnaireItemTypeDate
	QuestionnaireItemTypeDateTime
	QuestionnaireItemTypeTime
	QuestionnaireItemTypeString
	QuestionnaireItemTypeText
	QuestionnaireItemTypeUrl
	QuestionnaireItemTypeChoice
	QuestionnaireItemTypeOpenChoice
	QuestionnaireItemTypeAttachment
	QuestionnaireItemTypeReference
	QuestionnaireItemTypeQuantity
)

func (code QuestionnaireItemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *QuestionnaireItemType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "group":
		*code = QuestionnaireItemTypeGroup
	case "display":
		*code = QuestionnaireItemTypeDisplay
	case "question":
		*code = QuestionnaireItemTypeQuestion
	case "boolean":
		*code = QuestionnaireItemTypeBoolean
	case "decimal":
		*code = QuestionnaireItemTypeDecimal
	case "integer":
		*code = QuestionnaireItemTypeInteger
	case "date":
		*code = QuestionnaireItemTypeDate
	case "dateTime":
		*code = QuestionnaireItemTypeDateTime
	case "time":
		*code = QuestionnaireItemTypeTime
	case "string":
		*code = QuestionnaireItemTypeString
	case "text":
		*code = QuestionnaireItemTypeText
	case "url":
		*code = QuestionnaireItemTypeUrl
	case "choice":
		*code = QuestionnaireItemTypeChoice
	case "open-choice":
		*code = QuestionnaireItemTypeOpenChoice
	case "attachment":
		*code = QuestionnaireItemTypeAttachment
	case "reference":
		*code = QuestionnaireItemTypeReference
	case "quantity":
		*code = QuestionnaireItemTypeQuantity
	default:
		return fmt.Errorf("unknown QuestionnaireItemType code `%s`", s)
	}
	return nil
}
func (code QuestionnaireItemType) String() string {
	return code.Code()
}
func (code QuestionnaireItemType) Code() string {
	switch code {
	case QuestionnaireItemTypeGroup:
		return "group"
	case QuestionnaireItemTypeDisplay:
		return "display"
	case QuestionnaireItemTypeQuestion:
		return "question"
	case QuestionnaireItemTypeBoolean:
		return "boolean"
	case QuestionnaireItemTypeDecimal:
		return "decimal"
	case QuestionnaireItemTypeInteger:
		return "integer"
	case QuestionnaireItemTypeDate:
		return "date"
	case QuestionnaireItemTypeDateTime:
		return "dateTime"
	case QuestionnaireItemTypeTime:
		return "time"
	case QuestionnaireItemTypeString:
		return "string"
	case QuestionnaireItemTypeText:
		return "text"
	case QuestionnaireItemTypeUrl:
		return "url"
	case QuestionnaireItemTypeChoice:
		return "choice"
	case QuestionnaireItemTypeOpenChoice:
		return "open-choice"
	case QuestionnaireItemTypeAttachment:
		return "attachment"
	case QuestionnaireItemTypeReference:
		return "reference"
	case QuestionnaireItemTypeQuantity:
		return "quantity"
	}
	return "<unknown>"
}
func (code QuestionnaireItemType) Display() string {
	switch code {
	case QuestionnaireItemTypeGroup:
		return "Group"
	case QuestionnaireItemTypeDisplay:
		return "Display"
	case QuestionnaireItemTypeQuestion:
		return "Question"
	case QuestionnaireItemTypeBoolean:
		return "Boolean"
	case QuestionnaireItemTypeDecimal:
		return "Decimal"
	case QuestionnaireItemTypeInteger:
		return "Integer"
	case QuestionnaireItemTypeDate:
		return "Date"
	case QuestionnaireItemTypeDateTime:
		return "Date Time"
	case QuestionnaireItemTypeTime:
		return "Time"
	case QuestionnaireItemTypeString:
		return "String"
	case QuestionnaireItemTypeText:
		return "Text"
	case QuestionnaireItemTypeUrl:
		return "Url"
	case QuestionnaireItemTypeChoice:
		return "Choice"
	case QuestionnaireItemTypeOpenChoice:
		return "Open Choice"
	case QuestionnaireItemTypeAttachment:
		return "Attachment"
	case QuestionnaireItemTypeReference:
		return "Reference"
	case QuestionnaireItemTypeQuantity:
		return "Quantity"
	}
	return "<unknown>"
}
func (code QuestionnaireItemType) Definition() string {
	switch code {
	case QuestionnaireItemTypeGroup:
		return "An item with no direct answer but should have at least one child item."
	case QuestionnaireItemTypeDisplay:
		return "Text for display that will not capture an answer or have child items."
	case QuestionnaireItemTypeQuestion:
		return "An item that defines a specific answer to be captured, and may have child items.\n(the answer provided in the QuestionnaireResponse should be of the defined datatype)"
	case QuestionnaireItemTypeBoolean:
		return "Question with a yes/no answer (valueBoolean)"
	case QuestionnaireItemTypeDecimal:
		return "Question with is a real number answer (valueDecimal)"
	case QuestionnaireItemTypeInteger:
		return "Question with an integer answer (valueInteger)"
	case QuestionnaireItemTypeDate:
		return "Question with a date answer (valueDate)"
	case QuestionnaireItemTypeDateTime:
		return "Question with a date and time answer (valueDateTime)"
	case QuestionnaireItemTypeTime:
		return "Question with a time (hour:minute:second) answer independent of date. (valueTime)"
	case QuestionnaireItemTypeString:
		return "Question with a short (few words to short sentence) free-text entry answer (valueString)"
	case QuestionnaireItemTypeText:
		return "Question with a long (potentially multi-paragraph) free-text entry answer (valueString)"
	case QuestionnaireItemTypeUrl:
		return "Question with a URL (website, FTP site, etc.) answer (valueUri)"
	case QuestionnaireItemTypeChoice:
		return "Question with a Coding drawn from a list of options (specified in either the option property, or via the valueset referenced in the options property) as an answer (valueCoding)"
	case QuestionnaireItemTypeOpenChoice:
		return "Answer is a Coding drawn from a list of options (as with the choice type) or a free-text entry in a string (valueCoding or valueString)"
	case QuestionnaireItemTypeAttachment:
		return "Question with binary content such as a image, PDF, etc. as an answer (valueAttachment)"
	case QuestionnaireItemTypeReference:
		return "Question with a reference to another resource (practitioner, organization, etc.) as an answer (valueReference)"
	case QuestionnaireItemTypeQuantity:
		return "Question with a combination of a numeric value and unit, potentially with a comparator (<, >, etc.) as an answer. (valueQuantity)\nThere is an extension 'http://hl7.org/fhir/StructureDefinition/questionnaire-unit' that can be used to define what unit whould be captured (or the a unit that has a ucum conversion from the provided unit)"
	}
	return "<unknown>"
}
