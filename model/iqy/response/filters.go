package response

type IqiyiMyYear struct {
	Cid  []string `json:"cid"`
	List []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"list"`
}

type IqiyiFirstCategoryList struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Order int    `json:"order"`
	Child []struct {
		Name  string        `json:"name"`
		Id    int           `json:"id"`
		Order int           `json:"order"`
		Child []interface{} `json:"child,omitempty"`
	} `json:"child"`
}

type IqiyiOrderlist struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Filter struct {
	Key   string             `json:"key"`  // 过滤器的key，用于拼接url时使用
	Name  string             `json:"name"` // 过滤器的名字，用于显示给人看
	Value []FilterValueItems `json:"value"`
}

// 所以一个影视的过滤器类型是一个数组: []Filter
type Filters []Filter

// 而所有影视的过滤器类型是一个map: map[string]Filters
type FilterMap map[string]Filters

type FilterValueItems struct {
	ShowName string `json:"n"` // 选项展示的名字
	UrlValue string `json:"v"` // 选项的值，用于拼接url时使用
}
