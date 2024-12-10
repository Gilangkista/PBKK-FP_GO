package controller

import (
	"FP_GO_PBKK-D/internal/domain"
	"FP_GO_PBKK-D/internal/usecases"
	"net/http"
	"strconv"
	"strings"
	"unicode"

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

	// Buat slug berdasarkan nama playlist
	playlist.Slug = generateSlug(playlist.Name)

	if err := c.Usecase.CreatePlaylist(&playlist); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create playlist"})
		return
	}

	ctx.JSON(http.StatusOK, playlist)
}

// Fungsi untuk mengupdate playlist berdasarkan slug
func (c *PlaylistController) UpdatePlaylistBySlug(ctx *gin.Context) {
	// Mengambil slug dari URL parameter
	slug := ctx.Param("slug")

	// Mengambil data playlist yang akan di-update dari request body
	var playlist domain.Playlist
	if err := ctx.ShouldBindJSON(&playlist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memanggil usecase untuk mendapatkan playlist berdasarkan slug
	existingPlaylist, err := c.Usecase.GetPlaylistBySlug(slug)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Playlist not found"})
		return
	}

	// Mengupdate field playlist yang ada
	existingPlaylist.Name = playlist.Name
	existingPlaylist.Description = playlist.Description

	// Memanggil usecase untuk menyimpan perubahan
	err = c.Usecase.UpdatePlaylist(existingPlaylist)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update playlist"})
		return
	}

	// Mengembalikan playlist yang diperbarui
	ctx.JSON(http.StatusOK, existingPlaylist)
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

// Fungsi untuk membuat slug
func generateSlug(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' {
			return r
		}
		return -1
	}, name)
}

// Fungsi untuk menambahkan lagu ke playlist
func (c *PlaylistController) AddSongToPlaylist(ctx *gin.Context) {
	playlistSlug := ctx.Param("slug")
	songSlug := ctx.Param("songSlug")

	err := c.Usecase.AddSongToPlaylist(playlistSlug, songSlug)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to add song to playlist"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Song added to playlist"})
}

// Fungsi untuk menghapus lagu dari playlist
func (c *PlaylistController) RemoveSongFromPlaylist(ctx *gin.Context) {
	playlistSlug := ctx.Param("slug")
	songSlug := ctx.Param("songSlug")

	err := c.Usecase.RemoveSongFromPlaylist(playlistSlug, songSlug)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to remove song from playlist"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Song removed from playlist"})
}
