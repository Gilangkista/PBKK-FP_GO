package database

import (
	"FP_GO_PBKK-D/internal/domain"
	"log"

	"gorm.io/gorm"
)

func SeedArtists(db *gorm.DB) error {
	// Data artists untuk di-seed
	artists := []domain.Artist{
		{
			Name: "Bernadya",
			Slug: "bernadya",
		},
		{
			Name: "Adele",
			Slug: "adele",
		},
		{
			Name: "NIKI",
			Slug: "niki",
		},
		{
			Name: "Adera",
			Slug: "adera",
		},
		{
			Name: "Arctic Monkeys",
			Slug: "arctic-monkeys",
		},
	}

	// Insert data ke dalam database
	for _, artist := range artists {
		// Cek apakah data sudah ada berdasarkan nama
		var existingArtist domain.Artist
		err := db.Where("name = ?", artist.Name).First(&existingArtist).Error
		if err == nil {
			log.Printf("Artist '%s' already exists, skipping...", artist.Name)
			continue
		}

		// Tambahkan jika belum ada
		if err := db.Create(&artist).Error; err != nil {
			return err
		}
		log.Printf("Seeded artist: %s", artist.Name)
	}

	return nil
}

func SeedCategories(db *gorm.DB) error {
	// Data categories untuk di-seed
	categories := []domain.Category{
		{
			Name: "Pop",
			Slug: "pop",
		},
		{
			Name: "Rock",
			Slug: "rock",
		},
		{
			Name: "Jazz",
			Slug: "jazz",
		},
		{
			Name: "Classic",
			Slug: "classic",
		},
		{
			Name: "Hip-hop",
			Slug: "hip-hop",
		},
	}

	// Insert data ke dalam database
	for _, category := range categories {
		// Cek apakah data sudah ada berdasarkan nama
		var existingCategory domain.Category
		err := db.Where("name = ?", category.Name).First(&existingCategory).Error
		if err == nil {
			log.Printf("Category '%s' already exists, skipping...", category.Name)
			continue
		}

		// Tambahkan jika belum ada
		if err := db.Create(&category).Error; err != nil {
			return err
		}
		log.Printf("Seeded category: %s", category.Name)
	}

	return nil
}

func SeedSongs(db *gorm.DB) error {
	// Data songs untuk di-seed
	songs := []domain.Song{
		{
			Title:      "Someone Like You",
			Slug:       "someone-like-you",
			ArtistID:   1,
			CategoryID: 1,
		},
		{
			Title:      "Bad Habits",
			Slug:       "bad-habits",
			ArtistID:   2,
			CategoryID: 2,
		},
		{
			Title:      "High School Musical",
			Slug:       "high-school-musical",
			ArtistID:   3,
			CategoryID: 3,
		},
		{
			Title:      "Take Me To Church",
			Slug:       "take-me-to-church",
			ArtistID:   4,
			CategoryID: 4,
		},
		{
			Title:      "Billionaire",
			Slug:       "billionaire",
			ArtistID:   5, // Ganti dengan ID artist yang sesuai
			CategoryID: 5,
		},
	}

	// Insert data ke dalam database
	for _, song := range songs {
		// Cek apakah data sudah ada berdasarkan slug
		var existingSong domain.Song
		err := db.Where("slug = ?", song.Slug).First(&existingSong).Error
		if err == nil {
			log.Printf("Song '%s' already exists, skipping...", song.Title)
			continue
		}

		// Tambahkan jika belum ada
		if err := db.Create(&song).Error; err != nil {
			return err
		}
		log.Printf("Seeded song: %s", song.Title)
	}

	return nil
}
