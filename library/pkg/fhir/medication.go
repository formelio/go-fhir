package fhir

import "encoding/json"

// Medication is documented here http://hl7.org/fhir/StructureDefinition/Medication
type Medication struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              *CodeableConcept       `bson:"code,omitempty" json:"code,omitempty"`
	Status            *string                `bson:"status,omitempty" json:"status,omitempty"`
	IsBrand           *bool                  `bson:"isBrand,omitempty" json:"isBrand,omitempty"`
	IsOverTheCounter  *bool                  `bson:"isOverTheCounter,omitempty" json:"isOverTheCounter,omitempty"`
	Manufacturer      *Reference             `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	Form              *CodeableConcept       `bson:"form,omitempty" json:"form,omitempty"`
	Ingredient        []MedicationIngredient `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
	Package           *MedicationPackage     `bson:"package,omitempty" json:"package,omitempty"`
	Image             []Attachment           `bson:"image,omitempty" json:"image,omitempty"`
}
type MedicationIngredient struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	IsActive          *bool       `bson:"isActive,omitempty" json:"isActive,omitempty"`
	Amount            *Ratio      `bson:"amount,omitempty" json:"amount,omitempty"`
}
type MedicationPackage struct {
	Id                *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Container         *CodeableConcept           `bson:"container,omitempty" json:"container,omitempty"`
	Content           []MedicationPackageContent `bson:"content,omitempty" json:"content,omitempty"`
	Batch             []MedicationPackageBatch   `bson:"batch,omitempty" json:"batch,omitempty"`
}
type MedicationPackageContent struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Amount            *Quantity   `bson:"amount,omitempty" json:"amount,omitempty"`
}
type MedicationPackageBatch struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LotNumber         *string     `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	ExpirationDate    *string     `bson:"expirationDate,omitempty" json:"expirationDate,omitempty"`
}
type OtherMedication Medication

// MarshalJSON marshals the given Medication as JSON into a byte slice
func (r Medication) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedication
		ResourceType string `json:"resourceType"`
	}{
		OtherMedication: OtherMedication(r),
		ResourceType:    "Medication",
	})
}

// UnmarshalMedication unmarshals a Medication.
func UnmarshalMedication(b []byte) (Medication, error) {
	var medication Medication
	if err := json.Unmarshal(b, &medication); err != nil {
		return medication, err
	}
	return medication, nil
}
