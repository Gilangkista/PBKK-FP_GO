package migrations

import (
	"FP_GO_PBKK-D/internal/domain"

	"gorm.io/gorm"
)

func MigrateArtist(db *gorm.DB) error {
	return db.AutoMigrate(&domain.Artist{})
}
