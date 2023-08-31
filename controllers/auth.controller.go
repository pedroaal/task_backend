package controllers

import (
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)
	err := models.FindUserByMail(&user)

	if err != nil {
		utils.RespondJSON(ctx, 404, user)
		return
	}

	utils.RespondJSON(ctx, 200, user)
	return
}

func Logout(ctx *gin.Context) {
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

func Me(ctx *gin.Context) {
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
