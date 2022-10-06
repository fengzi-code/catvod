package rxzh

import (
	"catvod/global"
	"catvod/model"
	"catvod/utils"
)

const rxzhHomeApi = "http://103.222.188.156:99/rongxingapi.php/provide/vod/?ac=list"

type RXZH struct {
	VodClass  []model.VodClass
	Filters   string
	Channel   string //就是vodClass里的TypeId，只包含当前的
	FilterMap model.FilterMap
}

func (this *RXZH) GetHome() (res model.HomeContent) {
	//client := resty.New()
	//get, err := client.R().
	//	SetResult(rongxin.RxHome{}).
	//	ForceContentType(global.JsonType).
	//	SetHeaders(global.Headers).
	//	Get(rxzhHomeApi)
	//if err != nil {
	//	return
	//}
	//c := get.Result().(*rongxin.RxHome)

	res.VodClass = make([]model.VodClass, 0)
	classJsonFile := global.RxStaticDir + "/class.json"
	exist, _ := utils.PathExists(classJsonFile)
	if !exist {
		// TODO: 补充不存在时从网络上获取并写到本地的逻辑
		return
	}
	this.VodClass = utils.LoadClassJson(classJsonFile)
	res.VodClass = this.VodClass
	this.FilterMap = make(model.FilterMap)
	filterJsonFile := global.RxStaticDir + "/filters.json"
	exist, _ = utils.PathExists(filterJsonFile)
	if !exist {
		// TODO: 补充不存在时从网络上获取并写到本地的逻辑
		return
	}
	this.FilterMap = utils.LoadFilterJson(filterJsonFile)
	res.Filters = this.FilterMap
	category := this.GetCategory("1", 1)
	res.VodList = category.VodList
	return
}
