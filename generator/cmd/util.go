package cmd

import (
	"fmt"
	"strings"

	"github.com/ivido/go-fhir-library/generator/fhir"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func FirstLower(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}

func Title(s string) string {
	return cases.Title(language.English).String(s[:1]) + s[1:]
}

func requiredValueSetBinding(elementDefinition fhir.ElementDefinition) *string {
	if elementDefinition.Binding != nil {
		binding := *elementDefinition.Binding
		if binding.Strength == fhir.BindingStrengthRequired {
			return binding.ValueSet
		}
	}
	return nil
}

func wasValueSetRequired(element fhir.ElementDefinition) (*string, error) {
	if url := requiredValueSetBinding(element); url != nil {
		if bytes := context.Resources["ValueSet"][*url]; bytes != nil {
			valueSet, err := fhir.UnmarshalValueSet(bytes)
			if err != nil {
				return nil, err
			}
			if name := valueSet.Name; name != nil {
				if !namePattern.MatchString(*name) {
					fmt.Printf("Skip generating an enum for a ValueSet binding to `%s` because the ValueSet has a non-conforming name.\n", *name)
				} else if len(valueSet.Compose.Include) > 1 {
					fmt.Printf("Skip generating an enum for a ValueSet binding to `%s` because the ValueSet includes more than one CodeSystem.\n", *valueSet.Name)
				} else if codeSystemUrl := canonical(valueSet.Compose.Include[0]); context.Resources["CodeSystem"][codeSystemUrl] == nil {
					fmt.Printf("Skip generating an enum for a ValueSet binding to `%s` because the ValueSet includes the non-existing CodeSystem with canonical URL `%s`.\n", *valueSet.Name, codeSystemUrl)
				} else {
					context.RequiredValueSetBindings[*url] = true
					return name, nil
				}
			} else {
				return nil, fmt.Errorf("missing name in ValueSet with canonical URL `%s`", *url)
			}
		}
	}
	return nil, nil
}

func typeCodeToTypeIdentifier(typeCode string) string {
	switch typeCode {
	case "base64Binary":
		return "string"
	case "boolean":
		return "bool"
	case "canonical":
		return "string"
	case "code":
		return "string"
	case "date":
		return "string"
	case "dateTime":
		return "string"
	case "decimal":
		return "string"
	case "id":
		return "string"
	case "instant":
		return "string"
	case "integer":
		return "int"
	case "markdown":
		return "string"
	case "oid":
		return "string"
	case "positiveInt":
		return "int"
	case "string":
		return "string"
	case "time":
		return "string"
	case "unsignedInt":
		return "int"
	case "uri":
		return "string"
	case "url":
		return "string"
	case "uuid":
		return "string"
	case "xhtml":
		return "string"
	case "http://hl7.org/fhirpath/System.String":
		return "string"
	default:
		return typeCode
	}
}
