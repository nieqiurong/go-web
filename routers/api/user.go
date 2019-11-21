package api

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"go-web/entity"
	"go-web/model"
	"net/http"
	"strconv"
)

type AddUser struct {
	CmbUid string `form:"cmbUid" binding:"required"`
	WxName string `form:"wxName" binding:"required"`
}

func SaveUser(ctx *gin.Context) {
	var user AddUser
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusOK, model.BaseResponse{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	err = entity.SaveUser(user.CmbUid, user.WxName)
	if err != nil {
		ctx.JSON(http.StatusOK, model.BaseResponse{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, model.BaseResponse{
		Code: http.StatusOK,
		Msg:  "操作成功",
	})
}

func Test(ctx *gin.Context) {
	num, err := strconv.Atoi(ctx.DefaultQuery("num", "10000"))
	if err != nil {
		ctx.JSON(http.StatusOK, model.BaseResponse{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	for j := 0; j < num; j++ {
		go entity.SaveUser(uuid.NewV4().String(), "靓仔")
	}
}
