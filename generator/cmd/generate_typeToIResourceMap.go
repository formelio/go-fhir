package cmd

import (
	"sort"

	"github.com/dave/jennifer/jen"
)

func generateResourceTypeToIResourceMap() {
	// Create new file
	file := jen.NewFile("fhir")

	// Sort the resourceTypes
	sort.Strings(context.GeneratedResourceTypes)

	file.Func().Id("ResourceTypeToIResource").Params(jen.Id("resourceType").Id("ResourceType")).Id("IResource").BlockFunc(func(g *jen.Group) {
		g.Switch(jen.Id("resourceType")).BlockFunc(func(g *jen.Group) {
			for _, resourceType := range context.GeneratedResourceTypes {
				g.Case(jen.Id("ResourceType" + resourceType)).Block(jen.Return(jen.Op("&").Id(resourceType).Values()))
			}
			g.Default().Block(jen.Return(jen.Nil()))
		})
	})
	err := file.Save("resourceType_to_iResource.go")
	if err != nil {
		panic(err)
	}
}
