package iqytv

import (
	"catvod/global"
	"catvod/model"
	"catvod/model/iqy/response"
	"catvod/utils"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
)

var Filters model.Filters

func getiqiyiFilter(t, url string, comic, year bool) (res model.Filters) {
	client := resty.New()
	get, err := client.R().
		SetHeaders(global.Headers).
		Get(url)
	if err != nil {
		fmt.Println("ddddddd")
	}
	c := get.Body()
	b := string(c)
	d := utils.GetBetweenStr(b, "qy-list-category", "category-content")
	var iqiyiOrderlist []response.IqiyiOrderlist
	var iqiyiMyYear response.IqiyiMyYear
	var iqiyiFirstCategoryList []response.IqiyiFirstCategoryList
	var filtersTmp model.Filters
	Filters = filtersTmp
	e := utils.GetBetweenStr(d, "order-list='", "'") //取排序
	appenData([]byte(e), "排序", "mode", iqiyiOrderlist)
	if year {
		e = utils.GetBetweenStr(d, ":my-year='", "'") //取年份
		appenData([]byte(e), "年份", "market_release_date_level", iqiyiMyYear)
	}
	if comic {
		e = utils.GetBetweenStr(d, ":comic-status='", "'") //取状态
		appenData([]byte(e), "状态", "is_album_finished", iqiyiMyYear)
	}
	e = utils.GetBetweenStr(d, ":first-category-list='", "'") //取地区
	appenData([]byte(e), "", "", iqiyiFirstCategoryList)
	e = utils.GetBetweenStr(d, ":pay-status='", "'") //取资费
	appenData([]byte(e), "资费", "is_purchase", iqiyiMyYear)
	res = Filters
	return
}

func appenData(data []byte, filterName, filterKey string, v interface{}) {
	fmt.Println(filterKey, filterName)
	var filter model.Filter
	var filterTmp model.Filter
	var filterValueItems model.FilterValueItems
	var filterValueItemsTmp model.FilterValueItems
	var mustId int
	switch v.(type) {
	case []response.IqiyiOrderlist:

		tmp, _ := v.([]response.IqiyiOrderlist)
		json.Unmarshal(data, &tmp)
		for _, xxxx := range tmp {
			filterValueItems.UrlValue = strconv.Itoa(xxxx.Id)
			filterValueItems.ShowName = xxxx.Name
			filter.Value = append(filter.Value, filterValueItems)
		}
		filter.Name = filterName
		filter.Key = filterKey

	case response.IqiyiMyYear:
		tmp, _ := v.(response.IqiyiMyYear)
		json.Unmarshal(data, &tmp)
		for _, xxxx := range tmp.List {
			filterValueItems.UrlValue = xxxx.Id
			filterValueItems.ShowName = xxxx.Name
			filter.Value = append(filter.Value, filterValueItems)
		}
		filter.Name = filterName
		filter.Key = filterKey

	case []response.IqiyiFirstCategoryList:
		tmp, _ := v.([]response.IqiyiFirstCategoryList)
		json.Unmarshal(data, &tmp)
		fmt.Println(tmp)
		for _, xxxx := range tmp {
			fmt.Println(xxxx.Name, len(xxxx.Child))
			filter = filterTmp
			filterValueItems = filterValueItemsTmp
			ddd := xxxx.Name
			filterValueItems.UrlValue = ""
			filterValueItems.ShowName = "全部" + ddd
			filter.Value = append(filter.Value, filterValueItems)
			switch ddd {
			case "类型", "版本", "语种", "课程类型":
				filter.Name = ddd
				filter.Key = "must" + strconv.Itoa(mustId)
				mustId++
			case "地区", "风格", "新类型", "规格", "题材", "年龄段":
				filter.Name = ddd
				filter.Key = "three_category_id"
			case "状态":
				filter.Name = "状态"
				filter.Key = "is_album_finished"

			}
			for _, zz := range xxxx.Child {
				filterValueItems.UrlValue = strconv.Itoa(zz.Id)
				filterValueItems.ShowName = zz.Name
				filter.Value = append(filter.Value, filterValueItems)
			}
			Filters = append(Filters, filter)
		}
		return
	default:
		fmt.Println("unknown type")
	}
	Filters = append(Filters, filter)
}
