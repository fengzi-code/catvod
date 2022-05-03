package utils

import (
	"catvod/global"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"time"
)

type Player struct {
	Flag     string `json:"flag"`
	Encrypt  int    `json:"encrypt"`
	Trysee   int    `json:"trysee"`
	Points   int    `json:"points"`
	Link     string `json:"link"`
	LinkNext string `json:"link_next"`
	LinkPre  string `json:"link_pre"`
	Url      string `json:"url"`
	UrlNext  string `json:"url_next"`
	From     string `json:"from"`
	Server   string `json:"server"`
	Note     string `json:"note"`
	Id       string `json:"id"`
	Sid      int    `json:"sid"`
	Nid      int    `json:"nid"`
}

func GetMiaoUrl(url string) (playerUrl string) {

	client := resty.New()
	client.SetRetryWaitTime(time.Second * 5) //设置超时时间
	get, _ := client.R().
		SetHeaders(global.Headers).
		Get(url)
	fmt.Printf("请求地址: %s, 请求状态码: %d, 请求结果: %+v\n", url, get.StatusCode(), get.Result())
	body := string(get.Body())
	player := GetBetweenStr(body, `player_aaaa=`, `<`)
	var playerJson Player
	err := json.Unmarshal([]byte(player), &playerJson)
	if err != nil {
		log.Println("json unmarshal error:", err)
	}
	return playerJson.Url
}
