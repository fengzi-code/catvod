package route

import (
	v1 "catvod/api/v1"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	// Router.Use(cors.Cors())
	// Router.Use(gin.Logger())
	// Router.Use(gin.Recovery())
	Router.Static("/static", "./static")

}

func SetupRouter() *gin.Engine {
	Router.GET("/", v1.GetIndex)
	Router.GET("/config", v1.GetConfig) // 动态配置接口
	Router.GET("/api/v1/:tv", v1.CoreHandler)
	return Router
}
