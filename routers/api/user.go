package api

import (
	"github.com/gin-gonic/gin"
	"go-web/entity"
	"go-web/model"
	"net/http"
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
	_, err = entity.SaveUser(user.CmbUid, user.WxName)
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
