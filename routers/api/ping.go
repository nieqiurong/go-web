package api

import "github.com/gin-gonic/gin"

func Ping(ctx *gin.Context) {
	ctx.String(200, "pong")
}
