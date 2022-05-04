package utils

import (
	"catvod/global"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
	"time"
)

type Player struct {
	Url string `json:"url"`
}

var cookies []*http.Cookie

func GetMiaoUrl(url string) (playerUrl string) {
	headers := global.Headers
	headers["origin"] = "zjmiao.com"
	client := resty.New()
	client.SetRetryWaitTime(time.Second * 15) //设置超时时间
	get, _ := client.R().
		SetHeaders(headers).
		SetCookies(cookies).
		Get(url)
	fmt.Printf("请求地址: %s, 请求状态码: %d, 请求结果: %+v\n", url, get.StatusCode(), get.Result())
	cookies = get.Cookies()
	fmt.Println(cookies)
	body := string(get.Body())
	player := GetBetweenStr(body, `player_aaaa=`, `<`)
	fmt.Println("354444444444", player)
	var playerJson Player
	err := json.Unmarshal([]byte(player), &playerJson)
	if err != nil {
		log.Println("json unmarshal error:", err)
		return
	}
	fmt.Println(playerJson.Url)
	return playerJson.Url
}
