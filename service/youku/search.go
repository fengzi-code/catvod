package youku

import (
	"catvod/model"
	"catvod/model/youku/response"
	"catvod/utils"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const searchUrl = "https://so.youku.com/search_video/q_%s"

var cookies []*http.Cookie

func (this *YOUKU) Search(wd string) (res []model.VodInfo) {
	surl := fmt.Sprintf(searchUrl, url.QueryEscape(wd))
	client := resty.New()
	fmt.Println("cookies:  ", cookies)
	client.SetRetryWaitTime(time.Second * 15) //设置超时时间
	get, _ := client.R().
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
	var youkuSearch response.YoukuSearch1
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
