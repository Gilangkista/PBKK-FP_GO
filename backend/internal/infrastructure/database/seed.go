package database

import (
	"FP_GO_PBKK-D/internal/domain"
	"log"

	"gorm.io/gorm"
)

func SeedArtists(db *gorm.DB) error {
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

	for _, artist := range artists {
		var existingArtist domain.Artist
		err := db.Where("name = ?", artist.Name).First(&existingArtist).Error
		if err == nil {
			log.Printf("Artist '%s' already exists, skipping...", artist.Name)
			continue
		}

		if err := db.Create(&artist).Error; err != nil {
			return err
		}
		log.Printf("Seeded artist: %s", artist.Name)
	}

	return nil
}

func SeedCategories(db *gorm.DB) error {
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
			Name: "Indie",
			Slug: "indie",
		},
	}

	for _, category := range categories {
		var existingCategory domain.Category
		err := db.Where("name = ?", category.Name).First(&existingCategory).Error
		if err == nil {
			log.Printf("Category '%s' already exists, skipping...", category.Name)
			continue
		}

		if err := db.Create(&category).Error; err != nil {
			return err
		}
		log.Printf("Seeded category: %s", category.Name)
	}

	return nil
}

func SeedSongs(db *gorm.DB) error {
	songs := []domain.Song{
		{
			Title:      "Satu Bulan",
			Slug:       "satu-bulan",
			ArtistID:   1,
			CategoryID: 1,
		},
		{
			Title:      "Untungnya, Hidup Harus Tetap Berjalan",
			Slug:       "untungnya,-hidup-harus-tetap-berjalan",
			ArtistID:   1,
			CategoryID: 1,
		},
		{
			Title:      "Someone Like You",
			Slug:       "someone-like-you",
			ArtistID:   2,
			CategoryID: 1,
		},
		{
			Title:      "All I Ask",
			Slug:       "all-i-ask",
			ArtistID:   2,
			CategoryID: 1,
		},
		{
			Title:      "Chasing Pavements",
			Slug:       "chasing-pavements",
			ArtistID:   2,
			CategoryID: 1,
		},
		{
			Title:      "Backburner",
			Slug:       "backburner",
			ArtistID:   3,
			CategoryID: 5,
		},
		{
			Title:      "lowkey",
			Slug:       "lowkey",
			ArtistID:   3,
			CategoryID: 5,
		},
		{
			Title:      "505",
			Slug:       "505",
			ArtistID:   5,
			CategoryID: 2,
		},
		{
			Title:      "Do I Wanna Know?",
			Slug:       "do-i-wanna-know?",
			ArtistID:   5,
			CategoryID: 2,
		},
		{
			Title:      "Lebih Indah",
			Slug:       "lebih-indah",
			ArtistID:   4,
			CategoryID: 1,
		},
	}

	for _, song := range songs {
		var existingSong domain.Song
		err := db.Where("slug = ?", song.Slug).First(&existingSong).Error
		if err == nil {
			log.Printf("Song '%s' already exists, skipping...", song.Title)
			continue
		}

		if err := db.Create(&song).Error; err != nil {
			return err
		}
		log.Printf("Seeded song: %s", song.Title)
	}

	return nil
}

func SeedPlaylists(db *gorm.DB) error {
	playlists := []domain.Playlist{
		{
			Name:        "Top Hits",
			Slug:        "top-hits",
			Description: StringPtr("The top hits of the year"),
		},
		{
			Name:        "Chill Vibes",
			Slug:        "chill-vibes",
			Description: StringPtr("Relaxing music for a chill atmosphere"),
		},
		{
			Name:        "Workout Playlist",
			Slug:        "workout-playlist",
			Description: StringPtr("Energetic tracks for workout sessions"),
		},
	}

	for _, playlist := range playlists {
		var existingPlaylist domain.Playlist
		err := db.Where("slug = ?", playlist.Slug).First(&existingPlaylist).Error
		if err == nil {
			log.Printf("Playlist '%s' already exists, skipping...", playlist.Name)
			continue
		}

		if err := db.Create(&playlist).Error; err != nil {
			return err
		}
		log.Printf("Seeded playlist: %s", playlist.Name)

		var createdPlaylist domain.Playlist
		err = db.Where("slug = ?", playlist.Slug).First(&createdPlaylist).Error
		if err != nil {
			return err
		}

		var songs []domain.Song
		err = db.Limit(3).Find(&songs).Error
		if err != nil {
			return err
		}

		for _, song := range songs {
			err := db.Model(&createdPlaylist).Association("Songs").Append(&song)
			if err != nil {
				return err
			}
			log.Printf("Added song '%s' to playlist '%s'", song.Title, playlist.Name)
		}
	}

	return nil
}

func StringPtr(s string) *string {
	return &s
}
