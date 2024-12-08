package controller

import (
	"FP_GO_PBKK-D/internal/usecases"
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