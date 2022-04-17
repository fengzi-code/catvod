package iqytv

import (
	"catvod/global"
	"catvod/model"
	"catvod/model/iqy/response"
	"encoding/base64"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/url"
	"strconv"
	"strings"
)

func (this *IQYTV) Search(wd string) (res []model.VodInfo) {
	wd = url.QueryEscape(wd)
	client := resty.New()

	get, err := client.R().
		SetResult(response.IqiyiSoJson{}). //model.Category结构体用来存储返回的json数据
		ForceContentType("application/json").
		SetHeaders(global.Headers).
		// 从芒果api中取分类页中的数据
		Get(global.IqiyiSobaseUrl + wd)
	if err != nil {
		fmt.Println("ddddddd")
	}
	c := get.Result().(*response.IqiyiSoJson)

	classListID := []string{"2", "1", "6", "4", "15"}
	if len(c.Data.FormatData.IntentList) != 0 {
		for _, x := range c.Data.FormatData.IntentList {

			b2 := x.Id
			b := strconv.FormatInt(b2, 10)
			cc := strconv.Itoa(x.Cid)

			for _, bb := range classListID { //判断是否在类型中
				if bb == cc {
					d := base64.StdEncoding.EncodeToString([]byte(b + "|" + cc + "|" + x.GTitle + "|" + x.GImg))
					res = append(res, model.VodInfo{
						VodId:      d,
						VodName:    x.GTitle,
						VodPic:     x.GImg,
						VodRemarks: x.GMeta,
					})
				}
			}

		}
	} else {
		for _, x := range c.Data.FormatData.List {
			b2 := x.Id
			b := strconv.FormatInt(b2, 10)
			cc := strconv.Itoa(x.Cid)
			for _, bb := range classListID { //判断是否在类型中
				if bb == cc {
					nInt := strings.Index(x.GTitle, "精彩")
					nInt2 := strings.Index(x.GTitle, "瞬间")
					if x.Uploader.Name != "" || nInt != -1 || nInt2 != -1 {
						break
					}
					d := base64.StdEncoding.EncodeToString([]byte(b + "|" + cc + "|" + x.GTitle + "|" + x.GImg))
					res = append(res, model.VodInfo{
						VodId:      d,
						VodName:    x.GTitle,
						VodPic:     x.GImg,
						VodRemarks: x.GMeta,
					})
				}
			}
		}
	}
	return
}
