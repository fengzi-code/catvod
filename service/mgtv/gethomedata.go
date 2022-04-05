package mgtv

import (
	global "catvod/global/mgtv"
	model "catvod/model/mgtv"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func GetHomeData() ([]model.HomeVideo, []model.HomeClass, int, model.FilterIds) {
	contentType := "application/x-www-form-urlencoded;charset=UTF-8"
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"
	headers := map[string]string{
		"User-Agent":   userAgent,
		"Content-Type": contentType,
	}

	client := resty.New()

	get, err := client.R().
		SetResult(model.HomeList{}).
		ForceContentType("application/json").
		SetHeaders(headers).
		Get("https://rc.mgtv.com/pc/rank?allowedRC=0&t=day&c=2,3&limit=20&_support=10000000")

	if err != nil {
		fmt.Println("ddddddd")
	}

	c := get.Result().(*model.HomeList)

	var homeclasslist []model.HomeClass
	var homeclass model.HomeClass

	classlist := []string{
		"电视剧", "电影", "综艺", "少儿", "动漫", "纪录片", "游戏",
	}
	classListID := []string{"2", "3", "1", "10", "50", "51", "116"}
	for y, v := range classlist {
		homeclass.Typename = v
		homeclass.Typeid = classListID[y]
		homeclasslist = append(homeclasslist, homeclass)
	}

	var homevideos []model.HomeVideo
	var homevideo model.HomeVideo
	for _, x := range c.Data {
		homevideo.VodId = x.VideoId
		homevideo.VodName = x.Name
		homevideo.Image = x.Image
		homevideo.VodRemarks = x.Desc

		homevideos = append(homevideos, homevideo)

	}
	fmt.Println(homeclasslist, homevideos)

	var filterIds model.FilterIds

	err = json.Unmarshal([]byte(global.Id1), &filterIds.Id1)
	err = json.Unmarshal([]byte(global.Id2), &filterIds.Id2)
	err = json.Unmarshal([]byte(global.Id3), &filterIds.Id3)
	err = json.Unmarshal([]byte(global.Id10), &filterIds.Id10)
	err = json.Unmarshal([]byte(global.Id50), &filterIds.Id50)
	err = json.Unmarshal([]byte(global.Id51), &filterIds.Id51)
	err = json.Unmarshal([]byte(global.Id116), &filterIds.Id116)

	if err != nil {

		panic(err)
	}

	fmt.Println("-----------------", filterIds.Id116, filterIds, "----------------------")

	return homevideos, homeclasslist, c.Code, filterIds

}
