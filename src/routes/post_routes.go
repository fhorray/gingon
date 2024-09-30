package routes

import (
	"github.com/fhorray/gingon/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetupPostRoutes (r *gin.Engine) {
	postRoutes := r.Group("/posts")
	{
		postRoutes.GET("/", controllers.GetPosts)
		postRoutes.POST("/", controllers.CreatePost)
	}
}