package routers

import (
	"github.com/gin-gonic/gin"
	"go-web/jwt"
	"go-web/model"
	"log"
	"net/http"
)

func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusOK, model.Response(http.StatusBadRequest, "token为空,请重新登录!"))
			ctx.Abort()
			return
		}
		log.Printf("开始验证jwt:%s", tokenString)
		token, err := jwt.ParseToken(tokenString)
		if err != nil || !token.Valid {
			log.Println("token校验失败或过期:", err)
			ctx.JSON(http.StatusOK, model.Response(http.StatusBadRequest, "老夫不管你什么错,去重新登录吧!"))
			ctx.Abort()
			return
		}
		//上面只验证了token的合法性(含过期),这里加上自定义验证token过期，不一定token合法就是没过期的.
		ctx.Next()
	}
}
