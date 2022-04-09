package qqtv

import (
	"catvod/global"
	"catvod/model"
	"fmt"
	"github.com/antchfx/htmlquery"
	"strings"
)

type QQTV struct {
	VodClass []model.VodClass
	Filters  string
	Channel  string //就是vodClass里的TypeId，只包含当前的
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
	vodClassNodes := htmlquery.Find(doc, "//div[@class='site_channel']/a[contains(@class, 'channel_nav')]")
	for _, vodClassNode := range vodClassNodes {
		vodClassUrl := htmlquery.SelectAttr(vodClassNode, "href")
		vodClassName := htmlquery.InnerText(vodClassNode)
		if vodClassName == "VIP会员" {
			continue
		}
		urlArr := strings.Split(vodClassUrl, "/")
		vodClassUrl = urlArr[len(urlArr)-1]
		if vodClassUrl == "" {
			continue
		}
		// 如果this.VodClass不存在此分类，则添加到this.VodClass数组中
		this.VodClass = append(this.VodClass, model.VodClass{
			TypeName: vodClassName,
			TypeId:   vodClassUrl,
		})
	}
	fmt.Printf("%+v\n", this.VodClass)
	list := htmlquery.Find(doc, "//div[@class='list_item']")
	res.VodClass = make([]model.VodClass, 0)
	res.VodList = make([]model.VodInfo, 0)
	res.VodClass = this.VodClass
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
			filterValues = append(filterValues, model.FilterValueItems{
				ShowName: filterValueName,
				UrlValue: filterValueId,
			})
		}
		filts = append(filts, model.Filter{
			Name:  filterName,
			Key:   filterKey,
			Value: filterValues,
		})
	}
	filterMap := make(model.FilterMap)
	filterMap[this.Channel] = filts
	res.Filters = filterMap
	for _, item := range list {
		a := htmlquery.FindOne(item, "//a[@class='figure']")
		vodPic := htmlquery.InnerText(htmlquery.FindOne(item, "//img[@class='figure_pic']/@src"))
		vodId := htmlquery.SelectAttr(a, "data-float")
		vodName := htmlquery.SelectAttr(a, "title")
		vodRemarks := htmlquery.InnerText(htmlquery.FindOne(item, "//div[@class='figure_caption']"))
		if strings.HasPrefix(vodPic, "//") {
			vodPic = strings.Replace(vodPic, "//", "https://", 1)
		} else if strings.TrimSpace(vodPic) == "" {
			vodPic = global.DefaultPic
		}
		// fmt.Printf("vod_id: %s, vod_name: %s, vod_img: %s\n", vodId, vodName, vodPic)
		res.VodList = append(res.VodList, model.VodInfo{
			VodId:      vodId,
			VodName:    vodName,
			VodPic:     vodPic,
			VodRemarks: vodRemarks,
		})
	}
	if len(res.VodList) > 0 {
		res.Code = 0
		res.Msg = "success"
	} else {
		res.Code = -1
		res.Msg = "error"
	}
	return

}
