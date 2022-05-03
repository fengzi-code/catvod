package zjmiao

import (
	"catvod/model"
	"catvod/utils"
	"fmt"
	"github.com/antchfx/htmlquery"
)

/*
这个函数的作用是通过访问不同类型影视的首页，获取filterMap
*/
func (this *ZJMIAO) GetFilterMap(t string) {
	url := fmt.Sprintf("%sindex.php/vod/show/id/%s", baseUrl, t)
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		fmt.Println("qqtv get home html error: ", err)
		return
	}
	// 取filter节点
	filterNodes := htmlquery.Find(doc, "//div[@class='filter_line']")
	//fmt.Printf(htmlquery.OutputHTML(filterNodes[1],true))
	var filts model.Filters
	for y, filterNode := range filterNodes {
		if y == 0 {
			continue
		}
		// 在这里定义filterValues变量，是为了每次都是新的变量，避免把之前的值也加入进来了filter_item zbs
		var filterValues []model.FilterValueItems
		filterNameNode := htmlquery.FindOne(filterNode, "/span/text()")
		filterName := htmlquery.InnerText(filterNameNode)
		filterKeyNode := htmlquery.Find(filterNode, "/a")
		filterKey := htmlquery.SelectAttr(filterKeyNode[1], "href")
		filterKey = utils.GetBetweenStr(filterKey, "vod/show/", "/")
		if filterKey == "id" {
			filterKey = "year"
		}
		filterValueNodes := htmlquery.Find(filterNode, "//a/text()")
		for _, filterValueNode := range filterValueNodes {
			filterValueName := htmlquery.InnerText(filterValueNode)
			filterValueId := htmlquery.InnerText(filterValueNode)
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
