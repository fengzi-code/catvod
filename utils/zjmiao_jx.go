package utils

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

func GetZJMiaoUrl(reqUrl string) (playUrl string) {
	// 加了headers会出现搜索后无法播放的情况，并且导致其他从分类页点进去的也无法播放
	// headers := global.Headers
	// headers["Host"] = "jx.zjmiao.com"
	// headers["origin"] = "jx.zjmiao.com"
	client := resty.New()
	client.SetRetryWaitTime(time.Second * 15) //设置超时时间
	get, _ := client.R().
		// SetHeaders(headers).
		Get(reqUrl)
	fmt.Printf("请求地址: %s, 请求状态码: %d, 请求结果: %+v\n", reqUrl, get.StatusCode(), get.Result())
	body := string(get.Body())
	btToken := GetBetweenStr(body, `bt_token = "`, `"`)
	url := GetBetweenStr(body, `getVideoInfo("`, `"`)
	aesPass := "63f49d21a0dccf3c"
	playUrl, _ = Decrypt(url, aesPass, btToken)
	return playUrl
}
