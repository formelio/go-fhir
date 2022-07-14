// Copyright 2019 - 2021 The Samply Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"regexp"

	"github.com/dave/jennifer/jen"
	"github.com/ivido/go-fhir-library/generator/fhir"
	"github.com/spf13/cobra"
)

var namePattern = regexp.MustCompile("^[A-Z]([A-Za-z0-9_]){0,254}$")

var context = newContext()

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates Go structs from FHIR resource structure definitions.",
	Run: func(cmd *cobra.Command, args []string) {

		// Walk the directory and process all JSON FHIR definitions into context.Resources.
		readSourcesToContext(args[0])

		// Generate All Resources and find the elements that are required.
		generateResources()

		// Generate Elements that were added to context.RequiredElements during the generation of Resources
		generateElements(make(map[string]bool), context.RequiredElements)

		// Generate ValueSets that were added to context.RequiredValueSetBindings during the generation of other structs.
		generateValueSets()

		// Generate IResource interface that is used to allow polymorphism in the generated code
		// generateIResource()
	},
}

/*
	// <definition.Name> is documented here <definition.Url>"
	type <definition.Name> struct {
		..fields
	}
*/
func generateStruct(file *jen.File, definition fhir.StructureDefinition, elementDefinitions []fhir.ElementDefinition, requiredElements map[string]bool) error {
	var err error
	file.Commentf("%s is documented here %s", definition.Name, definition.Url)
	file.Type().Id(definition.Name).StructFunc(func(rootStruct *jen.Group) {
		_, err = appendFields(requiredElements, file, rootStruct, definition.Name, elementDefinitions, 1, 1)
	})
	return err
}

/*
	// MarshalJSON marshals the given <resourceType> as JSON into a byte slice
	func (r <resourceType>) MarshalJSON() ([]byte, error) {
		return json.Marshal(struct {
			Other<resourceType>
			ResourceType string `json:"resourceType"`
		}{
			Other<resourceType>: Other<resourceType>(r),
			ResourceType:            "<resourceType>",
		})
	}
*/
func generateMarshall(file *jen.File, resourceType string) {
	file.Commentf("MarshalJSON marshals the given %s as JSON into a byte slice", resourceType)
	file.Func().Params(jen.Id("r").Id(resourceType)).Id("MarshalJSON").Params().
		Params(jen.Op("[]").Byte(), jen.Error()).Block(
		jen.Return().Qual("encoding/json", "Marshal").Call(jen.Struct(
			jen.Id("Other"+resourceType),
			jen.Id("ResourceType").String().Tag(map[string]string{"json": "resourceType"}),
		).Values(jen.Dict{
			jen.Id("Other" + resourceType): jen.Id("Other" + resourceType).Call(jen.Id("r")),
			jen.Id("ResourceType"):         jen.Lit(resourceType),
		})),
	)
}

/*
	// Unmarshal<resourceType> unmarshals a <resourceType>.
	func Unmarshal<resourceType>(b []byte) (<resourceType>, error) {
		var <resourceType> <resourceType>
		if err := json.Unmarshal(b, &<resourceType>); err != nil {
			return <resourceType>, err
		}
		return <resourceType>, nil
	}
*/
func generateUnmarshall(file *jen.File, resourceType string) {
	file.Commentf("Unmarshal%s unmarshals a %s.", resourceType, resourceType)
	file.Func().Id("Unmarshal"+resourceType).
		Params(jen.Id("b").Op("[]").Byte()).
		Params(jen.Id(resourceType), jen.Error()).
		Block(
			jen.Var().Id(FirstLower(resourceType)).Id(resourceType),
			jen.If(
				jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(
					jen.Id("b"),
					jen.Op("&").Id(FirstLower(resourceType)),
				),
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(jen.Id(FirstLower(resourceType)), jen.Err()),
			),
			jen.Return(jen.Id(FirstLower(resourceType)), jen.Nil()),
		)

}

func init() {
	rootCmd.AddCommand(generateCmd)
}
