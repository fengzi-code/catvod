package mgtv

type CategoryList struct {
	VodId      string `json:"vod_id"`
	VodName    string `json:"vod_name"`
	VodPic     string `json:"vod_pic"`
	VodRemarks string `json:"vod_remarks"`
}

type Category struct {
	Code int `json:"code"`
	Data struct {
		Template  string `json:"template"`
		TotalHits int    `json:"totalHits"`
		Advs      []struct {
			Row string `json:"row"`
			Id  string `json:"id"`
		} `json:"advs"`
		HitDocs []struct {
			Schema      string `json:"schema"`
			UpdateInfo  string `json:"updateInfo"`
			JumpKind    string `json:"jumpKind"`
			Img         string `json:"img"`
			RightCorner struct {
				Color string `json:"color,omitempty"`
				Text  string `json:"text,omitempty"`
			} `json:"rightCorner"`
			Year         string   `json:"year"`
			PartId       string   `json:"partId"`
			Kind         []string `json:"kind"`
			Prefix       string   `json:"prefix"`
			FstlvlId     string   `json:"fstlvlId"`
			Title        string   `json:"title"`
			Params       string   `json:"params"`
			SeUpdateTime string   `json:"se_updateTime"`
			PlayPartId   string   `json:"playPartId"`
			Subtitle     string   `json:"subtitle"`
			Rpt          string   `json:"rpt"`
			ClipId       string   `json:"clipId"`
			Ic           string   `json:"ic"`
			ZhihuScore   string   `json:"zhihuScore"`
			Views        string   `json:"views"`
			ImgUrlH      string   `json:"imgUrlH"`
			Story        string   `json:"story"`
		} `json:"hitDocs"`
		SearchKeys []interface{} `json:"searchKeys"`
	} `json:"data"`
	Msg   string `json:"msg"`
	Seqid string `json:"seqid"`
}

type Detail struct {
	Data struct {
		CategoryList []struct {
			Area         string  `json:"area"`
			PageRefresh  int     `json:"pageRefresh"`
			PlayNextType int     `json:"playNextType"`
			QuickLocate  int     `json:"quickLocate"`
			DataType     int     `json:"dataType"`
			NextDataType string  `json:"nextDataType"`
			PlayOrder    int     `json:"playOrder"`
			ObjectType   int     `json:"objectType"`
			DisplayType  int     `json:"displayType"`
			IsRefresh    int     `json:"isRefresh"`
			SkipDataType string  `json:"skipDataType"`
			Ltitle       string  `json:"ltitle"`
			IsMore       int     `json:"isMore"`
			Url          *string `json:"url"`
		} `json:"categoryList"`
		Info struct {
			PlImage      string `json:"plImage"`
			PlId         string `json:"plId"`
			FstlvlId     string `json:"fstlvlId"`
			VideoId      string `json:"videoId"`
			RelativeVid  string `json:"relativeVid"`
			ShowMode     int    `json:"showMode"`
			ClipName     string `json:"clipName"`
			Puuid        string `json:"puuid"`
			Title        string `json:"title"`
			SeriesId     string `json:"seriesId"`
			ClipImage    string `json:"clipImage"`
			PlName       string `json:"plName"`
			IsIntact     string `json:"isIntact"`
			VclassId     string `json:"vclassId"`
			VideoName    string `json:"videoName"`
			ClipId       string `json:"clipId"`
			PuuidType    string `json:"puuid_type"`
			ContentType  string `json:"contentType"`
			Uploadid     string `json:"uploadid"`
			FstlvlType   string `json:"fstlvlType"`
			VideoImage   string `json:"videoImage"`
			PlayPriority string `json:"playPriority"`
			Time         string `json:"time"`
			Detail       struct {
				Area        string `json:"area"`
				UpdateInfo  string `json:"updateInfo"`
				Leader      string `json:"leader"`
				PlayCnt     string `json:"playCnt"`
				ReleaseTime string `json:"releaseTime"`
				Presenter   string `json:"presenter"`
				Kind        string `json:"kind"`
				Director    string `json:"director"`
				Television  string `json:"television"`
				Language    string `json:"language"`
				Url         string `json:"url"`
				ClipUrl     string `json:"clipUrl"`
				PlayRight   string `json:"playRight"`
				FstlvlType  string `json:"fstlvlType"`
				Story       string `json:"story"`
			} `json:"detail"`
			ClipImage2 string `json:"clipImage2"`
			Desc       string `json:"desc"`
		} `json:"info"`
	} `json:"data"`
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Seqid     string `json:"seqid"`
	Timestamp int64  `json:"timestamp"`
}
