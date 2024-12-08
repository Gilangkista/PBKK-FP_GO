package migrations

import (
	"FP_GO_PBKK-D/internal/domain"

	"gorm.io/gorm"
)

func MigrateCategory(db *gorm.DB) error {
	return db.AutoMigrate(&domain.Category{})
}
