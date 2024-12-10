package routes

import (
	"FP_GO_PBKK-D/internal/controller"

	"github.com/gin-gonic/gin"
)

func PlaylistRoutes(r *gin.Engine, ctrl *controller.PlaylistController) {
	playlistGroupBySlug := r.Group("/playlists/detail")
	playlistGroupBySlug.GET("/:slug", ctrl.GetPlaylistBySlug)
	playlistGroupBySlug.PUT("/:slug", ctrl.UpdatePlaylistBySlug)
	playlistGroupBySlug.POST("/:slug/songs", ctrl.AddSongToPlaylist)
	playlistGroupBySlug.DELETE("/:slug/songs/:songSlug", ctrl.RemoveSongFromPlaylist)

	playlistGroupById := r.Group("/playlists")
	playlistGroupById.GET("/", ctrl.GetPlaylists)
	playlistGroupById.POST("/", ctrl.CreatePlaylist)
	playlistGroupById.DELETE("/:id", ctrl.DeletePlaylist)
}
