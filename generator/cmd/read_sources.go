package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ivido/go-fhir-library/generator/fhir"
)

func readSourcesToContext(dir string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(info.Name(), ".json") {
			return err
		}

		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		fmt.Printf("Generate Go sources from file: %s\n", path)
		resource, err := UnmarshalResource(bytes)
		if err != nil {
			return err
		}

		return processResource(resource, bytes)
	})

	if err != nil {
		panic(err)
	}
}

func processResource(resource Resource, resourceBytes []byte) error {
	switch resource.ResourceType {
	case "Bundle":
		bundle, err := fhir.UnmarshalBundle(resourceBytes)
		if err != nil {
			return err
		}
		for _, entry := range bundle.Entry {
			entryResource, err := UnmarshalResource(*entry.Resource)
			if err != nil {
				return err
			}
			processResource(entryResource, *entry.Resource)
		}
	case "StructureDefinition":
		if resource.Name != nil && *resource.Name != "Resource" && *resource.Name != "Element" && *resource.Name != "BackboneElement" {
			context.Resources[resource.ResourceType][*resource.Name] = resourceBytes
		}
	case "ValueSet", "CodeSystem":
		if resource.Url != nil {
			if resource.Version != nil {
				context.Resources[resource.ResourceType][*resource.Url+"|"+*resource.Version] = resourceBytes
				context.Resources[resource.ResourceType][*resource.Url] = resourceBytes
			} else {
				context.Resources[resource.ResourceType][*resource.Url] = resourceBytes
			}
		}
	}
	return nil
}
