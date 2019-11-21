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
		ctx.JSON(http.StatusOK, model.Response(http.StatusBadRequest, err.Error()))
		return
	}
	err = entity.Save(student.Name, student.Sex)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Response(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, model.Response(http.StatusOK, "保存成功"))
}

type DeleteStudent struct {
	Id int `json:"id" binding:"required"`
}

func Delete(ctx *gin.Context) {
	var student DeleteStudent
	err := ctx.ShouldBindJSON(&student)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Response(http.StatusBadRequest, err.Error()))
		return
	}
	err = entity.Delete(student.Id)
	if err != nil {
		model.Response(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, model.Response(http.StatusOK, "删除成功"))
}

type UpdateStudent struct {
	Name string `json:"name" binding:"required"`
	Id   int    `json:"id" binding:"required"`
}

func Update(ctx *gin.Context) {
	var student UpdateStudent
	err := ctx.ShouldBindJSON(&student)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Response(http.StatusBadRequest, err.Error()))
		return
	}
	err = entity.Update(student.Name, student.Id)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Response(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, model.Response(http.StatusOK, "更新成功"))
}

func SelectPage(ctx *gin.Context) {
	var page model.Page
	err := ctx.ShouldBind(&page)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Response(http.StatusBadRequest, err.Error()))
		return
	}
	students, err := entity.Page(page)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Response(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, model.ResponseData(http.StatusOK, "查询成功", students))
}
