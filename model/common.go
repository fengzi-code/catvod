package model

type CommonResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

/*
	VodClass 影视分类结构体，如电视剧、电影等
	TypeId: 分类ID，可以根据这个ID查询分类下的影视
	TypeName: 分类名称，会展示在前端。
*/

type VodClass struct {
	TypeId   string `json:"type_id"`
	TypeName string `json:"type_name"`
}

// VodInfo 影视信息结构体，包含影视的基本信息
type VodInfo struct {
	VodId      string `json:"vod_id"`      // 影视ID
	VodName    string `json:"vod_name"`    // 影视名称
	VodPic     string `json:"vod_pic"`     // 影视展示图片
	VodRemarks string `json:"vod_remarks"` // 影视标记，如多少集，展示在视频名上方
}
