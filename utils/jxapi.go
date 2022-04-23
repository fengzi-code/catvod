package utils

import (
	"catvod/global"
	"catvod/model"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"os"
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

func init() {
	filePtr, err := os.Open("jxapi/jxapi.json")
	if err != nil {
		fmt.Printf("Open file failed [Err:%s]\n", err.Error())
		return
	}
	defer func() {
		err = filePtr.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	bytes, err := ioutil.ReadAll(filePtr)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(bytes, &JxApiList)
	if err != nil {
		fmt.Println(err)
	} else {
		// 对读到的解析接口按权重倒序排序
		sort.Stable(sort.Reverse(JxApiList))
		fmt.Printf("解析Api列表: %+v\n", JxApiList)
	}
}

type JxApiResponse struct {
	Url string `json:"url"`
}

// GetPlayUrl 获取播放地址
func GetPlayUrl(url string) (res model.PlayResponse) {
	for _, v := range JxApiList { // 轮询法
		client := resty.New()
		reqUrl := fmt.Sprintf("%s%s", v.Url, url)
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

	return
}
