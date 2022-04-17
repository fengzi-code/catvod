package response

type IqiyiSoJson struct {
	Code string `json:"code"`
	Data struct {
		Base struct {
			IsOnline bool          `json:"isOnline"`
			Block    string        `json:"block"`
			Chunk    []interface{} `json:"chunk"`
		} `json:"base"`
		FormatData struct {
			HotList []string `json:"hotList"`
			List    []struct {
				Id           int64         `json:"id"`
				Cid          int           `json:"cid"`
				Idx          int           `json:"idx"`
				VideoDocType int           `json:"videoDocType"`
				SiteId       string        `json:"siteId"`
				SiteName     string        `json:"siteName"`
				StopPlayTime interface{}   `json:"stopPlayTime"`
				GTitle       string        `json:"g_title"`
				GMainLink    string        `json:"g_main_link"`
				Desc         string        `json:"desc"`
				GImg         string        `json:"g_img"`
				GImgSize     string        `json:"g_img_size"`
				GMeta        string        `json:"g_meta"`
				Region       string        `json:"region"`
				Director     []interface{} `json:"director"`
				ReleaseTime  string        `json:"releaseTime"`
				PublishTime  string        `json:"publishTime,omitempty"`
				Uploader     struct {
					Name string `json:"name"`
					Id   int64  `json:"id"`
					Mark int    `json:"mark"`
				} `json:"uploader,omitempty"`
				MoreLink      string `json:"moreLink"`
				IsSource      bool   `json:"isSource"`
				IsSeries      bool   `json:"isSeries"`
				IsLive        bool   `json:"isLive"`
				IsMoreList    bool   `json:"isMoreList"`
				DocId         string `json:"doc_id"`
				STRATEGYORDER int    `json:"STRATEGY_ORDER,omitempty"`
				Year          string `json:"year,omitempty"`
				Videoinfos    []struct {
					Date          string `json:"date"`
					Name          string `json:"name"`
					Url           string `json:"url"`
					Order         int    `json:"order"`
					Id            int64  `json:"id"`
					IsPay         bool   `json:"isPay"`
					STRATEGYORDER int    `json:"STRATEGY_ORDER,omitempty"`
				} `json:"videoinfos,omitempty"`
				Reason string `json:"reason,omitempty"`
			} `json:"list"`
			SearchTime int           `json:"search_time"`
			SameStar   []interface{} `json:"sameStar"`
			IntentStar []interface{} `json:"intentStar"`
			IntentList []struct {
				Id            int64  `json:"id"`
				Cid           int    `json:"cid"`
				GTitle        string `json:"g_title"`
				GMainLink     string `json:"g_main_link"`
				GImg          string `json:"g_img"`
				GImgSize      string `json:"g_img_size"`
				GCornerMark   string `json:"g_corner_mark,omitempty"`
				GMeta         string `json:"g_meta"`
				Desc          string `json:"desc"`
				DocId         string `json:"doc_id"`
				STRATEGYORDER int    `json:"STRATEGY_ORDER,omitempty"`
			} `json:"intentList"`
			IntentTotalCount int `json:"intentTotalCount"`
			TermQuery        struct {
				Mode       int      `json:"mode"`
				Site       []string `json:"site"`
				DataType   []string `json:"data_type"`
				EntityName []string `json:"entity_name"`
				GraphType  []string `json:"graph_type"`
				RealQuery  []string `json:"real_query"`
			} `json:"termQuery"`
			GraphCategories []struct {
				Desc  string `json:"desc"`
				Key   string `json:"key"`
				Value []struct {
					Id   string `json:"id"`
					Name string `json:"name"`
				} `json:"value"`
			} `json:"graphCategories"`
			ResultNum int      `json:"resultNum"`
			Terms     []string `json:"terms"`
			RealQuery string   `json:"realQuery"`
			GraphType struct {
				Type             int `json:"type"`
				IntentSubType    int `json:"intent_sub_type"`
				IntentActionType int `json:"intent_action_type"`
			} `json:"graphType"`
			Code       string `json:"code"`
			EventId    string `json:"eventId"`
			IndexLayer int    `json:"index_layer"`
			Bkt        string `json:"bkt"`
			SideParams struct {
				StarVideoId  int64  `json:"starVideoId"`
				HotQueryType string `json:"hot_query_type"`
			} `json:"sideParams"`
		} `json:"formatData"`
	} `json:"data"`
}

type IqiyiSoListJson []struct {
	GTitle    string `json:"g_title"`
	GMainLink string `json:"g_main_link"`
	GImg      string `json:"g_img"`
}
