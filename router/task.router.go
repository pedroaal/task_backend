package router

import (
	"server/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000", "https://todos-frontend.onrender.com"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS", "GET", "POST", "PUT", "DELETE")
	router.Use(cors.New(corsConfig))

	task := router.Group("/task")
	{
		task.GET("", controllers.FindTasks)
		task.POST("", controllers.CreateTask)
		task.GET("/:id", controllers.FindTaskById)
		task.PUT("/:id", controllers.UpdateTask)
		task.DELETE("/:id", controllers.DeleteTask)
	}

	return router
}
