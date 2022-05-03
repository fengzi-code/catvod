package zjmiao

import (
	"catvod/model"
	"catvod/utils"
	"fmt"
	"github.com/antchfx/htmlquery"
	"strconv"
)

const categoryUrl = "https://zjmiao.com/index.php/vod/show/id/"

func (this *ZJMIAO) GetCategory(typeId string, page int) (res model.Category) {
	url := fmt.Sprintf("%s%s/page/%d/%s", categoryUrl, typeId, page, this.Filters)
	fmt.Println("分页地址：", url)
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		fmt.Println("qqtv get home html error: ", err)
		return
	}
	res.Page = page
	res.Limit = 48
	// <div class="page_tip cor3">共15052条数据,当前1/314页</div>
	var totalstr, pageCount string
	totalstr = htmlquery.InnerText(htmlquery.FindOne(doc, "//div[@class='page_tip cor3']"))
	total, _ := strconv.Atoi(utils.GetBetweenStr(totalstr, "共", "条"))
	pageCount = utils.GetBetweenStr(totalstr, "/", "页")
	pageMax, _ := strconv.Atoi(pageCount)
	if total <= 0 {
		total = 5000
	}
	if pageMax <= 0 {
		pageMax = 167
	}
	res.Total = total
	res.PageCount = pageMax
	res.VodList = make([]model.VodInfo, 0)
	vodNodes := htmlquery.Find(doc, "//div[@class='pack-ykpack']/div[contains(@class,'pack-packcover')]/a")
	for _, vodNode := range vodNodes {
		vodName := htmlquery.SelectAttr(vodNode, "title")
		vodUrl := htmlquery.SelectAttr(vodNode, "href")
		if vodUrl == "" {
			continue
		}
		vodUrl = "https://zjmiao.com" + vodUrl
		vodImgNode := htmlquery.FindOne(vodNode, "//div[@class='bj eclazy']")
		if vodImgNode == nil {
			continue
		}
		vodImg := htmlquery.SelectAttr(vodImgNode, "data-original")
		vodRemarkNode := htmlquery.FindOne(vodNode, "//span[@class='pack-prb hidden']")
		var vodRemark string
		if vodRemarkNode == nil {
			vodRemark = "追剧喵"
		} else {
			vodRemark = htmlquery.InnerText(vodRemarkNode)
		}
		vodInfo := model.VodInfo{
			VodName:    vodName,
			VodId:      vodUrl,
			VodPic:     vodImg,
			VodRemarks: vodRemark,
		}
		res.VodList = append(res.VodList, vodInfo)
	}

	return
}
