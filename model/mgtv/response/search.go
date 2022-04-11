package response

type SoStruct struct { // mgtv 搜索页 返回的json
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
				Desc []struct {
					Label string `json:"label"`
					Text  string `json:"text"`
					Url   string `json:"url"`
				} `json:"desc,omitempty"`
				FantuanUrl     string `json:"fantuanUrl,omitempty"`
				HitTitle       string `json:"hitTitle,omitempty"`
				Pic            string `json:"pic,omitempty"`
				Vid            string `json:"vid,omitempty"`
				PlayTime       string `json:"playTime,omitempty"`
				RightTopCorner struct {
					Color string `json:"color,omitempty"`
					Text  string `json:"text,omitempty"`
				} `json:"rightTopCorner,omitempty"`
				Rpt        string `json:"rpt"`
				Score      string `json:"score,omitempty"`
				SourceList []struct {
					JumpKind  string `json:"jumpKind"`
					Name      string `json:"name"`
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
				} `json:"sourceList,omitempty"`
				Story   string `json:"story,omitempty"`
				Tablist []struct {
					Data []struct {
						Pic         string `json:"pic"`
						RightCorner struct {
						} `json:"rightCorner"`
						Subtitle   string `json:"subtitle"`
						Title      string `json:"title"`
						UpdateInfo string `json:"updateInfo"`
						Url        string `json:"url"`
						Vid        string `json:"vid"`
						VideoNum   string `json:"videoNum"`
					} `json:"data"`
					ImgType string `json:"imgType"`
					Label   string `json:"label"`
					TabId   string `json:"tabId"`
				} `json:"tablist"`
				Title      string        `json:"title,omitempty"`
				Uuid       string        `json:"uuid,omitempty"`
				Year       string        `json:"year,omitempty"`
				FmsHistory []interface{} `json:"fmsHistory,omitempty"`
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
				Url string `json:"url,omitempty"`
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
