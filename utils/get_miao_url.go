package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"time"
)

type Player struct {
	Url string `json:"url"`
}

func GetMiaoUrl(url string) (playerUrl string) {
	client := resty.New()
	client.SetRetryWaitTime(time.Second * 5) //设置超时时间
	get, _ := client.R().Get(url)
	fmt.Printf("请求地址: %s, 请求状态码: %d, 请求结果: %+v\n", url, get.StatusCode(), get.Result())
	body := string(get.Body())
	player := GetBetweenStr(body, `player_aaaa=`, `<`)
	var playerJson Player
	err := json.Unmarshal([]byte(player), &playerJson)
	if err != nil {
		log.Println("json unmarshal error:", err)
	}
	fmt.Println(playerJson.Url)
	return playerJson.Url
}
