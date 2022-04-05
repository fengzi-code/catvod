package qqtv

import (
	model "catvod/model/mgtv"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

var filtersPrefix, filtersDefault, filters string

func (this *QQTV) GetFilter(ext string) (filter string) {
	filtersPrefix = "&listpage=1&"
	filtersDefault = "_all=1&sort=19"
	ext = strings.Replace(ext, "0%3D", "", -1)
	filters = getChannelId(ext) //对过滤规则进行修理

	return filters
}

func getChannelId(ext string) string {
	if ext == "" {
		filters = filtersPrefix + filtersDefault
	} else {
		//{"area":"12"}
		decodeBytes, err := base64.StdEncoding.DecodeString(ext) //对ext解码
		if err != nil {
			log.Fatalln(err)
		}
		extDecode := string(decodeBytes) //对解码后的数据转换为字符串格式
		var extStruct model.ExtStruct

		err = json.Unmarshal([]byte(extDecode), &extStruct) //对收到的过滤规则反序列化并存到extStruct结构体
		if err != nil {
			fmt.Println(err)
		}
		area := extStruct.Area
		year := extStruct.Year
		kind := extStruct.Kind
		chargeInfo := extStruct.ChargeInfo
		sort := extStruct.Sort
		edition := extStruct.Edition
		feature := extStruct.Feature
		fitAge := extStruct.FitAge
		filters = filtersPrefix + "area=" + area + "&year=" + year + "&kind=" + kind + "&chargeInfo=" + chargeInfo + "&sort=" + sort + "&edition=" + edition + "&feature=" + feature + "&fitAge=" + fitAge

	}
	return filters

}
