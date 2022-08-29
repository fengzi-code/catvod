package dy555

import (
	"catvod/model"
	"catvod/utils"
	"fmt"
	"github.com/antchfx/htmlquery"
	"strconv"
)

const dy555CateGoryApi = "https://www.5dy6.cc/vodshow/%s--------%d---.html"

func (this *DY555) GetCategory(typeId string, page int) (res model.Category) {
	url := fmt.Sprintf(dy555CateGoryApi, typeId, page)
	fmt.Println(url)
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		fmt.Println("dy555 get home html error: ", err)
		return
	}
	pageCountNode := htmlquery.Find(doc, "//div[@class='module-main module-page ']/div[@id='page']/a")
	pageCount := htmlquery.SelectAttr(pageCountNode[len(pageCountNode)-1], "href")
	fmt.Println(pageCount)
	pageCount = utils.GetBetweenStr(pageCount, "--------", "-")
	res.Page = page
	res.Limit = 12
	res.PageCount, _ = strconv.Atoi(pageCount)
	res.Total = 12 * res.PageCount
	list := htmlquery.Find(
		doc, "//div[@class='module-items module-poster-items-base ']/a",
	)
	//fmt.Println("dy555 list: ", htmlquery.OutputHTML(list[0], true))
	vodInfo := make([]model.VodInfo, 0)
	for _, item := range list {
		//a := htmlquery.FindOne(item, "//a[@class='figure']")
		vodName := htmlquery.SelectAttr(item, "title")
		vodId := htmlquery.SelectAttr(item, "href")
		vodId = utils.GetBetweenStr(vodId, "tail/", ".")
		img := htmlquery.FindOne(item, "//img[@class='lazy lazyload']")
		vodPic := htmlquery.SelectAttr(img, "data-original")
		remarks := htmlquery.FindOne(item, "//div [@class='module-item-note']")
		vodRemarks := htmlquery.InnerText(remarks)
		fmt.Println(vodName, vodId, vodPic, vodRemarks)

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
