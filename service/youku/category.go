package youku

import (
	"catvod/global"
	"catvod/model"
	"catvod/model/youku/response"
	"fmt"
	"github.com/go-resty/resty/v2"
)

const cateGoryApi = "https://www.youku.com/category/data?c=%s&type=show&p=%d"

func (this *YOUKU) GetCategory(typeId string, page int) (res model.Category) {
	url := fmt.Sprintf(cateGoryApi, typeId, page)
	fmt.Println("分页地址：", url)
	headers := global.Headers
	headers["referer"] = "https://www.youku.com/category/show/" + typeId + ".html?theme=dark"
	client := resty.New()
	get, err := client.R().
		SetResult(response.YoukuCategory{}). //model.Category结构体用来存储返回的json数据
		ForceContentType(global.JsonType).
		SetHeaders(headers).
		// 从芒果api中取分类页中的数据
		Get(url)
	if err != nil {
		fmt.Println("获取分类页数据失败：", err)
		return
	}
	c := get.Result().(*response.YoukuCategory)
	res.Total = 16800
	res.Page = page
	res.Limit = 84
	res.PageCount = 200
	vodInfo := make([]model.VodInfo, 0)

	for _, v := range c.Data.CategoryVideos {
		fmt.Printf("视频名字: %s, 视频id: %s, 视频图片: %s, 视频评论: %s, \n", v.Title, v.VideoId, v.Img, v.SubTitle)
		vodInfo = append(
			vodInfo, model.VodInfo{
				VodId:      v.VideoId,
				VodName:    v.Title,
				VodPic:     "https://" + v.Img,
				VodRemarks: v.SubTitle,
			},
		)
	}
	res.VodList = vodInfo
	return
}
