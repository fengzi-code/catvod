package mgtv

import (
	"catvod/model/mgtv"
	"fmt"
	"github.com/go-resty/resty/v2"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var categoryList mgtv.RequestDetail
var categoryLists, categoryTemp mgtv.RequestDetails

func Getdetails(ids []string) mgtv.RequestDetails {
	contentType := "application/x-www-form-urlencoded;charset=UTF-8"
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"
	headers := map[string]string{
		"User-Agent":   userAgent,
		"Content-Type": contentType,
	}
	// 清空结构体
	categoryLists = categoryTemp // 每次查询详情时先清空旧数据
	var playurl string

	mgtvurl := "https://www.mgtv.com"

	// 循环读取所有的ids详情
	for _, x := range ids {
		client := resty.New()
		get, err := client.R().
			SetResult(mgtv.Detail{}). //返回的json存到model.Detail结构体
			ForceContentType("application/json").
			SetHeaders(headers).
			Get("https://pcweb.api.mgtv.com/video/info?vid=" + x)
		if err != nil {
			fmt.Println("ddddddd")
		}
		c := get.Result().(*mgtv.Detail)

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
		// 由于每次只能返回30个数据,所有要判断视频集数再循环多次取数据
		get, err = client.R().
			SetResult(mgtv.RequestPlayUrl{}).
			ForceContentType("application/json").
			SetHeaders(headers).
			Get("https://pcweb.api.mgtv.com/episode/list?&video_id=" + x + "&page=0&size=30")
		if err != nil {
			fmt.Println("ddddddd")
		}
		// 取页数
		b := get.Result().(*mgtv.RequestPlayUrl)
		pageTotal := b.Data.Count
		pageTotal = int(math.Ceil(float64(pageTotal) / 30))
		if pageTotal == 0 {
			pageTotal = pageTotal + 1
		}
		// 取页数结束

		// 循环页数取数据
		var playurlList []string

		// 由于每次只能返回30个数据,所有要判断视频集数再循环多次取数据
		for i := 1; i < pageTotal+1; i++ {
			y := strconv.Itoa(i)
			fmt.Println("+++++++++++", y)
			fmt.Println("https://pcweb.api.mgtv.com/episode/list?&video_id=" + x + "&page=" + y + "&size=30")
			get, err = client.R().
				SetResult(mgtv.RequestPlayUrl{}).
				ForceContentType("application/json").
				SetHeaders(headers).
				Get("https://pcweb.api.mgtv.com/episode/list?&video_id=" + x + "&page=" + y + "&size=30")
			if err != nil {
				fmt.Println("ddddddd")
			}
			too := get.Result().(*mgtv.RequestPlayUrl)

			for _, x := range too.Data.List {
				// 清除预告片
				hzRegexp, _ := regexp.Compile("([\u4e00-\u9fa5]+)")

				if hzRegexp.MatchString(x.T1) { //如果T1有中文则用T1
					if x.Corner[0].Font != "预" {
						playurlList = append(playurlList, x.T1+"$"+mgtvurl+x.Url+"#")
					} else {
						playurlList = append(playurlList, "[预告]"+x.T1+"$"+mgtvurl+x.Url+"#")
					}
				} else if hzRegexp.MatchString(x.T4) {
					if x.Corner[0].Font != "预" {
						playurlList = append(playurlList, x.T4+"$"+mgtvurl+x.Url+"#")
					} else {
						playurlList = append(playurlList, "[预告]"+x.T4+"$"+mgtvurl+x.Url+"#")
					}
				} else {
					// 取视频时间
					a := strings.Split(x.T4, ":")
					c, _ := strconv.Atoi(a[0])
					b := 45
					// 大于45取t1标题，否则取t2标题
					if c > b {
						if x.Corner[0].Font != "预" {
							playurlList = append(playurlList, x.T1+"$"+mgtvurl+x.Url+"#")
						} else {
							playurlList = append(playurlList, "[预告]"+x.T1+"$"+mgtvurl+x.Url+"#")
						}
					} else {
						playurlList = append(playurlList, x.T2+"$"+mgtvurl+x.Url+"#")
					}
				}

			}

		}
		// 由于每次只能返回30个数据,所有要判断视频集数再循环多次取数据结束

		// 官网返回的集数排序很乱,需要重新对集数进行集数进行排序
		var test [][]int
		defaultIndex := 99999 //预告和其它的排到最后面
		for x, y := range playurlList {
			b2 := regexp.MustCompile("^第(\\S+)集")
			q := b2.FindString(y)
			b2 = regexp.MustCompile("\\d+")
			q = b2.FindString(q) //正则取出带第集中间的数字
			var tmp []int
			if q == "" {
				defaultIndex++                     //如果不是第1集样式的序号99999+1
				tmp = append(tmp, defaultIndex, x) //[10000,0]
			} else {
				z, _ := strconv.Atoi(q) //带第1集样式的取出中间的数字
				fmt.Println(z)
				tmp = append(tmp, z, x) //  [1,0]
			}
			test = append(test, tmp) //将切片添加到二维数组
		}
		fmt.Println(test, "_____________________________")
		sort.Slice(
			test, func(i, j int) bool {
				return test[i][0] < test[j][0] // 二维数据按第一列进行排序
			},
		)

		playurl = ""
		for _, x := range test {
			playurl = playurl + playurlList[x[1]] // 装所有数据组合成playurl需要的文本
		}
		fmt.Println(test, "_____________________________++")
		// 官网返回的集数排序很乱,需要重新对集数进行集数进行排序结束

		categoryList.VodPlayUrl = playurl[:len(playurl)-1] //去除playurl最后的#号
		categoryLists = append(categoryLists, categoryList)
		fmt.Println(pageTotal, b.Data.Total)
	}
	return categoryLists

}
