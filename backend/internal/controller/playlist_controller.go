package controller

import (
	"FP_GO_PBKK-D/internal/domain"
	"FP_GO_PBKK-D/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlaylistController struct {
	Usecase *usecases.PlaylistUsecase
}

// Fungsi untuk mengambil semua playlist
func (c *PlaylistController) GetPlaylists(ctx *gin.Context) {
	playlists, err := c.Usecase.GetAllPlaylists()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, playlists)
}

// Fungsi untuk mendapatkan playlist berdasarkan slug
func (c *PlaylistController) GetPlaylistBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug") // Mendapatkan slug dari URL
	playlist, err := c.Usecase.GetPlaylistBySlug(slug)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Playlist not found"})
		return
	}
	ctx.JSON(http.StatusOK, playlist)
}

func (c *PlaylistController) CreatePlaylist(ctx *gin.Context) {
	var playlist domain.Playlist
	if err := ctx.ShouldBindJSON(&playlist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.Usecase.CreatePlaylist(&playlist)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, playlist)
}

func (c *PlaylistController) UpdatePlaylist(ctx *gin.Context) {
	var playlist domain.Playlist
	if err := ctx.ShouldBindJSON(&playlist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.Usecase.UpdatePlaylist(&playlist)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, playlist)
}

func (c *PlaylistController) DeletePlaylist(ctx *gin.Context) {
	// Mengambil id dari URL parameter dan mengonversinya ke uint
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32) // Mengonversi string ke uint
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}

	// Memanggil usecase dengan id yang sudah dikonversi
	err = c.Usecase.DeletePlaylist(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Playlist deleted"})
}
