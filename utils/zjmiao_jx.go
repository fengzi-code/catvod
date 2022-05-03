package utils

import (
	"catvod/global"
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

func GetZJMiaoUrl(reqUrl string) (playUrl string) {
	client := resty.New()
	client.SetRetryWaitTime(time.Second * 5) //设置超时时间
	get, _ := client.R().
		SetHeaders(global.Headers).
		Get(reqUrl)
	fmt.Printf("请求地址: %s, 请求状态码: %d, 请求结果: %+v\n", reqUrl, get.StatusCode(), get.Result())
	body := string(get.Body())
	//fmt.Println(body)
	btToken := GetBetweenStr(body, `bt_token = "`, `"`)
	url := GetBetweenStr(body, `getVideoInfo("`, `"`)
	aesPass := "63f49d21a0dccf3c"

	playUrl, _ = Decrypt(url, aesPass, btToken)
	return playUrl
}
