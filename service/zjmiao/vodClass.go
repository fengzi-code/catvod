package zjmiao

import (
	"catvod/model"
	"catvod/utils"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

/*
	GetVodClass 获取视频分类
@param doc *html.Node  // 使用htmlquery解析的html文档
获取视频分类后将其设置到当前结构体的vodClass字段中
*/

func (this *ZJMIAO) GetVodClass(doc *html.Node) {
	vodClassNodes := htmlquery.Find(doc, "//div[@class='show-filter back br top20']/div[@class='filter_line'][1]/a[contains(@class, 'filter_item')]")
	for _, vodClassNode := range vodClassNodes {
		vodClassUrl := htmlquery.SelectAttr(vodClassNode, "href")
		vodClassName := htmlquery.InnerText(vodClassNode)
		if vodClassName == "直播" {
			break
		}
		vodClassUrl = utils.GetBetweenStr(vodClassUrl, "id/", "/")
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
