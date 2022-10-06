package rxzh

import (
	"catvod/global"
	"catvod/model"
	"catvod/model/rongxin"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
)

const rxzhSearchUrl = "http://103.222.188.156:99/rongxingapi.php/provide/vod/?ac=list&wd=%s&ac=detail"

func (this *RXZH) Search(wd string) (res []model.VodInfo) {
	url := fmt.Sprintf(rxzhSearchUrl, wd)
	fmt.Println("搜索地址：", url)
	client := resty.New()
	get, err := client.R().
		SetResult(rongxin.RxSearch{}). //model.Category结构体用来存储返回的json数据
		ForceContentType(global.JsonType).
		//SetHeaders(global.Headers).
		// 从芒果api中取分类页中的数据
		Get(url)
	if err != nil {
		fmt.Println("获取搜索页数据失败：", err)
		return
	}
	c := get.Result().(*rongxin.RxSearch)

	for _, s := range c.List {
		res = append(
			res, model.VodInfo{
				VodId:      strconv.Itoa(s.VodId),
				VodName:    s.VodName,
				VodPic:     s.VodPic,
				VodRemarks: s.VodRemarks,
			},
		)
	}
	return
}
