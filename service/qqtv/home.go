package qqtv

import (
	"catvod/global"
	"catvod/model"
	"fmt"
	"github.com/antchfx/htmlquery"
	"strings"
)

type QQTV struct {
}

const baseUrl = "https://v.qq.com/channel"

func (this *QQTV) GetHome() (res model.HomeContent) {
	url := baseUrl + "/tv?listpage=1&channel=tv&_all=1&sort=19"
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		fmt.Println("qqtv get home html error: ", err)
		return
	}
	list := htmlquery.Find(doc, "//div[@class='list_item']")
	res.VodClass = make([]model.VodClass, 0)
	res.VodList = make([]model.VodInfo, 0)
	res.VodClass = append(res.VodClass, model.VodClass{
		TypeName: "电视剧",
		TypeId:   "tv",
	})
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
