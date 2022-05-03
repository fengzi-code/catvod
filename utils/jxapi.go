package utils

import (
	"catvod/global"
	"catvod/model"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"sort"
)

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
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&apis); err != nil {
			fmt.Println(err)
		}
		JxApiList = apis.Apis
		sort.Stable(sort.Reverse(JxApiList))
		fmt.Printf("文件改变后，解析Api列表: %+v\n", JxApiList)
	})
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

		reqUrl := fmt.Sprintf("%s%s", v.Url, url)
		switch reqUrl {
		case "https://jx.zjmiao.com/?url=" + url:
			r := GetZJMiaoUrl(reqUrl)
			if r != "" {
				res.Url = r
				res.Header = global.Headers
				return
			}
		default:
			client := resty.New()
			get, _ := client.R().
				SetResult(&JxApiResponse{}).
				SetHeaders(global.Headers).
				Get(reqUrl)
			fmt.Printf("请求地址: %s, 请求状态码: %d, 请求结果: %+v\n", reqUrl, get.StatusCode(), get.Result())
			if get.IsSuccess() {
				r := get.Result().(*JxApiResponse)
				if r.Url != "" {
					res.Url = r.Url
					res.Header = global.Headers
					return
				}
			}
		}

	}
	return
}
