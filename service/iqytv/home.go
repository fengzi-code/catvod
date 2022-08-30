package iqytv

import (
	"catvod/global"
	"catvod/model"
	"catvod/model/iqy/response"
	"catvod/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"strconv"
)

type IQYTV struct {
	VodClass  []model.VodClass
	Filters   string
	Channel   string //就是vodClass里的TypeId，只包含当前的
	FilterMap model.FilterMap
}

const iqyBaseApi = "https://pcw-api.iqiyi.com/search/recommend/list?channel_id=2&data_type=1&mode=11&page_id=1&ret_num=20"

func (this *IQYTV) GetHome() (res model.HomeContent) {
	client := resty.New()
	get, err := client.R().
		SetResult(response.IqiyiCategory{}).
		ForceContentType(global.JsonType).
		SetHeaders(global.Headers).
		Get(iqyBaseApi)
	if err != nil {
		fmt.Println(err)
		return
	}
	c := get.Result().(*response.IqiyiCategory)
	res.VodClass = make([]model.VodClass, 0)
	classJsonFile := global.IQYStaticDir + "/class.json"
	exist, err := utils.PathExists(classJsonFile)
	if !exist {
		// TODO: 补充不存在时从网络上获取并写到本地的逻辑
		return
	}
	this.VodClass = utils.LoadClassJson(classJsonFile)
	res.VodClass = this.VodClass
	this.FilterMap = make(model.FilterMap)

	filterJsonFile := global.IQYStaticDir + "/filters.json"
	exist, err = utils.PathExists(filterJsonFile)
	if !exist {
		//Iqiyiglobal.FilterMap = make(model.FilterMap)
		var year bool
		var comic bool
		FilterMap := make(model.FilterMap)
		for _, t := range res.VodClass {
			if t.TypeId == "2" || t.TypeId == "1" {
				year = true
			}
			if t.TypeId == "4" {
				comic = true
			}
			filters := GetFilterMap(
				t.TypeId, global.IqiyiFilterBaseurl+t.TypeId+global.IqiyiFilterbaseEndUrl, comic, year,
			)
			year = false
			comic = false
			FilterMap[t.TypeId] = filters
		}

		marshal, err := json.Marshal(FilterMap)
		if err != nil {
			fmt.Println("ddddddddd")
		}
		err = ioutil.WriteFile(filterJsonFile, marshal, 0644)
		if err != nil {
			panic(err)
		}
	}
	this.FilterMap = utils.LoadFilterJson(filterJsonFile)
	res.Filters = this.FilterMap
	res.VodList = make([]model.VodInfo, 0)
	for _, v := range c.Data.List {
		b2 := v.AlbumId
		b := strconv.FormatInt(b2, 10)
		cc := strconv.Itoa(v.ChannelId)
		d := base64.StdEncoding.EncodeToString([]byte(b + "|" + cc + "|" + v.Name + "|" + v.ImageUrl))
		res.VodList = append(
			res.VodList, model.VodInfo{
				VodId:      d,
				VodName:    v.Name,
				VodPic:     v.ImageUrl,
				VodRemarks: fmt.Sprintf("更新至%d集", v.LatestOrder),
			},
		)
	}
	return
}
