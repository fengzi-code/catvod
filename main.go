package main

import (
	"catvod/route"
	"github.com/gin-gonic/gin"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	// 1.设置为调试模式，默认就是这个模式
	gin.SetMode(gin.DebugMode)
	// 2.设置路由
	r := route.SetupRouter()
	// 3.启动服务
	r.Run(":8080")
}
