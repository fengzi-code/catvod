package qqtv

import (
	"catvod/model"
	"fmt"
	"github.com/antchfx/htmlquery"
	"strconv"
)

/*
分类页
*/
const cateGoryApi = "https://v.qq.com/x/bu/pagesheet/list?append=1&channel=%s&listpage=1&offset=%d&pagesize=30&%s"

func (this *QQTV) GetCategory(typeid string, page int) (res model.Category) {
	fmt.Println("GetCategory", this.Filters)
	url := fmt.Sprintf(cateGoryApi, typeid, (page-1)*30, this.Filters)
	fmt.Println("分页地址：", url)
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		fmt.Println("qqtv get home html error: ", err)
		return
	}

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
	res.VodList = GetVodInfo(doc)
	if len(res.VodList) > 0 {
		res.Code = 0
		res.Msg = "success"
	} else {
		res.Code = -1
		res.Msg = "error"
	}
	return
}
