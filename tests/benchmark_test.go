package tests

import (
	"testing"

	"github.com/apodicticscott/oaas/internal/causality"
	"github.com/apodicticscott/oaas/internal/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func BenchmarkCreateSubstance(b *testing.B) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		b.Fatalf("Failed to connect to database: %v", err)
	}
	
	err = db.AutoMigrate(&entities.Substance{})
	if err != nil {
		b.Fatalf("Failed to migrate: %v", err)
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		substance := entities.NewSubstance("Tree-001", "Oak", "Living organism")
		if err := db.Create(substance).Error; err != nil {
			b.Fatalf("Failed to create substance: %v", err)
		}
	}
}

func BenchmarkCreatePotentiality(b *testing.B) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		b.Fatalf("Failed to connect to database: %v", err)
	}
	
	err = db.AutoMigrate(&entities.Substance{}, &entities.Potentiality{})
	if err != nil {
		b.Fatalf("Failed to migrate: %v", err)
	}
	
	// Create a test substance
	substance := entities.NewSubstance("Tree-001", "Oak", "Living organism")
	if err := db.Create(substance).Error; err != nil {
		b.Fatalf("Failed to create substance: %v", err)
	}
	
	engine := causality.NewEngine(db)
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_, err := engine.CreatePotentiality("Grow Leaves", "Description", "", substance.ID)
		if err != nil {
			b.Fatalf("Failed to create potentiality: %v", err)
		}
	}
}

func BenchmarkCheckConditions(b *testing.B) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		b.Fatalf("Failed to connect to database: %v", err)
	}
	
	err = db.AutoMigrate(&entities.Substance{}, &entities.Attribute{}, &entities.Mode{}, &entities.Potentiality{})
	if err != nil {
		b.Fatalf("Failed to migrate: %v", err)
	}
	
	// Create test data
	substance := entities.NewSubstance("Tree-001", "Oak", "Living organism")
	if err := db.Create(substance).Error; err != nil {
		b.Fatalf("Failed to create substance: %v", err)
	}
	
	attribute := entities.NewAttribute("color", "Visual property", "string")
	if err := db.Create(attribute).Error; err != nil {
		b.Fatalf("Failed to create attribute: %v", err)
	}
	
	mode := entities.NewMode("green", substance.ID, attribute.ID)
	if err := db.Create(mode).Error; err != nil {
		b.Fatalf("Failed to create mode: %v", err)
	}
	
	conditions := `[{"type":"mode","name":"color","value":"green"}]`
	potentiality := entities.NewPotentiality("Grow Leaves", "Description", conditions, substance.ID)
	if err := db.Create(potentiality).Error; err != nil {
		b.Fatalf("Failed to create potentiality: %v", err)
	}
	
	engine := causality.NewEngine(db)
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_, _, err := engine.CheckConditions(potentiality.ID)
		if err != nil {
			b.Fatalf("Failed to check conditions: %v", err)
		}
	}
}

func BenchmarkActualizePotentiality(b *testing.B) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		b.Fatalf("Failed to connect to database: %v", err)
	}
	
	err = db.AutoMigrate(&entities.Substance{}, &entities.Attribute{}, &entities.Mode{}, &entities.Potentiality{}, &entities.Actuality{})
	if err != nil {
		b.Fatalf("Failed to migrate: %v", err)
	}
	
	// Create test data
	substance := entities.NewSubstance("Tree-001", "Oak", "Living organism")
	if err := db.Create(substance).Error; err != nil {
		b.Fatalf("Failed to create substance: %v", err)
	}
	
	attribute := entities.NewAttribute("color", "Visual property", "string")
	if err := db.Create(attribute).Error; err != nil {
		b.Fatalf("Failed to create attribute: %v", err)
	}
	
	mode := entities.NewMode("green", substance.ID, attribute.ID)
	if err := db.Create(mode).Error; err != nil {
		b.Fatalf("Failed to create mode: %v", err)
	}
	
	conditions := `[{"type":"mode","name":"color","value":"green"}]`
	potentiality := entities.NewPotentiality("Grow Leaves", "Description", conditions, substance.ID)
	if err := db.Create(potentiality).Error; err != nil {
		b.Fatalf("Failed to create potentiality: %v", err)
	}
	
	engine := causality.NewEngine(db)
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// Create a new potentiality for each iteration to avoid conflicts
		newPotentiality := entities.NewPotentiality("Grow Leaves", "Description", conditions, substance.ID)
		if err := db.Create(newPotentiality).Error; err != nil {
			b.Fatalf("Failed to create potentiality: %v", err)
		}
		
		_, err := engine.ActualizePotentiality(newPotentiality.ID, "Description")
		if err != nil {
			b.Fatalf("Failed to actualize potentiality: %v", err)
		}
	}
}
