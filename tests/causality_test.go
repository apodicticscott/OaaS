package tests

import (
	"encoding/json"
	"testing"

	"github.com/apodicticscott/oaas/internal/causality"
	"github.com/apodicticscott/oaas/internal/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	
	// Auto-migrate all entities
	err = db.AutoMigrate(
		&entities.Kind{},
		&entities.Attribute{},
		&entities.Substance{},
		&entities.Mode{},
		&entities.CausalRelation{},
		&entities.Potentiality{},
		&entities.Actuality{},
	)
	require.NoError(t, err)
	
	return db
}

func TestCausalityEngine_AddCausalRelation(t *testing.T) {
	db := setupTestDB(t)
	engine := causality.NewEngine(db)
	
	// Test valid cause types
	validTypes := []string{"material", "formal", "efficient", "final"}
	
	for _, causeType := range validTypes {
		relation, err := engine.AddCausalRelation("substance-1", "cause-entity", causeType)
		assert.NoError(t, err)
		assert.Equal(t, causeType, relation.CauseType)
		assert.Equal(t, "substance-1", relation.FromEntity)
		assert.Equal(t, "cause-entity", relation.ToEntity)
	}
	
	// Test invalid cause type
	_, err := engine.AddCausalRelation("substance-1", "cause-entity", "invalid")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid cause type")
}

func TestCausalityEngine_CreatePotentiality(t *testing.T) {
	db := setupTestDB(t)
	engine := causality.NewEngine(db)
	
	// Create a test substance
	substance := entities.NewSubstance("Tree-001", "Oak", "Living organism")
	err := db.Create(substance).Error
	require.NoError(t, err)
	
	// Test creating a potentiality
	conditions := `[{"type":"attribute","name":"season","value":"spring"}]`
	potentiality, err := engine.CreatePotentiality("Grow Leaves", "Tree can grow leaves", conditions, substance.ID)
	
	assert.NoError(t, err)
	assert.Equal(t, "Grow Leaves", potentiality.Name)
	assert.Equal(t, "Tree can grow leaves", potentiality.Description)
	assert.Equal(t, conditions, potentiality.Conditions)
	assert.Equal(t, substance.ID, potentiality.SubstanceID)
}

func TestCausalityEngine_CreatePotentiality_InvalidSubstance(t *testing.T) {
	db := setupTestDB(t)
	engine := causality.NewEngine(db)
	
	// Test creating potentiality for non-existent substance
	_, err := engine.CreatePotentiality("Grow Leaves", "Description", "", "non-existent-id")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "substance not found")
}

func TestCausalityEngine_CheckConditions(t *testing.T) {
	db := setupTestDB(t)
	engine := causality.NewEngine(db)
	
	// Create test data
	substance := entities.NewSubstance("Tree-001", "Oak", "Living organism")
	err := db.Create(substance).Error
	require.NoError(t, err)
	
	attribute := entities.NewAttribute("color", "Visual property", "string")
	err = db.Create(attribute).Error
	require.NoError(t, err)
	
	mode := entities.NewMode("green", substance.ID, attribute.ID)
	err = db.Create(mode).Error
	require.NoError(t, err)
	
	// Create potentiality with conditions
	conditions := `[{"type":"mode","name":"color","value":"green"}]`
	potentiality := entities.NewPotentiality("Grow Leaves", "Tree can grow leaves", conditions, substance.ID)
	err = db.Create(potentiality).Error
	require.NoError(t, err)
	
	// Test condition checking
	canActualize, unmetConditions, err := engine.CheckConditions(potentiality.ID)
	
	assert.NoError(t, err)
	assert.True(t, canActualize)
	assert.Empty(t, unmetConditions)
}

func TestCausalityEngine_CheckConditions_Unmet(t *testing.T) {
	db := setupTestDB(t)
	engine := causality.NewEngine(db)
	
	// Create test data
	substance := entities.NewSubstance("Tree-001", "Oak", "Living organism")
	err := db.Create(substance).Error
	require.NoError(t, err)
	
	attribute := entities.NewAttribute("color", "Visual property", "string")
	err = db.Create(attribute).Error
	require.NoError(t, err)
	
	// Create mode with different value
	mode := entities.NewMode("red", substance.ID, attribute.ID)
	err = db.Create(mode).Error
	require.NoError(t, err)
	
	// Create potentiality with conditions that won't be met
	conditions := `[{"type":"mode","name":"color","value":"green"}]`
	potentiality := entities.NewPotentiality("Grow Leaves", "Tree can grow leaves", conditions, substance.ID)
	err = db.Create(potentiality).Error
	require.NoError(t, err)
	
	// Test condition checking
	canActualize, unmetConditions, err := engine.CheckConditions(potentiality.ID)
	
	assert.NoError(t, err)
	assert.False(t, canActualize)
	assert.NotEmpty(t, unmetConditions)
}

func TestCausalityEngine_ActualizePotentiality(t *testing.T) {
	db := setupTestDB(t)
	engine := causality.NewEngine(db)
	
	// Create test data
	substance := entities.NewSubstance("Tree-001", "Oak", "Living organism")
	err := db.Create(substance).Error
	require.NoError(t, err)
	
	attribute := entities.NewAttribute("color", "Visual property", "string")
	err = db.Create(attribute).Error
	require.NoError(t, err)
	
	mode := entities.NewMode("green", substance.ID, attribute.ID)
	err = db.Create(mode).Error
	require.NoError(t, err)
	
	// Create potentiality with conditions that will be met
	conditions := `[{"type":"mode","name":"color","value":"green"}]`
	potentiality := entities.NewPotentiality("Grow Leaves", "Tree can grow leaves", conditions, substance.ID)
	err = db.Create(potentiality).Error
	require.NoError(t, err)
	
	// Test actualization
	actuality, err := engine.ActualizePotentiality(potentiality.ID, "Tree successfully grew green leaves")
	
	assert.NoError(t, err)
	assert.Equal(t, "Tree successfully grew green leaves", actuality.Description)
	assert.Equal(t, substance.ID, actuality.SubstanceID)
	assert.Equal(t, potentiality.ID, actuality.PotentialityID)
}

func TestCausalityEngine_ActualizePotentiality_UnmetConditions(t *testing.T) {
	db := setupTestDB(t)
	engine := causality.NewEngine(db)
	
	// Create test data
	substance := entities.NewSubstance("Tree-001", "Oak", "Living organism")
	err := db.Create(substance).Error
	require.NoError(t, err)
	
	// Create potentiality with conditions that won't be met
	conditions := `[{"type":"mode","name":"color","value":"green"}]`
	potentiality := entities.NewPotentiality("Grow Leaves", "Tree can grow leaves", conditions, substance.ID)
	err = db.Create(potentiality).Error
	require.NoError(t, err)
	
	// Test actualization with unmet conditions
	_, err = engine.ActualizePotentiality(potentiality.ID, "Tree grew leaves")
	
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot actualize potentiality")
}

func TestCausalityEngine_GetFourCauses(t *testing.T) {
	db := setupTestDB(t)
	engine := causality.NewEngine(db)
	
	// Create test substance
	substance := entities.NewSubstance("Tree-001", "Oak", "Living organism")
	err := db.Create(substance).Error
	require.NoError(t, err)
	
	// Add causal relations
	_, err = engine.AddCausalRelation(substance.ID, "wood_water", "material")
	require.NoError(t, err)
	
	_, err = engine.AddCausalRelation(substance.ID, "oak_essence", "formal")
	require.NoError(t, err)
	
	_, err = engine.AddCausalRelation(substance.ID, "sunlight_soil", "efficient")
	require.NoError(t, err)
	
	_, err = engine.AddCausalRelation(substance.ID, "provide_shade", "final")
	require.NoError(t, err)
	
	// Test getting four causes
	causes, err := engine.GetFourCauses(substance.ID)
	
	assert.NoError(t, err)
	assert.Equal(t, "wood_water", causes["material"])
	assert.Equal(t, "oak_essence", causes["formal"])
	assert.Equal(t, "sunlight_soil", causes["efficient"])
	assert.Equal(t, "provide_shade", causes["final"])
}

func TestCausalityEngine_GetSubstanceEvolution(t *testing.T) {
	db := setupTestDB(t)
	engine := causality.NewEngine(db)
	
	// Create test data
	substance := entities.NewSubstance("Tree-001", "Oak", "Living organism")
	err := db.Create(substance).Error
	require.NoError(t, err)
	
	// Create potentialities
	potentiality1 := entities.NewPotentiality("Grow Leaves", "Tree can grow leaves", "", substance.ID)
	err = db.Create(potentiality1).Error
	require.NoError(t, err)
	
	potentiality2 := entities.NewPotentiality("Produce Acorns", "Tree can produce acorns", "", substance.ID)
	err = db.Create(potentiality2).Error
	require.NoError(t, err)
	
	// Create actuality
	actuality := entities.NewActuality("Tree grew green leaves", substance.ID, potentiality1.ID)
	err = db.Create(actuality).Error
	require.NoError(t, err)
	
	// Test getting evolution
	evolution, err := engine.GetSubstanceEvolution(substance.ID)
	
	assert.NoError(t, err)
	assert.Equal(t, substance.ID, evolution.SubstanceID)
	assert.Len(t, evolution.Potentialities, 2)
	assert.Len(t, evolution.Actualities, 1)
	assert.Equal(t, "Tree grew green leaves", evolution.Actualities[0].Description)
}

func TestConditionParsing(t *testing.T) {
	// Test valid JSON conditions
	validConditions := `[{"type":"attribute","name":"season","value":"spring"},{"type":"mode","name":"health","value":"good"}]`
	
	var conditions []causality.Condition
	err := json.Unmarshal([]byte(validConditions), &conditions)
	assert.NoError(t, err)
	assert.Len(t, conditions, 2)
	assert.Equal(t, "attribute", conditions[0].Type)
	assert.Equal(t, "season", conditions[0].Name)
	assert.Equal(t, "spring", conditions[0].Value)
}

func TestConditionParsing_InvalidJSON(t *testing.T) {
	// Test invalid JSON
	invalidConditions := `[{"type":"attribute","name":"season","value":"spring"` // Missing closing bracket
	
	var conditions []causality.Condition
	err := json.Unmarshal([]byte(invalidConditions), &conditions)
	assert.Error(t, err)
}
