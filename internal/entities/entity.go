package entities

import (
	"time"

	"github.com/google/uuid"
)

// Substance = Neo-Aristotelian "independent entity"
type Substance struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Kind      string    `json:"kind"` // This will store the kind name, not ID for simplicity
	Essence   string    `json:"essence"`
	CreatedAt time.Time `json:"created_at"`

	// Relationships
	Attributes     []Attribute    `gorm:"many2many:substance_attributes;" json:"attributes,omitempty"`
	Modes          []Mode         `gorm:"foreignKey:SubstanceID" json:"modes,omitempty"`
	Potentialities []Potentiality `gorm:"foreignKey:SubstanceID" json:"potentialities,omitempty"`
	Actualities    []Actuality    `gorm:"foreignKey:SubstanceID" json:"actualities,omitempty"`
}

// Kind = Natural classification (e.g., oak, human)
type Kind struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex" json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`

	// Relationships
	Substances []Substance `gorm:"foreignKey:Kind" json:"substances,omitempty"`
}

// Attribute = General property (e.g., color, weight)
type Attribute struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex" json:"name"`
	Description string    `json:"description"`
	DataType    string    `json:"data_type"` // string, number, boolean, etc.
	CreatedAt   time.Time `json:"created_at"`

	// Relationships
	Substances []Substance `gorm:"many2many:substance_attributes;" json:"substances,omitempty"`
	Modes      []Mode      `gorm:"foreignKey:AttributeID" json:"modes,omitempty"`
}

// Mode = Particular way a substance instantiates an attribute (e.g., this tree's green leaf)
type Mode struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`

	// Foreign Keys
	SubstanceID string `gorm:"not null" json:"substance_id"`
	AttributeID string `gorm:"not null" json:"attribute_id"`

	// Relationships
	Substance *Substance `gorm:"foreignKey:SubstanceID" json:"substance,omitempty"`
	Attribute *Attribute `gorm:"foreignKey:AttributeID" json:"attribute,omitempty"`
}

// CausalRelation = Aristotelian causes (material, formal, efficient, final)
type CausalRelation struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	CauseType  string    `json:"cause_type"` // material, formal, efficient, final
	FromEntity string    `json:"from_entity"`
	ToEntity   string    `json:"to_entity"`
	CreatedAt  time.Time `json:"created_at"`
}

// Potentiality = What a substance can become
type Potentiality struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Conditions  string    `json:"conditions"` // JSON string of required conditions
	CreatedAt   time.Time `json:"created_at"`

	// Foreign Key
	SubstanceID string `gorm:"not null" json:"substance_id"`

	// Relationships
	Substance *Substance `gorm:"foreignKey:SubstanceID" json:"substance,omitempty"`
}

// Actuality = Realized potentiality
type Actuality struct {
	ID           string    `gorm:"primaryKey" json:"id"`
	Description  string    `json:"description"`
	ActualizedAt time.Time `json:"actualized_at"`

	// Foreign Keys
	SubstanceID    string `gorm:"not null" json:"substance_id"`
	PotentialityID string `gorm:"not null" json:"potentiality_id"`

	// Relationships
	Substance    *Substance    `gorm:"foreignKey:SubstanceID" json:"substance,omitempty"`
	Potentiality *Potentiality `gorm:"foreignKey:PotentialityID" json:"potentiality,omitempty"`
}

// NewSubstance creates a new substance with generated ID
func NewSubstance(name, kind, essence string) *Substance {
	return &Substance{
		ID:        uuid.New().String(),
		Name:      name,
		Kind:      kind,
		Essence:   essence,
		CreatedAt: time.Now(),
	}
}

// NewKind creates a new kind with generated ID
func NewKind(name, description string) *Kind {
	return &Kind{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
	}
}

// NewAttribute creates a new attribute with generated ID
func NewAttribute(name, description, dataType string) *Attribute {
	return &Attribute{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		DataType:    dataType,
		CreatedAt:   time.Now(),
	}
}

// NewMode creates a new mode with generated ID
func NewMode(value, substanceID, attributeID string) *Mode {
	return &Mode{
		ID:          uuid.New().String(),
		Value:       value,
		SubstanceID: substanceID,
		AttributeID: attributeID,
		CreatedAt:   time.Now(),
	}
}

// NewCausalRelation creates a new causal relation with generated ID
func NewCausalRelation(causeType, fromEntity, toEntity string) *CausalRelation {
	return &CausalRelation{
		ID:         uuid.New().String(),
		CauseType:  causeType,
		FromEntity: fromEntity,
		ToEntity:   toEntity,
		CreatedAt:  time.Now(),
	}
}

// NewPotentiality creates a new potentiality with generated ID
func NewPotentiality(name, description, conditions, substanceID string) *Potentiality {
	return &Potentiality{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Conditions:  conditions,
		SubstanceID: substanceID,
		CreatedAt:   time.Now(),
	}
}

// NewActuality creates a new actuality with generated ID
func NewActuality(description, substanceID, potentialityID string) *Actuality {
	return &Actuality{
		ID:             uuid.New().String(),
		Description:    description,
		ActualizedAt:   time.Now(),
		SubstanceID:    substanceID,
		PotentialityID: potentialityID,
	}
}
