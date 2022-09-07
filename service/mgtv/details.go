package mgtv

import (
	"catvod/global"
	"catvod/model"
	"catvod/model/mgtv/response"
	"fmt"
	"github.com/go-resty/resty/v2"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	detailApi   = "https://pcweb.api.mgtv.com/video/info?vid=%s"
	episodesApi = "https://pcweb.api.mgtv.com/episode/list?&video_id=%s&page=%d&page_size=%d"
)

func (this *MGTV) GetDetails(ids []string) (res []model.VodDetail) {
	mgtvurl := "https://www.mgtv.com"
	// 循环读取所有的ids详情
	for _, id := range ids {
		client := resty.New()
		get, err := client.R().
			SetResult(response.Detail{}). //返回的json存到model.Detail结构体
			ForceContentType(global.JsonType).
			SetHeaders(global.Headers).
			Get(fmt.Sprintf(detailApi, id))
		if err != nil {
			fmt.Println(err)
			continue
		}
		c := get.Result().(*response.Detail)
		var detail model.VodDetail
		detail.VodId = c.Data.Info.VideoId
		detail.VodName = c.Data.Info.ClipName
		detail.VodRemarks = c.Data.Info.Detail.UpdateInfo
		detail.TypeName = c.Data.Info.FstlvlType
		detail.VodActor = c.Data.Info.Detail.Leader
		detail.VodArea = c.Data.Info.Detail.Area
		detail.VodContent = c.Data.Info.Detail.Story
		detail.VodDirector = c.Data.Info.Detail.Director
		detail.VodYear = c.Data.Info.Detail.ReleaseTime
		detail.VodPic = c.Data.Info.VideoImage
		detail.VodPlayFrom = "mgtv"
		get, err = client.R().
			SetResult(response.PlayUrl{}).
			ForceContentType(global.JsonType).
			SetHeaders(global.Headers).
			Get(fmt.Sprintf(episodesApi, id, 0, 30))
		if err != nil {
			fmt.Println(err)
			continue
		}
		// 取页数
		b := get.Result().(*response.PlayUrl)
		pageTotal := b.Data.Count
		pageTotal = int(math.Ceil(float64(pageTotal) / 30))
		if pageTotal == 0 {
			pageTotal = pageTotal + 1
		}
		// 循环页数取数据
		var playurlList []string
		// 由于每次只能返回30个数据,所有要判断视频集数再循环多次取数据
		for i := 1; i < pageTotal+1; i++ {
			get, err = client.R().
				SetResult(response.PlayUrl{}).
				ForceContentType(global.JsonType).
				SetHeaders(global.Headers).
				Get(fmt.Sprintf(episodesApi, id, i, 30))
			if err != nil {
				continue
			}
			too := get.Result().(*response.PlayUrl)

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
			sort.Slice(
				test, func(i, j int) bool {
					return test[i][0] < test[j][0] // 二维数据按第一列进行排序
				},
			)

			for _, x := range test {
				detail.VodPlayUrl += playurlList[x[1]] // 装所有数据组合成playurl需要的文本
			}
			detail.VodPlayUrl = detail.VodPlayUrl[:len(detail.VodPlayUrl)-1]
		}
		res = append(res, detail)
	}
	return
}
