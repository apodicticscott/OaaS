package persistence

import (
	"github.com/apodicticscott/oaas/internal/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate tables
	db.AutoMigrate(
		&entities.Kind{},
		&entities.Attribute{},
		&entities.Substance{},
		&entities.Mode{},
		&entities.CausalRelation{},
		&entities.Potentiality{},
		&entities.Actuality{},
	)

	return db, nil
}
