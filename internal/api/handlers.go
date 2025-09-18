package api

import (
	"net/http"

	"github.com/apodicticscott/oaas/internal/causality"
	"github.com/apodicticscott/oaas/internal/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handler contains dependencies for API handlers
type Handler struct {
	DB              *gorm.DB
	CausalityEngine *causality.Engine
}

// NewHandler creates a new API handler
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		DB:              db,
		CausalityEngine: causality.NewEngine(db),
	}
}

// HealthCheck returns the health status of the API
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "OaaS API is running",
	})
}

// Substances handlers

// GetSubstances returns all substances
func (h *Handler) GetSubstances(c *gin.Context) {
	var substances []entities.Substance
	if err := h.DB.Preload("Attributes").Preload("Modes").Preload("Potentialities").Preload("Actualities").Find(&substances).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"substances": substances})
}

// GetSubstance returns a specific substance by ID
func (h *Handler) GetSubstance(c *gin.Context) {
	id := c.Param("id")
	var substance entities.Substance
	if err := h.DB.Preload("Attributes").Preload("Modes").Preload("Potentialities").Preload("Actualities").First(&substance, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "substance not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, substance)
}

// CreateSubstance creates a new substance
func (h *Handler) CreateSubstance(c *gin.Context) {
	var req struct {
		Name    string `json:"name" binding:"required"`
		Kind    string `json:"kind" binding:"required"`
		Essence string `json:"essence" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	substance := entities.NewSubstance(req.Name, req.Kind, req.Essence)
	if err := h.DB.Create(substance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, substance)
}

// UpdateSubstance updates an existing substance
func (h *Handler) UpdateSubstance(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Name    *string `json:"name"`
		Kind    *string `json:"kind"`
		Essence *string `json:"essence"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var substance entities.Substance
	if err := h.DB.First(&substance, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "substance not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Kind != nil {
		updates["kind"] = *req.Kind
	}
	if req.Essence != nil {
		updates["essence"] = *req.Essence
	}

	if err := h.DB.Model(&substance).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, substance)
}

// DeleteSubstance deletes a substance
func (h *Handler) DeleteSubstance(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.Delete(&entities.Substance{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "substance deleted"})
}

// Kinds handlers

// GetKinds returns all kinds
func (h *Handler) GetKinds(c *gin.Context) {
	var kinds []entities.Kind
	if err := h.DB.Find(&kinds).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"kinds": kinds})
}

// CreateKind creates a new kind
func (h *Handler) CreateKind(c *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kind := entities.NewKind(req.Name, req.Description)
	if err := h.DB.Create(kind).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, kind)
}

// Attributes handlers

// GetAttributes returns all attributes
func (h *Handler) GetAttributes(c *gin.Context) {
	var attributes []entities.Attribute
	if err := h.DB.Find(&attributes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"attributes": attributes})
}

// CreateAttribute creates a new attribute
func (h *Handler) CreateAttribute(c *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		DataType    string `json:"data_type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attribute := entities.NewAttribute(req.Name, req.Description, req.DataType)
	if err := h.DB.Create(attribute).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, attribute)
}

// Modes handlers

// GetModes returns all modes
func (h *Handler) GetModes(c *gin.Context) {
	var modes []entities.Mode
	if err := h.DB.Preload("Substance").Preload("Attribute").Find(&modes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"modes": modes})
}

// CreateMode creates a new mode
func (h *Handler) CreateMode(c *gin.Context) {
	var req struct {
		Value       string `json:"value" binding:"required"`
		SubstanceID string `json:"substance_id" binding:"required"`
		AttributeID string `json:"attribute_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mode := entities.NewMode(req.Value, req.SubstanceID, req.AttributeID)
	if err := h.DB.Create(mode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mode)
}

// Causality handlers

// GetCauses returns the four causes for a substance
func (h *Handler) GetCauses(c *gin.Context) {
	id := c.Param("id")
	causes, err := h.CausalityEngine.GetFourCauses(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, causes)
}

// AddCause adds a causal relation
func (h *Handler) AddCause(c *gin.Context) {
	var req struct {
		FromEntity string `json:"from_entity" binding:"required"`
		ToEntity   string `json:"to_entity" binding:"required"`
		CauseType  string `json:"cause_type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	relation, err := h.CausalityEngine.AddCausalRelation(req.FromEntity, req.ToEntity, req.CauseType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, relation)
}

// Potentialities handlers

// GetPotentialities returns all potentialities
func (h *Handler) GetPotentialities(c *gin.Context) {
	var potentialities []entities.Potentiality
	if err := h.DB.Preload("Substance").Find(&potentialities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"potentialities": potentialities})
}

// CreatePotentiality creates a new potentiality
func (h *Handler) CreatePotentiality(c *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Conditions  string `json:"conditions"`
		SubstanceID string `json:"substance_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	potentiality, err := h.CausalityEngine.CreatePotentiality(req.Name, req.Description, req.Conditions, req.SubstanceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, potentiality)
}

// ActualizePotentiality converts a potentiality to an actuality
func (h *Handler) ActualizePotentiality(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Description string `json:"description" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	actuality, err := h.CausalityEngine.ActualizePotentiality(id, req.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, actuality)
}

// GetSubstanceEvolution returns the evolution of a substance
func (h *Handler) GetSubstanceEvolution(c *gin.Context) {
	id := c.Param("id")
	evolution, err := h.CausalityEngine.GetSubstanceEvolution(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, evolution)
}

// CheckConditions checks if conditions for a potentiality are met
func (h *Handler) CheckConditions(c *gin.Context) {
	id := c.Param("id")
	canActualize, unmetConditions, err := h.CausalityEngine.CheckConditions(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"can_actualize":    canActualize,
		"unmet_conditions": unmetConditions,
	})
}
