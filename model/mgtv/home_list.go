package mgtv

type HomeList struct {
	ErrCode int        `json:"err_code"`
	Code    int        `json:"code"`
	Ver     string     `json:"ver"`
	Reqid   string     `json:"reqid"`
	Style   string     `json:"style"`
	Data    []HomeData `json:"data"`
	IsRec   int        `json:"isrec"`
}

type HomeData struct {
	PlayUrl    string `json:"play_url"`
	Vv         int    `json:"vv"`
	Type       int    `json:"type"`
	VideoId    int    `json:"videoId"`
	ClipId     int    `json:"clipId"`
	Plid       int    `json:"plid"`
	VideoIndex int    `json:"videoIndex"`
	Name       string `json:"name"`
	Title      string `json:"title"`
	Image      string `json:"image"`
	Desc       string `json:"desc"`
	Plcount    int    `json:"plcount"`
	Info       string `json:"info"`
	RcType     int    `json:"rcType"`
	Source     int    `json:"source"`
	Jumpkind   int    `json:"jumpkind"`
	Cid        int    `json:"cid"`
	PreviewId  int    `json:"previewId"`
	IsPreview  int    `json:"isPreview"`
}

type HomeVideo struct {
	VodId      int    `json:"vod_id"`      // 视频id
	VodName    string `json:"vod_name"`    // 视频名
	Image      string `json:"vod_pic"`     // 展示图片
	VodRemarks string `json:"vod_remarks"` // 视频信息 展示在 视频名上方
}

type HomeClass struct {
	Typeid   string `json:"type_id"`   //视频id
	Typename string `json:"type_name"` // 视频名
}
