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

// Fungsi untuk membuat playlist baru
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

// Fungsi untuk mengupdate playlist berdasarkan ID
func (c *PlaylistController) UpdatePlaylist(ctx *gin.Context) {
	// Mengambil ID dari URL parameter
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32) // Mengonversi string ID ke uint
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}

	// Mengambil data playlist yang akan di-update dari request body
	var playlist domain.Playlist
	if err := ctx.ShouldBindJSON(&playlist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Menambahkan ID yang didapat dari URL ke dalam data playlist yang akan di-update
	playlist.ID = uint(id)

	// Memanggil usecase untuk mengupdate playlist
	err = c.Usecase.UpdatePlaylist(&playlist)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan response dengan data playlist yang sudah terupdate
	ctx.JSON(http.StatusOK, playlist)
}

// Fungsi untuk menghapus playlist berdasarkan ID
func (c *PlaylistController) DeletePlaylist(ctx *gin.Context) {
	// Mengambil ID dari URL parameter dan mengonversinya ke uint
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32) // Mengonversi string ke uint
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}

	// Memanggil usecase dengan ID yang sudah dikonversi untuk menghapus playlist
	err = c.Usecase.DeletePlaylist(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan response bahwa playlist sudah terhapus
	ctx.JSON(http.StatusOK, gin.H{"message": "Playlist deleted"})
}
