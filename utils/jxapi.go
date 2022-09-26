package utils

import (
	"catvod/model"
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Dy555Url struct {
	Flag     string `json:"flag"`
	Encrypt  int    `json:"encrypt"`
	Trysee   int    `json:"trysee"`
	Points   int    `json:"points"`
	Link     string `json:"link"`
	LinkNext string `json:"link_next"`
	LinkPre  string `json:"link_pre"`
	VodData  struct {
		VodName     string `json:"vod_name"`
		VodActor    string `json:"vod_actor"`
		VodDirector string `json:"vod_director"`
		VodClass    string `json:"vod_class"`
	} `json:"vod_data"`
	Url     string `json:"url"`
	UrlNext string `json:"url_next"`
	From    string `json:"from"`
	Server  string `json:"server"`
	Note    string `json:"note"`
	Id      string `json:"id"`
	Sid     int    `json:"sid"`
	Nid     int    `json:"nid"`
}
type JxApi struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Scope  string `json:"scope"`
	Weight int    `json:"weight"`
	Parse  int    `json:"parse"`
}

type JxApis []JxApi

var JxApiList JxApis

func (s JxApis) Len() int {
	return len(s)
}

func (s JxApis) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s JxApis) Less(i, j int) bool {
	return s[i].Weight < s[j].Weight
}

type Apis struct {
	Apis JxApis `json:"apis"`
}

func init() {
	v := viper.New()
	v.SetConfigFile("jxapi/jxapi.json")
	v.SetConfigType("json")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	var apis Apis
	v.OnConfigChange(
		func(e fsnotify.Event) {
			fmt.Println("config file changed:", e.Name)
			if err := v.Unmarshal(&apis); err != nil {
				fmt.Println(err)
			}
			JxApiList = apis.Apis
			sort.Stable(sort.Reverse(JxApiList))
			fmt.Printf("文件改变后，解析Api列表: %+v\n", JxApiList)
		},
	)
	if err := v.Unmarshal(&apis); err != nil {
		fmt.Println(err)
	}
	JxApiList = apis.Apis
	sort.Stable(sort.Reverse(JxApiList))
	fmt.Printf("解析Api列表: %+v\n", JxApiList)

}

type JxApiResponse struct {
	Url string `json:"url"`
}

// GetPlayUrl 获取播放地址
func GetPlayUrl(url string) (res model.PlayResponse) {

	for _, v := range JxApiList { // 轮询法
		//reqUrl := fmt.Sprintf("%s%s", v.Url, url)
		n := v.Name
		if strings.Contains(url, "5dy6") || strings.Contains(url, "lgyy") || strings.Contains(url, "cokemv") {
			n = "555dy"
			//playMaoUrl := GetMiaoUrl(url)
			//reqUrl = fmt.Sprintf("%s%s", "https://jx.zjmiao.com/?url=", playMaoUrl)
		}

		switch n {
		case "555dy":
			fmt.Printf("555dy请求地址: %s \n", url)
			//playurl, cook, head := GetDy555Play(url)
			//res.Url = playurl

			res.Header = map[string]string{
				"User-Agent": "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36",
			}
			res.Url = url
			res.Parse = 1

			return

		default:
			fmt.Printf("请求地址: %s, 解析API: %s", url, v.Url)
			res.Url = url
			res.PlayUrl = v.Url
			res.Parse = 1
			return
		}

	}
	return
}

func GetDy555Play(url string) (playurl string, cookies []*http.Cookie, headers http.Header) {
	client := resty.New()
	//client.SetProxy("http://
	client.SetRetryWaitTime(time.Second * 15) //设置超时时间
	dyUrlGet, _ := client.R().
		SetHeaders(
			map[string]string{
				"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36 Edg/102.0.1245.33",
			},
		).
		Get(url)

	dyUrlStr := GetBetweenStr(dyUrlGet.String(), `player_aaaa=`, `</script>`)
	var dy555Url Dy555Url
	err := json.Unmarshal([]byte(dyUrlStr), &dy555Url)
	if err != nil {
		log.Fatal(err)
	}
	headers = make(map[string][]string)
	playurl = `https://player.sakurot.com:3458/?url=` + dy555Url.Url + `&jump=` + url
	//playurl = `https://www.5dy6.cc/static/player/duoduozy.js?v=202207055`
	headers.Set("Host", "www.5dy6.cc")
	headers.Set("Referer", url)

	return playurl, dyUrlGet.Cookies(), headers
}
