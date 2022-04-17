package qqtv

import (
	"catvod/global"
	"catvod/model"
	"catvod/model/qqtv/response"
	"catvod/utils"
	"encoding/json"
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
	var (
		vodId      string
		vodName    string
		vodPic     string
		vodRemarks string
	)
	// 取第一层结果
	resultNodes := htmlquery.Find(doc, "//div[@class='_infos']/div/a[@class='figure result_figure']")
	for _, resultNode := range resultNodes {
		vodUrl := htmlquery.SelectAttr(resultNode, "href")
		if strings.Contains(vodUrl, "v.qq.com/x/") && !strings.Contains(vodUrl, "redirect") {
			vodIdTmpStr := strings.Split(vodUrl, "/")[len(strings.Split(vodUrl, "/"))-1]
			vodId = strings.Split(vodIdTmpStr, ".html")[0]
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
			vodName = strings.Replace(vodName, "\u0005", "", -1)
			vodName = strings.Replace(vodName, "\u0006", "", -1)
			// fmt.Printf("vodId: %s, vodName: %s, vodPic: %s\n", vodId, vodName, vodPic)
			res = append(res, model.VodInfo{
				VodId:      vodId,
				VodName:    vodName,
				VodPic:     vodPic,
				VodRemarks: vodRemarks,
			})
		}
	}
	fmt.Printf("第一层结果总数: %d, 详情: %+v\n", len(res), res)
	// 取第二层结果，取出result_series_new节点，这个节点的r-props属性中包含了所有搜索结果
	resultN := htmlquery.FindOne(doc, "//div[@class='result_series_new']")
	if resultN == nil {
		return
	}
	rprops := htmlquery.SelectAttr(resultN, "r-props")
	if rprops == "" {
		return
	}
	// 从r-props中取出totalData
	rprops = utils.GetBetweenStr(rprops, "totalData: '", "';")
	// 将rprops进行url解码
	rprops, err = url.QueryUnescape(rprops)
	if err != nil {
		fmt.Println("url unescape rprops error: ", err)
		return
	}
	// 将rprops进行json解析
	var searchResult response.SearchTotalData
	err = json.Unmarshal([]byte(rprops), &searchResult)
	if err != nil {
		fmt.Println("qqtv search result json unmarshal error: ", err)
		return
	}

	for _, item := range searchResult.ItemList {

		vodId = item.Doc.Id
		vodName = item.VideoInfo.Title
		byName := item.VideoInfo.CoverDoc.ByName // byName是一个文本数组，存放影视别名
		skip := true
		// 检查byName里是否存在搜索关键词
		for _, v := range byName {
			if strings.Contains(v, wd) {
				skip = false
				break
			}
		}
		// 如果标题中和byName中都不存在关键词，则跳过
		if !strings.Contains(vodName, wd) && skip {
			continue
		}
		// 如果url中带redirect，则跳过
		if strings.Contains(item.VideoInfo.Url, "search_redirect") {
			continue
		}
		vodName = strings.Replace(vodName, "\u0005", "", -1)
		vodName = strings.Replace(vodName, "\u0006", "", -1)
		// 去掉空格
		if item.VideoInfo.Year != "" {
			vodName = vodName + "(" + item.VideoInfo.Year + ")"
		}
		if len(item.VideoInfo.Language) != 0 {
			vodName = fmt.Sprintf("%s-(%s)", vodName, item.VideoInfo.Language)
		}
		vodPic = item.VideoInfo.ImgUrl

		// TODO: vodRemarks原本是打算从item.VideoInfo.ImgTag中的Tag3.Text取，但是有的结果没有
		vodRemarks = "qqtv"
		res = append(res, model.VodInfo{
			VodId:      vodId,
			VodName:    vodName,
			VodPic:     vodPic,
			VodRemarks: vodRemarks,
		})

	}
	fmt.Printf("两层结果总数: %d, 详情: %+v\n", len(res), res)
	return
}
