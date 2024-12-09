package repositories

import (
	"FP_GO_PBKK-D/internal/domain"

	"gorm.io/gorm"
)

type PlaylistRepository struct {
	DB *gorm.DB
}

// Fungsi untuk mencari playlist berdasarkan slug
func (r *PlaylistRepository) FindBySlug(slug string) (*domain.Playlist, error) {
	var playlist domain.Playlist
	err := r.DB.Preload("Songs").Where("slug = ?", slug).First(&playlist).Error
	if err != nil {
		return nil, err
	}
	return &playlist, nil
}

func (r *PlaylistRepository) Create(playlist *domain.Playlist) error {
	return r.DB.Create(playlist).Error
}

func (r *PlaylistRepository) FindAll() ([]domain.Playlist, error) {
	var playlists []domain.Playlist
	err := r.DB.Preload("Songs").Find(&playlists).Error
	return playlists, err
}

func (r *PlaylistRepository) FindByID(id uint) (*domain.Playlist, error) {
	var playlist domain.Playlist
	err := r.DB.Preload("Songs").First(&playlist, id).Error
	return &playlist, err
}

func (r *PlaylistRepository) Update(playlist *domain.Playlist) error {
	return r.DB.Save(playlist).Error
}

func (r *PlaylistRepository) Delete(id uint) error {
	return r.DB.Delete(&domain.Playlist{}, id).Error
}
