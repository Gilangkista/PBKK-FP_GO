package migrations

import (
	"FP_GO_PBKK-D/internal/domain"

	"gorm.io/gorm"
)

func MigratePlaylist(db *gorm.DB) error {
	return db.AutoMigrate(&domain.Playlist{})
}
