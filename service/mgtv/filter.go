package mgtv

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func (this *MGTV) GetFilter(ext string) (filter string) {
	filtersPrefix := "pc=20&hudong=1&_support=10000000&" //固定不变格式
	filtersDefault := "kind=a1&edition=a1&year=all&chargeInfo=a1&sort=c1&feature=all"
	ext = strings.Replace(ext, "0%3D", "", -1)
	if ext == "" {
		filter = filtersPrefix + filtersDefault
	} else {
		// 对ext解码
		decodeBytes, err := base64.StdEncoding.DecodeString(ext)
		if err != nil {
			log.Println("base64 decode error:", err)
			return filtersPrefix + filtersDefault
		}
		// 将json字符串转换为map类型，并将key和value进行拼接成字符串
		var filterMap map[string]string
		fmt.Printf("decodeBytes: %s\n", decodeBytes)
		err = json.Unmarshal(decodeBytes, &filterMap)
		if err != nil {
			log.Println("json unmarshal error:", err)
			return filtersPrefix + filtersDefault
		}
		for k, v := range filterMap {
			filter += fmt.Sprintf("%s=%s&", k, v)
		}
		if filter != "" {
			filter = filtersPrefix + filter
		} else {
			filter = filtersPrefix + filtersDefault
		}
	}
	this.Filters = filter
	return
}
