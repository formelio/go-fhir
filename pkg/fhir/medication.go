package fhir

import "encoding/json"

// Medication is documented here http://hl7.org/fhir/StructureDefinition/Medication
type Medication struct {
	Id                *string                `bson:"id" json:"id"`
	Meta              *Meta                  `bson:"meta" json:"meta"`
	ImplicitRules     *string                `bson:"implicitRules" json:"implicitRules"`
	Language          *string                `bson:"language" json:"language"`
	Text              *Narrative             `bson:"text" json:"text"`
	RawContained      []json.RawMessage      `bson:"contained" json:"contained"`
	Contained         []IResource            `bson:"-" json:"-"`
	Extension         []Extension            `bson:"extension" json:"extension"`
	ModifierExtension []Extension            `bson:"modifierExtension" json:"modifierExtension"`
	Code              *CodeableConcept       `bson:"code" json:"code"`
	Status            *MedicationStatus      `bson:"status" json:"status"`
	IsBrand           *bool                  `bson:"isBrand" json:"isBrand"`
	IsOverTheCounter  *bool                  `bson:"isOverTheCounter" json:"isOverTheCounter"`
	Manufacturer      *Reference             `bson:"manufacturer" json:"manufacturer"`
	Form              *CodeableConcept       `bson:"form" json:"form"`
	Ingredient        []MedicationIngredient `bson:"ingredient" json:"ingredient"`
	Package           *MedicationPackage     `bson:"package" json:"package"`
	Image             []Attachment           `bson:"image" json:"image"`
}
type MedicationIngredient struct {
	Id                  *string          `bson:"id" json:"id"`
	Extension           []Extension      `bson:"extension" json:"extension"`
	ModifierExtension   []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	ItemCodeableConcept *CodeableConcept `bson:"itemCodeableConcept,omitempty" json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `bson:"itemReference,omitempty" json:"itemReference,omitempty"`
	IsActive            *bool            `bson:"isActive" json:"isActive"`
	Amount              *Ratio           `bson:"amount" json:"amount"`
}
type MedicationPackage struct {
	Id                *string                    `bson:"id" json:"id"`
	Extension         []Extension                `bson:"extension" json:"extension"`
	ModifierExtension []Extension                `bson:"modifierExtension" json:"modifierExtension"`
	Container         *CodeableConcept           `bson:"container" json:"container"`
	Content           []MedicationPackageContent `bson:"content" json:"content"`
	Batch             []MedicationPackageBatch   `bson:"batch" json:"batch"`
}
type MedicationPackageContent struct {
	Id                  *string          `bson:"id" json:"id"`
	Extension           []Extension      `bson:"extension" json:"extension"`
	ModifierExtension   []Extension      `bson:"modifierExtension" json:"modifierExtension"`
	ItemCodeableConcept *CodeableConcept `bson:"itemCodeableConcept,omitempty" json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `bson:"itemReference,omitempty" json:"itemReference,omitempty"`
	Amount              *Quantity        `bson:"amount" json:"amount"`
}
type MedicationPackageBatch struct {
	Id                *string     `bson:"id" json:"id"`
	Extension         []Extension `bson:"extension" json:"extension"`
	ModifierExtension []Extension `bson:"modifierExtension" json:"modifierExtension"`
	LotNumber         *string     `bson:"lotNumber" json:"lotNumber"`
	ExpirationDate    *string     `bson:"expirationDate" json:"expirationDate"`
}

// OtherMedication is a helper type to use the default implementations of Marshall and Unmarshal
type OtherMedication Medication

// MarshalJSON marshals the given Medication as JSON into a byte slice
func (r Medication) MarshalJSON() ([]byte, error) {
	// If the field has contained resources, we need to marshal them individually and store them in .RawContained
	if len(r.Contained) > 0 {
		var err error
		r.RawContained = make([]json.RawMessage, len(r.Contained))
		for i, contained := range r.Contained {
			r.RawContained[i], err = json.Marshal(contained)
			if err != nil {
				return nil, err
			}
		}
	}
	return json.Marshal(struct {
		OtherMedication
		ResourceType string `json:"resourceType"`
	}{
		OtherMedication: OtherMedication(r),
		ResourceType:    "Medication",
	})
}

// UnmarshalJSON unmarshals the given byte slice into Medication
func (r *Medication) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, (*OtherMedication)(r)); err != nil {
		return err
	}
	// If the field has contained resources, we need to unmarshal them individually and store them in .Contained
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
	return nil
}

// Returns the resourceType of the resource, makes this resource an instance of IResource
func (r Medication) GetResourceType() ResourceType {
	return ResourceTypeMedication
}
