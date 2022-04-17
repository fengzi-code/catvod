package response

type SearchTotalData struct {
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
	BoxId     string `json:"boxId"`
	BoxUID    string `json:"boxUID"`
	TotalNum  int    `json:"totalNum"`
	Filters   struct {
		FirstLine   interface{}   `json:"firstLine"`
		HigherLines []interface{} `json:"higherLines"`
	} `json:"filters"`
	BoxTitle struct {
		Type   int    `json:"type"`
		Title  string `json:"title"`
		UiType int    `json:"uiType"`
	} `json:"boxTitle"`
	PageContext   string `json:"pageContext"`
	SearchSession string `json:"searchSession"`
	ItemList      []struct {
		Doc struct {
			DataType     int           `json:"dataType"`
			Md           string        `json:"md"`
			Id           string        `json:"id"`
			ShowReport   interface{}   `json:"ShowReport"`
			ClickReport  interface{}   `json:"ClickReport"`
			AreaBoxIndex []interface{} `json:"areaBoxIndex"`
		} `json:"doc"`
		VideoInfo struct {
			VideoDoc interface{} `json:"videoDoc"`
			CoverDoc struct {
				TimeLong          int           `json:"timeLong"`
				ByName            []string      `json:"byName"`
				EnglishName       string        `json:"englishName"`
				Score             float64       `json:"score"`
				EditVersion       string        `json:"edit_version"`
				Roles             []string      `json:"roles"`
				SecondTitle       string        `json:"secondTitle"`
				TopicName         string        `json:"topicName"`
				ColumnTitle       string        `json:"columnTitle"`
				Stars             []interface{} `json:"stars"`
				ResolutionList    []string      `json:"resolution_list"`
				PositiveContentID string        `json:"positiveContentID"`
			} `json:"coverDoc"`
			ColumnDoc       interface{} `json:"columnDoc"`
			SubjectDoc      interface{} `json:"subjectDoc"`
			ViewType        int         `json:"viewType"`
			VideoType       int         `json:"videoType"`
			Title           string      `json:"title"`
			TypeName        string      `json:"typeName"`
			Year            string      `json:"year"`
			CheckupTime     string      `json:"checkupTime"`
			ImgTag          interface{} `json:"imgTag"`
			Area            string      `json:"area"`
			Language        []string    `json:"language"`
			Directors       []string    `json:"directors"`
			Actors          []string    `json:"actors"`
			FirstBlockSites []struct {
				UiType               int    `json:"uiType"`
				ShowName             string `json:"showName"`
				EnName               string `json:"enName"`
				IconUrl              string `json:"iconUrl"`
				FloatingLayerIconUrl string `json:"floatingLayerIconUrl"`
				PayType              int    `json:"payType"`
				Tabs                 []struct {
					Title       string `json:"title"`
					AsnycParams string `json:"asnycParams"`
					IsSelected  bool   `json:"isSelected"`
				} `json:"tabs"`
				PlaysoureType   int `json:"playsoureType"`
				IsDanger        int `json:"isDanger"`
				TotalEpisode    int `json:"totalEpisode"`
				EpisodeInfoList []struct {
					Id          string `json:"id"`
					DataType    int    `json:"dataType"`
					Url         string `json:"url"`
					Title       string `json:"title"`
					MarkLabel   string `json:"markLabel"`
					AsnycParams string `json:"asnycParams"`
					DisplayType int    `json:"displayType"`
					CheckUpTime string `json:"checkUpTime"`
					ImgUrl      string `json:"imgUrl"`
				} `json:"episodeInfoList"`
			} `json:"firstBlockSites"`
			SecondBlockSite interface{} `json:"secondBlockSite"`
			Descrip         string      `json:"descrip"`
			Views           string      `json:"views"`
			ImgUrl          string      `json:"imgUrl"`
			Url             string      `json:"url"`
			RecommendInfo   interface{} `json:"recommendInfo"`
			SubTitle        string      `json:"subTitle"`
			PayStatus       int         `json:"payStatus"`
		} `json:"videoInfo"`
		Live                   interface{} `json:"live"`
		StaticRelateSearchWord interface{} `json:"staticRelateSearchWord"`
		StarCard               interface{} `json:"starCard"`
		CpResult               interface{} `json:"cpResult"`
		DocReportInfo          struct {
			ReportInfo struct {
				TriggerIds string `json:"trigger_ids"`
			} `json:"report_info"`
		} `json:"doc_report_info"`
	} `json:"itemList"`
	ViewType int `json:"viewType"`
	PageSize int `json:"pageSize"`
}

type SearchResult struct {
	TotalData SearchTotalData `json:"totalData"`
}

type SearchTotalData2 struct {
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
	BoxId     string `json:"boxId"`
	BoxUID    string `json:"boxUID"`
	TotalNum  int    `json:"totalNum"`
	Filters   struct {
		FirstLine   interface{}   `json:"firstLine"`
		HigherLines []interface{} `json:"higherLines"`
	} `json:"filters"`
	BoxTitle struct {
		Type   int    `json:"type"`
		Title  string `json:"title"`
		UiType int    `json:"uiType"`
	} `json:"boxTitle"`
	PageContext   string `json:"pageContext"`
	SearchSession string `json:"searchSession"`
	ItemList      []struct {
		Doc struct {
			DataType     int           `json:"dataType"`
			Md           string        `json:"md"`
			Id           string        `json:"id"`
			ShowReport   interface{}   `json:"ShowReport"`
			ClickReport  interface{}   `json:"ClickReport"`
			AreaBoxIndex []interface{} `json:"areaBoxIndex"`
		} `json:"doc"`
		VideoInfo struct {
			VideoDoc interface{} `json:"videoDoc"`
			CoverDoc struct {
				TimeLong          int           `json:"timeLong"`
				ByName            []string      `json:"byName"`
				EnglishName       string        `json:"englishName"`
				Score             float64       `json:"score"`
				EditVersion       string        `json:"edit_version"`
				Roles             []interface{} `json:"roles"`
				SecondTitle       string        `json:"secondTitle"`
				TopicName         string        `json:"topicName"`
				ColumnTitle       string        `json:"columnTitle"`
				Stars             []interface{} `json:"stars"`
				ResolutionList    []string      `json:"resolution_list"`
				PositiveContentID string        `json:"positiveContentID"`
			} `json:"coverDoc"`
			ColumnDoc       interface{} `json:"columnDoc"`
			SubjectDoc      interface{} `json:"subjectDoc"`
			ViewType        int         `json:"viewType"`
			VideoType       int         `json:"videoType"`
			Title           string      `json:"title"`
			TypeName        string      `json:"typeName"`
			Year            string      `json:"year"`
			CheckupTime     string      `json:"checkupTime"`
			ImgTag          interface{} `json:"imgTag"`
			Area            string      `json:"area"`
			Language        []string    `json:"language"`
			Directors       []string    `json:"directors"`
			Actors          []string    `json:"actors"`
			FirstBlockSites []struct {
				UiType               int    `json:"uiType"`
				ShowName             string `json:"showName"`
				EnName               string `json:"enName"`
				IconUrl              string `json:"iconUrl"`
				FloatingLayerIconUrl string `json:"floatingLayerIconUrl"`
				PayType              int    `json:"payType"`
				Tabs                 []struct {
					Title       string `json:"title"`
					AsnycParams string `json:"asnycParams"`
					IsSelected  bool   `json:"isSelected"`
				} `json:"tabs"`
				PlaysoureType   int `json:"playsoureType"`
				IsDanger        int `json:"isDanger"`
				TotalEpisode    int `json:"totalEpisode"`
				EpisodeInfoList []struct {
					Id          string `json:"id"`
					DataType    int    `json:"dataType"`
					Url         string `json:"url"`
					Title       string `json:"title"`
					MarkLabel   string `json:"markLabel"`
					AsnycParams string `json:"asnycParams"`
					DisplayType int    `json:"displayType"`
					CheckUpTime string `json:"checkUpTime"`
					ImgUrl      string `json:"imgUrl"`
				} `json:"episodeInfoList"`
			} `json:"firstBlockSites"`
			SecondBlockSite interface{} `json:"secondBlockSite"`
			Descrip         string      `json:"descrip"`
			Views           string      `json:"views"`
			ImgUrl          string      `json:"imgUrl"`
			Url             string      `json:"url"`
			RecommendInfo   *struct {
				ShowTitle   string `json:"showTitle"`
				RequestInfo string `json:"requestInfo"`
				IntentId    string `json:"intentId"`
			} `json:"recommendInfo"`
			SubTitle  string `json:"subTitle"`
			PayStatus int    `json:"payStatus"`
		} `json:"videoInfo"`
		Live                   interface{} `json:"live"`
		StaticRelateSearchWord interface{} `json:"staticRelateSearchWord"`
		StarCard               interface{} `json:"starCard"`
		CpResult               interface{} `json:"cpResult"`
		DocReportInfo          struct {
			ReportInfo struct {
				TriggerIds string `json:"trigger_ids"`
			} `json:"report_info"`
		} `json:"doc_report_info"`
	} `json:"itemList"`
	ViewType int `json:"viewType"`
	PageSize int `json:"pageSize"`
}

type ImgTag struct {
	Tag1 struct {
		Id    string `json:"id"`
		Param string `json:"param"`
		Text  string `json:"text"`
	} `json:"tag_1"`
	Tag2 struct {
		Id    string `json:"id"`
		Param struct {
			X  string `json:"1X"`
			X1 string `json:"2X"`
		} `json:"param"`
		Text string `json:"text"`
	} `json:"tag_2"`
	Tag3 struct {
		Id    string `json:"id"`
		Param string `json:"param"`
		Text  string `json:"text"`
	} `json:"tag_3"`
	Tag4 struct {
		Id    string `json:"id"`
		Param string `json:"param"`
		Text  string `json:"text"`
	} `json:"tag_4"`
}
