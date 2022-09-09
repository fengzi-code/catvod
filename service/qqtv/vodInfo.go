package qqtv

import (
	"catvod/global"
	"catvod/model"
	"fmt"
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
	vodInfo = make([]model.VodInfo, 0)
	list := htmlquery.Find(doc, "//div[@class='list_item']")
	for _, item := range list {
		a := htmlquery.FindOne(item, "//a[@class='figure']")
		vodPic := htmlquery.InnerText(htmlquery.FindOne(item, "//img[@class='figure_pic']/@src"))
		vodId := htmlquery.SelectAttr(a, "data-float")
		vodName := htmlquery.SelectAttr(a, "title")
		// 在纪录片时运行到半路会报错，因为没有这个figure_caption节点
		// 所以在取InnerText时需要提前判断一下*html.Node是否为nil
		remarkNode := htmlquery.FindOne(item, "//div[@class='figure_caption']")
		var vodRemarks string
		if remarkNode != nil { // 这里需要判断一下，否则没有这个元素会报错
			vodRemarks = htmlquery.InnerText(remarkNode)
		} else {
			vodRemarks = ""
		}
		if strings.HasPrefix(vodPic, "//") {
			vodPic = strings.Replace(vodPic, "//", "https://", 1)
		} else if strings.TrimSpace(vodPic) == "" {
			vodPic = global.DefaultPic
		}
		fmt.Printf("视频名字: %s, 视频id: %s, 视频图片: %s, 视频描述: %s, \n", vodName, vodId, vodPic, vodRemarks)
		vodInfo = append(
			vodInfo, model.VodInfo{
				VodId:      vodId,
				VodName:    vodName,
				VodPic:     vodPic,
				VodRemarks: vodRemarks,
			},
		)
	}
	return
}
