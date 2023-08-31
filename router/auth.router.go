package router

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Auth(server *gin.Engine) {
	auth := server.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/login", controllers.Logout)
		auth.GET("/me", controllers.Me)
	}
}
