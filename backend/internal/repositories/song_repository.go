package repositories

import (
	"FP_GO_PBKK-D/internal/domain"

	"gorm.io/gorm"
)

type SongRepository struct {
	DB *gorm.DB
}

func (r *SongRepository) Create(song *domain.Song) error {
	return r.DB.Create(song).Error
}

func (r *SongRepository) FindAll() ([]domain.Song, error) {
	var songs []domain.Song
	err := r.DB.Preload("Artist").Preload("Category").Find(&songs).Error
	return songs, err
}

func (r *SongRepository) FindByID(id uint) (*domain.Song, error) {
	var song domain.Song
	err := r.DB.Preload("Artist").Preload("Category").First(&song, id).Error
	return &song, err
}

func (r *SongRepository) Update(song *domain.Song) error {
	return r.DB.Save(song).Error
}

func (r *SongRepository) Delete(id uint) error {
	return r.DB.Delete(&domain.Song{}, id).Error
}
