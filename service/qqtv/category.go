package qqtv

import (
	"catvod/global"
	"catvod/model"
	"fmt"
	"github.com/antchfx/htmlquery"
	"strconv"
	"strings"
)

/*
分类页
*/

func (this *QQTV) GetCategory(typeid string, page int) (res model.Category) {
	url := fmt.Sprintf(baseUrl+"/%s?listpage=%d&channel=%s&_all=1&pagesize=30&sort=19", typeid, page, typeid)
	fmt.Println("分页地址：", url)
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		fmt.Println("qqtv get home html error: ", err)
		return
	}
	list := htmlquery.Find(doc, "//div[@class='list_item']")
	res.Page = page
	res.Limit = 30
	totalText := htmlquery.FindOne(doc, "//div[@class='mod_list_filter']/div[@class='filter_result ']")
	var totalstr, pageCount string

	totalstr = htmlquery.SelectAttr(totalText, "data-total")
	pageCount = htmlquery.SelectAttr(totalText, "data-pagemax")
	total, _ := strconv.Atoi(totalstr)
	pageMax, _ := strconv.Atoi(pageCount)
	if total <= 0 {
		total = 5000
	}
	if pageMax <= 0 {
		pageMax = 167
	}
	res.Total = total
	res.PageCount = pageMax

	res.VodList = make([]model.VodInfo, 0)
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
