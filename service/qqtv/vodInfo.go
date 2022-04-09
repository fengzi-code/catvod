package qqtv

import (
	"catvod/global"
	"catvod/model"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strings"
)

/*
	GetVodInfo 获取视频信息
	@param doc *html.Node // 通过htmlquery解析出来的html文档
	@return []model.VodInfo	// 这个是从html文档中解析出来的视频信息数组
*/
func GetVodInfo(doc *html.Node) (vodInfo []model.VodInfo) {
	list := htmlquery.Find(doc, "//div[@class='list_item']")
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
		vodInfo = append(vodInfo, model.VodInfo{
			VodId:      vodId,
			VodName:    vodName,
			VodPic:     vodPic,
			VodRemarks: vodRemarks,
		})
	}
	return
}
