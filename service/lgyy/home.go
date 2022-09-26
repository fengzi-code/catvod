package lgyy

import (
	"catvod/global"
	"catvod/model"
	"catvod/utils"
	"encoding/json"
	"fmt"
	"github.com/antchfx/htmlquery"
	"io/ioutil"
)

const dyHomeApi = "https://www.lgyy.tv/label/new.html"

type LGYY struct {
	VodClass  []model.VodClass
	Filters   string
	Channel   string //就是vodClass里的TypeId，只包含当前的
	FilterMap model.FilterMap
}

func (this *LGYY) GetHome() (res model.HomeContent) {
	res.VodClass = make([]model.VodClass, 0)
	classJsonFile := global.LgyyStaticDir + "/class.json"
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
	var filterJsonFile string
	if global.AppMode == "f" {
		filterJsonFile = global.LgyyStaticDir + "/f_filters.json"
	} else {
		filterJsonFile = global.LgyyStaticDir + "/filters.json"
	}
	exist, err = utils.PathExists(filterJsonFile)
	FilterMap := make(model.FilterMap)
	if !exist {
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
	res.VodList = make([]model.VodInfo, 0)

	doc, err := htmlquery.LoadURL(dyHomeApi)
	if err != nil {
		fmt.Println("Lgyy get home html error: ", err)
		return
	}
	list := htmlquery.Find(
		doc, "//div[@class='module-main tab-list active']/div[@class='module-items module-poster-items']/a",
	)
	vodInfo := make([]model.VodInfo, 0)
	for _, item := range list {
		//a := htmlquery.FindOne(item, "//a[@class='figure']")
		vodName := htmlquery.SelectAttr(item, "title")
		vodId := htmlquery.SelectAttr(item, "href")
		vodId = utils.GetBetweenStr(vodId, "tail/", ".")
		img := htmlquery.FindOne(item, "//img[@class='lazy lazyload']")
		vodPic := htmlquery.SelectAttr(img, "data-original")
		remarks := htmlquery.FindOne(item, "//div [@class='module-item-note']")
		vodRemarks := htmlquery.InnerText(remarks)
		fmt.Printf("视频名字: %s, 视频id: %s, 视频图片: %s, 视频描述: %s, \n", vodName, vodId, vodPic, vodRemarks)
		vodInfo = append(
			vodInfo, model.VodInfo{
				VodId:      vodId,
				VodName:    vodName,
				VodPic:     vodPic,
				VodRemarks: vodRemarks,
			},
		)
	}
	res.VodList = vodInfo
	return
}
