package dy555

import (
	"catvod/model"
	"catvod/utils"
	"fmt"
	"github.com/antchfx/htmlquery"
	url2 "net/url"
	"regexp"
	"strconv"
	"strings"
)

//https://www.5dy6.cc/vodshow/类型-地区-排序-剧情-语言----页数---年份.html
const dy555CateGoryApi = "https://www.5dy6.cc/vodshow/%s-%s-%s-%s-%s----%d---%s.html"

func (this *DY555) GetCategory(typeId string, page int) (res model.Category) {
	lang := utils.GetBetweenStr(this.Filters, `lang=`, `&`)
	area := utils.GetBetweenStr(this.Filters, `area=`, `&`)
	item := utils.GetBetweenStr(this.Filters, `item=`, `&`)
	year := utils.GetBetweenStr(this.Filters, `year=`, `&`)
	sort := utils.GetBetweenStr(this.Filters, `sort=`, `&`)
	class := utils.GetBetweenStr(this.Filters, `class=`, `&`)

	var url string
	if class != "" {
		url = fmt.Sprintf(
			dy555CateGoryApi, class, url2.QueryEscape(area), sort, url2.QueryEscape(item), url2.QueryEscape(lang), page,
			year,
		)
	} else {
		url = fmt.Sprintf(
			dy555CateGoryApi, typeId, url2.QueryEscape(area), sort, url2.QueryEscape(item), url2.QueryEscape(lang),
			page, year,
		)
	}
	fmt.Printf("请求的分类地址: %s \n", url)
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		fmt.Println("dy555 get home html error: ", err)
		return
	}
	pageCountNode := htmlquery.Find(doc, "//div[@class='module-main module-page ']/div[@id='page']/a")
	var pageCount string
	if len(pageCountNode) == 0 {
		pageCount = "1"
	} else {
		pageCount = htmlquery.SelectAttr(pageCountNode[len(pageCountNode)-1], "href")
		pageCount = utils.GetBetweenStr(pageCount, "---", ".html")
		re := regexp.MustCompile(`\d+-`)
		pageCount = re.FindString(pageCount)
		pageCount = strings.Replace(pageCount, `-`, ``, -1)

	}
	res.Page = page
	res.Limit = 12
	res.PageCount, _ = strconv.Atoi(pageCount)
	res.Total = 12 * res.PageCount
	list := htmlquery.Find(
		doc, "//div[@class='module-items module-poster-items-base ']/a",
	)
	vodInfo := make([]model.VodInfo, 0)
	for _, item := range list {
		vodName := htmlquery.SelectAttr(item, "title")
		vodId := htmlquery.SelectAttr(item, "href")
		vodId = utils.GetBetweenStr(vodId, "tail/", ".")
		img := htmlquery.FindOne(item, "//img[@class='lazy lazyload']")
		vodPic := htmlquery.SelectAttr(img, "data-original")
		remarks := htmlquery.FindOne(item, "//div [@class='module-item-note']")
		vodRemarks := htmlquery.InnerText(remarks)
		fmt.Printf("视频名字: %s, 视频id: %s, 视频图片: %s, 视频评论: %s, \n", vodName, vodId, vodPic, vodRemarks)
		vodInfo = append(
			vodInfo, model.VodInfo{
				VodId:      vodId,
				VodName:    vodName,
				VodPic:     vodPic,
				VodRemarks: vodRemarks,
			},
		)
	}
	res.VodList = vodInfo

	return
}
