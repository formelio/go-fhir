package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// IssueType is documented here http://hl7.org/fhir/ValueSet/issue-type
type IssueType int

const (
	IssueTypeInvalid IssueType = iota
	IssueTypeStructure
	IssueTypeRequired
	IssueTypeValue
	IssueTypeInvariant
	IssueTypeSecurity
	IssueTypeLogin
	IssueTypeUnknown
	IssueTypeExpired
	IssueTypeForbidden
	IssueTypeSuppressed
	IssueTypeProcessing
	IssueTypeNotSupported
	IssueTypeDuplicate
	IssueTypeNotFound
	IssueTypeTooLong
	IssueTypeCodeInvalid
	IssueTypeExtension
	IssueTypeTooCostly
	IssueTypeBusinessRule
	IssueTypeConflict
	IssueTypeIncomplete
	IssueTypeTransient
	IssueTypeLockError
	IssueTypeNoStore
	IssueTypeException
	IssueTypeTimeout
	IssueTypeThrottled
	IssueTypeInformational
)

func (code IssueType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *IssueType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "invalid":
		*code = IssueTypeInvalid
	case "structure":
		*code = IssueTypeStructure
	case "required":
		*code = IssueTypeRequired
	case "value":
		*code = IssueTypeValue
	case "invariant":
		*code = IssueTypeInvariant
	case "security":
		*code = IssueTypeSecurity
	case "login":
		*code = IssueTypeLogin
	case "unknown":
		*code = IssueTypeUnknown
	case "expired":
		*code = IssueTypeExpired
	case "forbidden":
		*code = IssueTypeForbidden
	case "suppressed":
		*code = IssueTypeSuppressed
	case "processing":
		*code = IssueTypeProcessing
	case "not-supported":
		*code = IssueTypeNotSupported
	case "duplicate":
		*code = IssueTypeDuplicate
	case "not-found":
		*code = IssueTypeNotFound
	case "too-long":
		*code = IssueTypeTooLong
	case "code-invalid":
		*code = IssueTypeCodeInvalid
	case "extension":
		*code = IssueTypeExtension
	case "too-costly":
		*code = IssueTypeTooCostly
	case "business-rule":
		*code = IssueTypeBusinessRule
	case "conflict":
		*code = IssueTypeConflict
	case "incomplete":
		*code = IssueTypeIncomplete
	case "transient":
		*code = IssueTypeTransient
	case "lock-error":
		*code = IssueTypeLockError
	case "no-store":
		*code = IssueTypeNoStore
	case "exception":
		*code = IssueTypeException
	case "timeout":
		*code = IssueTypeTimeout
	case "throttled":
		*code = IssueTypeThrottled
	case "informational":
		*code = IssueTypeInformational
	default:
		return fmt.Errorf("unknown IssueType code `%s`", s)
	}
	return nil
}
func (code IssueType) String() string {
	return code.Code()
}
func (code IssueType) Code() string {
	switch code {
	case IssueTypeInvalid:
		return "invalid"
	case IssueTypeStructure:
		return "structure"
	case IssueTypeRequired:
		return "required"
	case IssueTypeValue:
		return "value"
	case IssueTypeInvariant:
		return "invariant"
	case IssueTypeSecurity:
		return "security"
	case IssueTypeLogin:
		return "login"
	case IssueTypeUnknown:
		return "unknown"
	case IssueTypeExpired:
		return "expired"
	case IssueTypeForbidden:
		return "forbidden"
	case IssueTypeSuppressed:
		return "suppressed"
	case IssueTypeProcessing:
		return "processing"
	case IssueTypeNotSupported:
		return "not-supported"
	case IssueTypeDuplicate:
		return "duplicate"
	case IssueTypeNotFound:
		return "not-found"
	case IssueTypeTooLong:
		return "too-long"
	case IssueTypeCodeInvalid:
		return "code-invalid"
	case IssueTypeExtension:
		return "extension"
	case IssueTypeTooCostly:
		return "too-costly"
	case IssueTypeBusinessRule:
		return "business-rule"
	case IssueTypeConflict:
		return "conflict"
	case IssueTypeIncomplete:
		return "incomplete"
	case IssueTypeTransient:
		return "transient"
	case IssueTypeLockError:
		return "lock-error"
	case IssueTypeNoStore:
		return "no-store"
	case IssueTypeException:
		return "exception"
	case IssueTypeTimeout:
		return "timeout"
	case IssueTypeThrottled:
		return "throttled"
	case IssueTypeInformational:
		return "informational"
	}
	return "<unknown>"
}
func (code IssueType) Display() string {
	switch code {
	case IssueTypeInvalid:
		return "Invalid Content"
	case IssueTypeStructure:
		return "Structural Issue"
	case IssueTypeRequired:
		return "Required element missing"
	case IssueTypeValue:
		return "Element value invalid"
	case IssueTypeInvariant:
		return "Validation rule failed"
	case IssueTypeSecurity:
		return "Security Problem"
	case IssueTypeLogin:
		return "Login Required"
	case IssueTypeUnknown:
		return "Unknown User"
	case IssueTypeExpired:
		return "Session Expired"
	case IssueTypeForbidden:
		return "Forbidden"
	case IssueTypeSuppressed:
		return "Information  Suppressed"
	case IssueTypeProcessing:
		return "Processing Failure"
	case IssueTypeNotSupported:
		return "Content not supported"
	case IssueTypeDuplicate:
		return "Duplicate"
	case IssueTypeNotFound:
		return "Not Found"
	case IssueTypeTooLong:
		return "Content Too Long"
	case IssueTypeCodeInvalid:
		return "Invalid Code"
	case IssueTypeExtension:
		return "Unacceptable Extension"
	case IssueTypeTooCostly:
		return "Operation Too Costly"
	case IssueTypeBusinessRule:
		return "Business Rule Violation"
	case IssueTypeConflict:
		return "Edit Version Conflict"
	case IssueTypeIncomplete:
		return "Incomplete Results"
	case IssueTypeTransient:
		return "Transient Issue"
	case IssueTypeLockError:
		return "Lock Error"
	case IssueTypeNoStore:
		return "No Store Available"
	case IssueTypeException:
		return "Exception"
	case IssueTypeTimeout:
		return "Timeout"
	case IssueTypeThrottled:
		return "Throttled"
	case IssueTypeInformational:
		return "Informational Note"
	}
	return "<unknown>"
}
func (code IssueType) Definition() string {
	switch code {
	case IssueTypeInvalid:
		return "Content invalid against the specification or a profile."
	case IssueTypeStructure:
		return "A structural issue in the content such as wrong namespace, or unable to parse the content completely, or invalid json syntax."
	case IssueTypeRequired:
		return "A required element is missing."
	case IssueTypeValue:
		return "An element value is invalid."
	case IssueTypeInvariant:
		return "A content validation rule failed - e.g. a schematron rule."
	case IssueTypeSecurity:
		return "An authentication/authorization/permissions issue of some kind."
	case IssueTypeLogin:
		return "The client needs to initiate an authentication process."
	case IssueTypeUnknown:
		return "The user or system was not able to be authenticated (either there is no process, or the proferred token is unacceptable)."
	case IssueTypeExpired:
		return "User session expired; a login may be required."
	case IssueTypeForbidden:
		return "The user does not have the rights to perform this action."
	case IssueTypeSuppressed:
		return "Some information was not or may not have been returned due to business rules, consent or privacy rules, or access permission constraints.  This information may be accessible through alternate processes."
	case IssueTypeProcessing:
		return "Processing issues. These are expected to be final e.g. there is no point resubmitting the same content unchanged."
	case IssueTypeNotSupported:
		return "The resource or profile is not supported."
	case IssueTypeDuplicate:
		return "An attempt was made to create a duplicate record."
	case IssueTypeNotFound:
		return "The reference provided was not found. In a pure RESTful environment, this would be an HTTP 404 error, but this code may be used where the content is not found further into the application architecture."
	case IssueTypeTooLong:
		return "Provided content is too long (typically, this is a denial of service protection type of error)."
	case IssueTypeCodeInvalid:
		return "The code or system could not be understood, or it was not valid in the context of a particular ValueSet.code."
	case IssueTypeExtension:
		return "An extension was found that was not acceptable, could not be resolved, or a modifierExtension was not recognized."
	case IssueTypeTooCostly:
		return "The operation was stopped to protect server resources; e.g. a request for a value set expansion on all of SNOMED CT."
	case IssueTypeBusinessRule:
		return "The content/operation failed to pass some business rule, and so could not proceed."
	case IssueTypeConflict:
		return "Content could not be accepted because of an edit conflict (i.e. version aware updates) (In a pure RESTful environment, this would be an HTTP 404 error, but this code may be used where the conflict is discovered further into the application architecture.)"
	case IssueTypeIncomplete:
		return "Not all data sources typically accessed could be reached, or responded in time, so the returned information may not be complete."
	case IssueTypeTransient:
		return "Transient processing issues. The system receiving the error may be able to resubmit the same content once an underlying issue is resolved."
	case IssueTypeLockError:
		return "A resource/record locking failure (usually in an underlying database)."
	case IssueTypeNoStore:
		return "The persistent store is unavailable; e.g. the database is down for maintenance or similar action."
	case IssueTypeException:
		return "An unexpected internal error has occurred."
	case IssueTypeTimeout:
		return "An internal timeout has occurred."
	case IssueTypeThrottled:
		return "The system is not prepared to handle this request due to load management."
	case IssueTypeInformational:
		return "A message unrelated to the processing success of the completed operation (examples of the latter include things like reminders of password expiry, system maintenance times, etc.)."
	}
	return "<unknown>"
}
