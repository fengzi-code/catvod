package mgtv

import (
	"catvod/global"
	"catvod/model"
	"catvod/model/mgtv/response"
	"fmt"
	"github.com/go-resty/resty/v2"
	"math"
)

const cateGoryApi = "https://pianku.api.mgtv.com/rider/list/pcweb/v3?platform=pcweb&channelId=%s&pn=%d&%s"

func (this *MGTV) GetCategory(typeId string, page int) (res model.Category) {
	url := fmt.Sprintf(cateGoryApi, typeId, (page-1)*30, this.Filters)
	fmt.Println("分页地址：", url)
	client := resty.New()
	get, err := client.R().
		SetResult(response.Category{}). //model.Category结构体用来存储返回的json数据
		ForceContentType(global.JsonType).
		SetHeaders(global.Headers).
		// 从芒果api中取分类页中的数据
		Get(url)
	if err != nil {
		fmt.Println("获取分类页数据失败：", err)
		return
	}
	c := get.Result().(*response.Category)
	res.Total = c.Data.TotalHits
	res.Page = page
	res.Limit = 10
	res.PageCount = int(math.Ceil(float64(c.Data.TotalHits) / 10))
	res.VodList = make([]model.VodInfo, 0)
	for _, v := range c.Data.HitDocs {
		var vodRemarks string
		if v.UpdateInfo == "" {
			vodRemarks = v.RightCorner.Text
		} else {
			vodRemarks = v.RightCorner.Text + "[" + v.UpdateInfo + "]"
		}
		res.VodList = append(res.VodList, model.VodInfo{
			VodId:      v.PlayPartId,
			VodName:    v.Title,
			VodPic:     v.Img,
			VodRemarks: vodRemarks,
		})
	}
	return
}
