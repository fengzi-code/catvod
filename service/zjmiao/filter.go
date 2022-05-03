package zjmiao

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func (this *ZJMIAO) GetFilter(ext string) (filter string) {
	//{"area":"中国大陆","year":"2022","class":"全部"}
	ext = strings.Replace(ext, "0%3D", "", -1)
	decodeBytes, err := base64.StdEncoding.DecodeString(ext) //对ext解码
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
	fmt.Println(`https://zjmiao.com/index.php/vod/show/id/1/` + filter)
	this.Filters = filter

	return
}
