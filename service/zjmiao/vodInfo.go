package zjmiao

import (
	"catvod/model"
	"fmt"
	"github.com/antchfx/htmlquery"
	"strings"
)

/*
	GetVodInfo 获取视频信息
	@param doc *html.Node // 通过htmlquery解析出来的html文档
	@return []model.VodInfo	// 这个是从html文档中解析出来的视频信息数组
*/
func GetVodInfo(url string) (vodInfo []model.VodInfo) {
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		return
	}
	vodInfo = make([]model.VodInfo, 0)
	list := htmlquery.Find(doc, "//div[contains(@class,'public wow')]//div[contains(@class,'pack-ykpack hide')]")

	for _, item := range list {
		a := htmlquery.FindOne(item, "//a[@class='aplus-exp ecimgbor']")
		vodPicNode := htmlquery.FindOne(a, "div[@class='bj eclazy']")
		vodPic := htmlquery.SelectAttr(vodPicNode, "data-original")
		vodId := htmlquery.SelectAttr(a, "href")
		vodName := htmlquery.SelectAttr(a, "title")
		vodPic = strings.Replace(vodPic, `/img.php?url=`, "", -1)
		fmt.Println(vodPic, vodId, vodName)
		// 在纪录片时运行到半路会报错，因为没有这个figure_caption节点
		// 所以在取InnerText时需要提前判断一下*html.Node是否为nil
		vodRemarksNode := htmlquery.FindOne(a, "span[@class='pack-prb hidden']")
		vodRemarks := htmlquery.InnerText(vodRemarksNode)

		fmt.Printf("vod_id: %s, vod_name: %s, vod_img: %s\n", vodId, vodName, vodPic)
		vodInfo = append(vodInfo, model.VodInfo{
			VodId:      vodId,
			VodName:    vodName,
			VodPic:     vodPic,
			VodRemarks: vodRemarks,
		})
	}
	return
}
