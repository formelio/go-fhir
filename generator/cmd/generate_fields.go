package cmd

import (
	"fmt"
	"log"
	"strings"
	"unicode"

	"github.com/dave/jennifer/jen"
	"github.com/ivido/go-fhir-library/generator/fhir"
)

type FieldType int

const (
	// For fields that are referenced, we only need to create a field with that type. Type itself will be created in the non-referenced definition of it
	FieldTypeReferenced FieldType = iota
	// For fields that have type Resource, we do not process values of this type
	// TODO: Process resources
	FieldTypeResource
	// For fields that have type Code, we only create structs for it if the value set binding is required.
	FieldTypeCode
	// For fields that only have a single type that does need to be generated
	FieldTypeSingularPrimitive
	// For fields that only have a single type that needs to be generated in the same file
	FieldTypeSingularComplex
	// For fields that have polymorphic types
	FieldTypePolymorphic
	// For fields that could not be processed to a known type
	FieldTypeUnknown
)

func generateFields(
	requiredTypes map[string]bool,
	file *jen.File, fields *jen.Group,
	parentName string,
	elementDefinitions []fhir.ElementDefinition,
	start, level int,
) (int, error) {
	for i := start; i < len(elementDefinitions); i++ {
		element := elementDefinitions[i]
		pathSplit := strings.Split(element.Path, ".")

		// If the path is not long enough, or too long, we return
		if len(pathSplit) != level+1 {
			return i, nil
		}

		name := FirstUpper(pathSplit[level])
		fieldType := findFieldType(element)

		// We find the type identifier based on the type of the field
		typeIdentifier, err := findTypeIdentifier(fieldType, element, parentName, name)
		if err != nil {
			return 0, err
		}

		if fieldType == FieldTypePolymorphic {
			generatePolymorphicType(requiredTypes, file, fields, typeIdentifier, element)
			continue
		}

		statement := fields.Id(name)
		// Add field operators to denote whether a field is a list or a nilable value
		generateOperators(statement, *element.Max == "*", *element.Min != 0)

		if fieldType == FieldTypeResource {
			// For Resource we set the field type to json.RawMessage
			statement.Qual("encoding/json", "RawMessage")
		} else {
			// If the field uses a complex element that needs to be generated, we add the complex element type to the required types
			if fieldType == FieldTypeSingularPrimitive && unicode.IsUpper(rune(typeIdentifier[0])) {
				requiredTypes[typeIdentifier] = true
			} else if fieldType == FieldTypeSingularComplex {
				// We need to generate a complex element type in the same file,
				i, err = generateElementType(requiredTypes, file, typeIdentifier, elementDefinitions, i, level)
				if err != nil {
					return 0, err
				}
			}

			// For other tpes of fields we set the field type to the type identifier found
			statement.Id(typeIdentifier)
		}

		// If the field is not required, we add add `omitempty` to the field tags
		generateJsonTags(statement, pathSplit[level], *element.Min == 0)
	}
	return 0, nil
}

// Finds the type of the field, based on this type different generation logic is used
func findFieldType(element fhir.ElementDefinition) FieldType {
	switch len(element.Type) {
	// The element doesn't have a type, check if it contains a reference
	case 0:
		if element.ContentReference != nil && (*element.ContentReference)[:1] == "#" {
			return FieldTypeReferenced
		} else {
			return FieldTypeUnknown
		}
	// The field has a single type, return the type based on the .code
	case 1:
		if elementTypeCode := element.Type[0].Code; elementTypeCode == "Resource" {
			return FieldTypeResource
		} else if elementTypeCode == "code" {
			return FieldTypeCode
		} else if (elementTypeCode == "Element" && element.Type[0].Code != "id") || elementTypeCode == "BackboneElement" {
			return FieldTypeSingularComplex
		} else {
			return FieldTypeSingularPrimitive
		}
	// The field has multiple types, it can be a polymorphic field or a single-type field that allows only certain values for the type.
	// For example only Patient or Practitioner references, we do not support this yet.
	default:
		elementTypeCode := element.Type[0].Code
		for _, elementType := range element.Type {
			if elementType.Code != elementTypeCode {
				fmt.Println("Multiple types for " + element.Path)
				return FieldTypePolymorphic
			}
		}
		// All of the typeCodes are the same, we treat this as a single-type field
		element.Type = element.Type[:1]
		return findFieldType(element)
	}
}

func findTypeIdentifier(fieldType FieldType, element fhir.ElementDefinition, parentName, name string) (string, error) {
	var typeIdentifier string
	switch fieldType {
	// For Referenced fields, we build the field Type identifier from the referenced path
	case FieldTypeReferenced:
		for _, pathPart := range strings.Split((*element.ContentReference)[1:], ".") {
			typeIdentifier = typeIdentifier + FirstUpper(pathPart)
		}
	// For code fields, check if value set is required, save it to context.RequiredValueSetBindings
	// Set the typeIdentifier to either the name of the value set or string if it is not required
	case FieldTypeCode:
		if name, err := getValueSetNameIfRequired(element); err != nil {
			return "", err
		} else if name != nil {
			typeIdentifier = *name
		} else {
			typeIdentifier = "string"
		}
	// For singular fields, we build the type identifier based on the type.Code and parentName
	case FieldTypeSingularPrimitive:
		typeIdentifier = typeCodeToTypeIdentifier(element.Type[0].Code)
	case FieldTypeSingularComplex:
		typeIdentifier = parentName + name
	case FieldTypePolymorphic:
		typeIdentifier = name
	}
	return typeIdentifier, nil
}

func generateElementType(
	requiredTypes map[string]bool,
	file *jen.File,
	typeIdentifier string,
	elementDefinitions []fhir.ElementDefinition,
	startIndex, level int,
) (int, error) {
	var err error
	var indexToContinueFrom int
	file.Type().Id(typeIdentifier).StructFunc(func(childFields *jen.Group) {
		indexToContinueFrom, err = generateFields(requiredTypes, file, childFields, typeIdentifier, elementDefinitions, startIndex+1, level+1)
	})
	// We substract 1 to the index to continue from, because it will be incremented in the next for loop
	indexToContinueFrom--
	return indexToContinueFrom, err
}

func generatePolymorphicType(
	requiredTypes map[string]bool,
	file *jen.File, fields *jen.Group,
	typeIdentifier string,
	element fhir.ElementDefinition) {

	if !strings.HasSuffix(typeIdentifier, "[x]") {
		log.Panicf("Polymorphic type does not end with [x]: %s", typeIdentifier)
	} else {
		typeIdentifier = typeIdentifier[:len(typeIdentifier)-3]
	}
	addedTypes := []string{}
	for _, elementType := range element.Type {
		elementTypeCode := elementType.Code
		if Contains(addedTypes, elementTypeCode) {
			continue
		}
		addedTypes = append(addedTypes, elementTypeCode)

		if unicode.IsUpper(rune(elementTypeCode[0])) {
			requiredTypes[elementType.Code] = true
		} else {
			elementTypeCode = typeCodeToTypeIdentifier(elementTypeCode)
		}

		fieldName := typeIdentifier + FirstUpper(elementType.Code)
		jsonFieldName := strings.ToLower(fieldName[:1]) + fieldName[1:]

		addPolymorphicStatement(fields, fieldName, jsonFieldName, elementTypeCode, false, *element.Max == "*")
	}

}

func addPolymorphicStatement(fields *jen.Group, fieldName, jsonFieldName, typeIdentifier string, required, list bool) {
	statement := fields.Id(fieldName)
	generateOperators(statement, list, required)
	statement.Id(typeIdentifier)
	generateJsonTags(statement, jsonFieldName, required)
}

// Checks whether the value set is required, return it's URL if it is
func requiredValueSetBinding(elementDefinition fhir.ElementDefinition) *string {
	if elementDefinition.Binding != nil {
		binding := *elementDefinition.Binding
		if binding.Strength == fhir.BindingStrengthRequired {
			return binding.ValueSet
		}
	}
	return nil
}

// Gets the name of the ValueSet if the element contains one that has a required binding.
// Saves the value set in context.RequiredValueSetBindings
func getValueSetNameIfRequired(element fhir.ElementDefinition) (*string, error) {
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
