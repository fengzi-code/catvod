package youku

import (
	"catvod/model"
	"catvod/model/youku/response"
	"catvod/utils"
	"crypto/md5"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
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

var cookies []*http.Cookie

func (this *YOUKU) Search(wd string) (res []model.VodInfo) {
	client := resty.New()
	client.SetRetryWaitTime(time.Second * 15) //设置超时时间
	data := strings.Replace(sData, `SoKeyword`, wd, 1)
	fmt.Println(sH5Url)
	get, _ := client.R().
		SetHeaders(
			map[string]string{
				"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36 Edg/102.0.1245.33",
			},
		).
		Get(sH5Url)
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
	fmt.Println(token, timeStamp, sAppkey)
	sign := token + "&" + timeStamp + "&" + sAppkey + "&" + data
	//md5加密
	signByte := md5.Sum([]byte(sign))
	sign = fmt.Sprintf("%x", signByte)
	api := searcapi1 + `&t=` + timeStamp + `&sign=` + sign + "&api=" + sApi + `&type=originaljson&v=2.0&ecode=1&dataType=json&jsonpIncPrefix=headerSearch&data=` + url.QueryEscape(data)
	fmt.Println(api)
	get, _ = client.R().
		SetHeaders(
			map[string]string{
				"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
				"Referer":    "https://so.youku.com",
				"Origin":     "https://so.youku.com",
			},
		).
		SetCookies(cookies).
		SetResult(response.YoukuSearch{}).
		Get(api)

	c := get.Result().(*response.YoukuSearch)
	for _, node := range c.Data.Nodes {
		for _, s := range node.Nodes {

			if s.Nodes[0].Data.TempTitle != "" && s.Nodes[0].Data.IsYouku == 1 {
				name := s.Nodes[0].Data.TempTitle
				showid := s.Nodes[0].Data.RealShowId
				resp, err := http.Get("https://v.youku.com/v_nextstage/id_" + showid + ".html?spm=a2h0c.8166622.PhoneSokuProgram_1.dposter")
				request := resp.Request.URL.String()
				if err != nil {
					fmt.Println(err)
					request = ""
				}
				pic := s.Nodes[0].Data.PosterDTO.VThumbUrl
				remark := s.Nodes[0].Data.StripeBottom
				vid := utils.GetBetweenStr(request, `id_`, `.html`)
				fmt.Println(name, showid, pic, remark, vid)
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
