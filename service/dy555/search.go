package dy555

import (
	"catvod/model"
	"catvod/utils"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
	"regexp"
)

const dy55SearchUrl = "https://www.555yy2.com/vodsearch/%s-------------.html"

func (this *DY555) Search(wd string) (res []model.VodInfo) {
	surl := fmt.Sprintf(dy55SearchUrl, url.QueryEscape(wd))
	client := resty.New()
	get, err := client.R().
		SetCookie(
			&http.Cookie{
				Name:  "searchneed",
				Value: "ok",
			},
		).
		Get(surl)
	if err != nil {
		fmt.Println("搜索数据失败：", err)
		return
	}
	reg := regexp.MustCompile(`html"><strong>(.*)</strong>`)
	names := reg.FindAllString(get.String(), -1)
	reg = regexp.MustCompile(`<a href="/vodplay/(.*)-`)
	ids := reg.FindAllString(get.String(), -1)
	reg = regexp.MustCompile(`module-item-note">(.*)</div>`)
	remarks := reg.FindAllString(get.String(), -1)
	reg = regexp.MustCompile(`data-original="(.*)"`)
	pics := reg.FindAllString(get.String(), -1)
	for i, name := range names {
		name = utils.GetBetweenStr(name, "<strong>", "</strong>")
		id := utils.GetBetweenStr(ids[i], `<a href="/vodplay/`, `-`)
		remark := utils.GetBetweenStr(remarks[i], `module-item-note">`, `</div>`)
		pic := utils.GetBetweenStr(pics[i], `data-original="`, `"`)
		fmt.Printf("视频名字: %s, 视频id: %s, 视频图片: %s, 视频描述: %s, \n", name, id, pic, remark)
		res = append(
			res, model.VodInfo{
				VodId:      id,
				VodName:    name,
				VodPic:     pic,
				VodRemarks: remark,
			},
		)
	}
	return
}
