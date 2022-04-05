package mgtv

import (
	"catvod/model/mgtv"
	"fmt"
	"github.com/go-resty/resty/v2"
)

/*
本模块用于获取芒果TV视频详情
*/

func Getdetail(ids string) mgtv.RequestDetail {
	contentType := "application/x-www-form-urlencoded;charset=UTF-8"
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"
	headers := map[string]string{
		"User-Agent":   userAgent,
		"Content-Type": contentType,
	}

	client := resty.New()

	get, err := client.R().
		SetResult(mgtv.Detail{}).
		ForceContentType("application/json").
		SetHeaders(headers).
		Get("https://pcweb.api.mgtv.com/video/info?vid=" + ids)

	if err != nil {
		fmt.Println("ddddddd")
	}

	c := get.Result().(*mgtv.Detail)
	fmt.Println(c.Data)

	var categoryList mgtv.RequestDetail
	categoryList.VodId = c.Data.Info.VideoId
	categoryList.VodName = c.Data.Info.ClipName
	categoryList.VodRemarks = c.Data.Info.Detail.UpdateInfo
	categoryList.TypeName = c.Data.Info.FstlvlType
	categoryList.VodActor = c.Data.Info.Detail.Leader
	categoryList.VodArea = c.Data.Info.Detail.Area
	categoryList.VodContent = c.Data.Info.Detail.Story
	categoryList.VodDirector = c.Data.Info.Detail.Director
	categoryList.VodYear = c.Data.Info.Detail.ReleaseTime
	categoryList.VodPic = c.Data.Info.VideoImage
	categoryList.VodPlayFrom = "mgtv"

	get, err = client.R().
		SetResult(mgtv.RequestPlayUrl{}).
		ForceContentType("application/json").
		SetHeaders(headers).
		Get("https://pcweb.api.mgtv.com/episode/list?&video_id=" + ids)

	if err != nil {
		fmt.Println("ddddddd")
	}
	var playurl string
	mgtvurl := "https://www.mgtv.com"
	b := get.Result().(*mgtv.RequestPlayUrl)
	fmt.Println(b.Data)
	for _, x := range b.Data.List {
		fmt.Println(x)
		playurl = playurl + x.T4 + "$" + mgtvurl + x.Url + "#"
		//playurl = playurl + x.T4+"$"+x.VideoId+"#"
	}
	fmt.Println(playurl)
	categoryList.VodPlayUrl = playurl[:len(playurl)-1]
	return categoryList
}
