package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// GuidePageKind is documented here http://hl7.org/fhir/ValueSet/guide-page-kind
type GuidePageKind int

const (
	GuidePageKindPage GuidePageKind = iota
	GuidePageKindExample
	GuidePageKindList
	GuidePageKindInclude
	GuidePageKindDirectory
	GuidePageKindDictionary
	GuidePageKindToc
	GuidePageKindResource
)

func (code GuidePageKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *GuidePageKind) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "page":
		*code = GuidePageKindPage
	case "example":
		*code = GuidePageKindExample
	case "list":
		*code = GuidePageKindList
	case "include":
		*code = GuidePageKindInclude
	case "directory":
		*code = GuidePageKindDirectory
	case "dictionary":
		*code = GuidePageKindDictionary
	case "toc":
		*code = GuidePageKindToc
	case "resource":
		*code = GuidePageKindResource
	default:
		return fmt.Errorf("unknown GuidePageKind code `%s`", s)
	}
	return nil
}
func (code GuidePageKind) String() string {
	return code.Code()
}
func (code GuidePageKind) Code() string {
	switch code {
	case GuidePageKindPage:
		return "page"
	case GuidePageKindExample:
		return "example"
	case GuidePageKindList:
		return "list"
	case GuidePageKindInclude:
		return "include"
	case GuidePageKindDirectory:
		return "directory"
	case GuidePageKindDictionary:
		return "dictionary"
	case GuidePageKindToc:
		return "toc"
	case GuidePageKindResource:
		return "resource"
	}
	return "<unknown>"
}
func (code GuidePageKind) Display() string {
	switch code {
	case GuidePageKindPage:
		return "Page"
	case GuidePageKindExample:
		return "Example"
	case GuidePageKindList:
		return "List"
	case GuidePageKindInclude:
		return "Include"
	case GuidePageKindDirectory:
		return "Directory"
	case GuidePageKindDictionary:
		return "Dictionary"
	case GuidePageKindToc:
		return "Table Of Contents"
	case GuidePageKindResource:
		return "Resource"
	}
	return "<unknown>"
}
func (code GuidePageKind) Definition() string {
	switch code {
	case GuidePageKindPage:
		return "This is a page of content that is included in the implementation guide. It has no particular function."
	case GuidePageKindExample:
		return "This is a page that represents a human readable rendering of an example."
	case GuidePageKindList:
		return "This is a page that represents a list of resources of one or more types."
	case GuidePageKindInclude:
		return "This is a page showing where an included guide is injected."
	case GuidePageKindDirectory:
		return "This is a page that lists the resources of a given type, and also creates pages for all the listed types as other pages in the section."
	case GuidePageKindDictionary:
		return "This is a page that creates the listed resources as a dictionary."
	case GuidePageKindToc:
		return "This is a generated page that contains the table of contents."
	case GuidePageKindResource:
		return "This is a page that represents a presented resource. This is typically used for generated conformance resource presentations."
	}
	return "<unknown>"
}
