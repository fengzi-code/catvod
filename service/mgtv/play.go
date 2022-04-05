package mgtv

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

/*
	获取播放地址
解析接口：
https://jx.blbo.cc:4433/?url=
https: //www.gai4.com/gai4/?url=https://www.mgtv.com/b/363286/15255617.html
https://json.pangujiexi.com:12345/json.php?url=https://v.qq.com/x/cover/mzc00200x0x6r30.html
https://jx.m3u8.tv/jiexi/?url=
https://jx.xmflv.com/?url=
*/

func Play(play string) string {
	playurl := "https://jx.blbo.cc:4433/analysis.php?v=" + play
	//playurl := "https://jx.blbo.cc:4433/?url="+ play

	fmt.Println(playurl)
	contentType := "application/x-www-form-urlencoded;charset=UTF-8"
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"
	headers := map[string]string{
		"User-Agent":   userAgent,
		"Content-Type": contentType,
	}
	client := resty.New()

	_, err := client.R().
		SetHeaders(headers).
		Get(playurl)

	if err != nil {
		fmt.Println("ddddddd")
	}

	return playurl
}
