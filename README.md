# go-fhir
Go FHIR library currently includes a STU3 FHIR library and a generator used to generate it. The original generator can be found here [golang-fhir-models](https://github.com/samply/golang-fhir-models)

# STU3 FHIR library

This library was generated using the definitions found in [HL7](https://www.hl7.org/fhir/STU3/). 

## Functionalities
1. All of the fields found in the definitions of the resources are supported. Including polymorphic value[x] fields.
2. Full JSON/BSON Marshal and Unmarshal support. Fields of type Resource are unmarshalled into an IResource interface, which each StructureDefinition implements.
3. ValueSets that have a required binding are implemented as structs for easier use. 

## Caveats
1. HTML Escape. By default in `encoding/json` "String values encode as JSON strings coerced to valid UTF-8, replacing invalid bytes with the Unicode replacement rune" ([source](https://pkg.go.dev/encoding/json#Marshal)). This escapes the Unicode values when marshaling using `json.Marshal()`. This is prevented in `<resourceType>.MarshalJSON()`, the bytes returned by the method will have the HTML unescaped. 

# FHIR Struct generator

This generator is used to generate the library. It generates all of the StructureDefinitions and it's backbone elements. It keeps track of all required Element types and generates them after the StructureDefinition generation is done. All the ValueSets that have a required binding are also generated. Finally, utility structs such as [resourceType_to_IResource.go](pkg/stu3/fhir/resourceType_to_iResource.go) are generated. 

## Caveats
1. Files used to generate the library located in `generator/fhir` are required for the generation and were generated using previous iterations of it. If another version of the FHIR library will be generated using those files, they will have to be manually tweaked to match the expected structure.

# Future work
1. Go FHIR Client, a client to retrieve and send resources from and to a FHIR server. 

## Issues
1. No XML support
2. No Testing/Validation of the library
3. No modular support of other FHIR versions
