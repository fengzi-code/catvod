package mgtv

import (
	model "catvod/model/mgtv"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

/*
	pg: 页码，为空时默认为1
	本模块是实现filter拼接的，所以需要把filter拼接成一个字符串
*/

var filtersPrefix, filtersDefault, filters string

func (m *MGTV) GetFilter(ext string) string {
	GetFilters(ext)
	m.Filters = filters
	return filters
}

func GetFilters(ext string) string {
	filtersPrefix = "&pc=20&hudong=1&_support=10000000&" //固定不变格式
	filtersDefault = "kind=a1&edition=a1&year=all&chargeInfo=a1&sort=c1&feature=all"
	ext = strings.Replace(ext, "0%3D", "", -1) // 由于app返回的时候最后有时会多出0%3D字符,导致base64解码失败,替换为空
	filters = getChannelId2(ext)               //对过滤规则进行修理

	return filters
}

func getChannelId2(ext string) string {

	if ext == "" {
		filters = filtersPrefix + filtersDefault
	} else {
		//{"area":"12"}
		decodeBytes, err := base64.StdEncoding.DecodeString(ext) //对ext解码
		if err != nil {
			log.Fatalln(err)
		}
		extDecode := string(decodeBytes) //对解码后的数据转换为字符串格式
		fmt.Println("getChannelId2", extDecode)
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
