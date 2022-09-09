package mgtv

import (
	"catvod/global"
	"catvod/model"
	"catvod/model/mgtv/response"
	"catvod/utils"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
)

type MGTV struct {
	VodClass  []model.VodClass
	Filters   string
	Channel   string //就是vodClass里的TypeId，只包含当前的
	FilterMap model.FilterMap
}

func (this *MGTV) GetHome() (res model.HomeContent) {
	this.Channel = "2"
	client := resty.New()
	get, err := client.R().
		SetResult(response.Homelist{}).
		ForceContentType(global.JsonType).
		SetHeaders(global.Headers).
		Get(global.MgHomeApi)
	if err != nil {
		return
	}
	c := get.Result().(*response.Homelist)
	res.VodClass = make([]model.VodClass, 0)
	classJsonFile := global.MgStaticDir + "/class.json"
	exist, err := utils.PathExists(classJsonFile)
	if !exist {
		// TODO: 补充不存在时从网络上获取并写到本地的逻辑
		return
	}
	this.VodClass = utils.LoadClassJson(classJsonFile)
	res.VodClass = this.VodClass
	this.FilterMap = make(model.FilterMap)
	filterJsonFile := global.MgStaticDir + "/filters.json"
	exist, err = utils.PathExists(filterJsonFile)
	if !exist {
		// TODO: 补充不存在时从网络上获取并写到本地的逻辑
		return
	}
	this.FilterMap = utils.LoadFilterJson(filterJsonFile)
	res.Filters = this.FilterMap
	res.VodList = make([]model.VodInfo, 0)

	for _, v := range c.Data {
		fmt.Printf("视频名字: %s, 视频id: %d, 视频图片: %s, 视频描述: %s, \n", v.Name, v.VideoId, v.Image, v.Desc)
		res.VodList = append(
			res.VodList, model.VodInfo{
				VodId:      strconv.Itoa(v.VideoId),
				VodName:    v.Name,
				VodPic:     v.Image,
				VodRemarks: v.Desc,
			},
		)
	}
	return
}
