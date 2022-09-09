package main

import (
	"catvod/global"
	"catvod/route"
	"catvod/utils"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	// 1.设置gin模式，默认debug模式
	// 从flag接收端口号
	var (
		port string
		mode string
		ip   string
	)

	flag.StringVar(&port, "port", "80", "端口号")
	flag.StringVar(&mode, "mode", "debug", "gin模式，可选debug,release, 默认debug")
	flag.Parse()
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
		ip = utils.GetExternalIp()
	} else {
		ip = utils.GetInternalIp()
	}
	fmt.Printf("当前模式：%s\n", gin.Mode())
	fmt.Printf("当前接口对外IP：%s\n", ip)
	// 2.设置路由
	r := route.SetupRouter()
	if ip != "" {
		global.BindAddr = fmt.Sprintf("%s:%s", ip, port)
	}
	// 3.启动服务
	r.Run(":" + port)
}
