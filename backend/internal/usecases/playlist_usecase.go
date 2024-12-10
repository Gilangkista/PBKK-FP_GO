package usecases

import (
	"FP_GO_PBKK-D/internal/domain"
	"FP_GO_PBKK-D/internal/repositories"
)

type PlaylistUsecase struct {
	Repo *repositories.PlaylistRepository
}

func (uc *PlaylistUsecase) GetPlaylistBySlug(slug string) (*domain.Playlist, error) {
	return uc.Repo.FindBySlug(slug)
}

func (uc *PlaylistUsecase) CreatePlaylist(playlist *domain.Playlist) error {
	return uc.Repo.Create(playlist)
}

func (uc *PlaylistUsecase) GetAllPlaylists() ([]domain.Playlist, error) {
	return uc.Repo.FindAll()
}

func (uc *PlaylistUsecase) UpdatePlaylist(playlist *domain.Playlist) error {
	return uc.Repo.Update(playlist)
}

func (uc *PlaylistUsecase) DeletePlaylist(id uint) error {
	return uc.Repo.Delete(id)
}

func (uc *PlaylistUsecase) AddSongToPlaylist(playlistSlug string, songSlug string) error {
	return uc.Repo.AddSongToPlaylist(playlistSlug, songSlug)
}

func (uc *PlaylistUsecase) RemoveSongFromPlaylist(playlistSlug string, songSlug string) error {
	return uc.Repo.RemoveSongFromPlaylist(playlistSlug, songSlug)
}
