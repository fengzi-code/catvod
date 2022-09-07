package iqytv

import (
	"catvod/utils"
	"encoding/base64"
	"log"
	"strconv"
	"strings"
)

func (this *IQYTV) GetFilter(filter string) (filters string) {

	//{"mode":"11","must0":"1260","three_category_id":"4493","must3":"28985","must1":"4488","must2":"31016","is_purchase":"2"}
	ext := strings.Replace(filter, "%3D", "", -1)            // 由于app返回的时候最后有时会多出0%3D字符,导致base64解码失败,替换为空
	decodeBytes, err := base64.StdEncoding.DecodeString(ext) //对ext解码
	if err != nil {
		log.Fatalln(err)
	}
	decodeData := string(decodeBytes) //对解码后的数据转换为字符串格式
	//{"mode":"11","must0":"20","market_release_date_level":"2022","three_category_id":"16","is_purchase":"2"}
	mode := utils.GetBetweenStr(decodeData, "mode\":\"", "\"")
	isPurchase := utils.GetBetweenStr(decodeData, "is_purchase\":\"", "\"")
	threeCategoryId := utils.GetBetweenStr(decodeData, "three_category_id\":\"", "\"")
	if threeCategoryId != "" {
		threeCategoryId = "=" + threeCategoryId
	}
	year := utils.GetBetweenStr(decodeData, "market_release_date_level\":\"", "\"")
	finished := utils.GetBetweenStr(decodeData, "is_album_finished\":\"", "\"")

	mustCount := strings.Count(decodeData, "must")
	must := ";must"

	for i := 0; i < mustCount+1; i++ {
		mustId := strconv.Itoa(i)
		mustTmp := utils.GetBetweenStr(decodeData, "must"+mustId+"\":\"", "\"")
		must = must + "," + mustTmp + ";"
	}
	//is_purchase=1&market_release_date_level=2020&mode=11&page_id=1&ret_num=48&three_category_id=16;must,20;must
	filters = "&is_purchase=" + isPurchase + "&is_album_finished=" + finished + "&market_release_date_level=" + year + "&mode=" + mode + "&three_category_id" + threeCategoryId + must
	this.Filters = filters
	return
}
