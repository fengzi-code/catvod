package mgtv

import (
	"catvod/global"
	"catvod/model"
	"catvod/model/mgtv/response"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/url"
	"strconv"
)

func (this *MGTV) Search(wd string) (res []model.VodInfo) {
	wd = url.QueryEscape(wd)
	headers := map[string]string{
		"User-Agent":   global.UserAgent,
		"Content-Type": global.ContentType,
	}
	client := resty.New()
	get, err := client.R().
		SetResult(response.SoStruct{}).
		ForceContentType("application/json").
		SetHeaders(headers).
		Get("https://mobileso.bz.mgtv.com/pc/search/v1?allowedRC=1&q=" + wd + "&pn=1&pc=10&uid=&_support=10000000&source=imgo")
	if err != nil {
		fmt.Println(err)
		return
	}
	c := get.Result().(*response.SoStruct)
	for _, v := range c.Data.Contents {
		vodId, _ := strconv.Atoi(v.Data.Vid)
		if vodId != 0 {
			res = append(res, model.VodInfo{
				VodId:      strconv.Itoa(vodId),
				VodName:    v.Data.Title,
				VodPic:     v.Data.Pic,
				VodRemarks: "芒果TV",
			})
		}
		for _, z := range v.Data.YearList {
			i, _ := strconv.Atoi(z.SourceList[0].Vid)
			if i != 0 {
				res = append(res, model.VodInfo{
					VodId:      strconv.Itoa(i),
					VodName:    z.Title,
					VodPic:     z.Pic,
					VodRemarks: z.SourceList[0].Name,
				})

			}
		}
		for _, x := range v.Data.SourceList {
			i, _ := strconv.Atoi(x.Vid)
			if i != 0 {
				res = append(res, model.VodInfo{
					VodId:      strconv.Itoa(i),
					VodName:    v.Data.Title,
					VodPic:     v.Data.Pic,
					VodRemarks: x.Name,
				})
			}
		}
	}
	return
}
