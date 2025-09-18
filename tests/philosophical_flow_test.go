package tests

import (
	"testing"

	"github.com/apodicticscott/oaas/internal/causality"
	"github.com/apodicticscott/oaas/internal/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TestCompleteNeoAristotelianFlow demonstrates the complete philosophical flow
// from E.J. Lowe's Four-Category Ontology implementation
func TestCompleteNeoAristotelianFlow(t *testing.T) {
	// Setup test database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	
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
	
	engine := causality.NewEngine(db)
	
	// Step 1: Create a Kind (Natural Classification)
	// In Neo-Aristotelian terms: "Human" as a natural kind
	kind := entities.NewKind("Human", "Rational animal with potentiality for virtue and wisdom")
	err = db.Create(kind).Error
	require.NoError(t, err)
	
	// Step 2: Create an Attribute (General Property)
	// In Neo-Aristotelian terms: "virtue" as a general property
	attribute := entities.NewAttribute("virtue", "Moral excellence and character", "string")
	err = db.Create(attribute).Error
	require.NoError(t, err)
	
	// Step 3: Create a Substance (Independent Entity)
	// In Neo-Aristotelian terms: "Socrates" as an independent entity
	substance := entities.NewSubstance("Socrates", "Human", "Rational being with potentiality for wisdom")
	err = db.Create(substance).Error
	require.NoError(t, err)
	
	// Step 4: Create a Mode (Particular Instantiation)
	// In Neo-Aristotelian terms: "Socrates' courage" as a particular way he instantiates virtue
	mode := entities.NewMode("courageous", substance.ID, attribute.ID)
	err = db.Create(mode).Error
	require.NoError(t, err)
	
	// Step 5: Add the Four Aristotelian Causes
	// Material Cause: What Socrates is made of
	_, err = engine.AddCausalRelation(substance.ID, "flesh_bone_soul", "material")
	require.NoError(t, err)
	
	// Formal Cause: The essence of being human
	_, err = engine.AddCausalRelation(substance.ID, "human_essence", "formal")
	require.NoError(t, err)
	
	// Efficient Cause: What brought Socrates about
	_, err = engine.AddCausalRelation(substance.ID, "philosophical_education", "efficient")
	require.NoError(t, err)
	
	// Final Cause: Socrates' purpose
	_, err = engine.AddCausalRelation(substance.ID, "pursuit_of_wisdom", "final")
	require.NoError(t, err)
	
	// Step 6: Create a Potentiality
	// In Neo-Aristotelian terms: What Socrates can become
	conditions := `[{"type":"mode","name":"virtue","value":"courageous"}]`
	potentiality, err := engine.CreatePotentiality(
		"Achieve Wisdom",
		"Socrates can achieve philosophical wisdom through inquiry",
		conditions,
		substance.ID,
	)
	require.NoError(t, err)
	
	// Step 7: Check if conditions are met for actualization
	canActualize, unmetConditions, err := engine.CheckConditions(potentiality.ID)
	require.NoError(t, err)
	assert.True(t, canActualize, "Conditions should be met for actualization")
	assert.Empty(t, unmetConditions, "No conditions should be unmet")
	
	// Step 8: Actualize the Potentiality
	// In Neo-Aristotelian terms: Convert potentiality to actuality
	actuality, err := engine.ActualizePotentiality(potentiality.ID, "Socrates achieved wisdom through philosophical inquiry")
	require.NoError(t, err)
	assert.Equal(t, "Socrates achieved wisdom through philosophical inquiry", actuality.Description)
	assert.Equal(t, substance.ID, actuality.SubstanceID)
	assert.Equal(t, potentiality.ID, actuality.PotentialityID)
	
	// Step 9: Verify the Four Causes
	causes, err := engine.GetFourCauses(substance.ID)
	require.NoError(t, err)
	assert.Equal(t, "flesh_bone_soul", causes["material"])
	assert.Equal(t, "human_essence", causes["formal"])
	assert.Equal(t, "philosophical_education", causes["efficient"])
	assert.Equal(t, "pursuit_of_wisdom", causes["final"])
	
	// Step 10: Get the Complete Evolution
	evolution, err := engine.GetSubstanceEvolution(substance.ID)
	require.NoError(t, err)
	assert.Equal(t, substance.ID, evolution.SubstanceID)
	assert.Len(t, evolution.Potentialities, 1)
	assert.Len(t, evolution.Actualities, 1)
	assert.Equal(t, "Achieve Wisdom", evolution.Potentialities[0].Name)
	assert.Equal(t, "Socrates achieved wisdom through philosophical inquiry", evolution.Actualities[0].Description)
	
	// Philosophical Validation
	t.Log("✅ Complete Neo-Aristotelian Flow Validated:")
	t.Log("   - Kind (Human): Natural classification")
	t.Log("   - Attribute (virtue): General property")
	t.Log("   - Substance (Socrates): Independent entity")
	t.Log("   - Mode (courageous): Particular instantiation")
	t.Log("   - Four Causes: Material, Formal, Efficient, Final")
	t.Log("   - Potentiality → Actuality: Philosophical transformation")
	t.Log("   - Evolution: Complete ontological development")
}

// TestPhilosophicalCorrectness validates that the implementation
// correctly follows E.J. Lowe's Four-Category Ontology principles
func TestPhilosophicalCorrectness(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	
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
	
	// Test 1: Substances are independent entities
	substance := entities.NewSubstance("Tree-001", "Oak", "Independent living organism")
	err = db.Create(substance).Error
	require.NoError(t, err)
	
	var retrievedSubstance entities.Substance
	err = db.First(&retrievedSubstance, "id = ?", substance.ID).Error
	require.NoError(t, err)
	assert.Equal(t, substance.Name, retrievedSubstance.Name)
	assert.Equal(t, substance.Essence, retrievedSubstance.Essence)
	
	// Test 2: Modes depend on substances and attributes
	attribute := entities.NewAttribute("color", "Visual property", "string")
	err = db.Create(attribute).Error
	require.NoError(t, err)
	
	mode := entities.NewMode("green", substance.ID, attribute.ID)
	err = db.Create(mode).Error
	require.NoError(t, err)
	
	// Mode should reference both substance and attribute
	assert.Equal(t, substance.ID, mode.SubstanceID)
	assert.Equal(t, attribute.ID, mode.AttributeID)
	
	// Test 3: Causal relations follow Aristotelian principles
	engine := causality.NewEngine(db)
	
	validCauseTypes := []string{"material", "formal", "efficient", "final"}
	for _, causeType := range validCauseTypes {
		_, err := engine.AddCausalRelation(substance.ID, "cause-entity", causeType)
		assert.NoError(t, err, "Should accept valid cause type: %s", causeType)
	}
	
	// Test 4: Potentialities can only be actualized when conditions are met
	conditions := `[{"type":"mode","name":"color","value":"green"}]`
	potentiality, err := engine.CreatePotentiality("Grow Leaves", "Tree can grow leaves", conditions, substance.ID)
	require.NoError(t, err)
	
	canActualize, _, err := engine.CheckConditions(potentiality.ID)
	require.NoError(t, err)
	assert.True(t, canActualize, "Conditions should be met for actualization")
	
	// Test 5: Actualities represent realized potentialities
	actuality, err := engine.ActualizePotentiality(potentiality.ID, "Tree successfully grew green leaves")
	require.NoError(t, err)
	assert.Equal(t, potentiality.ID, actuality.PotentialityID)
	assert.Equal(t, substance.ID, actuality.SubstanceID)
	
	t.Log("✅ Philosophical Correctness Validated:")
	t.Log("   - Substances are independent entities")
	t.Log("   - Modes depend on substances and attributes")
	t.Log("   - Causal relations follow Aristotelian principles")
	t.Log("   - Potentialities require proper conditions for actualization")
	t.Log("   - Actualities represent realized potentialities")
}

// TestOntologicalRelationships validates the relationships between
// the four categories of E.J. Lowe's ontology
func TestOntologicalRelationships(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	
	err = db.AutoMigrate(
		&entities.Kind{},
		&entities.Attribute{},
		&entities.Substance{},
		&entities.Mode{},
	)
	require.NoError(t, err)
	
	// Create the four categories
	kind := entities.NewKind("Oak", "Type of tree")
	attribute := entities.NewAttribute("height", "Vertical measurement", "number")
	substance := entities.NewSubstance("Tree-001", "Oak", "Living organism")
	mode := entities.NewMode("10 meters", substance.ID, attribute.ID)
	
	// Save all entities
	err = db.Create(kind).Error
	require.NoError(t, err)
	err = db.Create(attribute).Error
	require.NoError(t, err)
	err = db.Create(substance).Error
	require.NoError(t, err)
	err = db.Create(mode).Error
	require.NoError(t, err)
	
	// Test relationships
	// 1. Substance belongs to a Kind (by name reference)
	assert.Equal(t, "Oak", substance.Kind)
	
	// 2. Mode belongs to a Substance and an Attribute
	assert.Equal(t, substance.ID, mode.SubstanceID)
	assert.Equal(t, attribute.ID, mode.AttributeID)
	
	// 3. Mode represents a particular instantiation of an attribute by a substance
	assert.Equal(t, "10 meters", mode.Value)
	
	// 4. The relationship between substance and attribute through mode
	var retrievedMode entities.Mode
	err = db.Preload("Substance").Preload("Attribute").First(&retrievedMode, "id = ?", mode.ID).Error
	require.NoError(t, err)
	
	assert.Equal(t, substance.Name, retrievedMode.Substance.Name)
	assert.Equal(t, attribute.Name, retrievedMode.Attribute.Name)
	
	t.Log("✅ Ontological Relationships Validated:")
	t.Log("   - Substance belongs to Kind")
	t.Log("   - Mode belongs to Substance and Attribute")
	t.Log("   - Mode represents particular instantiation")
	t.Log("   - Relationships properly established through foreign keys")
}
