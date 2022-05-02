package utils

import (
	"catvod/global"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func GetZJMiaoUrl(reqUrl string) (playUrl string) {
	client := resty.New()
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
	fmt.Println("22222222222222", btToken, url, playUrl)
	return playUrl
}

func Decrypt(text, key, iv string) (string, error) {
	decode_data, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", nil
	}
	//生成密码数据块cipher.Block
	block, _ := aes.NewCipher([]byte(key))
	//解密模式
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	//输出到[]byte数组
	origin_data := make([]byte, len(decode_data))
	blockMode.CryptBlocks(origin_data, decode_data)
	//去除填充,并返回
	return string(unpad(origin_data)), nil
}

func unpad(ciphertext []byte) []byte {
	length := len(ciphertext)
	//去掉最后一次的padding
	unpadding := int(ciphertext[length-1])
	return ciphertext[:(length - unpadding)]
}
