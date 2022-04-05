package qqtv

import (
	"catvod/model"
	"fmt"
	"github.com/antchfx/htmlquery"
)

const detailUrlPrefix = "https://v.qq.com/x/cover"

func (this *QQTV) GetDetails(ids []string) (res []model.VodDetail) {
	for _, id := range ids {
		url := fmt.Sprintf("%s/%s.html", detailUrlPrefix, id)
		doc, err := htmlquery.LoadURL(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		var detail model.VodDetail
		detail.VodId = id
		detail.VodName = htmlquery.InnerText(htmlquery.FindOne(doc, "//div[@class='mod_intro']/h1"))
		res = append(res, detail)
	}
	return
}
