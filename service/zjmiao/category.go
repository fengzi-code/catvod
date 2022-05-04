package zjmiao

import (
	"catvod/model"
	"catvod/utils"
	"fmt"
	"github.com/antchfx/htmlquery"
	"strconv"
	"strings"
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
	vodNodes := htmlquery.Find(doc, "//div[@class='pack-packcover returl list-top-b']/a")
	fmt.Printf("节点数量：%d\n", len(vodNodes))
	for _, vodNode := range vodNodes {
		vodName := htmlquery.SelectAttr(vodNode, "title")
		vodUrl := htmlquery.SelectAttr(vodNode, "href")
		if vodUrl == "" {
			continue
		}
		vodId := utils.GetBetweenStr(vodUrl, "id/", "/")
		vodImgNode := htmlquery.FindOne(vodNode, "//div[@class='bj eclazy']")
		if vodImgNode == nil {
			continue
		}
		vodImg := htmlquery.SelectAttr(vodImgNode, "data-original")
		// 处理有些图片不显示问题，原因是/img.php?url=的这种
		if strings.HasPrefix(vodImg, "/img.php?url=") {
			vodImg = strings.TrimPrefix(vodImg, "/img.php?url=")
		}
		if !strings.HasPrefix(vodImg, "http") {
			fmt.Printf("可能异常的图片链接:%s\n", vodImg)
		}

		vodRemarkNode := htmlquery.FindOne(vodNode, "//span[@class='pack-prb hidden']")
		var vodRemark string
		if vodRemarkNode == nil {
			vodRemark = "追剧喵"
		} else {
			vodRemark = htmlquery.InnerText(vodRemarkNode)
		}
		vodInfo := model.VodInfo{
			VodName:    vodName,
			VodId:      vodId,
			VodPic:     vodImg,
			VodRemarks: vodRemark,
		}
		res.VodList = append(res.VodList, vodInfo)
	}

	return
}
