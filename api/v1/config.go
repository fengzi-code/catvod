package v1

import (
	"catvod/service"
	"github.com/gin-gonic/gin"
)

func GetConfig(c *gin.Context) {
	data := service.GetConfig("static/1.json")
	c.JSON(200, data)
}
