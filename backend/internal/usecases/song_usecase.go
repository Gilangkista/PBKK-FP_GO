package usecases

import (
	"FP_GO_PBKK-D/internal/domain"
	"FP_GO_PBKK-D/internal/repositories"
)

type SongUsecase struct {
	Repo *repositories.SongRepository
}

func (u *SongUsecase) CreateSong(title, slug string, artistID, categoryID uint) (*domain.Song, error) {
	song := &domain.Song{
		Title:      title,
		Slug:       slug,
		ArtistID:   artistID,
		CategoryID: categoryID,
	}
	err := u.Repo.Create(song)
	return song, err
}

func (u *SongUsecase) GetAllSongs() ([]domain.Song, error) {
	return u.Repo.FindAll()
}
