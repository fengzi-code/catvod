package zjmiao

import (
	"catvod/model"
	"fmt"
	"github.com/antchfx/htmlquery"
	"strings"
)

const detailUrlPrefix = "https://zjmiao.com/index.php/vod/detail/id/%s/"

func (this *ZJMIAO) GetDetails(ids []string) (res []model.VodDetail) {
	for _, id := range ids {
		url := fmt.Sprintf(detailUrlPrefix, id)
		fmt.Printf("detail url: %s\n", url)
		doc, err := htmlquery.LoadURL(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		var detail model.VodDetail
		detail.VodId = id
		vodNameNode := htmlquery.FindOne(doc, "//div[@class='play-detail']/h1/span[@class='name']")
		if vodNameNode == nil {
			fmt.Println("vod name node is nil")
			continue
		}
		detail.VodName = htmlquery.SelectAttr(vodNameNode, "title")
		vodPicNode := htmlquery.FindOne(doc, "//div[@class='s-cover box']/a/img")
		if vodPicNode != nil {
			detail.VodPic = htmlquery.SelectAttr(vodPicNode, "src")
		}

		vodRemarksNode := htmlquery.FindOne(doc, "//div[@class='play-detail']/h1/span[@class='remarks']")
		if vodRemarksNode != nil {
			detail.VodRemarks = htmlquery.InnerText(vodRemarksNode)
		}
		directorNode := htmlquery.FindOne(doc, "//div[contains(@class,'cf fyy')]/p[@class='item-desc js-open-wrap hidden'][2]/a")
		if directorNode != nil {
			detail.VodDirector = htmlquery.InnerText(directorNode)
		}

		actorNodes := htmlquery.Find(doc, "//div[@class='actorlist_tit anthology-wrap']/a")
		for _, actorNode := range actorNodes {
			detail.VodActor += htmlquery.InnerText(actorNode) + ","
		}
		detail.VodActor = strings.TrimSuffix(detail.VodActor, ",")
		areaNode := htmlquery.FindOne(doc, "//div[contains(@class,'cf fyy')]/p[@class='item']/a")
		if areaNode != nil {
			detail.VodArea = htmlquery.InnerText(areaNode)
		}
		detail.VodContent = htmlquery.InnerText(htmlquery.FindOne(doc, "//div[@class='desc_txt cf ec-palytcji cor3']/span"))
		yearNode := htmlquery.FindOne(doc, "//div[contains(@class,'cf fyy')]/p[@class='item item-actor'][2]/a")
		if yearNode != nil {
			detail.VodYear = htmlquery.InnerText(yearNode)
		}
		srcNodes := htmlquery.Find(doc, "//a[contains(@class, 'channelname')]")
		for _, srcNode := range srcNodes {
			src := htmlquery.InnerText(srcNode)
			src = strings.Trim(src, "&nbsp;")
			detail.VodPlayFrom += src + "$$$"
		}
		detail.VodPlayFrom = strings.TrimSuffix(detail.VodPlayFrom, "$$$")
		ulNodes := htmlquery.Find(doc, "//ul[@class='content_playlist cf']")
		for _, ulNode := range ulNodes {
			liNodes := htmlquery.Find(ulNode, "//li/a")
			var playUrls string
			for _, liNode := range liNodes {
				playUrl := "https://zjmiao.com" + htmlquery.SelectAttr(liNode, "href")
				ep := htmlquery.InnerText(liNode)
				playUrls += ep + "$" + playUrl + "#"
			}
			playUrls = strings.TrimSuffix(playUrls, "#")
			detail.VodPlayUrl += playUrls + "$$$"
		}
		detail.VodPlayUrl = strings.TrimSuffix(detail.VodPlayUrl, "$$$")
		res = append(res, detail)
	}
	return
}
