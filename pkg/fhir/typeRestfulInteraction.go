package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// TypeRestfulInteraction is documented here http://hl7.org/fhir/ValueSet/type-restful-interaction
type TypeRestfulInteraction int

const (
	TypeRestfulInteractionRead TypeRestfulInteraction = iota
	TypeRestfulInteractionVread
	TypeRestfulInteractionUpdate
	TypeRestfulInteractionPatch
	TypeRestfulInteractionDelete
	TypeRestfulInteractionHistory
	TypeRestfulInteractionHistoryInstance
	TypeRestfulInteractionHistoryType
	TypeRestfulInteractionHistorySystem
	TypeRestfulInteractionCreate
	TypeRestfulInteractionSearch
	TypeRestfulInteractionSearchType
	TypeRestfulInteractionSearchSystem
	TypeRestfulInteractionCapabilities
	TypeRestfulInteractionTransaction
	TypeRestfulInteractionBatch
	TypeRestfulInteractionOperation
)

func (code TypeRestfulInteraction) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *TypeRestfulInteraction) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "read":
		*code = TypeRestfulInteractionRead
	case "vread":
		*code = TypeRestfulInteractionVread
	case "update":
		*code = TypeRestfulInteractionUpdate
	case "patch":
		*code = TypeRestfulInteractionPatch
	case "delete":
		*code = TypeRestfulInteractionDelete
	case "history":
		*code = TypeRestfulInteractionHistory
	case "history-instance":
		*code = TypeRestfulInteractionHistoryInstance
	case "history-type":
		*code = TypeRestfulInteractionHistoryType
	case "history-system":
		*code = TypeRestfulInteractionHistorySystem
	case "create":
		*code = TypeRestfulInteractionCreate
	case "search":
		*code = TypeRestfulInteractionSearch
	case "search-type":
		*code = TypeRestfulInteractionSearchType
	case "search-system":
		*code = TypeRestfulInteractionSearchSystem
	case "capabilities":
		*code = TypeRestfulInteractionCapabilities
	case "transaction":
		*code = TypeRestfulInteractionTransaction
	case "batch":
		*code = TypeRestfulInteractionBatch
	case "operation":
		*code = TypeRestfulInteractionOperation
	default:
		return fmt.Errorf("unknown TypeRestfulInteraction code `%s`", s)
	}
	return nil
}
func (code TypeRestfulInteraction) String() string {
	return code.Code()
}
func (code TypeRestfulInteraction) Code() string {
	switch code {
	case TypeRestfulInteractionRead:
		return "read"
	case TypeRestfulInteractionVread:
		return "vread"
	case TypeRestfulInteractionUpdate:
		return "update"
	case TypeRestfulInteractionPatch:
		return "patch"
	case TypeRestfulInteractionDelete:
		return "delete"
	case TypeRestfulInteractionHistory:
		return "history"
	case TypeRestfulInteractionHistoryInstance:
		return "history-instance"
	case TypeRestfulInteractionHistoryType:
		return "history-type"
	case TypeRestfulInteractionHistorySystem:
		return "history-system"
	case TypeRestfulInteractionCreate:
		return "create"
	case TypeRestfulInteractionSearch:
		return "search"
	case TypeRestfulInteractionSearchType:
		return "search-type"
	case TypeRestfulInteractionSearchSystem:
		return "search-system"
	case TypeRestfulInteractionCapabilities:
		return "capabilities"
	case TypeRestfulInteractionTransaction:
		return "transaction"
	case TypeRestfulInteractionBatch:
		return "batch"
	case TypeRestfulInteractionOperation:
		return "operation"
	}
	return "<unknown>"
}
func (code TypeRestfulInteraction) Display() string {
	switch code {
	case TypeRestfulInteractionRead:
		return "read"
	case TypeRestfulInteractionVread:
		return "vread"
	case TypeRestfulInteractionUpdate:
		return "update"
	case TypeRestfulInteractionPatch:
		return "patch"
	case TypeRestfulInteractionDelete:
		return "delete"
	case TypeRestfulInteractionHistory:
		return "history"
	case TypeRestfulInteractionHistoryInstance:
		return "history-instance"
	case TypeRestfulInteractionHistoryType:
		return "history-type"
	case TypeRestfulInteractionHistorySystem:
		return "history-system"
	case TypeRestfulInteractionCreate:
		return "create"
	case TypeRestfulInteractionSearch:
		return "search"
	case TypeRestfulInteractionSearchType:
		return "search-type"
	case TypeRestfulInteractionSearchSystem:
		return "search-system"
	case TypeRestfulInteractionCapabilities:
		return "capabilities"
	case TypeRestfulInteractionTransaction:
		return "transaction"
	case TypeRestfulInteractionBatch:
		return "batch"
	case TypeRestfulInteractionOperation:
		return "operation"
	}
	return "<unknown>"
}
func (code TypeRestfulInteraction) Definition() string {
	switch code {
	case TypeRestfulInteractionRead:
		return "Read the current state of the resource."
	case TypeRestfulInteractionVread:
		return "Read the state of a specific version of the resource."
	case TypeRestfulInteractionUpdate:
		return "Update an existing resource by its id (or create it if it is new)."
	case TypeRestfulInteractionPatch:
		return "Update an existing resource by posting a set of changes to it."
	case TypeRestfulInteractionDelete:
		return "Delete a resource."
	case TypeRestfulInteractionHistory:
		return "Retrieve the change history for a particular resource, type of resource, or the entire system."
	case TypeRestfulInteractionHistoryInstance:
		return "Retrieve the change history for a particular resource."
	case TypeRestfulInteractionHistoryType:
		return "Retrieve the change history for all resources of a particular type."
	case TypeRestfulInteractionHistorySystem:
		return "Retrieve the change history for all resources on a system."
	case TypeRestfulInteractionCreate:
		return "Create a new resource with a server assigned id."
	case TypeRestfulInteractionSearch:
		return "Search a resource type or all resources based on some filter criteria."
	case TypeRestfulInteractionSearchType:
		return "Search all resources of the specified type based on some filter criteria."
	case TypeRestfulInteractionSearchSystem:
		return "Search all resources based on some filter criteria."
	case TypeRestfulInteractionCapabilities:
		return "Get a Capability Statement for the system."
	case TypeRestfulInteractionTransaction:
		return "Update, create or delete a set of resources as a single transaction."
	case TypeRestfulInteractionBatch:
		return "perform a set of a separate interactions in a single http operation"
	case TypeRestfulInteractionOperation:
		return "Perform an operation as defined by an OperationDefinition."
	}
	return "<unknown>"
}
