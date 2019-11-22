package api

import (
	"github.com/gin-gonic/gin"
	"go-web/entity"
	"go-web/jwt"
	"go-web/model"
	"log"
	"net/http"
)

type UserLogin struct {
	Account  string `form:"account" binding:"required"`
	PassWord string `form:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	var login UserLogin
	err := ctx.ShouldBind(&login)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Response(http.StatusBadRequest, err.Error()))
		return
	}
	user, err := entity.FindUserByAccount(login.Account)
	if err != nil {
		log.Printf("查询用户%s信息出现异常:%s", login.Account, err)
		ctx.JSON(http.StatusOK, model.Response(http.StatusBadRequest, err.Error()))
		return
	}
	if user == nil {
		log.Printf("用户%s未注册", login.Account)
		ctx.JSON(http.StatusOK, model.Response(model.RegisterCode, "用户不存在,快去注册吧!"))
		return
	}
	//这里假装密码验证对了
	cla := make(map[string]interface{})
	cla["test"] = "测试用户"
	token, err := jwt.CreateToken(user.Id, user.Account, cla)
	if err != nil {
		log.Printf("生成用户%stoken失败,错误原因:%s", user.Account, err)
		ctx.JSON(http.StatusOK, model.Response(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, model.Response(http.StatusOK, token))
}
