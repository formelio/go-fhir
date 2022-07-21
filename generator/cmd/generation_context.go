package cmd

type GenerationContext struct {
	Resources                ResourceMap
	RequiredElements         map[string]bool
	RequiredValueSetBindings map[string]bool
	GeneratedResourceTypes   []string
}

func newContext() *GenerationContext {
	return &GenerationContext{
		Resources:                newResourceMap(),
		RequiredElements:         make(map[string]bool),
		RequiredValueSetBindings: make(map[string]bool),
		GeneratedResourceTypes:   make([]string, 0),
	}
}

type ResourceMap = map[string]map[string][]byte

func newResourceMap() ResourceMap {
	return map[string]map[string][]byte{
		"StructureDefinition": make(map[string][]byte),
		"ValueSet":            make(map[string][]byte),
		"CodeSystem":          make(map[string][]byte),
	}
}
