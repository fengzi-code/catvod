package qqtv

import (
	"catvod/model"
	"fmt"
	"github.com/antchfx/htmlquery"
)

/*
这个函数的作用是通过访问不同类型影视的首页，获取filterMap
*/
func (this *QQTV) GetFilterMap(t string) {
	url := fmt.Sprintf("%s/%s?listpage=1&_all=1&channel=%s", baseUrl, t, t)
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		fmt.Println("qqtv get home html error: ", err)
		return
	}
	// 取filter节点
	filterNodes := htmlquery.Find(doc, "//div[@class='mod_list_filter']/div[contains(@class, 'filter_line')]")
	var filts model.Filters
	for _, filterNode := range filterNodes {
		// 在这里定义filterValues变量，是为了每次都是新的变量，避免把之前的值也加入进来了
		var filterValues []model.FilterValueItems
		filterName := htmlquery.InnerText(htmlquery.FindOne(filterNode, "//span[@class='filter_label']"))
		filterKey := htmlquery.SelectAttr(filterNode, "data-key")
		filterValueNodes := htmlquery.Find(filterNode, "//a[contains(@class, 'filter_item')]")
		for _, filterValueNode := range filterValueNodes {
			filterValueName := htmlquery.InnerText(filterValueNode)
			filterValueId := htmlquery.SelectAttr(filterValueNode, "data-value")
			filterValues = append(
				filterValues, model.FilterValueItems{
					ShowName: filterValueName,
					UrlValue: filterValueId,
				},
			)
		}
		filts = append(
			filts, model.Filter{
				Name:  filterName,
				Key:   filterKey,
				Value: filterValues,
			},
		)
	}

	this.FilterMap[t] = filts
}
