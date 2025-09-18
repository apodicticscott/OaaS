package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apodicticscott/oaas/internal/api"
	"github.com/apodicticscott/oaas/internal/entities"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestAPI(t *testing.T) (*gin.Engine, *gorm.DB) {
	// Setup in-memory database
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

	// Setup API handler
	handler := api.NewHandler(db)

	// Setup Gin router
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Health check
	router.GET("/health", handler.HealthCheck)

	// API routes
	api := router.Group("/api/v1")
	{
		// Substances
		api.GET("/substances", handler.GetSubstances)
		api.GET("/substances/:id", handler.GetSubstance)
		api.POST("/substances", handler.CreateSubstance)
		api.PUT("/substances/:id", handler.UpdateSubstance)
		api.DELETE("/substances/:id", handler.DeleteSubstance)

		// Kinds
		api.GET("/kinds", handler.GetKinds)
		api.POST("/kinds", handler.CreateKind)

		// Attributes
		api.GET("/attributes", handler.GetAttributes)
		api.POST("/attributes", handler.CreateAttribute)

		// Modes
		api.GET("/modes", handler.GetModes)
		api.POST("/modes", handler.CreateMode)

		// Causality
		api.GET("/substances/:id/causes", handler.GetCauses)
		api.POST("/causes", handler.AddCause)

		// Potentialities
		api.GET("/potentialities", handler.GetPotentialities)
		api.POST("/potentialities", handler.CreatePotentiality)
		api.POST("/potentialities/:id/actualize", handler.ActualizePotentiality)
		api.GET("/potentialities/:id/conditions", handler.CheckConditions)

		// Evolution
		api.GET("/substances/:id/evolution", handler.GetSubstanceEvolution)
	}

	return router, db
}

func TestHealthCheck(t *testing.T) {
	router, _ := setupTestAPI(t)

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.Equal(t, "healthy", response["status"])
}

func TestCreateSubstance(t *testing.T) {
	router, _ := setupTestAPI(t)

	substanceData := map[string]string{
		"name":    "Tree-001",
		"kind":    "Oak",
		"essence": "Living organism with potentiality for growth",
	}

	jsonData, _ := json.Marshal(substanceData)
	req, _ := http.NewRequest("POST", "/api/v1/substances", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response entities.Substance
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.Equal(t, "Tree-001", response.Name)
	assert.Equal(t, "Oak", response.Kind)
	assert.Equal(t, "Living organism with potentiality for growth", response.Essence)
	assert.NotEmpty(t, response.ID)
}

func TestCreateSubstance_InvalidData(t *testing.T) {
	router, _ := setupTestAPI(t)

	// Missing required fields
	substanceData := map[string]string{
		"name": "Tree-001",
		// Missing kind and essence
	}

	jsonData, _ := json.Marshal(substanceData)
	req, _ := http.NewRequest("POST", "/api/v1/substances", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetSubstances(t *testing.T) {
	router, db := setupTestAPI(t)

	// Create test substances
	substance1 := entities.NewSubstance("Tree-001", "Oak", "Essence 1")
	substance2 := entities.NewSubstance("Tree-002", "Pine", "Essence 2")

	db.Create(&substance1)
	db.Create(&substance2)

	req, _ := http.NewRequest("GET", "/api/v1/substances", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string][]entities.Substance
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.Len(t, response["substances"], 2)
}

func TestCreateKind(t *testing.T) {
	router, _ := setupTestAPI(t)

	kindData := map[string]string{
		"name":        "Human",
		"description": "Rational animal with potentiality for virtue",
	}

	jsonData, _ := json.Marshal(kindData)
	req, _ := http.NewRequest("POST", "/api/v1/kinds", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response entities.Kind
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.Equal(t, "Human", response.Name)
	assert.Equal(t, "Rational animal with potentiality for virtue", response.Description)
	assert.NotEmpty(t, response.ID)
}

func TestCreateAttribute(t *testing.T) {
	router, _ := setupTestAPI(t)

	attributeData := map[string]string{
		"name":        "color",
		"description": "Visual property of substances",
		"data_type":   "string",
	}

	jsonData, _ := json.Marshal(attributeData)
	req, _ := http.NewRequest("POST", "/api/v1/attributes", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response entities.Attribute
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.Equal(t, "color", response.Name)
	assert.Equal(t, "Visual property of substances", response.Description)
	assert.Equal(t, "string", response.DataType)
	assert.NotEmpty(t, response.ID)
}

func TestCreateMode(t *testing.T) {
	router, db := setupTestAPI(t)

	// Create test substance and attribute
	substance := entities.NewSubstance("Tree-001", "Oak", "Essence")
	db.Create(&substance)

	attribute := entities.NewAttribute("color", "Visual property", "string")
	db.Create(&attribute)

	modeData := map[string]string{
		"value":        "green",
		"substance_id": substance.ID,
		"attribute_id": attribute.ID,
	}

	jsonData, _ := json.Marshal(modeData)
	req, _ := http.NewRequest("POST", "/api/v1/modes", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response entities.Mode
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.Equal(t, "green", response.Value)
	assert.Equal(t, substance.ID, response.SubstanceID)
	assert.Equal(t, attribute.ID, response.AttributeID)
	assert.NotEmpty(t, response.ID)
}

func TestAddCause(t *testing.T) {
	router, _ := setupTestAPI(t)

	causeData := map[string]string{
		"from_entity": "substance-1",
		"to_entity":   "wood_water",
		"cause_type":  "material",
	}

	jsonData, _ := json.Marshal(causeData)
	req, _ := http.NewRequest("POST", "/api/v1/causes", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response entities.CausalRelation
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.Equal(t, "material", response.CauseType)
	assert.Equal(t, "substance-1", response.FromEntity)
	assert.Equal(t, "wood_water", response.ToEntity)
	assert.NotEmpty(t, response.ID)
}

func TestAddCause_InvalidType(t *testing.T) {
	router, _ := setupTestAPI(t)

	causeData := map[string]string{
		"from_entity": "substance-1",
		"to_entity":   "wood_water",
		"cause_type":  "invalid",
	}

	jsonData, _ := json.Marshal(causeData)
	req, _ := http.NewRequest("POST", "/api/v1/causes", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCreatePotentiality(t *testing.T) {
	router, db := setupTestAPI(t)

	// Create test substance
	substance := entities.NewSubstance("Tree-001", "Oak", "Essence")
	db.Create(&substance)

	potentialityData := map[string]string{
		"name":         "Grow Leaves",
		"description":  "Tree can grow leaves in spring",
		"conditions":   `[{"type":"attribute","name":"season","value":"spring"}]`,
		"substance_id": substance.ID,
	}

	jsonData, _ := json.Marshal(potentialityData)
	req, _ := http.NewRequest("POST", "/api/v1/potentialities", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response entities.Potentiality
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.Equal(t, "Grow Leaves", response.Name)
	assert.Equal(t, "Tree can grow leaves in spring", response.Description)
	assert.Equal(t, substance.ID, response.SubstanceID)
	assert.NotEmpty(t, response.ID)
}

func TestGetSubstanceNotFound(t *testing.T) {
	router, _ := setupTestAPI(t)

	req, _ := http.NewRequest("GET", "/api/v1/substances/non-existent-id", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateSubstance(t *testing.T) {
	router, db := setupTestAPI(t)

	// Create test substance
	substance := entities.NewSubstance("Tree-001", "Oak", "Original essence")
	db.Create(&substance)

	updateData := map[string]string{
		"name":    "Tree-001-Updated",
		"essence": "Updated essence",
	}

	jsonData, _ := json.Marshal(updateData)
	req, _ := http.NewRequest("PUT", "/api/v1/substances/"+substance.ID, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response entities.Substance
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.Equal(t, "Tree-001-Updated", response.Name)
	assert.Equal(t, "Updated essence", response.Essence)
	assert.Equal(t, "Oak", response.Kind) // Should remain unchanged
}

func TestDeleteSubstance(t *testing.T) {
	router, db := setupTestAPI(t)

	// Create test substance
	substance := entities.NewSubstance("Tree-001", "Oak", "Essence")
	db.Create(&substance)

	req, _ := http.NewRequest("DELETE", "/api/v1/substances/"+substance.ID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Verify substance is deleted
	var count int64
	db.Model(&entities.Substance{}).Where("id = ?", substance.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}
