package rxlg

import (
	"catvod/global"
	"catvod/model"
	"catvod/model/rongxin"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
)

const rxlgDetailsUrl = "https://caiji.rx.rongxingvr.cc:7777/rongxingapi.php/provide/vod/from/rx/?ac=list&ac=detail&ids=%s"

func (this *RXLG) GetDetails(ids []string) (res []model.VodDetail) {
	for _, id := range ids {
		url := fmt.Sprintf(rxlgDetailsUrl, id)
		fmt.Println("详情地址：", url)
		client := resty.New()
		get, err := client.R().
			SetResult(rongxin.RxDetail{}). //model.Category结构体用来存储返回的json数据
			ForceContentType(global.JsonType).
			Get(url)
		if err != nil {
			fmt.Println("获取详情页数据失败：", err)
			return
		}
		c := get.Result().(*rongxin.RxDetail)
		var detail model.VodDetail
		for _, s := range c.List {
			detail.VodId = strconv.Itoa(s.VodId)
			detail.VodName = s.VodName
			detail.VodRemarks = s.VodRemarks
			detail.TypeName = s.TypeName
			detail.VodActor = s.VodActor
			detail.VodArea = s.VodArea
			detail.VodContent = s.VodContent
			detail.VodDirector = s.VodDirector
			detail.VodYear = s.VodYear
			detail.VodPic = s.VodPic
			detail.VodPlayFrom = s.VodPlayFrom
			detail.VodPlayUrl = s.VodPlayUrl
		}
		res = append(res, detail)
	}

	return
}
