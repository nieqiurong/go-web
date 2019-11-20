package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/cache"
	"go-web/entity"
	"go-web/routers"
	"go-web/setting"
	"log"
	"net/http"
)

func init() {
	setting.InitConfig()
	entity.InitDb()
	//schedule.InitSchedule()
	cache.InitRedis()
}

func main() {
	serverConfig := setting.Application.Server
	gin.SetMode(serverConfig.Mode)
	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", serverConfig.Port)
	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
	}
	log.Printf("start http server listening %s", endPoint)
	_ = server.ListenAndServe()
}
