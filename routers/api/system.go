package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/mem"
	"go-web/setting"
	"net/http"
	"runtime"
)

func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func Info(ctx *gin.Context) {
	system := make(map[string]interface{})
	system["os"] = runtime.GOOS
	system["goVersion"] = runtime.Version()
	system["cpu"] = runtime.NumCPU()
	v, _ := mem.VirtualMemory()
	memory := make(map[string]interface{})
	memory["count"] = v.Total
	memory["free"] = v.Free
	memory["usedPercent"] = fmt.Sprintf("%f%%", v.UsedPercent)
	ctx.JSON(http.StatusOK, gin.H{
		"app":    setting.Application.App,
		"system": system,
		"memory": memory,
	})
}
