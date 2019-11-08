package api

import (
	"github.com/gin-gonic/gin"
	"go-web/entity"
)

type AddStudent struct {
	Name string `form:"name"`
	Sex  int    `form:"sex"`
}

func Insert(ctx *gin.Context) {
	var student AddStudent
	_ = ctx.ShouldBind(&student)
	entity.Save(student.Name, student.Sex)
}

type DeleteStudent struct {
	Id int `form:"id"`
}

func Delete(ctx *gin.Context) {
	var student DeleteStudent
	_ = ctx.ShouldBind(&student)
	entity.Delete(student.Id)
}
