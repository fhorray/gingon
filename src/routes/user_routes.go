package routes

import (
	"github.com/fhorray/gingon/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes (r *gin.Engine) {
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUser)
	}
}