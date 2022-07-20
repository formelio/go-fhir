package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ObservationRelationshipType is documented here http://hl7.org/fhir/ValueSet/observation-relationshiptypes
type ObservationRelationshipType int

const (
	ObservationRelationshipTypeHasMember ObservationRelationshipType = iota
	ObservationRelationshipTypeDerivedFrom
	ObservationRelationshipTypeSequelTo
	ObservationRelationshipTypeReplaces
	ObservationRelationshipTypeQualifiedBy
	ObservationRelationshipTypeInterferedBy
)

func (code ObservationRelationshipType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ObservationRelationshipType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "has-member":
		*code = ObservationRelationshipTypeHasMember
	case "derived-from":
		*code = ObservationRelationshipTypeDerivedFrom
	case "sequel-to":
		*code = ObservationRelationshipTypeSequelTo
	case "replaces":
		*code = ObservationRelationshipTypeReplaces
	case "qualified-by":
		*code = ObservationRelationshipTypeQualifiedBy
	case "interfered-by":
		*code = ObservationRelationshipTypeInterferedBy
	default:
		return fmt.Errorf("unknown ObservationRelationshipType code `%s`", s)
	}
	return nil
}
func (code ObservationRelationshipType) String() string {
	return code.Code()
}
func (code ObservationRelationshipType) Code() string {
	switch code {
	case ObservationRelationshipTypeHasMember:
		return "has-member"
	case ObservationRelationshipTypeDerivedFrom:
		return "derived-from"
	case ObservationRelationshipTypeSequelTo:
		return "sequel-to"
	case ObservationRelationshipTypeReplaces:
		return "replaces"
	case ObservationRelationshipTypeQualifiedBy:
		return "qualified-by"
	case ObservationRelationshipTypeInterferedBy:
		return "interfered-by"
	}
	return "<unknown>"
}
func (code ObservationRelationshipType) Display() string {
	switch code {
	case ObservationRelationshipTypeHasMember:
		return "Has Member"
	case ObservationRelationshipTypeDerivedFrom:
		return "Derived From"
	case ObservationRelationshipTypeSequelTo:
		return "Sequel To"
	case ObservationRelationshipTypeReplaces:
		return "Replaces"
	case ObservationRelationshipTypeQualifiedBy:
		return "Qualified By"
	case ObservationRelationshipTypeInterferedBy:
		return "Interfered By"
	}
	return "<unknown>"
}
func (code ObservationRelationshipType) Definition() string {
	switch code {
	case ObservationRelationshipTypeHasMember:
		return "This observation is a group observation (e.g. a battery, a panel of tests, a set of vital sign measurements) that includes the target as a member of the group."
	case ObservationRelationshipTypeDerivedFrom:
		return "The target resource (Observation or QuestionnaireResponse) is part of the information from which this observation value is derived. (e.g. calculated anion gap, Apgar score)  NOTE:  \"derived-from\" is the only logical choice when referencing QuestionnaireResponse."
	case ObservationRelationshipTypeSequelTo:
		return "This observation follows the target observation (e.g. timed tests such as Glucose Tolerance Test)."
	case ObservationRelationshipTypeReplaces:
		return "This observation replaces a previous observation (i.e. a revised value). The target observation is now obsolete."
	case ObservationRelationshipTypeQualifiedBy:
		return "The value of the target observation qualifies (refines) the semantics of the source observation (e.g. a lipemia measure target from a plasma measure)."
	case ObservationRelationshipTypeInterferedBy:
		return "The value of the target observation interferes (degrades quality, or prevents valid observation) with the semantics of the source observation (e.g. a hemolysis measure target from a plasma potassium measure, which has no value)."
	}
	return "<unknown>"
}
