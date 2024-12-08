package routes

import (
	"FP_GO_PBKK-D/internal/controller"

	"github.com/gin-gonic/gin"
)

func ArtistRoutes(r *gin.Engine, ctrl *controller.ArtistController) {
	artistGroup := r.Group("/artists")
	artistGroup.GET("/", ctrl.GetArtists)
}
