package api

import (
	"github.com/gin-gonic/gin"
	"go-web/entity"
	"go-web/model"
	"net/http"
)

type AddStudent struct {
	Name string `json:"name" binding:"required"`
	Sex  int    `json:"sex" binding:"required"`
}

func Insert(ctx *gin.Context) {
	var student AddStudent
	err := ctx.ShouldBindJSON(&student)
	if err != nil {
		ctx.JSON(http.StatusOK, model.BaseResponse{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	entity.Save(student.Name, student.Sex)
	ctx.JSON(http.StatusOK, model.BaseResponse{
		Code: http.StatusOK,
		Msg:  "保存成功",
	})
}

type DeleteStudent struct {
	Id int `json:"id" binding:"required"`
}

func Delete(ctx *gin.Context) {
	var student DeleteStudent
	err := ctx.ShouldBindJSON(&student)
	if err != nil {
		ctx.JSON(http.StatusOK, model.BaseResponse{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	entity.Delete(student.Id)
	ctx.JSON(http.StatusOK, model.BaseResponse{
		Code: http.StatusOK,
		Msg:  "删除成功",
	})
}

func SelectPage(ctx *gin.Context) {
	var page model.Page
	err := ctx.ShouldBind(&page)
	if err != nil {
		ctx.JSON(http.StatusOK, model.BaseResponse{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	students, err := entity.Page(page)
	if err != nil {
		ctx.JSON(http.StatusOK, model.BaseResponse{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{
		BaseResponse: model.BaseResponse{
			Code: http.StatusOK,
			Msg:  "查询成功",
		},
		Data: students,
	})
}
