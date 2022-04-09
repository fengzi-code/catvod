package qqtv

import (
	"catvod/model"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strings"
)

/*
	GetVodClass 获取视频分类
@param doc *html.Node  // 使用htmlquery解析的html文档
获取视频分类后将其设置到当前结构体的vodClass字段中
*/
func (this *QQTV) GetVodClass(doc *html.Node) {
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
}
