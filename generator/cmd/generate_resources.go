package cmd

import (
	"fmt"
	"os"

	"github.com/dave/jennifer/jen"
	"github.com/ivido/go-fhir-stu3/generator/fhir"
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

	fmt.Printf("Generate Go sources for Resource: %s\n", definition.Name)

	var err error

	// Create new file
	file := jen.NewFile("fhir")

	// Generate the struct
	generateStruct(file, definition, elementDefinitions, context.RequiredElements)
	if err != nil {
		return nil, err
	}

	/* generate OtherType that is used in the Marshall function:
	type Other<resourceType> <resourceType>
	*/
	file.Type().Id("Other" + definition.Name).Id(definition.Name)

	// generate marshal
	generateMarshall(file, definition.Name)

	// generate unmarshal
	generateUnmarshall(file, definition.Name)
	return file, nil
}
