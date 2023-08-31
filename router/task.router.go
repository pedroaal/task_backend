package router

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Task(server *gin.Engine) {
	task := server.Group("/task")
	{
		task.GET("", controllers.FindTasks)
		task.POST("", controllers.CreateTask)
		task.GET("/:id", controllers.FindTaskById)
		task.PUT("/:id", controllers.UpdateTask)
		task.DELETE("/:id", controllers.DeleteTask)
	}
}
