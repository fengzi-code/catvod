package v1

import (
	"catvod/global"
	"catvod/model"
	"catvod/service"
	"github.com/gin-gonic/gin"
)

func GetConfig(c *gin.Context) {
	m, ok1 := c.GetQuery("m") // 取appmode,f=family
	//k, ok2 := c.GetQuery("k") // 取key,没有key则跳过
	////http://xxxxxxx/config?m=f&k=key
	//if ok2 == false {
	//	c.JSON(
	//		200, gin.H{
	//			"message": "pong",
	//			"key":     k,
	//		},
	//	)
	//	return
	//}
	global.AppMode = m
	var data model.ServerConfig
	switch {
	case ok1:
		if global.AppMode == "f" {
			data = service.GetConfig("static/2.json")
		}
	default:
		data = service.GetConfig("static/1.json")
	}

	c.JSON(200, data)
}
