package repositories

import (
	"FP_GO_PBKK-D/internal/domain"

	"gorm.io/gorm"
)

type SongRepository struct {
	DB *gorm.DB
}

func (r *SongRepository) FindAll() ([]domain.Song, error) {
	var songs []domain.Song
	err := r.DB.Preload("Artist").Preload("Category").Find(&songs).Error
	return songs, err
}

func (r *SongRepository) FindBySlug(slug string) (*domain.Song, error) {
	var song domain.Song
	err := r.DB.Preload("Artist").Preload("Category").Where("slug = ?", slug).First(&song).Error
	if err != nil {
		return nil, err
	}
	return &song, nil
}
