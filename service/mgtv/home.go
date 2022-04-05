package mgtv

import (
	"catvod/global"
	"catvod/model"
	"catvod/model/mgtv"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
)

type MGTV struct {
}

func (m *MGTV) GetHome() (res model.HomeContent) {
	headers := map[string]string{
		"User-Agent":   global.UserAgent,
		"Content-Type": global.ContentType,
	}
	client := resty.New()
	get, err := client.R().
		SetResult(mgtv.HomeList{}).
		ForceContentType("application/json").
		SetHeaders(headers).
		Get("https://rc.mgtv.com/pc/rank?allowedRC=0&t=day&c=2,3&limit=20&_support=10000000")
	if err != nil {
		fmt.Println("ddddddd")
	}
	c := get.Result().(*mgtv.HomeList)

	res.VodClass = []model.VodClass{
		{
			TypeName: "电视剧",
			TypeId:   "2",
		},
		{
			TypeName: "电影",
			TypeId:   "3",
		},
		{
			TypeName: "综艺",
			TypeId:   "1",
		},
		{
			TypeName: "少儿",
			TypeId:   "10",
		},
		{
			TypeName: "动漫",
			TypeId:   "50",
		},
		{
			TypeName: "纪录片",
			TypeId:   "51",
		},
		{
			TypeName: "游戏",
			TypeId:   "116",
		},
	}
	res.VodList = make([]model.VodInfo, 0)
	for _, item := range c.Data {
		res.VodList = append(res.VodList, model.VodInfo{
			VodId:      strconv.Itoa(item.VideoId),
			VodName:    item.Name,
			VodPic:     item.Image,
			VodRemarks: item.Desc,
		})
	}
	res.Code = c.Code
	if len(res.VodList) == 0 {
		res.Msg = "暂无数据"
	} else {
		res.Msg = "获取成功"
	}
	return
}
