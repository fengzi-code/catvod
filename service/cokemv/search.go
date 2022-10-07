package cokemv

import (
	"catvod/global"
	"catvod/model"
	"catvod/utils"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

const cokemvSearchUrl = "https://cokemv.me/vodsearch/%s-------------.html"
const imgUrl = "https://cokemv.me/index.php/verify/index.html?"
const verifyUrl = "https://cokemv.me/index.php/ajax/verify_check?type=search&verify="

type verifyJson struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (this *COKEMV) Search(wd string) (res []model.VodInfo) {

	var imgCode string
	var cookies []*http.Cookie
	client := resty.New()
	var i int
	for {
		// 取图片验证码
		i = i + 1
		imgCode, cookies = utils.GetImgCode(imgUrl)
		fmt.Println("imcode............", imgCode)
		// 提交图片验证码
		post, _ := client.R().
			SetHeaders(
				map[string]string{
					"x-requested-with": "XMLHttpRequest",
				},
			).
			SetCookies(cookies).
			SetResult(verifyJson{}).
			ForceContentType(global.JsonType).
			Post(verifyUrl + imgCode)
		json := post.Result().(*verifyJson)
		//fmt.Println(
		//	"cokes........................", post.Cookies(), cookies, "code:", json.Code, "msg:"+json.Msg+"22222222",
		//)
		if json.Msg == "ok" || i > 5 {
			break
		}
		time.Sleep(time.Second / 10)
	}

	surl := fmt.Sprintf(cokemvSearchUrl, url.QueryEscape(wd))

	get, err := client.R().
		SetCookies(cookies).
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
