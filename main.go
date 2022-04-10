package main

import (
	"catvod/route"
	"flag"
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
	// 从flag接收端口号
	var port string
	flag.StringVar(&port, "port", "8080", "端口号")
	flag.Parse()
	// 3.启动服务
	r.Run(":" + port)
}
