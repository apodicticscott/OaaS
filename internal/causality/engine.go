package causality

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/apodicticscott/oaas/internal/entities"
	"gorm.io/gorm"
)

// Engine handles potentiality → actuality transitions
type Engine struct {
	db *gorm.DB
}

// NewEngine creates a new causality engine
func NewEngine(db *gorm.DB) *Engine {
	return &Engine{db: db}
}

// Condition represents a condition that must be met for actualization
type Condition struct {
	Type  string      `json:"type"`  // "attribute", "mode", "external"
	Name  string      `json:"name"`  // attribute name or condition name
	Value interface{} `json:"value"` // expected value
}

// CheckConditions verifies if all conditions for a potentiality are met
func (e *Engine) CheckConditions(potentialityID string) (bool, []string, error) {
	var potentiality entities.Potentiality
	if err := e.db.Preload("Substance").First(&potentiality, "id = ?", potentialityID).Error; err != nil {
		return false, nil, fmt.Errorf("potentiality not found: %w", err)
	}

	// Parse conditions from JSON
	var conditions []Condition
	if err := json.Unmarshal([]byte(potentiality.Conditions), &conditions); err != nil {
		return false, nil, fmt.Errorf("invalid conditions format: %w", err)
	}

	var unmetConditions []string
	allMet := true

	for _, condition := range conditions {
		met, reason := e.checkSingleCondition(potentiality.SubstanceID, condition)
		if !met {
			allMet = false
			unmetConditions = append(unmetConditions, reason)
		}
	}

	return allMet, unmetConditions, nil
}

// checkSingleCondition checks if a single condition is met
func (e *Engine) checkSingleCondition(substanceID string, condition Condition) (bool, string) {
	switch condition.Type {
	case "attribute":
		return e.checkAttributeCondition(substanceID, condition)
	case "mode":
		return e.checkModeCondition(substanceID, condition)
	case "external":
		// External conditions would be checked against external systems
		// For now, we'll assume they're always met
		return true, ""
	default:
		return false, fmt.Sprintf("unknown condition type: %s", condition.Type)
	}
}

// checkAttributeCondition checks if a substance has a specific attribute value
func (e *Engine) checkAttributeCondition(substanceID string, condition Condition) (bool, string) {
	var mode entities.Mode
	err := e.db.Joins("JOIN attributes ON modes.attribute_id = attributes.id").
		Where("modes.substance_id = ? AND attributes.name = ?", substanceID, condition.Name).
		First(&mode).Error

	if err != nil {
		return false, fmt.Sprintf("attribute '%s' not found for substance", condition.Name)
	}

	// Simple string comparison for now
	if mode.Value != fmt.Sprintf("%v", condition.Value) {
		return false, fmt.Sprintf("attribute '%s' has value '%s', expected '%v'", condition.Name, mode.Value, condition.Value)
	}

	return true, ""
}

// checkModeCondition checks if a substance has a specific mode
func (e *Engine) checkModeCondition(substanceID string, condition Condition) (bool, string) {
	var count int64
	err := e.db.Model(&entities.Mode{}).
		Joins("JOIN attributes ON modes.attribute_id = attributes.id").
		Where("modes.substance_id = ? AND attributes.name = ? AND modes.value = ?",
			substanceID, condition.Name, condition.Value).
		Count(&count).Error

	if err != nil {
		return false, fmt.Sprintf("error checking mode condition: %v", err)
	}

	if count == 0 {
		return false, fmt.Sprintf("mode condition not met: %s = %v", condition.Name, condition.Value)
	}

	return true, ""
}

// ActualizePotentiality converts a potentiality to an actuality
func (e *Engine) ActualizePotentiality(potentialityID, description string) (*entities.Actuality, error) {
	// Check if conditions are met
	canActualize, unmetConditions, err := e.CheckConditions(potentialityID)
	if err != nil {
		return nil, err
	}

	if !canActualize {
		return nil, fmt.Errorf("cannot actualize potentiality: conditions not met: %v", unmetConditions)
	}

	// Get the potentiality
	var potentiality entities.Potentiality
	if err := e.db.First(&potentiality, "id = ?", potentialityID).Error; err != nil {
		return nil, fmt.Errorf("potentiality not found: %w", err)
	}

	// Create the actuality
	actuality := entities.NewActuality(description, potentiality.SubstanceID, potentialityID)

	// Save to database
	if err := e.db.Create(actuality).Error; err != nil {
		return nil, fmt.Errorf("failed to create actuality: %w", err)
	}

	log.Printf("Potentiality '%s' actualized as '%s'", potentiality.Name, description)
	return actuality, nil
}

// GetFourCauses returns the four Aristotelian causes for a substance
func (e *Engine) GetFourCauses(substanceID string) (map[string]string, error) {
	causes := make(map[string]string)

	// Get all causal relations for this substance
	var relations []entities.CausalRelation
	if err := e.db.Where("from_entity = ? OR to_entity = ?", substanceID, substanceID).Find(&relations).Error; err != nil {
		return nil, fmt.Errorf("failed to get causal relations: %w", err)
	}

	// Group by cause type
	for _, relation := range relations {
		causes[relation.CauseType] = relation.ToEntity
	}

	return causes, nil
}

// AddCausalRelation adds a new causal relation
func (e *Engine) AddCausalRelation(fromEntity, toEntity, causeType string) (*entities.CausalRelation, error) {
	// Validate cause type
	validTypes := map[string]bool{
		"material":  true,
		"formal":    true,
		"efficient": true,
		"final":     true,
	}

	if !validTypes[causeType] {
		return nil, fmt.Errorf("invalid cause type: %s. Must be one of: material, formal, efficient, final", causeType)
	}

	relation := entities.NewCausalRelation(causeType, fromEntity, toEntity)

	if err := e.db.Create(relation).Error; err != nil {
		return nil, fmt.Errorf("failed to create causal relation: %w", err)
	}

	return relation, nil
}

// GetPotentialitiesForSubstance returns all potentialities for a substance
func (e *Engine) GetPotentialitiesForSubstance(substanceID string) ([]entities.Potentiality, error) {
	var potentialities []entities.Potentiality
	if err := e.db.Where("substance_id = ?", substanceID).Find(&potentialities).Error; err != nil {
		return nil, fmt.Errorf("failed to get potentialities: %w", err)
	}
	return potentialities, nil
}

// GetActualitiesForSubstance returns all actualities for a substance
func (e *Engine) GetActualitiesForSubstance(substanceID string) ([]entities.Actuality, error) {
	var actualities []entities.Actuality
	if err := e.db.Preload("Potentiality").Where("substance_id = ?", substanceID).Find(&actualities).Error; err != nil {
		return nil, fmt.Errorf("failed to get actualities: %w", err)
	}
	return actualities, nil
}

// CreatePotentiality creates a new potentiality for a substance
func (e *Engine) CreatePotentiality(name, description, conditions string, substanceID string) (*entities.Potentiality, error) {
	// Validate that the substance exists
	var substance entities.Substance
	if err := e.db.First(&substance, "id = ?", substanceID).Error; err != nil {
		return nil, fmt.Errorf("substance not found: %w", err)
	}

	// Validate conditions JSON format
	var testConditions []Condition
	if conditions != "" {
		if err := json.Unmarshal([]byte(conditions), &testConditions); err != nil {
			return nil, fmt.Errorf("invalid conditions JSON format: %w", err)
		}
	}

	potentiality := entities.NewPotentiality(name, description, conditions, substanceID)

	if err := e.db.Create(potentiality).Error; err != nil {
		return nil, fmt.Errorf("failed to create potentiality: %w", err)
	}

	return potentiality, nil
}

// GetSubstanceEvolution returns the evolution of a substance from potentialities to actualities
func (e *Engine) GetSubstanceEvolution(substanceID string) (*SubstanceEvolution, error) {
	potentialities, err := e.GetPotentialitiesForSubstance(substanceID)
	if err != nil {
		return nil, err
	}

	actualities, err := e.GetActualitiesForSubstance(substanceID)
	if err != nil {
		return nil, err
	}

	return &SubstanceEvolution{
		SubstanceID:    substanceID,
		Potentialities: potentialities,
		Actualities:    actualities,
		Timestamp:      time.Now(),
	}, nil
}

// SubstanceEvolution represents the evolution of a substance
type SubstanceEvolution struct {
	SubstanceID    string                  `json:"substance_id"`
	Potentialities []entities.Potentiality `json:"potentialities"`
	Actualities    []entities.Actuality    `json:"actualities"`
	Timestamp      time.Time               `json:"timestamp"`
}
