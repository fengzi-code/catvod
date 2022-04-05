package model

// Category 影视分类结构体，单个分类页会用到这个结构体
type Category struct {
	CommonResponse
	Page      int       `json:"page"`      // 当前页
	PageCount int       `json:"pagecount"` // 总页数
	Limit     int       `json:"limit"`     // 每页显示数量
	Total     int       `json:"total"`     // 总记录数
	VodList   []VodInfo `json:"list"`      // 视频列表 同homeContent中的
}
