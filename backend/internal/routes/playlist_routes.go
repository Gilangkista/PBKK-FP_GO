package routes

import (
	"FP_GO_PBKK-D/internal/controller"

	"github.com/gin-gonic/gin"
)

func PlaylistRoutes(r *gin.Engine, ctrl *controller.PlaylistController) {
	playlistGroup := r.Group("/playlists")
	playlistGroup.GET("/", ctrl.GetPlaylists) // Pastikan ini ada
	playlistGroup.GET("/:slug", ctrl.GetPlaylistBySlug)
	playlistGroup.POST("/", ctrl.CreatePlaylist)
	playlistGroup.PUT("/:id", ctrl.UpdatePlaylist)
	playlistGroup.DELETE("/:id", ctrl.DeletePlaylist)
}
