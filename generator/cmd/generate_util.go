package cmd

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/formelio/go-fhir/generator/fhir"
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
		IF(HAS_CONTAINED):
		------------------
			// If the field has contained resources, we need to marshal them individually and store them in RawContained
			if len(r.Contained) > 0 {
				r.RawContained = make([]json.RawMessage, len(r.Contained))
				var err error
				for i, contained := range r.Contained {
					r.RawContained[i], err = json.Marshal(contained)
					if err != nil {
						return nil, err
					}
				}
			}
		------------------

		buffer := bytes.NewBuffer([]byte{})
		jsonEncoder := json.NewEncoder(buffer)
		jsonEncoder.SetEscapeHTML(false)
		err := jsonEncoder.Encode(struct {
			ResourceType string `json:"resourceType"`
			Other<resourceType>
		}{
			Other<resourceType>: Other<resourceType>(r),
			ResourceType: "<resourceType>",
		})
		return buffer.Bytes(), err
	}
*/
func generateResourceMarshal(file *jen.File, resourceType string, hasContained bool) {
	file.Commentf("MarshalJSON marshals the given %s as JSON into a byte slice", resourceType)
	file.Func().Params(jen.Id("r").Id(resourceType)).Id("MarshalJSON").Params().Params(jen.Op("[]").Byte(), jen.Error()).BlockFunc(func(marshalBlock *jen.Group) {
		marshalBlock.Comment("If the field has contained resources, we need to marshal them individually and store them in .RawContained")
		if hasContained {
			marshalBlock.If(jen.Len(jen.Id("r").Dot("Contained")).Op(">").Lit(0)).BlockFunc(func(containedBlock *jen.Group) {
				containedBlock.Var().Id("err").Error()
				containedBlock.Id("r").Dot("RawContained").Op("=").Make(jen.Op("[]").Qual("encoding/json", "RawMessage"), jen.Len(jen.Id("r").Dot("Contained")))
				containedBlock.For(jen.List(jen.Id("i"), jen.Id("contained")).Op(":=").Range().Id("r").Dot("Contained")).BlockFunc(func(containedLoopBlock *jen.Group) {
					containedLoopBlock.List(jen.Id("r").Dot("RawContained").Index(jen.Id("i")), jen.Id("err")).Op("=").Qual("encoding/json", "Marshal").Call(jen.Id("contained"))
					containedLoopBlock.If(jen.Id("err").Op("!=").Nil()).Block(jen.Return(jen.List(jen.Nil(), jen.Id("err"))))
				})
			})
		}

		marshalBlock.Id("buffer").Op(":=").Qual("bytes", "NewBuffer").Call(jen.Op("[]").Byte().Op("{}"))
		marshalBlock.Id("jsonEncoder").Op(":=").Qual("encoding/json", "NewEncoder").Call(jen.Id("buffer"))
		marshalBlock.Id("jsonEncoder").Dot("SetEscapeHTML").Call(jen.False())
		marshalBlock.Id("err").Op(":=").Id("jsonEncoder").Dot("Encode").Call(jen.Struct(
			jen.Id("ResourceType").String().Tag(map[string]string{"json": "resourceType"}),
			jen.Id("Other"+resourceType),
		).Values(jen.Dict{
			jen.Id("ResourceType"):         jen.Lit(resourceType),
			jen.Id("Other" + resourceType): jen.Id("Other" + resourceType).Call(jen.Id("r")),
		}))
		marshalBlock.Return(jen.List(jen.Id("buffer").Dot("Bytes").Call(), jen.Id("err")))
	})
}

/*
	func (r *<resourceType>) UnmarshalJSON(data []byte) error {
		if err := json.Unmarshal(data, (*Other<resourceType>)(r)); err != nil {
			return err
		}
		if len(r.RawContained) > 0 {
			var err error
			r.Contained = make([]IResource, len(r.RawContained))
			for i, rawContained := range r.RawContained {
				r.Contained[i], err = UnmarshalResource(rawContained)
				if err != nil {
					return err
				}
			}
		}
		r.RawContained = nil
		return nil
	}
*/
func generateResourceUnmarshal(file *jen.File, resourceType string, hasContained bool) {
	file.Commentf("UnmarshalJSON unmarshals the given byte slice into %s", resourceType)
	file.Func().Params(jen.Id("r").Op("*").Id(resourceType)).Id("UnmarshalJSON").Params(jen.Id("data").Op("[]").Byte()).Error().BlockFunc(func(unmarshalBlock *jen.Group) {
		unmarshalBlock.If(jen.Id("err").Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Params(jen.Op("*").Id("Other"+resourceType)).Params(jen.Id("r"))), jen.Id("err").Op("!=").Nil()).Block(
			jen.Return(jen.Id("err")),
		)

		if hasContained {
			unmarshalBlock.Comment("If the field has contained resources, we need to unmarshal them individually and store them in .Contained")
			unmarshalBlock.If(jen.Len(jen.Id("r").Dot("RawContained")).Op(">").Lit(0)).BlockFunc(func(containedBlock *jen.Group) {
				containedBlock.Var().Id("err").Error()
				containedBlock.Id("r").Dot("Contained").Op("=").Make(jen.Op("[]").Id("IResource"), jen.Len(jen.Id("r").Dot("RawContained")))
				containedBlock.For(jen.List(jen.Id("i"), jen.Id("rawContained")).Op(":=").Range().Id("r").Dot("RawContained")).BlockFunc(func(containedLoopBlock *jen.Group) {
					containedLoopBlock.List(jen.Id("r").Dot("Contained").Index(jen.Id("i")), jen.Id("err")).Op("=").Id("UnmarshalResource").Call(jen.Id("rawContained"))
					containedLoopBlock.If(jen.Id("err").Op("!=").Nil()).Block(jen.Return(jen.Id("err")))
				})
			})
		}
		unmarshalBlock.Return(jen.Nil())
	})
}

/*
	generates the json tags for the field

`bson:"<jsonTag>,omitempty" json:"<jsonTag>,omitempty"`
*/
func generateTags(statement *jen.Statement, jsonTag string) {
	statement.Tag(map[string]string{"json": jsonTag + ",omitempty", "bson": jsonTag + ",omitempty"})
}

// generates either a slice operator or a pointer operator or both
func generateOperators(statement *jen.Statement, list, required bool) {
	if list {
		statement.Op("[]")
	}
	if !required {
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
		var err error
		r.Resource, err = UnmarshalResource(data)
		if err != nil {
			return err
		}
		r.RawResource = nil
		return nil
	}
*/
func generateElementUnmarshal(file *jen.File, resourceFields []string, typeIdentifier string) {
	file.Func().Params(jen.Id("r").Id("*" + typeIdentifier)).Id("UnmarshalJSON").Params(jen.Id("data").Op("[]").Byte()).Error().BlockFunc(func(unmarshal *jen.Group) {
		unmarshal.If(jen.Id("err").Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("data"), jen.Params(jen.Op("*").Id("Other"+typeIdentifier)).Params(jen.Id("r"))), jen.Id("err").Op("!=").Nil()).Block(
			jen.Return(jen.Id("err")),
		)
		if len(resourceFields) > 0 {
			unmarshal.Var().Id("err").Error()

			for _, resourceField := range resourceFields {
				unmarshal.List(jen.Id("r").Dot(resourceField), jen.Id("err")).Op("=").Id("UnmarshalResource").Call(jen.Id("r").Dot("Raw" + resourceField))
				unmarshal.If(jen.Id("err").Op("!=").Nil()).Block(jen.Return(jen.Id("err")))
				unmarshal.Id("r").Dot("Raw" + resourceField).Op("=").Nil()
			}
		}
		unmarshal.Return(jen.Nil())
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
func generateElementMarshal(file *jen.File, resourceFields []string, typeIdentifier string) {
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
