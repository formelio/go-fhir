package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// StructureMapTransform is documented here http://hl7.org/fhir/ValueSet/map-transform
type StructureMapTransform int

const (
	StructureMapTransformCreate StructureMapTransform = iota
	StructureMapTransformCopy
	StructureMapTransformTruncate
	StructureMapTransformEscape
	StructureMapTransformCast
	StructureMapTransformAppend
	StructureMapTransformTranslate
	StructureMapTransformReference
	StructureMapTransformDateOp
	StructureMapTransformUuid
	StructureMapTransformPointer
	StructureMapTransformEvaluate
	StructureMapTransformCc
	StructureMapTransformC
	StructureMapTransformQty
	StructureMapTransformId
	StructureMapTransformCp
)

func (code StructureMapTransform) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *StructureMapTransform) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "create":
		*code = StructureMapTransformCreate
	case "copy":
		*code = StructureMapTransformCopy
	case "truncate":
		*code = StructureMapTransformTruncate
	case "escape":
		*code = StructureMapTransformEscape
	case "cast":
		*code = StructureMapTransformCast
	case "append":
		*code = StructureMapTransformAppend
	case "translate":
		*code = StructureMapTransformTranslate
	case "reference":
		*code = StructureMapTransformReference
	case "dateOp":
		*code = StructureMapTransformDateOp
	case "uuid":
		*code = StructureMapTransformUuid
	case "pointer":
		*code = StructureMapTransformPointer
	case "evaluate":
		*code = StructureMapTransformEvaluate
	case "cc":
		*code = StructureMapTransformCc
	case "c":
		*code = StructureMapTransformC
	case "qty":
		*code = StructureMapTransformQty
	case "id":
		*code = StructureMapTransformId
	case "cp":
		*code = StructureMapTransformCp
	default:
		return fmt.Errorf("unknown StructureMapTransform code `%s`", s)
	}
	return nil
}
func (code StructureMapTransform) String() string {
	return code.Code()
}
func (code StructureMapTransform) Code() string {
	switch code {
	case StructureMapTransformCreate:
		return "create"
	case StructureMapTransformCopy:
		return "copy"
	case StructureMapTransformTruncate:
		return "truncate"
	case StructureMapTransformEscape:
		return "escape"
	case StructureMapTransformCast:
		return "cast"
	case StructureMapTransformAppend:
		return "append"
	case StructureMapTransformTranslate:
		return "translate"
	case StructureMapTransformReference:
		return "reference"
	case StructureMapTransformDateOp:
		return "dateOp"
	case StructureMapTransformUuid:
		return "uuid"
	case StructureMapTransformPointer:
		return "pointer"
	case StructureMapTransformEvaluate:
		return "evaluate"
	case StructureMapTransformCc:
		return "cc"
	case StructureMapTransformC:
		return "c"
	case StructureMapTransformQty:
		return "qty"
	case StructureMapTransformId:
		return "id"
	case StructureMapTransformCp:
		return "cp"
	}
	return "<unknown>"
}
func (code StructureMapTransform) Display() string {
	switch code {
	case StructureMapTransformCreate:
		return "create"
	case StructureMapTransformCopy:
		return "copy"
	case StructureMapTransformTruncate:
		return "truncate"
	case StructureMapTransformEscape:
		return "escape"
	case StructureMapTransformCast:
		return "cast"
	case StructureMapTransformAppend:
		return "append"
	case StructureMapTransformTranslate:
		return "translate"
	case StructureMapTransformReference:
		return "reference"
	case StructureMapTransformDateOp:
		return "dateOp"
	case StructureMapTransformUuid:
		return "uuid"
	case StructureMapTransformPointer:
		return "pointer"
	case StructureMapTransformEvaluate:
		return "evaluate"
	case StructureMapTransformCc:
		return "cc"
	case StructureMapTransformC:
		return "c"
	case StructureMapTransformQty:
		return "qty"
	case StructureMapTransformId:
		return "id"
	case StructureMapTransformCp:
		return "cp"
	}
	return "<unknown>"
}
func (code StructureMapTransform) Definition() string {
	switch code {
	case StructureMapTransformCreate:
		return "create(type : string) - type is passed through to the application on the standard API, and must be known by it"
	case StructureMapTransformCopy:
		return "copy(source)"
	case StructureMapTransformTruncate:
		return "truncate(source, length) - source must be stringy type"
	case StructureMapTransformEscape:
		return "escape(source, fmt1, fmt2) - change source from one kind of escaping to another (plain, java, xml, json). note that this is for when the string itself is escaped"
	case StructureMapTransformCast:
		return "cast(source, type?) - case source from one type to another. target type can be left as implicit if there is one and only one target type known"
	case StructureMapTransformAppend:
		return "append(source...) - source is element or string"
	case StructureMapTransformTranslate:
		return "translate(source, uri_of_map) - use the translate operation"
	case StructureMapTransformReference:
		return "reference(source : object) - return a string that references the provided tree properly"
	case StructureMapTransformDateOp:
		return "Perform a date operation. *Parameters to be documented*"
	case StructureMapTransformUuid:
		return "Generate a random UUID (in lowercase). No Parameters"
	case StructureMapTransformPointer:
		return "Return the appropriate string to put in a reference that refers to the resource provided as a parameter"
	case StructureMapTransformEvaluate:
		return "Execute the supplied fluentpath expression and use the value returned by that"
	case StructureMapTransformCc:
		return "Create a CodeableConcept. Parameters = (text) or (system. Code[, display])"
	case StructureMapTransformC:
		return "Create a Coding. Parameters = (system. Code[, display])"
	case StructureMapTransformQty:
		return "Create a quantity. Parameters = (text) or (value, unit, [system, code]) where text is the natural representation e.g. [comparator]value[space]unit"
	case StructureMapTransformId:
		return "Create an identifier. Parameters = (system, value[, type]) where type is a code from the identifier type value set"
	case StructureMapTransformCp:
		return "Create a contact details. Parameters = (value) or (system, value). If no system is provided, the system should be inferred from the content of the value"
	}
	return "<unknown>"
}
