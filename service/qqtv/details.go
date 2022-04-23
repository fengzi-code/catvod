package qqtv

import (
	"catvod/model"
	"catvod/model/qqtv/response"
	"encoding/json"
	"fmt"
	"github.com/antchfx/htmlquery"
	"regexp"
	"strings"
)

const detailUrlPrefix = "https://v.qq.com/x/cover"

func (this *QQTV) GetDetails(ids []string) (res []model.VodDetail) {
	for _, id := range ids {
		url := fmt.Sprintf("%s/%s.html", detailUrlPrefix, id)
		doc, err := htmlquery.LoadURL(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		html := htmlquery.OutputHTML(doc, true)
		coverInfoStrArr := regexp.MustCompile("COVER_INFO = (.*)").FindStringSubmatch(html)
		if len(coverInfoStrArr) < 2 {
			return
		}
		// json文件转结构体
		var coverInfo response.CoverInfo
		err = json.Unmarshal([]byte(coverInfoStrArr[1]), &coverInfo)
		if err != nil {
			fmt.Println(err)
			return
		}
		var detail model.VodDetail
		detail.VodId = coverInfo.Id
		detail.VodName = coverInfo.Title
		detail.VodPic = coverInfo.NewPicHz
		detail.VodRemarks = coverInfo.EpisodeUpdated
		detail.TypeName = coverInfo.TypeName
		detail.VodYear = coverInfo.Year
		detail.VodArea = coverInfo.AreaName
		detail.VodActor = strings.Join(coverInfo.LeadingActor, ",")
		detail.VodDirector = strings.Join(coverInfo.Director, ",")
		detail.VodContent = coverInfo.Description
		detail.VodPlayFrom = "qqtv"
		count := 0
		for _, v := range coverInfo.NomalIds {
			if v.F == 0 || v.F == 4 { // F=0或4的为预告，2为免费，7为VIP
				continue
			}
			count += 1 // 去掉预告来计数，避免集数对不上
			detail.VodPlayUrl += fmt.Sprintf("第%d集$%s/%s/%s.html#", count, detailUrlPrefix, id, v.V)
		}
		if strings.HasSuffix(detail.VodPlayUrl, "#") {
			detail.VodPlayUrl = detail.VodPlayUrl[:len(detail.VodPlayUrl)-1]
		}
		res = append(res, detail)
	}
	return
}
