package usecases

import (
	"FP_GO_PBKK-D/internal/domain"
	"FP_GO_PBKK-D/internal/repositories"
)

type SongUsecase struct {
	Repo *repositories.SongRepository
}

func (u *SongUsecase) GetAllSongs() ([]domain.Song, error) {
	return u.Repo.FindAll()
}

func (u *SongUsecase) GetSongBySlug(slug string) (*domain.Song, error) {
	return u.Repo.FindBySlug(slug)
}
