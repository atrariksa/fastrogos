package migrations

import (
	"github.com/atrariksa/fastrogos/rula/models"
	"gorm.io/gorm"
)

type Migrator struct {
	DB *gorm.DB
}

func (m *Migrator) MigrateUp() {
	m.DB.Debug().AutoMigrate(
		&models.User{},
	)
}
