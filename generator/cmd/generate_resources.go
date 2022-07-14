package cmd

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/dave/jennifer/jen"
	"github.com/ivido/go-fhir-library/generator/fhir"
)

func generateResources() {
	for _, bytes := range context.Resources["StructureDefinition"] {
		structureDefinition, err := fhir.UnmarshalStructureDefinition(bytes)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if structureDefinition.Kind == fhir.StructureDefinitionKindResource {
			goFile, err := generateResource(structureDefinition)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = goFile.Save(FirstLower(structureDefinition.Name) + ".go")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}

func generateResource(definition fhir.StructureDefinition) (*jen.File, error) {
	elementDefinitions := definition.Snapshot.Element
	if len(elementDefinitions) == 0 {
		return nil, fmt.Errorf("missing element definitions in structure definition `%s`", definition.Name)
	}

	fmt.Printf("Generate Go sources for R: %s\n", definition.Name)

	var err error

	// Create new file
	file := jen.NewFile("fhir")

	// Generate the struct
	generateStruct(file, definition, elementDefinitions, context.RequiredElements)
	if err != nil {
		return nil, err
	}

	/* generate OtherType that is used in the Marshall function
	type Other<resourceType> <resourceType>
	*/
	file.Type().Id("Other" + definition.Name).Id(definition.Name)

	// generate marshal
	generateMarshall(file, definition.Name)

	// generate unmarshal
	generateUnmarshall(file, definition.Name)
	return file, nil
}

func appendFields(requiredTypes map[string]bool, file *jen.File, fields *jen.Group, parentName string, elementDefinitions []fhir.ElementDefinition, start, level int) (int, error) {
	for i := start; i < len(elementDefinitions); i++ {
		element := elementDefinitions[i]
		pathParts := strings.Split(element.Path, ".")
		if len(pathParts) != level+1 {

			return i, nil
		}

		// direct childs
		name := Title(pathParts[level])

		// support contained resources later
		if name == "Contained" {
			continue
		}
		switch len(element.Type) {
		case 0:
			if element.ContentReference != nil && (*element.ContentReference)[:1] == "#" {
				statement := fields.Id(name)

				if *element.Max == "*" {
					statement.Op("[]")
				} else if *element.Min == 0 {
					statement.Op("*")
				}

				typeIdentifier := ""
				for _, pathPart := range strings.Split((*element.ContentReference)[1:], ".") {
					typeIdentifier = typeIdentifier + Title(pathPart)
				}
				statement.Id(typeIdentifier).Tag(map[string]string{"json": pathParts[level] + ",omitempty", "bson": pathParts[level] + ",omitempty"})
			}
		// support polymorphic elements later
		case 1:
			statement := fields.Id(name)
			switch element.Type[0].Code {
			case "code":
				if *element.Max == "*" {
					statement.Op("[]")
				} else if *element.Min == 0 {
					statement.Op("*")
				}

				if name, err := wasValueSetRequired(element); err != nil {
					return 0, err
				} else if name != nil {
					statement.Id(*name)
				} else {
					statement.Id("string")
				}
			case "Resource":
				statement.Qual("encoding/json", "RawMessage")
			default:
				if *element.Max == "*" {
					statement.Op("[]")
				} else if *element.Min == 0 {
					statement.Op("*")
				}

				var typeIdentifier string
				if parentName == "Element" && name == "Id" ||
					parentName == "Extension" && name == "Url" {
					typeIdentifier = "string"
				} else {
					typeIdentifier = typeCodeToTypeIdentifier(element.Type[0].Code)
				}
				if typeIdentifier == "Element" || typeIdentifier == "BackboneElement" {
					backboneElementName := parentName + name
					statement.Id(backboneElementName)
					var err error
					file.Type().Id(backboneElementName).StructFunc(func(childFields *jen.Group) {
						//var err error
						i, err = appendFields(requiredTypes, file, childFields, backboneElementName, elementDefinitions, i+1, level+1)
					})
					if err != nil {
						return 0, err
					}
					i--
				} else {
					if unicode.IsUpper(rune(typeIdentifier[0])) {
						requiredTypes[typeIdentifier] = true
					}
					statement.Id(typeIdentifier)
				}
			}

			if *element.Min == 0 {
				statement.Tag(map[string]string{"json": pathParts[level] + ",omitempty", "bson": pathParts[level] + ",omitempty"})
			} else {
				statement.Tag(map[string]string{"json": pathParts[level], "bson": pathParts[level]})
			}
		}

	}
	return 0, nil
}
