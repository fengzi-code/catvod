package qqtv

import (
	"catvod/global"
	"catvod/model"
	"catvod/utils"
	"fmt"
	"github.com/antchfx/htmlquery"
)

type QQTV struct {
	VodClass  []model.VodClass
	Filters   string
	Channel   string //就是vodClass里的TypeId，只包含当前的
	FilterMap model.FilterMap
}

const baseUrl = "https://v.qq.com/channel"

/*
https://v.qq.com/channel/tv?listpage=1&_all=1&channel=tv&sort=19
上面那个接口地址不太靠谱，换成下面这个会好点
https://v.qq.com/x/bu/pagesheet/list?_all=1&channel=tv&listpage=1&offset=30&pagesize=30&sort=19
改动offset值即可
*/

func (this *QQTV) GetHome() (res model.HomeContent) {
	this.Channel = "tv"
	url := baseUrl + "/tv?listpage=1&channel=tv&_all=1&sort=19"
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		fmt.Println("qqtv get home html error: ", err)
		return
	}

	res.VodClass = make([]model.VodClass, 0)
	res.VodList = make([]model.VodInfo, 0)
	classJsonFile := global.QQStaticDir + "/class.json"

	exist, err := utils.PathExists(classJsonFile)
	if !exist {
		this.GetVodClass(doc)
		err = utils.SaveClassJson(classJsonFile, this.VodClass)
	}
	if err != nil {
		return
	}
	this.VodClass = utils.LoadClassJson(classJsonFile)
	res.VodClass = this.VodClass
	this.FilterMap = make(model.FilterMap)
	// 如果filter文件存在就读文件，不存在则创建并写入
	filterJsonFile := global.QQStaticDir + "/filters.json"
	exist, err = utils.PathExists(filterJsonFile)
	if !exist {
		for _, t := range this.VodClass {
			this.GetFilterMap(t.TypeId)
		}
		err = utils.SaveFilterJson(filterJsonFile, this.FilterMap)
	}
	if err != nil {
		return
	}
	this.FilterMap = utils.LoadFilterJson(filterJsonFile)
	fmt.Printf("%+v\n", this.FilterMap)
	res.Filters = this.FilterMap
	// 从页面获取VodInfo
	res.VodList = GetVodInfo(doc)
	if len(res.VodList) > 0 {
		res.Code = 0
		res.Msg = "success"
	} else {
		res.Code = -1
		res.Msg = "error"
	}
	return

}
