package youku

import (
	"catvod/model"
	"catvod/model/youku/response"
	"crypto/md5"
	"fmt"
	"github.com/go-resty/resty/v2"
	url3 "net/url"
	"strconv"
	"strings"
	"time"
)

const (
	appkey    = "24679788"
	detailApi = "https://v.youku.com/v_show/id_"
	h5Api     = "mtop.youku.columbus.gateway.new.execute"
	h5Url     = "https://acs.youku.com/h5/" + h5Api + "/1.0/?jsv=2.6.2&appKey=" + appkey + "&api=" + h5Api
	desData   = `{"ms_codes":"2019030100","params":"{\"biz\":\"new_detail_web\",\"videoId\":\"iiiddd\",\"scene\":\"web_page\"}","system_info":"{\"os\":\"pc\",\"device\":\"pc\",\"ver\":\"1.0.0\",\"appPackageKey\":\"pcweb\",\"appPackageId\":\"pcweb\"}"}`
	playData  = `{"ms_codes":"2019030100","params":"{\"biz\":true,\"scene\":\"component\",\"nextSession\":\"{\\\"componentIndex\\\":\\\"3\\\",\\\"componentId\\\":\\\"61518\\\",\\\"level\\\":\\\"2\\\",\\\"itemPageNo\\\":\\\"0\\\",\\\"lastItemIndex\\\":\\\"0\\\",\\\"pageKey\\\":\\\"LOGICSHOW_LOGICTV_DEFAULT\\\",\\\"group\\\":\\\"0\\\",\\\"itemStartStage\\\":1,\\\"itemEndStage\\\":2000}\",\"videoId\":\"iiiddd\",\"showId\":\"\"}","system_info":"{\"os\":\"pc\",\"device\":\"pc\",\"ver\":\"1.0.0\",\"appPackageKey\":\"pcweb\",\"appPackageId\":\"pcweb\"}"}`
)

func (this *YOUKU) GetDetails(ids []string) (res []model.VodDetail) {
	for _, id := range ids {
		desDataR := strings.ReplaceAll(desData, "iiiddd", id)
		playDataR := strings.ReplaceAll(playData, "iiiddd", id)
		url := fmt.Sprintf("%s%s.html", detailApi, id)
		client := resty.New()
		client.SetRetryWaitTime(time.Second * 15) //设置超时时间
		get, _ := client.R().
			SetHeaders(
				map[string]string{
					"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36 Edg/102.0.1245.33",
				},
			).
			Get(h5Url)

		cookies := get.Cookies()
		fmt.Println("cookies: ", cookies)
		var m_h5_tk string
		for _, v := range cookies {
			if v.Name == "_m_h5_tk" {
				m_h5_tk = v.Value
			}

		}

		tokens := strings.Split(m_h5_tk, "_")
		//取时间戳
		timeStamp := tokens[1]
		token := tokens[0]
		sign := token + "&" + timeStamp + "&" + appkey + "&" + desDataR
		//md5加密
		signByte := md5.Sum([]byte(sign))
		sign = fmt.Sprintf("%x", signByte)
		api := h5Url + `&t=` + timeStamp + `&sign=` + sign + `&type=originaljson&v=1.0&ecode=1&dataType=json&data=` + url3.QueryEscape(desDataR)
		var describe response.YoukuDescribe
		get, _ = client.R().
			SetHeaders(
				map[string]string{
					"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
					"Referer":    url,
					"Origin":     "https://v.youku.com",
				},
			).
			SetResult(&describe).
			SetCookies(cookies).
			Get(api)
		var director string
		var actors string
		var vodContent string
		var vodYear string
		//var detail response.DetailResponse
		var detail model.VodDetail
		//JSON.data[2019030100].data.nodes[0].nodes[0].nodes[0].data.desc
		for _, node := range describe.Data.Field1.Data.Nodes {
			for _, nod := range node.Nodes {
				for _, no := range nod.Nodes {
					if no.Data.SubtitleType == "GENERAL" {
						if no.Data.Subtitle == "导演" {
							director += no.Data.Title + " "
						} else {
							actors += no.Data.Title + " "
						}
					} else if no.Data.SubtitleType == "" {
						if vodContent == "" {
							vodContent = no.Data.Desc
							vodYear = strconv.Itoa(no.Data.ShowReleaseYear)
						}
					}

				}
			}
		}
		detail.VodId = id
		detail.VodName = describe.Data.Field1.Data.Data.Extra.ShowName
		detail.VodRemarks = describe.Data.Field1.Data.Data.Extra.ShowSubtitle
		detail.TypeName = describe.Data.Field1.Data.Data.Extra.ShowCategory //类型
		detail.VodActor = actors                                            //主演
		//detail.Data.VodInfo.VodArea = ""           //地区
		detail.VodContent = vodContent //简介
		detail.VodDirector = director  //导演
		detail.VodYear = vodYear       //年份
		detail.VodPlayFrom = "youku"
		detail.VodPic = describe.Data.Field1.Data.Data.Extra.ShowImgV
		sign = token + "&" + timeStamp + "&" + appkey + "&" + playDataR
		signByte = md5.Sum([]byte(sign))
		sign = fmt.Sprintf("%x", signByte)
		api = h5Url + `&t=` + timeStamp + `&sign=` + sign + `&type=originaljson&v=1.0&ecode=1&dataType=json&data=` + url3.QueryEscape(playDataR)
		var youkuPlay response.YoukuPlay
		get, _ = client.R().
			SetHeaders(
				map[string]string{
					"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
					"Referer":    url,
					"Origin":     "https://v.youku.com",
				},
			).
			SetResult(&youkuPlay).
			SetCookies(cookies).
			Get(api)
		var vodPlayUrl string
		for _, node := range youkuPlay.Data.Field1.Data.Nodes {
			vname := node.Data.Title
			vid := node.Data.Action.Value
			vid = fmt.Sprintf("%s%s.html", detailApi, vid)
			vname = strings.ReplaceAll(vname, " ", "")
			if vname != detail.VodName {
				vname = strings.ReplaceAll(vname, detail.VodName, "")
			}
			vodPlayUrl = vodPlayUrl + vname + "$" + vid + "#"

		}
		vodPlayUrl = vodPlayUrl[:len(vodPlayUrl)-1]
		detail.VodPlayUrl = vodPlayUrl
		res = append(res, detail)
	}
	return
}
