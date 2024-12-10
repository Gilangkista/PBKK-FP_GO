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

func (c *PlaylistController) GetPlaylists(ctx *gin.Context) {
	playlists, err := c.Usecase.GetAllPlaylists()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, playlists)
}

func (c *PlaylistController) GetPlaylistBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
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

	playlist.Slug = generateSlug(playlist.Name)

	if err := c.Usecase.CreatePlaylist(&playlist); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create playlist"})
		return
	}

	ctx.JSON(http.StatusOK, playlist)
}

func (c *PlaylistController) UpdatePlaylistBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")

	var playlist domain.Playlist
	if err := ctx.ShouldBindJSON(&playlist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingPlaylist, err := c.Usecase.GetPlaylistBySlug(slug)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Playlist not found"})
		return
	}

	existingPlaylist.Name = playlist.Name
	existingPlaylist.Description = playlist.Description

	err = c.Usecase.UpdatePlaylist(existingPlaylist)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update playlist"})
		return
	}

	ctx.JSON(http.StatusOK, existingPlaylist)
}

func (c *PlaylistController) DeletePlaylist(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}

	err = c.Usecase.DeletePlaylist(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Playlist deleted"})
}

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

func (c *PlaylistController) AddSongToPlaylist(ctx *gin.Context) {
	playlistSlug := ctx.Param("slug")
	songSlug := ctx.Query("songSlug") // Mengambil dari query parameter

	if songSlug == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "songSlug is required"})
		return
	}

	err := c.Usecase.AddSongToPlaylist(playlistSlug, songSlug)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to add song to playlist"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Song added to playlist"})
}

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
