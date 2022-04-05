package mgtv

type RequestCategory struct {
	Page      int `json:"page"`
	PageCount int `json:"pagecount"`
	Limit     int `json:"limit"`
	Total     int `json:"total"`
}

type RequestDetail struct {
	VodId       string `json:"vod_id"`
	VodName     string `json:"vod_name"`
	VodPic      string `json:"vod_pic"`
	TypeName    string `json:"type_name"`
	VodYear     string `json:"vod_year"`
	VodArea     string `json:"vod_area"`
	VodRemarks  string `json:"vod_remarks"`
	VodActor    string `json:"vod_actor"`
	VodDirector string `json:"vod_director"`
	VodContent  string `json:"vod_content"`
	VodPlayFrom string `json:"vod_play_from"`
	VodPlayUrl  string `json:"vod_play_url"`
}

type RequestDetails []RequestDetail

type RequestPlayUrl struct {
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

type SoStruct struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Cost int    `json:"cost"`
	Data struct {
		RelativeArtists interface{} `json:"relativeArtists"`
		ListItems       []struct {
			Items []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"items"`
			Name  string `json:"name"`
			Param string `json:"param"`
		} `json:"listItems"`
		Pc       int `json:"pc"`
		Contents []struct {
			Data struct {
				FantuanUrl string        `json:"fantuanUrl,omitempty"`
				FmsHistory []interface{} `json:"fmsHistory,omitempty"`
				Rpt        string        `json:"rpt"`
				Tablist    []interface{} `json:"tablist,omitempty"`
				YearList   []struct {
					Desc []struct {
						Label string `json:"label"`
						Text  string `json:"text"`
						Url   string `json:"url"`
					} `json:"desc"`
					HitTitle       string `json:"hitTitle"`
					Pic            string `json:"pic"`
					PlayTime       string `json:"playTime"`
					RightTopCorner struct {
					} `json:"rightTopCorner"`
					Selected   bool `json:"selected"`
					SourceList []struct {
						HasMore   bool   `json:"hasMore"`
						JumpKind  string `json:"jumpKind"`
						MoreUrl   string `json:"moreUrl"`
						Name      string `json:"name"`
						PlayUrl   string `json:"playUrl"`
						Rpt       string `json:"rpt"`
						Source    string `json:"source"`
						Url       string `json:"url"`
						Vid       string `json:"vid"`
						VideoList []struct {
							Pic         string `json:"pic"`
							RightCorner struct {
								Color string `json:"color,omitempty"`
								Text  string `json:"text,omitempty"`
							} `json:"rightCorner"`
							Subtitle   string `json:"subtitle"`
							Title      string `json:"title"`
							UpdateInfo string `json:"updateInfo"`
							Url        string `json:"url"`
							Vid        string `json:"vid"`
							VideoNum   string `json:"videoNum"`
						} `json:"videoList"`
					} `json:"sourceList"`
					Story   string `json:"story"`
					TabId   string `json:"tabId"`
					TabName string `json:"tabName"`
					Title   string `json:"title"`
					Year    string `json:"year"`
				} `json:"yearList,omitempty"`
				BtnText string `json:"btnText,omitempty"`
				Desc    []struct {
					Label string `json:"label"`
					Text  string `json:"text"`
					Url   string `json:"url"`
				} `json:"desc,omitempty"`
				HitTitle       string `json:"hitTitle,omitempty"`
				Pic            string `json:"pic,omitempty"`
				RightTopCorner struct {
				} `json:"rightTopCorner,omitempty"`
				Story      string `json:"story,omitempty"`
				SubTitle   string `json:"subTitle,omitempty"`
				Title      string `json:"title,omitempty"`
				UpdateInfo string `json:"updateInfo,omitempty"`
				Url        string `json:"url,omitempty"`
				Vid        string `json:"vid,omitempty"`
			} `json:"data"`
			Idx   int    `json:"idx"`
			Name  string `json:"name"`
			Title string `json:"title"`
			Type  string `json:"type"`
		} `json:"contents"`
		TotalHits int    `json:"totalHits"`
		Query     string `json:"query"`
		Rpt       string `json:"rpt"`
		HasMore   bool   `json:"hasMore"`
		TopList   []struct {
			Data []struct {
				Icon  string `json:"icon"`
				Index int    `json:"index"`
				Name  string `json:"name"`
				Rpt   string `json:"rpt"`
			} `json:"data"`
			Label string `json:"label"`
		} `json:"topList"`
		RelativeStars interface{} `json:"relativeStars"`
		Qcorr         interface{} `json:"qcorr"`
	} `json:"data"`
	Pt         interface{} `json:"pt"`
	Ty         interface{} `json:"ty"`
	Sort       interface{} `json:"sort"`
	Version    string      `json:"version"`
	Du         interface{} `json:"du"`
	Port       int         `json:"port"`
	FmsId      interface{} `json:"fmsId"`
	ServerTime int64       `json:"serverTime"`
	Seqid      string      `json:"seqid"`
}
