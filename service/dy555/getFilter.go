package dy555

import (
	"catvod/global"
	"catvod/model"
	"catvod/utils"
	"fmt"
	"github.com/go-resty/resty/v2"
	"regexp"
	"strings"
)

const dy555FilterUrl = "https://www.555yy2.com/vodshow/%s-----------.html"

func GetFilterMap(t string) (res model.Filters) {
	url := fmt.Sprintf(dy555FilterUrl, t)
	client := resty.New()
	get, err := client.R().
		SetHeaders(global.Headers).
		Get(url)

	if err != nil {
		fmt.Println(err)
	}

	str := utils.GetBetweenStr(get.String(), "module-class-item scroll-content", `<!--电脑广告-->`)

	urls := strings.Split(str, `<div class="module-item-t`)
	var filter model.Filter
	for i, s := range urls {
		if i == 0 {
			continue
		}
		reg := regexp.MustCompile(`itle">(.*)<i`)
		names := reg.FindAllString(s, -1)
		//name = utils.GetBetweenStr(name, `>`, `<`)
		reg = regexp.MustCompile(`title=(.*)</a>`)
		showNames := reg.FindAllString(s, -1)

		s = strings.Replace(s, `href="/vodshow/`+t+`-----------.html`, ``, -1)
		reg = regexp.MustCompile(`"/vodshow(.*)html`)
		urlValues := reg.FindAllString(s, -1)
		var filterValueItems model.FilterValueItems
		var filterValueItemsTmp []model.FilterValueItems
		var fName = map[string]string{
			"类型": "class",
			"排序": "sort",
			"地区": "area",
			"年份": "year",
			"剧情": "item",
			"语言": "lang",
		}
		for _, name := range names {
			filter.Name = utils.GetBetweenStr(name, `>`, `<`)
			filter.Key = fName[filter.Name]
			filter.Value = filterValueItemsTmp
			for i3, showName := range showNames {
				filterValueItems.ShowName = utils.GetBetweenStr(showName, `>`, `<`)
				if filter.Name == "类型" {
					v := utils.GetBetweenStr(urlValues[i3], `show/`, `-`)
					filterValueItems.UrlValue = v
				} else if filter.Name == "排序" {
					v := utils.GetBetweenStr(urlValues[i3], `--`, `-`)
					filterValueItems.UrlValue = v
				} else {
					filterValueItems.UrlValue = filterValueItems.ShowName
				}
				filter.Value = append(filter.Value, filterValueItems)
			}
			res = append(
				res, model.Filter{
					Name:  filter.Name,
					Key:   filter.Key,
					Value: filter.Value,
				},
			)
		}
	}
	return
}
