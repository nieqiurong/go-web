package api

import (
	"github.com/gin-gonic/gin"
	"go-web/setting"
	"net/http"
)

func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func Info(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, setting.Application.App)
}
