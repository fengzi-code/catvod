package youku

import (
	"catvod/global"
	"catvod/model"
	"catvod/model/youku/response"
	"catvod/utils"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/url"
)

const s = `{"scene":"search_component_paging","id":227939,"category":"%s"}`
const param = `{"type":"%s"}`
const cateGoryApi = "https://www.youku.com/category/data?session=%s&params=%s&pageNo=%d"

func (this *YOUKU) GetCategory(typeId string, page int) (res model.Category) {
	vodClass := utils.LoadClassJson(global.YOUKUStaticDir + "/class.json")
	var typeName string
	for _, class := range vodClass {
		if class.TypeId == typeId {
			typeName = class.TypeName
		}
	}
	session := fmt.Sprintf(s, typeName)
	params := fmt.Sprintf(param, typeName)
	cateUrl := fmt.Sprintf(cateGoryApi, url.QueryEscape(session), url.QueryEscape(params), page)
	fmt.Println("分页地址：", cateUrl)
	headers := global.Headers
	headers["referer"] = "https://www.youku.com/category/show/" + typeId + ".html?theme=dark"
	client := resty.New()
	get, err := client.R().
		SetResult(response.YoukuCategory{}). //model.Category结构体用来存储返回的json数据
		ForceContentType(global.JsonType).
		SetHeaders(headers).
		// 从芒果api中取分类页中的数据
		Get(cateUrl)
	if err != nil {
		fmt.Println("获取分类页数据失败：", err)
		return
	}
	c := get.Result().(*response.YoukuCategory)

	res.Total = 84000
	res.Page = page
	res.Limit = 20
	res.PageCount = 420
	vodInfo := make([]model.VodInfo, 0)

	for _, v := range c.Data.FilterData.ListData {
		fmt.Printf(
			"视频名字: %s, 视频id: %s, 视频图片: %s, 视频描述: %s, \n", v.Title, utils.GetBetweenStr(v.VideoLink, "id_", ".html"),
			v.Img, v.Summary,
		)
		vodInfo = append(
			vodInfo, model.VodInfo{
				VodId:      utils.GetBetweenStr(v.VideoLink, "id_", ".html"),
				VodName:    v.Title,
				VodPic:     v.Img,
				VodRemarks: v.Summary,
			},
		)
	}
	res.VodList = vodInfo
	return
}
