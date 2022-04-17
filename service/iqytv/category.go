package iqytv

import (
	"catvod/global"
	"catvod/model"
	"catvod/model/iqy/response"
	"encoding/base64"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
)

func (this *IQYTV) GetCategory(typeId string, page int) (res model.Category) {
	//url := fmt.Sprintf(global.IqiyiCatBaseUrl, typeId, page, this.Filters)
	url := global.IqiyiCatBaseUrl + typeId + "&page_id=" + strconv.Itoa(page) + this.Filters
	fmt.Println(url)
	client := resty.New()
	get, err := client.R().
		SetResult(response.IqiyiCategory{}). //model.Category结构体用来存储返回的json数据
		ForceContentType("application/json").
		SetHeaders(global.Headers).
		// 从芒果api中取分类页中的数据
		Get(url)

	if err != nil {
		fmt.Println("ddddddd")
	}

	c := get.Result().(*response.IqiyiCategory)

	res.Total = 16000    //总共多少数据
	res.Page = page      //当前页
	res.Limit = 10       //每页多少条
	res.PageCount = 1600 //共有多少页
	//res.VodList = make([]model.VodInfo, 0)
	var vodList model.VodInfo
	for _, x := range c.Data.List {
		b2 := int(x.AlbumId)
		c := strconv.Itoa(x.ChannelId)
		// base64编码
		d := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(b2) + "|" + c + "|" + x.Name + "|" + x.ImageUrl))
		vodList.VodId = d
		vodList.VodName = x.Title   //视频名
		vodList.VodPic = x.ImageUrl // 图片
		if typeId == "6" {
			vodList.VodRemarks = x.Period
		} else if typeId == "1" {
			vodList.VodRemarks = x.Period
		} else {
			vodList.VodRemarks = "更新至 " + strconv.Itoa(x.LatestOrder) + "集"
		}
		res.VodList = append(res.VodList, vodList)
	}
	return
}
