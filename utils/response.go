package utils

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
}

func RespondJSON(ctx *gin.Context, status int, payload interface{}) {
	ctx.JSON(status, payload)
}