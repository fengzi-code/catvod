package rxlg

import (
	"catvod/global"
	"catvod/model"
	"catvod/model/rongxin"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
)

//https://www.5dy6.cc/vodshow/类型-地区-排序-剧情-语言----页数---年份.html
const rxlgCateGoryApi = "https://caiji.rx.rongxingvr.cc:7777/rongxingapi." +
	"php/provide/vod/from/rx/?ac=list&ac=detail&t=%s&pg=%d"

func (this *RXLG) GetCategory(typeId string, page int) (res model.Category) {
	url := fmt.Sprintf(rxlgCateGoryApi, typeId, page)
	fmt.Println("分页地址：", url)
	client := resty.New()
	get, err := client.R().
		SetResult(rongxin.RxCate{}). //model.Category结构体用来存储返回的json数据
		ForceContentType(global.JsonType).
		//SetHeaders(global.Headers).
		// 从芒果api中取分类页中的数据
		Get(url)
	if err != nil {
		fmt.Println("获取分类页数据失败：", err)
		return
	}
	c := get.Result().(*rongxin.RxCate)
	res.Total = c.Total
	res.Page = page
	res.Limit = 20
	res.PageCount = c.Pagecount
	res.VodList = make([]model.VodInfo, 0)
	for _, v := range c.List {
		fmt.Printf("视频名字: %s, 视频id: %d, 视频图片: %s, 视频描述: %s, \n", v.VodName, v.VodId, v.VodPic, v.VodRemarks)
		res.VodList = append(
			res.VodList, model.VodInfo{
				VodId:      strconv.Itoa(v.VodId),
				VodName:    v.VodName,
				VodPic:     v.VodPic,
				VodRemarks: v.VodRemarks,
			},
		)
	}
	return
}
