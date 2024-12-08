package usecases

import (
	"FP_GO_PBKK-D/internal/domain"
	"FP_GO_PBKK-D/internal/repositories"
)

type PlaylistUsecase struct {
	Repo *repositories.PlaylistRepository
}

// Fungsi untuk mengambil playlist berdasarkan slug
func (uc *PlaylistUsecase) GetPlaylistBySlug(slug string) (*domain.Playlist, error) {
	return uc.Repo.FindBySlug(slug)
}

// Fungsi untuk membuat playlist baru
func (uc *PlaylistUsecase) CreatePlaylist(playlist *domain.Playlist) error {
	return uc.Repo.Create(playlist)
}

// Fungsi untuk mengambil semua playlist
func (uc *PlaylistUsecase) GetAllPlaylists() ([]domain.Playlist, error) {
	return uc.Repo.FindAll()
}

// Fungsi untuk mengambil playlist berdasarkan ID
func (uc *PlaylistUsecase) GetPlaylistByID(id uint) (*domain.Playlist, error) {
	return uc.Repo.FindByID(id)
}

// Fungsi untuk memperbarui playlist
func (uc *PlaylistUsecase) UpdatePlaylist(playlist *domain.Playlist) error {
	return uc.Repo.Update(playlist)
}

// Fungsi untuk menghapus playlist
func (uc *PlaylistUsecase) DeletePlaylist(id uint) error {
	return uc.Repo.Delete(id)
}
