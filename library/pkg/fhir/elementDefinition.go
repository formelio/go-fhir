package fhir

// ElementDefinition is documented here http://hl7.org/fhir/StructureDefinition/ElementDefinition
type ElementDefinition struct {
	Id                 *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	Path               string                        `bson:"path" json:"path"`
	Representation     []string                      `bson:"representation,omitempty" json:"representation,omitempty"`
	SliceName          *string                       `bson:"sliceName,omitempty" json:"sliceName,omitempty"`
	Label              *string                       `bson:"label,omitempty" json:"label,omitempty"`
	Code               []Coding                      `bson:"code,omitempty" json:"code,omitempty"`
	Slicing            *ElementDefinitionSlicing     `bson:"slicing,omitempty" json:"slicing,omitempty"`
	Short              *string                       `bson:"short,omitempty" json:"short,omitempty"`
	Definition         *string                       `bson:"definition,omitempty" json:"definition,omitempty"`
	Comment            *string                       `bson:"comment,omitempty" json:"comment,omitempty"`
	Requirements       *string                       `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Alias              []string                      `bson:"alias,omitempty" json:"alias,omitempty"`
	Min                *int                          `bson:"min,omitempty" json:"min,omitempty"`
	Max                *string                       `bson:"max,omitempty" json:"max,omitempty"`
	Base               *ElementDefinitionBase        `bson:"base,omitempty" json:"base,omitempty"`
	ContentReference   *string                       `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
	Type               []ElementDefinitionType       `bson:"type,omitempty" json:"type,omitempty"`
	MeaningWhenMissing *string                       `bson:"meaningWhenMissing,omitempty" json:"meaningWhenMissing,omitempty"`
	OrderMeaning       *string                       `bson:"orderMeaning,omitempty" json:"orderMeaning,omitempty"`
	Example            []ElementDefinitionExample    `bson:"example,omitempty" json:"example,omitempty"`
	MaxLength          *int                          `bson:"maxLength,omitempty" json:"maxLength,omitempty"`
	Condition          []string                      `bson:"condition,omitempty" json:"condition,omitempty"`
	Constraint         []ElementDefinitionConstraint `bson:"constraint,omitempty" json:"constraint,omitempty"`
	MustSupport        *bool                         `bson:"mustSupport,omitempty" json:"mustSupport,omitempty"`
	IsModifier         *bool                         `bson:"isModifier,omitempty" json:"isModifier,omitempty"`
	IsSummary          *bool                         `bson:"isSummary,omitempty" json:"isSummary,omitempty"`
	Binding            *ElementDefinitionBinding     `bson:"binding,omitempty" json:"binding,omitempty"`
	Mapping            []ElementDefinitionMapping    `bson:"mapping,omitempty" json:"mapping,omitempty"`
}
type ElementDefinitionSlicing struct {
	Id            *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension     []Extension                             `bson:"extension,omitempty" json:"extension,omitempty"`
	Discriminator []ElementDefinitionSlicingDiscriminator `bson:"discriminator,omitempty" json:"discriminator,omitempty"`
	Description   *string                                 `bson:"description,omitempty" json:"description,omitempty"`
	Ordered       *bool                                   `bson:"ordered,omitempty" json:"ordered,omitempty"`
	Rules         string                                  `bson:"rules" json:"rules"`
}
type ElementDefinitionSlicingDiscriminator struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Type      string      `bson:"type" json:"type"`
	Path      string      `bson:"path" json:"path"`
}
type ElementDefinitionBase struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Path      string      `bson:"path" json:"path"`
	Min       int         `bson:"min" json:"min"`
	Max       string      `bson:"max" json:"max"`
}
type ElementDefinitionType struct {
	Id            *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension     []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Code          string      `bson:"code" json:"code"`
	Profile       *string     `bson:"profile,omitempty" json:"profile,omitempty"`
	TargetProfile *string     `bson:"targetProfile,omitempty" json:"targetProfile,omitempty"`
	Aggregation   []string    `bson:"aggregation,omitempty" json:"aggregation,omitempty"`
	Versioning    *string     `bson:"versioning,omitempty" json:"versioning,omitempty"`
}
type ElementDefinitionExample struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Label     string      `bson:"label" json:"label"`
}
type ElementDefinitionConstraint struct {
	Id           *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension    []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Key          string      `bson:"key" json:"key"`
	Requirements *string     `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Severity     string      `bson:"severity" json:"severity"`
	Human        string      `bson:"human" json:"human"`
	Expression   string      `bson:"expression" json:"expression"`
	Xpath        *string     `bson:"xpath,omitempty" json:"xpath,omitempty"`
	Source       *string     `bson:"source,omitempty" json:"source,omitempty"`
}
type ElementDefinitionBinding struct {
	Id          *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension   []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Strength    string      `bson:"strength" json:"strength"`
	Description *string     `bson:"description,omitempty" json:"description,omitempty"`
}
type ElementDefinitionMapping struct {
	Id        *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Identity  string      `bson:"identity" json:"identity"`
	Language  *string     `bson:"language,omitempty" json:"language,omitempty"`
	Map       string      `bson:"map" json:"map"`
	Comment   *string     `bson:"comment,omitempty" json:"comment,omitempty"`
}
