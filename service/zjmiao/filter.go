package zjmiao

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
)

func (this *ZJMIAO) GetFilter(ext string) (filter string) {
	ext, _ = url.PathUnescape(ext)
	decodeBytes, err := base64.URLEncoding.DecodeString(ext) //对ext解码,这里需要使用urlencoding
	if err != nil {
		log.Fatalln(err)
	}
	var filterMap map[string]string
	err = json.Unmarshal(decodeBytes, &filterMap)
	if err != nil {
		log.Println("json unmarshal error:", err)
	}
	fmt.Printf("%#v\n", filterMap)
	for k, v := range filterMap {
		filter += fmt.Sprintf("%s/%s/", k, v)
	}
	filter = strings.Replace(filter, "area/全部/", "", -1)
	filter = strings.Replace(filter, "year/全部/", "", -1)
	filter = strings.Replace(filter, "class/全部/", "", -1)
	fmt.Println(`https://zjmiao.com/index.php/vod/show/id/1/` + filter)
	this.Filters = filter
	return
}
