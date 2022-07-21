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
func generateStruct(file *jen.File, definition fhir.StructureDefinition, elementDefinitions []fhir.ElementDefinition, requiredElements map[string]bool) ([]string, error) {
	var err error
	var resourceFields []string
	file.Commentf("%s is documented here %s", definition.Name, definition.Url)
	file.Type().Id(definition.Name).StructFunc(func(rootStruct *jen.Group) {
		_, resourceFields, err = generateFields(requiredElements, file, rootStruct, definition.Name, elementDefinitions, 1, 1)
	})
	return resourceFields, err
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

/* generates either of the following two, depending if it is required or not
`bson:"<jsonTag>"           json:"<jsonTag>"`
`bson:"<jsonTag>,omitempty" json:"<jsonTag>,omitempty"`
*/
func generateJsonTags(statement *jen.Statement, jsonTag string, required bool) {
	// If the field is not required, we add add `omitempty` to the field tags
	if !required {
		statement.Tag(map[string]string{"json": jsonTag + ",omitempty", "bson": jsonTag + ",omitempty"})
	} else {
		statement.Tag(map[string]string{"json": jsonTag, "bson": jsonTag})
	}
}

// generates either a slice operator or a pointer operator
func generateOperators(statement *jen.Statement, list, required bool) {
	if list {
		statement.Op("[]")
	} else if !required {
		statement.Op("*")
	}
}

/*
// Returns the resourceType of the resource, makes this resource an instance of IResource
func (r <resourceType>) GetResourceType() ResourceType {
	return ResourceType<resourceType
}
*/
func generateGetResourceType(file *jen.File, resourceType string) {
	file.Commentf("Returns the resourceType of the resource, makes this resource an instance of IResource")
	file.Func().Params(jen.Id("r").Id(resourceType)).Id("GetResourceType").Params().Id("ResourceType").Block(
		jen.Return().Id("ResourceType" + resourceType),
	)
}

/*
// Other<resourceType> is a helper type to use the default implementations of Marshall and Unmarshal
type Other<resourceType> <resourceType>
*/
func generateOtherType(file *jen.File, resourceType string) {
	file.Commentf("Other%s is a helper type to use the default implementations of Marshall and Unmarshal", resourceType)
	file.Type().Id("Other" + resourceType).Id(resourceType)
}

/*
func (b *BundleEntry) UnmarshalJSON(data []byte) error {
	// Unmarshal into all defined fields
	err := json.Unmarshal(bytes, (*otherBundleEntry)(bundleEntry))
	if err != nil {
		return err
	}

	var resource Resource
	if err := json.Unmarshal(data, &resource); err != nil {
		return err
	}
	// Get the correct implementation of IResource from the ResourceType
	var resourceInterface = ResourceTypeToIResource(resource.ResourceType)
	err := json.Unmarshal(*b.RawResource, resourceInterface)
	if err != nil {
		return err
	}

	b.Resource = &resourceInterface
	b.RawResource = nil
	return nil
}
*/
func generateElementUnmarshall(file *jen.File, resourceFields []string, typeIdentifier string) {
	file.Func().Params(jen.Id("r").Id("*" + typeIdentifier)).Id("UnmarshalJSON").Params(jen.Id("data").Op("[]").Byte()).Error().BlockFunc(func(unmarshal *jen.Group) {
		unmarshal.If(jen.Id("err").Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Params(jen.Op("*").Id("Other"+typeIdentifier)).Params(jen.Id("r"))), jen.Id("err").Op("!=").Nil()).Block(
			jen.Return(jen.Id("err")),
		)

		for _, resourceField := range resourceFields {
			unmarshal.Var().Id(resourceField).Id("Resource")
			unmarshal.If(jen.Id("err").Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("r").Dot("Raw"+resourceField), jen.Op("&").Id(resourceField)), jen.Id("err").Op("!=").Nil()).Block(
				jen.Return(jen.Id("err")),
			)
			unmarshal.Comment("Get the correct implementation of IResource from the ResourceType")
			unmarshal.Var().Id(resourceField + "Interface").Op("=").Id("ResourceTypeToIResource").Call(jen.Id(resourceField).Dot("ResourceType"))
			unmarshal.Comment("Unmarshal the raw resource into the correct implementation of IResource")
			unmarshal.If(jen.Id("err").Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("r").Dot("Raw"+resourceField), jen.Id(resourceField+"Interface")), jen.Id("err").Op("!=").Nil()).Block(
				jen.Return(jen.Id("err")),
			)
			unmarshal.Id("r").Dot(resourceField).Op("=").Id(resourceField + "Interface")
			unmarshal.Id("r").Dot("Raw" + resourceField).Op("=").Nil()
			unmarshal.Return(jen.Nil())
		}
	})
}

/*
func (r *<typeIdentifier>) MarshalJSON() ([]byte, error) {
	if <typeIdentifier>.<resourceField> != nil {
		r, err := json.Marshal(<typeIdentifier>.<resourceField>)
		if err != nil {
			return nil, err
		}
		<typeIdentifier>.Raw<resourceField> = r
	}
	return json.Marshal((*Other<typeIdentifier>)(<typeIdentifier>))
}
*/
func generateElementMarshall(file *jen.File, resourceFields []string, typeIdentifier string) {
	file.Func().Params(jen.Id("r").Op("*").Id(typeIdentifier)).Id("MarshalJSON").Params().Params(jen.Op("[]").Byte(), jen.Error()).BlockFunc(func(marshal *jen.Group) {
		for _, resourceField := range resourceFields {
			marshal.If(jen.Id("r").Dot(resourceField).Op("!=").Nil()).Block(
				jen.Id(resourceField).Op(",").Id("err").Op(":=").Qual("encoding/json", "Marshal").Call(jen.Id("r").Dot(resourceField)),
				jen.If(jen.Id("err").Op("!=").Nil()).Block(
					jen.Return(jen.Nil(), jen.Id("err")),
				),
				jen.Id("r").Dot("Raw"+resourceField).Op("=").Id(resourceField),
			)
		}
		// TODO Decide how to handle RawResource, always keep empty? NO IDEA.
		marshal.Return(jen.Qual("encoding/json", "Marshal").Call(jen.Params(jen.Op("*").Id("Other" + typeIdentifier)).Params(jen.Id("r"))))
	})
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
