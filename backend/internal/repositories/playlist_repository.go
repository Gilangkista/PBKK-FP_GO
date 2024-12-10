package repositories

import (
	"FP_GO_PBKK-D/internal/domain"

	"gorm.io/gorm"
)

type PlaylistRepository struct {
	DB *gorm.DB
}

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

func (r *PlaylistRepository) Update(playlist *domain.Playlist) error {
	return r.DB.Save(playlist).Error
}

func (r *PlaylistRepository) Delete(id uint) error {
	return r.DB.Delete(&domain.Playlist{}, id).Error
}

func (r *PlaylistRepository) AddSongToPlaylist(playlistSlug string, songSlug string) error {
	var playlist domain.Playlist
	var song domain.Song

	err := r.DB.Preload("Songs").Where("slug = ?", playlistSlug).First(&playlist).Error
	if err != nil {
		return err
	}

	err = r.DB.Where("slug = ?", songSlug).First(&song).Error
	if err != nil {
		return err
	}

	playlist.Songs = append(playlist.Songs, song)

	return r.DB.Save(&playlist).Error
}

func (r *PlaylistRepository) RemoveSongFromPlaylist(playlistSlug string, songSlug string) error {
	var playlist domain.Playlist
	var song domain.Song

	err := r.DB.Preload("Songs").Where("slug = ?", playlistSlug).First(&playlist).Error
	if err != nil {
		return err
	}

	err = r.DB.Where("slug = ?", songSlug).First(&song).Error
	if err != nil {
		return err
	}

	assoc := r.DB.Model(&playlist).Association("Songs")
	if err := assoc.Delete(&song); err != nil {
		return err
	}
	return nil
}
