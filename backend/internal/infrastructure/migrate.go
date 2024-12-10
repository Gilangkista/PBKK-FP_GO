package infrastructure

import (
	"FP_GO_PBKK-D/internal/infrastructure/database"
	"FP_GO_PBKK-D/internal/infrastructure/migrations"
	"log"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	log.Println("Starting database migration...")

	migrations.MigrateArtist(db) // For artist
	database.SeedArtists(db)
	migrations.MigrateCategory(db) // For category
	database.SeedCategories(db)
	migrations.MigrateSong(db) // For song
	database.SeedSongs(db)
	migrations.MigratePlaylist(db) // For playlist

	log.Println("Database migration completed successfully.")
}
