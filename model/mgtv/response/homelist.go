package response

type Homelist struct { //mgtv返回的首页json
	ErrCode int    `json:"err_code"`
	Code    int    `json:"code"`
	Ver     string `json:"ver"`
	Reqid   string `json:"reqid"`
	Style   string `json:"style"`
	Data    []struct {
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
	} `json:"data"`
	Isrec int `json:"isrec"`
}
