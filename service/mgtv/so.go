package mgtv

import (
	model "catvod/model/mgtv"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
)

func So(wd string) []model.HomeVideo {
	//https://mobileso.bz.mgtv.com/pc/search/v1?allowedRC=1&q=%E7%A5%9E%E6%8E%A2%E5%A4%8F%E6%B4%9B%E5%85%8B%E5%85%A8%E5%AD%A3&pn=1&pc=10&uid=&_support=10000000

	fmt.Println(wd)
	contentType := "application/x-www-form-urlencoded"
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36"
	headers := map[string]string{
		"User-Agent":   userAgent,
		"Content-Type": contentType,
		"origin":       "https://so.mgtv.com",
		"referer":      "https://so.mgtv.com",
	}
	client := resty.New()

	get, err := client.R().
		SetResult(model.SoStruct{}).
		ForceContentType("application/json").
		SetHeaders(headers).
		Get("https://mobileso.bz.mgtv.com/pc/search/v1?allowedRC=1&q=" + wd + "&pn=1&pc=10&uid=&_support=10000000")
	if err != nil {
		fmt.Println("ddddddd")
	}
	fmt.Println("https://mobileso.bz.mgtv.com/pc/search/v1?allowedRC=1&q=" + wd + "&pn=1&pc=10&uid=&_support=10000000")
	//var solist model.SoStruct
	var sovideos []model.HomeVideo
	var sovideo model.HomeVideo

	c := get.Result().(*model.SoStruct)
	fmt.Println(c)

	for x, y := range c.Data.Contents[0].Data.YearList {

		sovideo.Image = y.Pic
		sovideo.VodName = y.Title
		sovideo.VodRemarks = y.SourceList[x].VideoList[0].RightCorner.Text
		sovideo.VodId, _ = strconv.Atoi(y.SourceList[x].VideoList[0].Vid)
		sovideos = append(sovideos, sovideo)

	}
	return sovideos
}
