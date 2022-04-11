package response

type PlayUrl struct { //mgtv 视频集数/播放地址 返回的json
	Data struct {
		Total     int           `json:"total"`
		Series    []interface{} `json:"series"`
		Count     int           `json:"count"`
		TotalPage int           `json:"total_page"`
		Short     []interface{} `json:"short"`
		List      []struct {
			T4        string `json:"t4"`
			Img       string `json:"img"`
			SrcClipId string `json:"src_clip_id"`
			Isnew     string `json:"isnew"`
			Isvip     string `json:"isvip"`
			Url       string `json:"url"`
			IsIntact  string `json:"isIntact"`
			Corner    []struct {
				Area  string `json:"area"`
				Flag  string `json:"flag"`
				Color string `json:"color"`
				Font  string `json:"font"`
			} `json:"corner"`
			ClipId      string `json:"clip_id"`
			Time        string `json:"time"`
			ContentType string `json:"contentType"`
			T1          string `json:"t1"`
			T2          string `json:"t2"`
			NextId      string `json:"next_id"`
			T3          string `json:"t3"`
			Ts          string `json:"ts"`
			VideoId     string `json:"video_id"`
		} `json:"list"`
		CurrentPage int `json:"current_page"`
		Info        struct {
			Title string `json:"title"`
			Type  string `json:"type"`
			Isvip string `json:"isvip"`
			Desc  string `json:"desc"`
		} `json:"info"`
	} `json:"data"`
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Seqid     string `json:"seqid"`
	Timestamp int64  `json:"timestamp"`
}
