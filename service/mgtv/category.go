package mgtv

import (
	"catvod/global"
	"catvod/model"
	"catvod/model/mgtv"
	"fmt"
	"github.com/go-resty/resty/v2"
	"math"
)

const mgtvCategoryApi = "https://pianku.api.mgtv.com/rider/list/pcweb/v3"

func (m *MGTV) GetCategory(typeid string, page int) (res model.Category) {
	headers := map[string]string{
		"User-Agent":   global.UserAgent,
		"Content-Type": global.ContentType,
	}
	client := resty.New()
	url := fmt.Sprintf("%s?platform=pcweb&channelId=%s&pn=%d", mgtvCategoryApi, typeid, page)
	get, err := client.R().
		SetResult(mgtv.Category{}). //model.Category结构体用来存储返回的json数据
		ForceContentType("application/json").
		SetHeaders(headers).
		// 从芒果api中取分类页中的数据
		Get(url)
	if err != nil {
		fmt.Println("get mgtv category error:", err)
	}
	c := get.Result().(*mgtv.Category)
	res.Total = c.Data.TotalHits
	res.Page = page
	res.Limit = 30
	res.PageCount = int(math.Ceil(float64(c.Data.TotalHits) / 10))
	res.VodList = make([]model.VodInfo, 0)
	for _, item := range c.Data.HitDocs {
		var vodRemarks string
		if item.UpdateInfo == "" {
			vodRemarks = item.RightCorner.Text
		} else {
			vodRemarks = item.RightCorner.Text + " [" + item.UpdateInfo + "]"
		}
		res.VodList = append(res.VodList, model.VodInfo{
			VodId:      item.PlayPartId,
			VodName:    item.Title,
			VodPic:     item.Img,
			VodRemarks: vodRemarks,
		})
	}
	return

}
