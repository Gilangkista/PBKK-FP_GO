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

// Update playlist
func (r *PlaylistRepository) Update(playlist *domain.Playlist) error {
	return r.DB.Save(playlist).Error
}

// Delete playlist
func (r *PlaylistRepository) Delete(id uint) error {
	return r.DB.Delete(&domain.Playlist{}, id).Error
}

// Fungsi untuk menambahkan lagu ke playlist
func (r *PlaylistRepository) AddSongToPlaylist(playlistSlug string, songSlug string) error {
	var playlist domain.Playlist
	var song domain.Song

	// Cari playlist berdasarkan slug
	err := r.DB.Preload("Songs").Where("slug = ?", playlistSlug).First(&playlist).Error
	if err != nil {
		return err
	}

	// Cari lagu berdasarkan slug
	err = r.DB.Where("slug = ?", songSlug).First(&song).Error
	if err != nil {
		return err
	}

	// Tambahkan lagu ke playlist
	playlist.Songs = append(playlist.Songs, song)

	// Simpan perubahan playlist ke database
	return r.DB.Save(&playlist).Error
}

// Fungsi untuk menghapus lagu dari playlist
func (r *PlaylistRepository) RemoveSongFromPlaylist(playlistSlug string, songSlug string) error {
	var playlist domain.Playlist
	var song domain.Song

	// Cari playlist berdasarkan slug
	err := r.DB.Preload("Songs").Where("slug = ?", playlistSlug).First(&playlist).Error
	if err != nil {
		return err
	}

	// Cari lagu berdasarkan slug
	err = r.DB.Where("slug = ?", songSlug).First(&song).Error
	if err != nil {
		return err
	}

	// Hapus lagu dari playlist
	var updatedSongs []domain.Song
	for _, s := range playlist.Songs {
		if s.Slug != songSlug {
			updatedSongs = append(updatedSongs, s)
		}
	}

	// Perbarui playlist
	playlist.Songs = updatedSongs

	// Simpan perubahan ke database
	return r.DB.Save(&playlist).Error
}
