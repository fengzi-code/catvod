package mgtv

import (
	"catvod/global"
	"catvod/model/mgtv"
	"fmt"
	"github.com/go-resty/resty/v2"
	"math"
	"strconv"
)

func GetCategory(t, pg, filters string) (mgtv.RequestCategory, []mgtv.CategoryList) {
	fmt.Println(pg, t)
	headers := map[string]string{
		"User-Agent":   global.UserAgent,
		"Content-Type": global.ContentType,
	}
	// 筛选分析
	fmt.Println("+++++++ " + "https://pianku.api.mgtv.com/rider/list/pcweb/v3?platform=pcweb&channelId=" + t + "&pn=" + pg + filters + "+++++++++++++++")
	client := resty.New()
	get, err := client.R().
		SetResult(mgtv.Category{}). //model.Category结构体用来存储返回的json数据
		ForceContentType("application/json").
		SetHeaders(headers).
		// 从芒果api中取分类页中的数据
		Get("https://pianku.api.mgtv.com/rider/list/pcweb/v3?platform=pcweb&channelId=" + t + "&pn=" + pg + filters)

	if err != nil {
		fmt.Println("ddddddd")
	}

	c := get.Result().(*mgtv.Category)
	fmt.Println(c.Data)
	var requstcategory mgtv.RequestCategory
	requstcategory.Total = c.Data.TotalHits                                   //总共多少数据
	requstcategory.Page, _ = strconv.Atoi(pg)                                 //当前页
	requstcategory.Limit = 10                                                 //每页多少条
	requstcategory.PageCount = int(math.Ceil(float64(c.Data.TotalHits) / 10)) //共有多少页

	var categoryLists []mgtv.CategoryList
	var categoryList mgtv.CategoryList
	for _, x := range c.Data.HitDocs {
		categoryList.VodId = x.PlayPartId //视频ID
		categoryList.VodName = x.Title    //视频名
		categoryList.VodPic = x.Img       // 图片
		if x.UpdateInfo == "" {
			categoryList.VodRemarks = x.RightCorner.Text //更新集数
		} else {
			categoryList.VodRemarks = x.RightCorner.Text + " [" + x.UpdateInfo + "]" //更新集数
		}
		categoryLists = append(categoryLists, categoryList)

	}
	return requstcategory, categoryLists
}
