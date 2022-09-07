package qqtv

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

/*
	qqtv filter
这个函数的作用是将前端传过来的ext参数解析成filter文本
*/
func (this *QQTV) GetFilter(ext string) (filter string) {
	filtersPrefix := "_all=1&pagesize=30&"
	filtersDefault := "sort=19"
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
		//fmt.Printf("decodeBytes: %s\n", decodeBytes)
		err = json.Unmarshal(decodeBytes, &filterMap)
		if err != nil {
			log.Println("json unmarshal error:", err)
			return filtersPrefix + filtersDefault
		}
		//fmt.Printf("%#v\n", filterMap)
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
	return filter
}
