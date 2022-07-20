package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SystemRestfulInteraction is documented here http://hl7.org/fhir/ValueSet/system-restful-interaction
type SystemRestfulInteraction int

const (
	SystemRestfulInteractionRead SystemRestfulInteraction = iota
	SystemRestfulInteractionVread
	SystemRestfulInteractionUpdate
	SystemRestfulInteractionPatch
	SystemRestfulInteractionDelete
	SystemRestfulInteractionHistory
	SystemRestfulInteractionHistoryInstance
	SystemRestfulInteractionHistoryType
	SystemRestfulInteractionHistorySystem
	SystemRestfulInteractionCreate
	SystemRestfulInteractionSearch
	SystemRestfulInteractionSearchType
	SystemRestfulInteractionSearchSystem
	SystemRestfulInteractionCapabilities
	SystemRestfulInteractionTransaction
	SystemRestfulInteractionBatch
	SystemRestfulInteractionOperation
)

func (code SystemRestfulInteraction) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SystemRestfulInteraction) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "read":
		*code = SystemRestfulInteractionRead
	case "vread":
		*code = SystemRestfulInteractionVread
	case "update":
		*code = SystemRestfulInteractionUpdate
	case "patch":
		*code = SystemRestfulInteractionPatch
	case "delete":
		*code = SystemRestfulInteractionDelete
	case "history":
		*code = SystemRestfulInteractionHistory
	case "history-instance":
		*code = SystemRestfulInteractionHistoryInstance
	case "history-type":
		*code = SystemRestfulInteractionHistoryType
	case "history-system":
		*code = SystemRestfulInteractionHistorySystem
	case "create":
		*code = SystemRestfulInteractionCreate
	case "search":
		*code = SystemRestfulInteractionSearch
	case "search-type":
		*code = SystemRestfulInteractionSearchType
	case "search-system":
		*code = SystemRestfulInteractionSearchSystem
	case "capabilities":
		*code = SystemRestfulInteractionCapabilities
	case "transaction":
		*code = SystemRestfulInteractionTransaction
	case "batch":
		*code = SystemRestfulInteractionBatch
	case "operation":
		*code = SystemRestfulInteractionOperation
	default:
		return fmt.Errorf("unknown SystemRestfulInteraction code `%s`", s)
	}
	return nil
}
func (code SystemRestfulInteraction) String() string {
	return code.Code()
}
func (code SystemRestfulInteraction) Code() string {
	switch code {
	case SystemRestfulInteractionRead:
		return "read"
	case SystemRestfulInteractionVread:
		return "vread"
	case SystemRestfulInteractionUpdate:
		return "update"
	case SystemRestfulInteractionPatch:
		return "patch"
	case SystemRestfulInteractionDelete:
		return "delete"
	case SystemRestfulInteractionHistory:
		return "history"
	case SystemRestfulInteractionHistoryInstance:
		return "history-instance"
	case SystemRestfulInteractionHistoryType:
		return "history-type"
	case SystemRestfulInteractionHistorySystem:
		return "history-system"
	case SystemRestfulInteractionCreate:
		return "create"
	case SystemRestfulInteractionSearch:
		return "search"
	case SystemRestfulInteractionSearchType:
		return "search-type"
	case SystemRestfulInteractionSearchSystem:
		return "search-system"
	case SystemRestfulInteractionCapabilities:
		return "capabilities"
	case SystemRestfulInteractionTransaction:
		return "transaction"
	case SystemRestfulInteractionBatch:
		return "batch"
	case SystemRestfulInteractionOperation:
		return "operation"
	}
	return "<unknown>"
}
func (code SystemRestfulInteraction) Display() string {
	switch code {
	case SystemRestfulInteractionRead:
		return "read"
	case SystemRestfulInteractionVread:
		return "vread"
	case SystemRestfulInteractionUpdate:
		return "update"
	case SystemRestfulInteractionPatch:
		return "patch"
	case SystemRestfulInteractionDelete:
		return "delete"
	case SystemRestfulInteractionHistory:
		return "history"
	case SystemRestfulInteractionHistoryInstance:
		return "history-instance"
	case SystemRestfulInteractionHistoryType:
		return "history-type"
	case SystemRestfulInteractionHistorySystem:
		return "history-system"
	case SystemRestfulInteractionCreate:
		return "create"
	case SystemRestfulInteractionSearch:
		return "search"
	case SystemRestfulInteractionSearchType:
		return "search-type"
	case SystemRestfulInteractionSearchSystem:
		return "search-system"
	case SystemRestfulInteractionCapabilities:
		return "capabilities"
	case SystemRestfulInteractionTransaction:
		return "transaction"
	case SystemRestfulInteractionBatch:
		return "batch"
	case SystemRestfulInteractionOperation:
		return "operation"
	}
	return "<unknown>"
}
func (code SystemRestfulInteraction) Definition() string {
	switch code {
	case SystemRestfulInteractionRead:
		return "Read the current state of the resource."
	case SystemRestfulInteractionVread:
		return "Read the state of a specific version of the resource."
	case SystemRestfulInteractionUpdate:
		return "Update an existing resource by its id (or create it if it is new)."
	case SystemRestfulInteractionPatch:
		return "Update an existing resource by posting a set of changes to it."
	case SystemRestfulInteractionDelete:
		return "Delete a resource."
	case SystemRestfulInteractionHistory:
		return "Retrieve the change history for a particular resource, type of resource, or the entire system."
	case SystemRestfulInteractionHistoryInstance:
		return "Retrieve the change history for a particular resource."
	case SystemRestfulInteractionHistoryType:
		return "Retrieve the change history for all resources of a particular type."
	case SystemRestfulInteractionHistorySystem:
		return "Retrieve the change history for all resources on a system."
	case SystemRestfulInteractionCreate:
		return "Create a new resource with a server assigned id."
	case SystemRestfulInteractionSearch:
		return "Search a resource type or all resources based on some filter criteria."
	case SystemRestfulInteractionSearchType:
		return "Search all resources of the specified type based on some filter criteria."
	case SystemRestfulInteractionSearchSystem:
		return "Search all resources based on some filter criteria."
	case SystemRestfulInteractionCapabilities:
		return "Get a Capability Statement for the system."
	case SystemRestfulInteractionTransaction:
		return "Update, create or delete a set of resources as a single transaction."
	case SystemRestfulInteractionBatch:
		return "perform a set of a separate interactions in a single http operation"
	case SystemRestfulInteractionOperation:
		return "Perform an operation as defined by an OperationDefinition."
	}
	return "<unknown>"
}
