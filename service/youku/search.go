package youku

import (
	"catvod/model"
	"catvod/model/youku/response"
	"catvod/utils"
	"crypto/md5"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	sApi      = `mtop.youku.soku.yksearch`
	sAppkey   = `23774304`
	sData     = `{"searchType":1,"keyword":"SoKeyword","pg":1,"pz":10,"site":1,"appCaller":"pc","appScene":"mobile_multi","userTerminal":2,"sdkver":313,"userFrom":1,"noqc":0,"aaid":"bf533a34a7532e5d1d0893b9e2e3bdef","ftype":0,"duration":"","categories":"","ob":"","utdId":"tP/hF4OrRWkCAXL6lgpWdwlE","userType":"guest","userNumId":0,"searchFrom":"1","sourceFrom":"home"}`
	searcapi1 = "https://acs.youku.com/h5/mtop.youku.soku.yksearch/2.0/?jsv=2.5.1&appKey=" + sAppkey
	sH5Url    = "https://acs.youku.com/h5/" + sApi + "/2.0/?jsv=2.5.1&appKey=" + "23774304" + "&api=" + sApi
)

func (this *YOUKU) Search(wd string) (res []model.VodInfo) {
	client := resty.New()
	//client.SetProxy("http://
	client.SetRetryWaitTime(time.Second * 15) //设置超时时间
	data := strings.Replace(sData, `SoKeyword`, wd, 1)
	ySoget, _ := client.R().
		SetHeaders(
			map[string]string{
				"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36 Edg/102.0.1245.33",
			},
		).
		Get(sH5Url)

	youkuCookies := ySoget.Cookies()
	fmt.Println("cookies: ", youkuCookies)
	var m_h5_tk string
	for _, v := range youkuCookies {
		if v.Name == "_m_h5_tk" {
			m_h5_tk = v.Value
		}
	}

	tokens := strings.Split(m_h5_tk, "_")
	//取时间戳
	timeStamp := tokens[1]
	token := tokens[0]
	sign := token + "&" + timeStamp + "&" + sAppkey + "&" + data
	//md5加密
	signByte := md5.Sum([]byte(sign))
	sign = fmt.Sprintf("%x", signByte)
	api := searcapi1 + `&t=` + timeStamp + `&sign=` + sign + "&api=" + sApi + `&type=originaljson&v=2.0&ecode=1&dataType=json&jsonpIncPrefix=headerSearch&data=` + url.QueryEscape(data)
	ySoget, err := client.R().
		SetHeaders(
			map[string]string{
				"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
				"Referer":    "https://so.youku.com",
				"Origin":     "https://so.youku.com",
			},
		).
		SetCookies(youkuCookies).
		SetResult(response.YoukuSearch{}).
		Get(api)
	if err != nil {
		print(err)
	}
	c := ySoget.Result().(*response.YoukuSearch)
	//fmt.Println(ySoget.String())
	for _, node := range c.Data.Nodes {
		for _, s := range node.Nodes {
			if s.Nodes[0].Data.TempTitle != "" && s.Nodes[0].Data.IsYouku == 1 {
				name := s.Nodes[0].Data.TempTitle
				showid := s.Nodes[0].Data.RealShowId
				vid := s.Nodes[0].Data.VideoId
				if vid == "" {
					for i := 0; i < 10; i++ {
						fmt.Println("第", i, "次匹配")
						if i == 9 {
							return
						}
						ySoget2, _ := client.R().
							SetHeaders(
								map[string]string{
									"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
									"Referer":    "https://v.youku.com",
									"Origin":     "https://v.youku.com",
								},
							).
							Get("https://v.youku.com/v_nextstage/id_" + showid + ".html?spm=a2ha1.14919748_WEBMOVIE_JINGXUAN.drawer6.d_zj1_3&s=" + showid + "&scm=20140719.manual.4423.show_" + showid)
						vid = utils.GetBetweenStr(ySoget2.RawResponse.Request.URL.String(), `id_`, `.html`)
						fmt.Println("第", i, "次匹配", vid)
						if vid == "" || vid == showid {
							continue
						}
						id, err := strconv.Atoi(vid) // vid不为数字则匹配成功,否则继续匹配
						if err != nil {
							fmt.Println("匹配成功: ", id, vid)
							break
						}

						time.Sleep(time.Second / 5)

					}
				}
				pic := s.Nodes[0].Data.PosterDTO.VThumbUrl
				remark := s.Nodes[0].Data.StripeBottom
				fmt.Printf("视频名字: %s, 视频id: %s, 视频图片: %s, 视频描述: %s, \n", name, vid, pic, remark)
				res = append(
					res, model.VodInfo{
						VodId:      vid,
						VodName:    name,
						VodPic:     pic,
						VodRemarks: remark,
					},
				)
			}
		}
	}
	return
}
