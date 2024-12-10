package controller

import (
	"FP_GO_PBKK-D/internal/usecases"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SongController struct {
	Usecase *usecases.SongUsecase
}

func (c *SongController) GetSongs(ctx *gin.Context) {
	songs, err := c.Usecase.GetAllSongs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, songs)
}

func (c *SongController) GetSongBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	log.Printf("Received slug: %s", slug)
	song, err := c.Usecase.GetSongBySlug(slug)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}
	ctx.JSON(http.StatusOK, song)
}
