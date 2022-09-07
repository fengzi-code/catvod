package dy555

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func (this *DY555) GetFilter(ext string) (filter string) {
	filtersPrefix := ""
	filtersDefault := ""
	ext = strings.Replace(ext, "%3D", "", -1)
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
	return filter
}
