package controller

import (
	"FP_GO_PBKK-D/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArtistController struct {
	Usecase *usecases.ArtistUsecase
}

func (c *ArtistController) GetArtists(ctx *gin.Context) {
	artists, err := c.Usecase.GetAllArtists()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, artists)
}
