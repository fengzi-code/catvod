package response

type YoukuCategory1 struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		CategoryVideos []struct {
			SummaryType  string `json:"summaryType"`
			Access       string `json:"access"`
			Type         string `json:"type"`
			ShowThumb    string `json:"showThumb"`
			Img          string `json:"img"`
			Summary      string `json:"summary"`
			Title        string `json:"title"`
			SubTitle     string `json:"subTitle"`
			RightTagText string `json:"rightTagText"`
			VideoId      string `json:"videoId"`
			VideoLink    string `json:"videoLink"`
		} `json:"categoryVideos"`
		CategoryFilter struct {
			Category []struct {
				Name        string `json:"name"`
				Key         string `json:"key"`
				Type        string `json:"type,omitempty"`
				ChannelPage string `json:"channelPage,omitempty"`
				Hide        bool   `json:"hide,omitempty"`
			} `json:"category"`
			Filter struct {
				Filter []struct {
					Label string `json:"label"`
					Key   string `json:"key"`
					List  []struct {
						Name string      `json:"name"`
						Key  interface{} `json:"key"`
						Data struct {
							Spm string `json:"spm"`
							Scm string `json:"scm"`
						} `json:"data,omitempty"`
						Hoverd bool `json:"hoverd,omitempty"`
					} `json:"list,omitempty"`
					Type string `json:"type,omitempty"`
				} `json:"filter"`
			} `json:"filter"`
			CommonFilter struct {
				R []struct {
					Name string      `json:"name"`
					Key  interface{} `json:"key"`
				} `json:"r"`
				U []struct {
					Name string      `json:"name"`
					Key  interface{} `json:"key"`
				} `json:"u"`
				Pt []struct {
					Name string      `json:"name"`
					Key  interface{} `json:"key"`
				} `json:"pt"`
				S []struct {
					Name string `json:"name"`
					Key  int    `json:"key"`
				} `json:"s"`
				Vip []struct {
					Name string `json:"name"`
					Key  int    `json:"key"`
					Href string `json:"href"`
				} `json:"vip"`
			} `json:"commonFilter"`
			Url         string `json:"url"`
			CateName    string `json:"cateName"`
			ListType    string `json:"listType"`
			ChannelPage string `json:"channelPage"`
			CateKey     string `json:"cateKey"`
		} `json:"categoryFilter"`
		Theme     string `json:"theme"`
		AdContent []struct {
			ModuleId       int    `json:"moduleId"`
			Type           string `json:"type"`
			TabType        string `json:"tabType"`
			IsHiddenHeader bool   `json:"isHiddenHeader"`
			IsLayoutHori   bool   `json:"isLayoutHori"`
			IsIgnoreBottom bool   `json:"isIgnoreBottom"`
			ContainerType  string `json:"containerType"`
			ShowContainer  bool   `json:"showContainer"`
			MainTitleLinks []struct {
				Title        string      `json:"title"`
				Logo         string      `json:"logo"`
				IsOpenNewTab bool        `json:"isOpenNewTab"`
				Link         interface{} `json:"link"`
				Spm          string      `json:"spm"`
				Scm          string      `json:"scm"`
				TrackInfo    interface{} `json:"trackInfo"`
			} `json:"mainTitleLinks"`
			SubTitleLinks    []interface{} `json:"subTitleLinks"`
			TabTitle         []interface{} `json:"tabTitle"`
			SubRightLinkInfo interface{}   `json:"subRightLinkInfo"`
			IsdispUnderLine  bool          `json:"isdisp_under_line"`
			ChangeButtonInfo interface{}   `json:"changeButtonInfo"`
			AdId             string        `json:"adId"`
			AdType           int           `json:"adType"`
			IsAd             bool          `json:"isAd"`
		} `json:"adContent"`
	} `json:"data"`
	Code    interface{} `json:"code"`
	JumpUrl interface{} `json:"jumpUrl"`
}

type YoukuCategory struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		FilterData struct {
			Param struct {
				Type string `json:"type"`
			} `json:"param"`
			ListData []struct {
				SummaryType  string `json:"summaryType"`
				Type         string `json:"type"`
				Img          string `json:"img"`
				Summary      string `json:"summary"`
				Title        string `json:"title"`
				SubTitle     string `json:"subTitle,omitempty"`
				VideoLink    string `json:"videoLink"`
				RightTagText string `json:"rightTagText,omitempty"`
			} `json:"listData"`
			Session struct {
				SubIndex  int `json:"subIndex"`
				TrackInfo struct {
					Parentdrawerid string `json:"parentdrawerid"`
				} `json:"trackInfo"`
				SpmA     string `json:"spmA"`
				Level    int    `json:"level"`
				SpmC     string `json:"spmC"`
				SpmB     string `json:"spmB"`
				Index    int    `json:"index"`
				PageName string `json:"pageName"`
				Scene    string `json:"scene"`
				ScmB     string `json:"scmB"`
				Path     string `json:"path"`
				ScmA     string `json:"scmA"`
				ScmC     string `json:"scmC"`
				From     string `json:"from"`
				Id       int    `json:"id"`
				Category string `json:"category"`
			} `json:"session"`
			HasMore bool `json:"hasMore"`
		} `json:"filterData"`
		Theme string `json:"theme"`
	} `json:"data"`
	Code    interface{} `json:"code"`
	JumpUrl interface{} `json:"jumpUrl"`
}
