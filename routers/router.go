package routers

import (
	"database/sql"
	"database/sql/driver"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-web/routers/api"
	"go-web/setting"
	"gopkg.in/go-playground/validator.v8"
	"reflect"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(Cors())
	if setting.Application.IsDebug() {
		r.Use(gin.Logger())
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r.GET("/ping", api.Ping)
	r.GET("/info", api.Info)
	student := r.Group("/student").Use(Jwt())
	{
		student.POST("/save", api.Insert)
		student.POST("/delete", api.Delete)
		student.POST("/update", api.Update)
		student.GET("/page", api.SelectPage)
	}
	user := r.Group("/")
	{
		user.POST("/login", api.Login)
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterCustomTypeFunc(func(field reflect.Value) interface{} {
			if valuer, ok := field.Interface().(driver.Valuer); ok {
				val, err := valuer.Value()
				if err == nil {
					return val
				}
			}
			return nil
		}, sql.NullString{}, sql.NullInt64{}, sql.NullBool{}, sql.NullFloat64{})
	}
	return r
}
