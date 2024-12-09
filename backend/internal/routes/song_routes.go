package routes

import (
	"FP_GO_PBKK-D/internal/controller"

	"github.com/gin-gonic/gin"
)

func SongRoutes(r *gin.Engine, ctrl *controller.SongController) {
	songGroup := r.Group("/songs")
	songGroup.GET("/", ctrl.GetSongs)
	songGroup.GET("/:slug", ctrl.GetSongBySlug)
}
