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
		if vodPicNode == nil {
			fmt.Println("vod pic node is nil")
			continue
		}
		detail.VodPic = htmlquery.SelectAttr(vodPicNode, "src")
		vodRemarksNode := htmlquery.FindOne(doc, "//div[@class='play-detail']/h1/span[@class='remarks']")
		if vodRemarksNode == nil {
			fmt.Println("vod remarks node is nil")
		} else {
			detail.VodRemarks = htmlquery.InnerText(vodRemarksNode)
		}

		detail.VodDirector = ""
		actorNodes := htmlquery.Find(doc, "//div[@class='actorlist_tit anthology-wrap']/a")
		for _, actorNode := range actorNodes {
			detail.VodActor += htmlquery.InnerText(actorNode) + ","
		}
		detail.VodActor = strings.TrimSuffix(detail.VodActor, ",")
		detail.VodArea = ""
		detail.VodContent = htmlquery.InnerText(htmlquery.FindOne(doc, "//div[@class='desc_txt cf ec-palytcji cor3']/span"))
		detail.VodYear = ""
		detail.VodPlayFrom = ""
		detail.VodPlayUrl = ""
		srcNodes := htmlquery.Find(doc, "//a[contains(@class, 'channelname')]")
		var srcs []string
		for _, srcNode := range srcNodes {
			src := htmlquery.InnerText(srcNode)
			src = strings.Trim(src, "&nbsp;")
			srcs = append(srcs, src)
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
