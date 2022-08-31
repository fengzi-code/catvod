package youku

import (
	"catvod/global"
	"catvod/model"
	"catvod/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type YOUKU struct {
	VodClass  []model.VodClass
	Filters   string
	Channel   string //就是vodClass里的TypeId，只包含当前的
	FilterMap model.FilterMap
}

func (this *YOUKU) GetHome() (res model.HomeContent) {
	res.VodClass = make([]model.VodClass, 0)
	classJsonFile := global.YOUKUStaticDir + "/class.json"
	exist, err := utils.PathExists(classJsonFile)
	if err != nil {
		return
	}
	if !exist {
		// TODO: 补充不存在时从网络上获取并写到本地的逻辑
		return
	}
	this.VodClass = utils.LoadClassJson(classJsonFile)
	res.VodClass = this.VodClass
	this.FilterMap = make(model.FilterMap)
	filterJsonFile := global.YOUKUStaticDir + "/filters.json"
	exist, err = utils.PathExists(filterJsonFile)
	FilterMap := make(model.FilterMap)
	if !exist {
		// TODO: 补充不存在时从网络上获取并写到本地的逻辑
		fmt.Println("补充")
		for _, t := range res.VodClass {
			filters := GetFilterMap(t.TypeId)
			FilterMap[t.TypeId] = filters
		}
		marshal, err := json.Marshal(FilterMap)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(filterJsonFile, marshal, 0644)
		if err != nil {
			panic(err)
		}
	}
	this.FilterMap = utils.LoadFilterJson(filterJsonFile)
	res.Filters = this.FilterMap

	return
}
