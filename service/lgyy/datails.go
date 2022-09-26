package lgyy

import (
	"catvod/global"
	"catvod/model"
	"fmt"
	"github.com/antchfx/htmlquery"
	"regexp"
	"strings"
)

const lgyyDetailsUrl = "https://www.lgyy.tv/voddetail/%s.html"
const lgyyUrl = "https://www.lgyy.tv"

func (this *LGYY) GetDetails(ids []string) (res []model.VodDetail) {
	for _, id := range ids {
		url := fmt.Sprintf(lgyyDetailsUrl, id)
		doc, err := htmlquery.LoadURL(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		node := htmlquery.FindOne(doc, "//div[@class='module-info-main']")
		var detail model.VodDetail
		detail.TypeName = htmlquery.InnerText(
			htmlquery.FindOne(
				node, "//div[@class='module-info-heading']//div[@class='module-info-tag-link'][3]/a",
			),
		)

		detail.VodId = id
		detail.VodRemarks = htmlquery.InnerText(
			htmlquery.FindOne(
				node, "//div [@class='module-info-content']//div [@class='module-info-item-content'][4]",
			),
		)
		detail.VodName = htmlquery.InnerText(htmlquery.FindOne(node, "//div [@class='module-info-heading']/h1"))
		detail.VodActor = htmlquery.InnerText(
			htmlquery.FindOne(
				node, "//div [@class='module-info-content']//div [@class='module-info-item-content'][1]",
			),
		)
		detail.VodActor, _ = ReplaceStringByRegex(detail.VodActor, "\t|\n| ", "")
		detail.VodArea = htmlquery.InnerText(
			htmlquery.FindOne(
				node, "//div[@class='module-info-heading']//div[@class='module-info-tag-link'][2]/a",
			),
		)
		vodContent := htmlquery.FindOne(
			node,
			"//div[@class='module-info-introduction-content']",
		)

		detail.VodContent = htmlquery.InnerText(vodContent)

		detail.VodContent, _ = ReplaceStringByRegex(detail.VodContent, "\t|\n| ", "")
		detail.VodDirector = htmlquery.InnerText(
			htmlquery.FindOne(
				node, "//div[@class='module-info-item-content'][3]",
			),
		)
		detail.VodDirector, _ = ReplaceStringByRegex(detail.VodDirector, "\t|\n| ", "")
		detail.VodYear = htmlquery.InnerText(
			htmlquery.FindOne(
				node, "//div[@class='module-info-heading']//div[@class='module-info-tag-link'][1]/a",
			),
		)
		vodPic := htmlquery.FindOne(
			doc,
			"//div[@class='module-info-poster']//div[@class='module-item-pic']//img[@class='ls-is-cached lazy lazyload']",
		)
		detail.VodPic = htmlquery.SelectAttr(vodPic, "data-original")
		playNodes := htmlquery.Find(
			doc,
			"//div[@class='module']//div[@class='module-tab-items-box hisSwiper']//div[@class='module-tab-item tab-item']",
		)
		var playNames []string
		for _, item := range playNodes {
			playName := htmlquery.SelectAttr(item, "data-dropdown-value")
			playNames = append(playNames, playName)
		}
		playUrlNodes := htmlquery.Find(
			doc,
			"//div[@id='panel1']//div[contains(@class, 'module-play-list-content')]",
		)
		var playurls []string

		for _, item := range playUrlNodes {

			html := htmlquery.OutputHTML(item, true)
			reg := regexp.MustCompile("<span>(.*)</span>")
			urls := reg.FindAllString(html, -1)

			reg = regexp.MustCompile(`/vodplay/(.*)html`)
			hrefs := reg.FindAllString(html, -1)
			var playurl string
			// 过滤家庭模式中的欧美电视剧
			if len(urls) > 2 && global.AppMode == "f" {
				if detail.VodArea == "美国" || detail.VodArea == "法国" || detail.VodArea == "英国" || detail.
					VodArea == "德国" || detail.VodArea == "其他" || detail.TypeName == "美剧" || detail.TypeName == "欧美剧" {
					fmt.Printf("app运行于家庭模式中,过滤对--[ %s ]--的请求!", detail.VodName)
					return nil
				}
			}
			for i, u := range urls {
				u = strings.Replace(u, "<span>", "", -1)
				u = strings.Replace(u, "</span>", "", -1)
				//href := strings.Replace(hrefs[i], `href="`, "", -1)
				//href = strings.Replace(hrefs[i], `"`, "", -1)
				//href := utils.GetBetweenStr(hrefs[i], `href="`, `.`)
				playurl += u + "$" + lgyyUrl + hrefs[i] + "#"
			}
			playurls = append(playurls, playurl[:len(playurl)-1])
		}
		for i, name := range playNames {
			detail.VodPlayFrom += name + "$$$"
			detail.VodPlayUrl += playurls[i] + "$$$"
		}
		res = append(res, detail)
	}

	return
}

func ReplaceStringByRegex(str, rule, replace string) (string, error) {
	reg, err := regexp.Compile(rule)
	if reg == nil || err != nil {
		return "", err
	}
	return reg.ReplaceAllString(str, replace), nil
}
