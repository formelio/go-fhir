package cmd

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/ivido/go-fhir-library/generator/fhir"
)

func generateElements(alreadyGeneratedTypes, typesToGenerate map[string]bool) error {
	moreRequiredTypes := make(map[string]bool, 0)
	for name := range typesToGenerate {
		bytes := context.Resources["StructureDefinition"][name]
		if bytes == nil {
			return fmt.Errorf("missing StructureDefinition with name `%s`", name)
		}
		structureDefinition, err := fhir.UnmarshalStructureDefinition(bytes)
		if err != nil {
			return err
		}
		goFile, err := generateElement(moreRequiredTypes, structureDefinition)
		if err != nil {
			return err
		}
		err = goFile.Save(FirstLower(structureDefinition.Name) + ".go")
		if err != nil {
			return err
		}
		alreadyGeneratedTypes[name] = true
	}
	for name := range alreadyGeneratedTypes {
		delete(moreRequiredTypes, name)
	}
	if len(moreRequiredTypes) > 0 {
		return generateElements(alreadyGeneratedTypes, moreRequiredTypes)
	}
	return nil
}

func generateElement(requiredTypes map[string]bool, definition fhir.StructureDefinition) (*jen.File, error) {
	elementDefinitions := definition.Snapshot.Element
	if len(elementDefinitions) == 0 {
		return nil, fmt.Errorf("missing element definitions in structure definition `%s`", definition.Name)
	}

	fmt.Printf("Generate Go sources for Element: %s\n", definition.Name)
	file := jen.NewFile("fhir")

	// generate struct
	err := generateStruct(file, definition, elementDefinitions, requiredTypes)
	if err != nil {
		return nil, err
	}

	return file, nil
}
