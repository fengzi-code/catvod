package model

/*
每个过滤器都可能有很多个值，比如年份过滤器，key为year, label为年份，value为
[{"key": "2018", "label": "2018"}, {"key": "2019", "key": "2019"}]
key是用于拼接url查询的，label是用于显示给人看的
*/

type FilterValueItems struct {
	ShowName string `json:"n"` // 选项展示的名字
	UrlValue string `json:"v"` // 选项的值，用于拼接url时使用
}

/*
Filter就代表一个过滤器，像上面提到的年份，区域，这样的就是过滤器
一个影视类型可能会有很多Filter
*/
type Filter struct {
	Key   string             `json:"key"`  // 过滤器的key，用于拼接url时使用
	Name  string             `json:"name"` // 过滤器的名字，用于显示给人看
	Value []FilterValueItems `json:"value"`
}

// 所以一个影视的过滤器类型是一个数组: []Filter
type Filters []Filter

// 而所有影视的过滤器类型是一个map: map[string]Filters
type FilterMap map[string]Filters
