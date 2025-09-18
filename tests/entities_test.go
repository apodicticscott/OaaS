package tests

import (
	"testing"
	"time"

	"github.com/apodicticscott/oaas/internal/entities"
	"github.com/stretchr/testify/assert"
)

func TestNewSubstance(t *testing.T) {
	substance := entities.NewSubstance("Tree-001", "Oak", "Living organism with potentiality for growth")
	
	assert.NotEmpty(t, substance.ID)
	assert.Equal(t, "Tree-001", substance.Name)
	assert.Equal(t, "Oak", substance.Kind)
	assert.Equal(t, "Living organism with potentiality for growth", substance.Essence)
	assert.WithinDuration(t, time.Now(), substance.CreatedAt, time.Second)
}

func TestNewKind(t *testing.T) {
	kind := entities.NewKind("Human", "Rational animal with potentiality for virtue")
	
	assert.NotEmpty(t, kind.ID)
	assert.Equal(t, "Human", kind.Name)
	assert.Equal(t, "Rational animal with potentiality for virtue", kind.Description)
	assert.WithinDuration(t, time.Now(), kind.CreatedAt, time.Second)
}

func TestNewAttribute(t *testing.T) {
	attribute := entities.NewAttribute("color", "Visual property", "string")
	
	assert.NotEmpty(t, attribute.ID)
	assert.Equal(t, "color", attribute.Name)
	assert.Equal(t, "Visual property", attribute.Description)
	assert.Equal(t, "string", attribute.DataType)
	assert.WithinDuration(t, time.Now(), attribute.CreatedAt, time.Second)
}

func TestNewMode(t *testing.T) {
	substanceID := "substance-123"
	attributeID := "attr-456"
	mode := entities.NewMode("green", substanceID, attributeID)
	
	assert.NotEmpty(t, mode.ID)
	assert.Equal(t, "green", mode.Value)
	assert.Equal(t, substanceID, mode.SubstanceID)
	assert.Equal(t, attributeID, mode.AttributeID)
	assert.WithinDuration(t, time.Now(), mode.CreatedAt, time.Second)
}

func TestNewCausalRelation(t *testing.T) {
	relation := entities.NewCausalRelation("material", "substance-1", "wood_water")
	
	assert.NotEmpty(t, relation.ID)
	assert.Equal(t, "material", relation.CauseType)
	assert.Equal(t, "substance-1", relation.FromEntity)
	assert.Equal(t, "wood_water", relation.ToEntity)
	assert.WithinDuration(t, time.Now(), relation.CreatedAt, time.Second)
}

func TestNewPotentiality(t *testing.T) {
	substanceID := "substance-123"
	conditions := `[{"type":"attribute","name":"season","value":"spring"}]`
	potentiality := entities.NewPotentiality("Grow Leaves", "Tree can grow leaves", conditions, substanceID)
	
	assert.NotEmpty(t, potentiality.ID)
	assert.Equal(t, "Grow Leaves", potentiality.Name)
	assert.Equal(t, "Tree can grow leaves", potentiality.Description)
	assert.Equal(t, conditions, potentiality.Conditions)
	assert.Equal(t, substanceID, potentiality.SubstanceID)
	assert.WithinDuration(t, time.Now(), potentiality.CreatedAt, time.Second)
}

func TestNewActuality(t *testing.T) {
	substanceID := "substance-123"
	potentialityID := "potentiality-456"
	actuality := entities.NewActuality("Tree grew green leaves", substanceID, potentialityID)
	
	assert.NotEmpty(t, actuality.ID)
	assert.Equal(t, "Tree grew green leaves", actuality.Description)
	assert.Equal(t, substanceID, actuality.SubstanceID)
	assert.Equal(t, potentialityID, actuality.PotentialityID)
	assert.WithinDuration(t, time.Now(), actuality.ActualizedAt, time.Second)
}

func TestEntityIDUniqueness(t *testing.T) {
	// Test that multiple entities have unique IDs
	substance1 := entities.NewSubstance("Tree-1", "Oak", "Essence 1")
	substance2 := entities.NewSubstance("Tree-2", "Pine", "Essence 2")
	
	assert.NotEqual(t, substance1.ID, substance2.ID)
	
	kind1 := entities.NewKind("Human", "Description 1")
	kind2 := entities.NewKind("Animal", "Description 2")
	
	assert.NotEqual(t, kind1.ID, kind2.ID)
}

func TestCausalRelationTypes(t *testing.T) {
	validTypes := []string{"material", "formal", "efficient", "final"}
	
	for _, causeType := range validTypes {
		relation := entities.NewCausalRelation(causeType, "from", "to")
		assert.Equal(t, causeType, relation.CauseType)
	}
}
