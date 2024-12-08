package repositories

import (
	"FP_GO_PBKK-D/internal/domain"

	"gorm.io/gorm"
)

type ArtistRepository struct {
	DB *gorm.DB
}

func (r *ArtistRepository) Create(artist *domain.Artist) error {
	return r.DB.Create(artist).Error
}

func (r *ArtistRepository) FindAll() ([]domain.Artist, error) {
	var artists []domain.Artist
	err := r.DB.Find(&artists).Error
	return artists, err
}

func (r *ArtistRepository) FindByID(id uint) (*domain.Artist, error) {
	var artist domain.Artist
	err := r.DB.First(&artist, id).Error
	return &artist, err
}

func (r *ArtistRepository) Update(artist *domain.Artist) error {
	return r.DB.Save(artist).Error
}

func (r *ArtistRepository) Delete(id uint) error {
	return r.DB.Delete(&domain.Artist{}, id).Error
}
