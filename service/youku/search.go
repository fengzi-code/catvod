package youku

import (
	"catvod/model"
	"catvod/model/youku/response"
	"catvod/utils"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const searchUrl = "https://so.youku.com/search_video/q_%s"
const (
	sApi      = `mtop.youku.soku.yksearch`
	sAppkey   = `23774304`
	sData     = `{"searchType":1,"keyword":"女人","pg":1,"pz":5,"site":1,"appCaller":"pc","appScene":"mobile_multi","userTerminal":2,"sdkver":313,"userFrom":1,"noqc":0,"aaid":"bf533a34a7532e5d1d0893b9e2e3bdef","ftype":0,"duration":"","categories":"","ob":"","utdId":"tP/hF4OrRWkCAXL6lgpWdwlE","userType":"guest","userNumId":0,"searchFrom":"1","sourceFrom":"home"}`
	searcapi1 = "https://acs.youku.com/h5/mtop.youku.soku.yksearch/2.0/?jsv=2.5.1&appKey=" + sAppkey
	sH5Url    = "https://acs.youku.com/h5/" + sApi + "/2.0/?jsv=2.5.1&appKey=" + "23774304" + "&api=" + sApi
)

var cookies []*http.Cookie

func (this *YOUKU) Search(wd string) (res []model.VodInfo) {

	client := resty.New()
	client.SetRetryWaitTime(time.Second * 15) //设置超时时间

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

	sign := token + "&" + timeStamp + "&" + sAppkey + "&" + sData
	//md5加密
	signByte := md5.Sum([]byte(sign))
	sign = fmt.Sprintf("%x", signByte)

	api := searcapi1 + `&t=` + timeStamp + `&sign=` + sign + "&api=" + sApi + `&type=originaljson&v=2.0&ecode=1&dataType=json&jsonpIncPrefix=headerSearch&data=` + url.QueryEscape(sData)
	fmt.Println(api)
	var youkuSearch response.YoukuSearch
	get, _ = client.R().
		SetHeaders(
			map[string]string{
				"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
				"Referer":    "https://so.youku.com",
				"Origin":     "https://so.youku.com",
			},
		).
		SetCookies(cookies).
		SetResult(&youkuSearch).
		Get(api)
	fmt.Println(get.String())
	os.Exit(0)

	surl := fmt.Sprintf(searchUrl, url.QueryEscape(wd))
	client = resty.New()
	fmt.Println("cookies:  ", cookies)
	client.SetRetryWaitTime(time.Second * 15) //设置超时时间
	get, _ = client.R().
		SetHeaders(
			map[string]string{
				"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
				"Referer":    surl,
				"Origin":     "https://v.youku.com",
			},
		).
		SetCookies(cookies).
		Get(surl)
	cookies = get.Cookies()
	body := string(get.Body())

	bodyJson := utils.GetBetweenStr(body, `__INITIAL_DATA__ =`, `;window`)
	fmt.Println(bodyJson)
	bodyJson = utils.Unicode2utf8(bodyJson)
	//var youkuSearch response.YoukuSearch1
	err := json.Unmarshal([]byte(bodyJson), &youkuSearch)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, node := range youkuSearch.Data.Nodes {
		for _, nod := range node.Nodes {
			if nod.Data.FeatureDTO.Text != "" && nod.Nodes[0].Data.VideoId != "" {
				remark := strings.ReplaceAll(nod.Data.FeatureDTO.Text, `<em>`, ``)
				remark = strings.ReplaceAll(remark, `</em>`, ``)
				res = append(
					res, model.VodInfo{
						VodId:      nod.Nodes[0].Data.VideoId,
						VodName:    nod.Data.TempTitle,
						VodPic:     nod.Data.PosterDTO.VThumbUrl,
						VodRemarks: remark,
					},
				)
			}
		}
	}
	return
}
