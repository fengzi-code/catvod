package zjmiao

import (
	"catvod/model"
	"catvod/utils"
	"fmt"
	"github.com/antchfx/htmlquery"
	"net/url"
	"strings"
)

const searchUrl = `https://zjmiao.com/index.php/vod/search/?wd=`

func (this *ZJMIAO) Search(wd string) (res []model.VodInfo) {
	sUrl := searchUrl + url.QueryEscape(wd)
	fmt.Println(sUrl)

	doc, err := htmlquery.LoadURL(sUrl)
	if err != nil {
		fmt.Println("qqtv get home html error: ", err)
		return
	}
	var (
		vodId      string
		vodName    string
		vodPic     string
		vodRemarks string
	) //li[@class='list_item']
	nodes := htmlquery.Find(doc, "//div[@class='ec-search']//li[@class='search-list cf']")
	fmt.Println(htmlquery.OutputHTML(doc, true))
	for _, node := range nodes {
		a := htmlquery.FindOne(node, "//a[@class='aplus-exp ecimgbor']")
		vodPicNode := htmlquery.FindOne(a, "//div[@class='bj eclazy']")
		vodPic = htmlquery.SelectAttr(vodPicNode, "data-original")
		vodPic = strings.Replace(vodPic, `/img.php?url=`, "", -1)
		vodId = htmlquery.SelectAttr(a, "href")
		vodId = utils.GetBetweenStr(vodId, "id/", "/")
		vodName = htmlquery.SelectAttr(a, "title")
		vodRemarksNode := htmlquery.FindOne(a, "span[@class='pack-prb']")
		vodRemarks = htmlquery.InnerText(vodRemarksNode)

		res = append(res, model.VodInfo{
			VodId:      vodId,
			VodName:    vodName,
			VodPic:     vodPic,
			VodRemarks: vodRemarks,
		})

	}

	return
}
