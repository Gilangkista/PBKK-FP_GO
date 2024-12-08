package usecases

import (
	"FP_GO_PBKK-D/internal/domain"
	"FP_GO_PBKK-D/internal/repositories"
)

type ArtistUsecase struct {
	Repo *repositories.ArtistRepository
}

func (u *ArtistUsecase) CreateArtist(name string) (*domain.Artist, error) {
	artist := &domain.Artist{Name: name}
	err := u.Repo.Create(artist)
	return artist, err
}

func (u *ArtistUsecase) GetAllArtists() ([]domain.Artist, error) {
	return u.Repo.FindAll()
}
