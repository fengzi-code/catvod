package v1

import (
	"catvod/global"
	"catvod/model"
	"catvod/service"
	"catvod/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func GetConfig(c *gin.Context) {
	m, ok1 := c.GetQuery("m") // 取appmode,f=family
	k, ok2 := c.GetQuery("k") // 取key,没有key则跳过
	//http://xxxxxxx/config?m=f&k=key
	if ok2 {
		var appkey model.AppKey
		tmp := utils.LoadJson("static/appkey.json", appkey)
		//使用mapstructure.Decode()方法
		err := mapstructure.Decode(tmp, &appkey)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(k, appkey.Name, appkey.Key)
		if k == appkey.Name {
			global.IsRongxin = true
			global.RongXinKey = appkey.Key
		}
	}
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
