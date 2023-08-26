package controllers

import (
	"server/utils"
	"server/models"
	"github.com/gin-gonic/gin"
)

func FindTasks(ctx *gin.Context) {
	var task []models.Task
	err := models.FindTasks(&task)

	if err != nil {
		utils.RespondJSON(ctx, 404, task)
		return
	}

	utils.RespondJSON(ctx, 200, task)
	return
}

func CreateTask(ctx *gin.Context) {
	var task models.Task
	ctx.BindJSON(&task)
	err := models.CreateTask(&task)

	if err != nil {
		utils.RespondJSON(ctx, 404, task)
		return
	}
	
	utils.RespondJSON(ctx, 200, task)
	return
}

func FindTaskById(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	
	var task models.Task
	err := models.FindTaskById(&task, id)

	if err != nil {
		utils.RespondJSON(ctx, 404, task)
		return
	}
	
	utils.RespondJSON(ctx, 200, task)
	return
}

func UpdateTask(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	var task models.Task
	err := models.FindTaskById(&task, id)

	if err != nil {
		utils.RespondJSON(ctx, 404, task)
		return
	}
	
	ctx.BindJSON(&task)
	err = models.UpdateTask(&task, id)

	if err != nil {
		utils.RespondJSON(ctx, 404, task)
		return
	}
	
	utils.RespondJSON(ctx, 200, task)
	return
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	var task models.Task
	err := models.DeleteTask(&task, id)
	
	if err != nil {
		utils.RespondJSON(ctx, 404, task)
		return
	}
	
	utils.RespondJSON(ctx, 200, task)
	return
}