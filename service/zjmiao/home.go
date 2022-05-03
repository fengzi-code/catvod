package zjmiao

import (
	"catvod/global"
	"catvod/model"
	"catvod/utils"
	"fmt"
	"github.com/antchfx/htmlquery"
)

type ZJMIAO struct {
	VodClass  []model.VodClass
	Filters   string
	Channel   string //就是vodClass里的TypeId，只包含当前的
	FilterMap model.FilterMap
}

const baseUrl = "https://zjmiao.com/"

func (this *ZJMIAO) GetHome() (res model.HomeContent) {

	this.Channel = "index.php/vod/show/id/1/"
	url := baseUrl + this.Channel
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		fmt.Println("qqtv get home html error: ", err)
		return
	}

	res.VodClass = make([]model.VodClass, 0)
	res.VodList = make([]model.VodInfo, 0)
	classJsonFile := global.ZJStaticDir + "/class.json"
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
	filterJsonFile := global.ZJStaticDir + "/filters.json"
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
	//fmt.Printf("%+v\n", this.FilterMap)
	res.Filters = this.FilterMap
	// 从页面获取VodInfo
	res.VodList = GetVodInfo(baseUrl)
	if len(res.VodList) > 0 {
		res.Code = 0
		res.Msg = "success"
	} else {
		res.Code = -1
		res.Msg = "error"
	}
	return

}
