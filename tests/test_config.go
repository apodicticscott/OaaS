package tests

import (
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TestConfig holds configuration for tests
type TestConfig struct {
	DB *gorm.DB
}

// SetupTestConfig initializes test configuration
func SetupTestConfig(t *testing.T) *TestConfig {
	// Use in-memory SQLite for tests
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	return &TestConfig{
		DB: db,
	}
}

// CleanupTestConfig cleans up test resources
func CleanupTestConfig(t *testing.T, config *TestConfig) {
	if config.DB != nil {
		sqlDB, err := config.DB.DB()
		if err != nil {
			t.Logf("Warning: Failed to get underlying DB: %v", err)
			return
		}
		sqlDB.Close()
	}
}

// SkipIfShort skips the test if the short flag is set
func SkipIfShort(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test")
	}
}

// SkipIfNoDB skips the test if database is not available
func SkipIfNoDB(t *testing.T) {
	if os.Getenv("SKIP_DB_TESTS") == "true" {
		t.Skip("Skipping database tests")
	}
}
