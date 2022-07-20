package fhir

import "encoding/json"

// Medication is documented here http://hl7.org/fhir/StructureDefinition/Medication
type Medication struct {
	Id                *string                `bson:"id" json:"id"`
	Meta              *Meta                  `bson:"meta" json:"meta"`
	ImplicitRules     *string                `bson:"implicitRules" json:"implicitRules"`
	Language          *string                `bson:"language" json:"language"`
	Text              *Narrative             `bson:"text" json:"text"`
	Contained         []json.RawMessage      `bson:"contained" json:"contained"`
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

// UnmarshalMedication unmarshalls a Medication.
func UnmarshalMedication(b []byte) (Medication, error) {
	var medication Medication
	if err := json.Unmarshal(b, &medication); err != nil {
		return medication, err
	}
	return medication, nil
}
