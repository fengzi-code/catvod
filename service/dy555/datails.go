package dy555

import (
	"catvod/model"
	"fmt"
	"github.com/antchfx/htmlquery"
	"regexp"
	"strings"
)

const dy555DetailsUrl = "https://www.5dy6.cc/voddetail/%s.html"
const dy555Url = "https://www.5dy6.cc"

func (this *DY555) GetDetails(ids []string) (res []model.VodDetail) {
	for _, id := range ids {
		url := fmt.Sprintf(dy555DetailsUrl, id)
		fmt.Println("url: ", url, "444")
		doc, err := htmlquery.LoadURL(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		node := htmlquery.FindOne(doc, "//div[@class='module-info-main']")
		//fmt.Println(htmlquery.OutputHTML(node, true))
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
				node, "//div [@class='module-info-content']//div [@class='module-info-item-content'][2]",
			),
		)
		detail.VodArea = htmlquery.InnerText(
			htmlquery.FindOne(
				node, "//div[@class='module-info-heading']//div[@class='module-info-tag-link'][2]/a",
			),
		)
		vodContent := htmlquery.FindOne(
			node,
			"//div[@class='module-info-item module-info-introduction']",
		)

		detail.VodContent = htmlquery.InnerText(vodContent)
		detail.VodDirector = htmlquery.InnerText(
			htmlquery.FindOne(
				node, "//div [@class='module-info-content']//div[@class='module-info-item-content'][1]",
			),
		)
		detail.VodYear = htmlquery.InnerText(
			htmlquery.FindOne(
				node, "//div[@class='module-info-heading']//div[@class='module-info-tag-link'][1]/a",
			),
		)
		vodPic := htmlquery.FindOne(
			doc,
			"//div[@class='module-info-poster']//div[@class='module-item-pic']//img[@class='ls-is-cached lazy lazyload']",
		)
		fmt.Println(htmlquery.OutputHTML(vodPic, true))
		detail.VodPic = htmlquery.SelectAttr(vodPic, "data-original")
		fmt.Println(detail)
		playNodes := htmlquery.Find(
			doc,
			"//div[@class='module']//div[@class='module-tab-items-box hisSwiper']//div[@class='module-tab-item tab-item']",
		)
		var playNames []string
		for _, item := range playNodes {
			playName := htmlquery.SelectAttr(item, "data-dropdown-value")
			playNames = append(playNames, playName)
			fmt.Println(playNames)
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
			fmt.Println(urls, hrefs)
			var playurl string
			for i, u := range urls {
				u = strings.Replace(u, "<span>", "", -1)
				u = strings.Replace(u, "</span>", "", -1)
				//href := strings.Replace(hrefs[i], `href="`, "", -1)
				//href = strings.Replace(hrefs[i], `"`, "", -1)
				//href := utils.GetBetweenStr(hrefs[i], `href="`, `.`)
				playurl += u + "$" + dy555Url + hrefs[i] + "#"
			}
			playurls = append(playurls, playurl[:len(playurl)-1])
		}
		fmt.Println(len(playUrlNodes), playurls)
		for i, name := range playNames {
			detail.VodPlayFrom += name + "$$$"
			detail.VodPlayUrl += playurls[i] + "$$$"
		}
		res = append(res, detail)
	}

	return
}