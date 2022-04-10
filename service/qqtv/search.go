package qqtv

import (
	"catvod/global"
	"catvod/model"
	"catvod/utils"
	"fmt"
	"github.com/antchfx/htmlquery"
	"net/url"
	"strings"
)

const searchUrl = "https://v.qq.com/x/search/?q=%s&stag=0&smartbox_ab="

func (this *QQTV) Search(wd string) (res []model.VodInfo) {
	// 将传入的wd转换为url编码
	surl := fmt.Sprintf(searchUrl, url.QueryEscape(wd))
	doc, err := htmlquery.LoadURL(surl)
	if err != nil {
		fmt.Println("qqtv get home html error: ", err)
		return
	}
	resultNodes := htmlquery.Find(doc, "//div[@class='_infos']/div/a[@class='figure result_figure']")
	for _, resultNode := range resultNodes {
		var (
			vodId      string
			vodName    string
			vodPic     string
			vodRemarks string
		)
		vodUrl := htmlquery.SelectAttr(resultNode, "href")
		if strings.Contains(vodUrl, "v.qq.com/x/cover") {
			vodId = utils.GetBetweenStr(vodUrl, "v.qq.com/x/cover/", ".html")
			vodPicNode := htmlquery.FindOne(resultNode, "//img[@class='figure_pic']")
			if vodPicNode != nil {
				vodPic = htmlquery.SelectAttr(vodPicNode, "src")
				vodName = htmlquery.SelectAttr(vodPicNode, "alt")
			}
			if strings.HasPrefix(vodPic, "//") {
				vodPic = strings.Replace(vodPic, "//", "https://", 1)
			} else if strings.TrimSpace(vodPic) == "" {
				vodPic = global.DefaultPic
			}
			fmt.Printf("vodId: %s, vodName: %s, vodPic: %s\n", vodId, vodName, vodPic)
			res = append(res, model.VodInfo{
				VodId:      vodId,
				VodName:    vodName,
				VodPic:     vodPic,
				VodRemarks: vodRemarks,
			})
		}
	}
	return
}
