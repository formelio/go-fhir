package cmd

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/ivido/go-fhir-stu3/generator/fhir"
)

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
		_, err = generateFields(requiredElements, file, rootStruct, definition.Name, elementDefinitions, 1, 1)
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
	file.Commentf("Unmarshal%s unmarshalls a %s.", resourceType, resourceType)
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

func generateJsonTags(statement *jen.Statement, jsonTag string, required bool) {
	// If the field is not required, we add add `omitempty` to the field tags
	if !required {
		statement.Tag(map[string]string{"json": jsonTag + ",omitempty", "bson": jsonTag + ",omitempty"})
	} else {
		statement.Tag(map[string]string{"json": jsonTag, "bson": jsonTag})
	}
}

func generateOperators(statement *jen.Statement, list, required bool) {
	if list {
		statement.Op("[]")
	} else if !required {
		statement.Op("*")
	}
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
		return "float64"
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
		if typeCode != FirstUpper(typeCode) {
			fmt.Println("Unknown type code:", typeCode)
		}
		return typeCode
	}
}
