package iqytv

import (
	"catvod/global"
	"catvod/model"
	"catvod/model/iqy/response"
	"encoding/base64"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
	"strings"
)

var categoryList model.VodDetail
var categoryTemp model.VodDetail

func (this *IQYTV) GetDetails(ids []string) (res []model.VodDetail) {
	//categoryList = categoryTemp // 每次查询详情时先清空旧数据
	for _, x := range ids {
		xx, _ := base64.StdEncoding.DecodeString(x)
		xxx := string(xx[:])
		fmt.Println("xxx================", xxx)
		xs := strings.Split(xxx, "|")
		id := xs[0]
		t := xs[1]
		vodName := xs[2]
		pic := xs[3]
		switch t {
		case "15", "4", "2": //电视剧,动漫,儿童
			fmt.Println("aid =============", id, t, vodName, pic)
			res = getId2(t, id, vodName, pic, x)
		case "1": //电影
			res = getId1(t, id, vodName, pic, x)
		case "6": //综艺
			res = getId6(t, id, vodName, pic, x)
		}

	}
	return

}
func getId2(t, aid, vodName, pic, id string) (res []model.VodDetail) {
	//https://pcw-api.iqiyi.com/albums/album/avlistinfo?aid=4378722550512701&page=1&size=30 电视剧/动漫/儿童/详情页
	fmt.Println(t)
	client := resty.New()
	get, err := client.R().
		SetResult(response.Avlistinfo{}).
		ForceContentType("application/json").
		SetHeaders(global.Headers).
		Get("https://pcw-api.iqiyi.com/albums/album/avlistinfo?aid=" + aid + "&page=1&size=100")

	if err != nil {
		fmt.Println("ddddddd")
	}
	c := get.Result().(*response.Avlistinfo)

	categoryList.VodId = id                                                //视频ID
	categoryList.VodName = vodName                                         //视频名字
	categoryList.VodRemarks = "更新至" + strconv.Itoa(len(c.Data.Epsodelist)) //更新状态
	categoryList.TypeName = ""                                             // 分类
	vodActor := ""
	for _, d := range c.Data.Epsodelist[0].People.MainCharactor {
		vodActor = vodActor + d.Name + " "
	}
	categoryList.VodActor = vodActor                           //主演
	categoryList.VodArea = ""                                  //地区
	categoryList.VodContent = c.Data.Epsodelist[0].Description //简介
	vodDirector := ""
	for _, d := range c.Data.Epsodelist[0].People.Director {
		vodDirector = vodDirector + d.Name + " "

	}
	categoryList.VodDirector = vodDirector             //导演
	categoryList.VodYear = c.Data.Epsodelist[0].Period // 年份
	categoryList.VodPic = pic                          //图片
	categoryList.VodPlayFrom = "qiyi"                  //播放源

	// 取视频集数和播放地址开始
	var vodPlayList string
	if c.Data.Page == 1 {
		for _, m := range c.Data.Epsodelist {
			vodPlayList = vodPlayList + strconv.Itoa(m.Order) + "$" + m.PlayUrl + "#"
		}
	} else {

		for i := 1; i < c.Data.Page+1; i++ {
			client := resty.New()
			get, err := client.R().
				SetResult(response.Avlistinfo{}).
				ForceContentType("application/json").
				SetHeaders(global.Headers).
				Get("https://pcw-api.iqiyi.com/albums/album/avlistinfo?aid=" + aid + "&page=" + strconv.Itoa(i) + "&size=100")

			if err != nil {
				fmt.Println("ddddddd")
			}
			c := get.Result().(*response.Avlistinfo)
			for _, m := range c.Data.Epsodelist {
				vodPlayList = vodPlayList + strconv.Itoa(m.Order) + "$" + m.PlayUrl + "#"
			}
		}

	}
	categoryList.VodPlayUrl = vodPlayList[:len(vodPlayList)-1]
	res = append(res, categoryList)
	return
	// 取视频集数和播放地址结束

}

func getId1(t, aid, vodName, pic, id string) (res []model.VodDetail) {
	//https://pcw-api.iqiyi.com/video/video/baseinfo/94453400  电影详情
	client := resty.New()
	get, err := client.R().
		SetResult(response.Videoinfo{}).
		ForceContentType("application/json").
		SetHeaders(global.Headers).
		Get("https://pcw-api.iqiyi.com/video/video/baseinfo/" + aid)

	if err != nil {
		fmt.Println("ddddddd")
	}
	c := get.Result().(*response.Videoinfo)

	categoryList.VodId = id                                            //视频ID
	categoryList.VodName = vodName                                     //视频名字
	categoryList.VodRemarks = "更新至" + strconv.Itoa(c.Data.LatestOrder) //更新状态
	categoryList.TypeName = ""                                         // 分类
	for _, b := range c.Data.Categories {
		categoryList.TypeName = categoryList.TypeName + " " + b.Name
	}
	vodDirector := "" //导演

	for _, b := range c.Data.People.Director {
		vodDirector = vodDirector + " " + b.Name
	}

	vodActor := "" //主演
	for _, b := range c.Data.People.MainCharactor {
		vodActor = vodActor + " " + b.Name
	}

	categoryList.VodActor = vodActor             //主演
	categoryList.VodArea = ""                    //地区
	categoryList.VodContent = c.Data.Description //简介

	categoryList.VodDirector = vodDirector //导演
	categoryList.VodYear = c.Data.Period   // 年份
	categoryList.VodPic = pic              //图片
	categoryList.VodPlayFrom = "qiyi"      //播放源

	// 取视频集数和播放地址开始

	categoryList.VodPlayUrl = strconv.Itoa(c.Data.Order) + "$" + c.Data.AlbumUrl
	res = append(res, categoryList)
	return
	// 取视频集数和播放地址结束

}

func getId6(t, aid, vodName, pic, id string) (res []model.VodDetail) {
	//https://pcw-api.iqiyi.com/album/album/baseinfo/238934101 综艺详情 先取tvid
	//https://pcw-api.iqiyi.com/album/source/listByNumber/238934101?include=true&number=100&seq=true&tvId=3922660400

	client := resty.New()
	get, err := client.R().
		SetResult(response.Tvidinfo{}).
		ForceContentType("application/json").
		SetHeaders(global.Headers).
		Get("https://pcw-api.iqiyi.com/album/album/baseinfo/" + aid)

	if err != nil {
		fmt.Println("ddddddd")
	}
	c := get.Result().(*response.Tvidinfo)

	tvid := strconv.FormatInt(c.Data.LatestVideo.TvId, 10)
	categoryList.VodId = id                                 //视频ID
	categoryList.VodName = vodName                          //视频名字
	categoryList.VodRemarks = c.Data.LatestVideo.ShortTitle //更新状态
	categoryList.TypeName = ""                              // 分类
	for _, b := range c.Data.Categories {
		categoryList.TypeName = categoryList.TypeName + " " + b.Name
	}
	vodDirector := "" //导演

	vodActor := "" //主演
	for _, b := range c.Data.People.Host {
		vodActor = vodActor + " " + b.Name
	}
	vodArea := ""
	for _, b := range c.Data.Areas {
		vodArea = vodArea + " " + b
	}
	categoryList.VodActor = vodActor             //主演
	categoryList.VodArea = vodArea               //地区
	categoryList.VodContent = c.Data.Description //简介

	categoryList.VodDirector = vodDirector //导演
	categoryList.VodYear = c.Data.Period   // 年份
	categoryList.VodPic = pic              //图片
	categoryList.VodPlayFrom = "qiyi"      //播放源
	// 取视频集数和播放地址开始
	client = resty.New()
	get, err = client.R().
		SetResult(response.ListByNumber{}).
		ForceContentType("application/json").
		SetHeaders(global.Headers).
		Get("https://pcw-api.iqiyi.com/album/source/listByNumber/" + aid + "?include=true&number=100&seq=true&tvId=" + tvid)

	if err != nil {
		fmt.Println("ListByNumber")
	}
	cc := get.Result().(*response.ListByNumber)
	vodurl := ""
	for _, b := range cc.Data {
		vodurl = b.ShortTitle + "$" + b.PlayUrl + "#" + vodurl
	}
	categoryList.VodPlayUrl = vodurl[:len(vodurl)-1]
	res = append(res, categoryList)
	return
}
