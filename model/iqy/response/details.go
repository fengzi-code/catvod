package response

type Avlistinfo struct { //电视剧,动漫,儿童,详情api返回的json
	Code string `json:"code"`
	Data struct {
		AlbumId    string `json:"albumId"`
		Epsodelist []struct {
			TvId         int64  `json:"tvId"`
			Description  string `json:"description"`
			Subtitle     string `json:"subtitle"`
			Vid          string `json:"vid"`
			Name         string `json:"name"`
			PlayUrl      string `json:"playUrl"`
			IssueTime    int64  `json:"issueTime"`
			PublishTime  int64  `json:"publishTime"`
			ContentType  int    `json:"contentType"`
			PayMark      int    `json:"payMark"`
			PayMarkUrl   string `json:"payMarkUrl"`
			ImageUrl     string `json:"imageUrl"`
			Duration     string `json:"duration"`
			Period       string `json:"period"`
			Exclusive    bool   `json:"exclusive"`
			Order        int    `json:"order"`
			Effective    bool   `json:"effective"`
			QiyiProduced bool   `json:"qiyiProduced"`
			Focus        string `json:"focus"`
			ShortTitle   string `json:"shortTitle"`
			People       struct {
				Director []struct {
					Id       int    `json:"id"`
					Name     string `json:"name"`
					ImageUrl string `json:"image_url"`
				} `json:"director"`
				MainCharactor []struct {
					Id            int      `json:"id"`
					Name          string   `json:"name"`
					ImageUrl      string   `json:"image_url"`
					Character     []string `json:"character"`
					PaopaoSummary struct {
						CircleId int64 `json:"circle_id"`
					} `json:"paopao_summary,omitempty"`
				} `json:"main_charactor"`
				ScreenWriter []struct {
					Id   int    `json:"id"`
					Name string `json:"name"`
				} `json:"screen_writer"`
			} `json:"people"`
			InteractionType      int      `json:"interactionType"`
			IsEnabledInteraction int      `json:"isEnabledInteraction"`
			ImageSize            []string `json:"imageSize"`
			ImageProductionType  []string `json:"imageProductionType"`
			OrderName            string   `json:"orderName"`
		} `json:"epsodelist"`
		BeforeEpisodeList []interface{} `json:"beforeEpisodeList"`
		AfterEpisodeList  []interface{} `json:"afterEpisodeList"`
		PreEpisodeList    []interface{} `json:"preEpisodeList"`
		StarEpisodeList   []interface{} `json:"starEpisodeList"`
		Updateprevuelist  []interface{} `json:"updateprevuelist"`
		Vipprevuelist     []struct {
			TvId         int64  `json:"tvId"`
			Description  string `json:"description"`
			Subtitle     string `json:"subtitle"`
			Vid          string `json:"vid"`
			Name         string `json:"name"`
			PlayUrl      string `json:"playUrl"`
			IssueTime    int64  `json:"issueTime"`
			PublishTime  int64  `json:"publishTime"`
			ContentType  int    `json:"contentType"`
			PayMark      int    `json:"payMark"`
			PayMarkUrl   string `json:"payMarkUrl"`
			ImageUrl     string `json:"imageUrl"`
			Duration     string `json:"duration"`
			Period       string `json:"period"`
			Exclusive    bool   `json:"exclusive"`
			Order        int    `json:"order"`
			Effective    bool   `json:"effective"`
			QiyiProduced bool   `json:"qiyiProduced"`
			Focus        string `json:"focus"`
			ShortTitle   string `json:"shortTitle"`
			People       struct {
				Director []struct {
					Id       int    `json:"id"`
					Name     string `json:"name"`
					ImageUrl string `json:"image_url"`
				} `json:"director"`
				MainCharactor []struct {
					Id            int      `json:"id"`
					Name          string   `json:"name"`
					ImageUrl      string   `json:"image_url"`
					Character     []string `json:"character"`
					PaopaoSummary struct {
						CircleId int64 `json:"circle_id"`
					} `json:"paopao_summary,omitempty"`
				} `json:"main_charactor"`
			} `json:"people"`
			InteractionType      int      `json:"interactionType"`
			IsEnabledInteraction int      `json:"isEnabledInteraction"`
			ImageSize            []string `json:"imageSize"`
			ImageProductionType  []string `json:"imageProductionType"`
			OrderName            string   `json:"orderName"`
		} `json:"vipprevuelist"`
		PrePrevueList  []interface{} `json:"prePrevueList"`
		StarPrevueList []interface{} `json:"starPrevueList"`
		Size           int           `json:"size"`
		Page           int           `json:"page"`
		Total          int           `json:"total"`
		Part           int           `json:"part"`
		LatestOrder    int           `json:"latestOrder"`
		VideoCount     int           `json:"videoCount"`
		HasMore        bool          `json:"hasMore"`
	} `json:"data"`
}

type Videoinfo struct { //电影api返回的json
	Code string `json:"code"`
	Data struct {
		TvId        int    `json:"tvId"`
		AlbumId     int    `json:"albumId"`
		ChannelId   int    `json:"channelId"`
		Description string `json:"description"`
		Subtitle    string `json:"subtitle"`
		Circle      struct {
			CircleId   int    `json:"circle_id"`
			CircleType string `json:"circle_type"`
		} `json:"circle"`
		RewardAllowed  bool     `json:"rewardAllowed"`
		Vid            string   `json:"vid"`
		Name           string   `json:"name"`
		PlayUrl        string   `json:"playUrl"`
		IssueTime      int64    `json:"issueTime"`
		PublishTime    int64    `json:"publishTime"`
		FeatureAlbumId int      `json:"featureAlbumId"`
		ContentType    int      `json:"contentType"`
		DisplayCircle  bool     `json:"displayCircle"`
		IsAdvance      bool     `json:"isAdvance"`
		PayMark        int      `json:"payMark"`
		PayMarkUrl     string   `json:"payMarkUrl"`
		DisplayUpDown  bool     `json:"displayUpDown"`
		VipType        []string `json:"vipType"`
		ImageUrl       string   `json:"imageUrl"`
		VideoType      string   `json:"videoType"`
		Duration       string   `json:"duration"`
		PpsUrl         string   `json:"ppsUrl"`
		CommentAllowed bool     `json:"commentAllowed"`
		VideoCount     int      `json:"videoCount"`
		LatestOrder    int      `json:"latestOrder"`
		BossMixerAlbum int      `json:"bossMixerAlbum"`
		PublicLevel    string   `json:"publicLevel"`
		StartTime      int      `json:"startTime"`
		EndTime        int      `json:"endTime"`
		Categories     []struct {
			Id             int    `json:"id"`
			QipuId         int    `json:"qipuId"`
			Name           string `json:"name"`
			SubType        int    `json:"subType"`
			LastUpdateTime int64  `json:"lastUpdateTime"`
			Display        int    `json:"display"`
			SubName        string `json:"subName"`
			ParentId       int    `json:"parentId"`
			Url            string `json:"url"`
		} `json:"categories"`
		AlbumName         string        `json:"albumName"`
		BaikeUrl          string        `json:"baikeUrl"`
		AlbumImageUrl     string        `json:"albumImageUrl"`
		UserId            int           `json:"userId"`
		Period            string        `json:"period"`
		Exclusive         bool          `json:"exclusive"`
		Order             int           `json:"order"`
		PlaylistReason    string        `json:"playlistReason"`
		Effective         bool          `json:"effective"`
		PagePublishStatus string        `json:"pagePublishStatus"`
		QiyiProduced      bool          `json:"qiyiProduced"`
		AlbumUrl          string        `json:"albumUrl"`
		SourceId          int           `json:"sourceId"`
		Focus             string        `json:"focus"`
		ShortTitle        string        `json:"shortTitle"`
		Solo              bool          `json:"solo"`
		AlbumFocus        string        `json:"albumFocus"`
		QitanId           int           `json:"qitanId"`
		DownloadAllowed   bool          `json:"downloadAllowed"`
		SeoKeyword        []interface{} `json:"seoKeyword"`
		SupportedDrmTypes []interface{} `json:"supportedDrmTypes"`
		FeatureKeyword    string        `json:"featureKeyword"`
		People            struct {
			Director []struct {
				Id       int    `json:"id"`
				Name     string `json:"name"`
				ImageUrl string `json:"image_url"`
			} `json:"director"`
			MainCharactor []struct {
				Id            int      `json:"id"`
				Name          string   `json:"name"`
				ImageUrl      string   `json:"image_url"`
				Character     []string `json:"character"`
				PaopaoSummary struct {
					CircleId int `json:"circle_id"`
				} `json:"paopao_summary,omitempty"`
			} `json:"main_charactor"`
			ScreenWriter []struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
			} `json:"screen_writer"`
		} `json:"people"`
		StarNumberInfo struct {
			StarTotalNumber int `json:"star_total_number"`
			FiveStarNumber  int `json:"five_star_number"`
			FourStarNumber  int `json:"four_star_number"`
			ThreeStarNumber int `json:"three_star_number"`
			TwoStarNumber   int `json:"two_star_number"`
			OneStarNumber   int `json:"one_star_number"`
		} `json:"starNumberInfo"`
		Cast struct {
		} `json:"cast"`
		VideoFocuses         []interface{} `json:"videoFocuses"`
		DisplayBarrage       bool          `json:"displayBarrage"`
		PreviewImageUrl      string        `json:"previewImageUrl"`
		PublishDate          string        `json:"publishDate"`
		Score                float64       `json:"score"`
		CutLimitStatus       string        `json:"cutLimitStatus"`
		HideGif              bool          `json:"hideGif"`
		PoliticallySensitive bool          `json:"politicallySensitive"`
		InteractionType      int           `json:"interactionType"`
		IsEnabledInteraction int           `json:"isEnabledInteraction"`
		ImageSize            []string      `json:"imageSize"`
		ImageProductionType  []string      `json:"imageProductionType"`
		DynamicImageUrl      string        `json:"dynamicImageUrl"`
		DynamicImageSizeInfo []interface{} `json:"dynamicImageSizeInfo"`
		FormatIssueTime      string        `json:"formatIssueTime"`
		FeatureVideoId       int           `json:"featureVideoId"`
		DurationSec          int           `json:"durationSec"`
		CopyrightStatus      string        `json:"copyrightStatus"`
		OrderName            string        `json:"orderName"`
		PartyTemplateUsed    bool          `json:"partyTemplateUsed"`
		ControlInfo          struct {
			Comment struct {
				DisplayStatus bool `json:"displayStatus"`
				DisplayType   int  `json:"displayType"`
			} `json:"comment"`
			Barrage struct {
				DisplayStatus bool `json:"displayStatus"`
				DisplayType   int  `json:"displayType"`
			} `json:"barrage"`
			UpDown struct {
				DisplayStatus bool `json:"displayStatus"`
				DisplayType   int  `json:"displayType"`
			} `json:"upDown"`
		} `json:"controlInfo"`
		SuperAlbumCollection []struct {
			Id                 int    `json:"id"`
			CollectionType     string `json:"collectionType"`
			CollectionSubType  string `json:"collectionSubType"`
			Region             string `json:"region"`
			LastUpdateTime     int64  `json:"lastUpdateTime"`
			SelectedSuperAlbum bool   `json:"selectedSuperAlbum"`
		} `json:"superAlbumCollection"`
	} `json:"data"`
}

type Tvidinfo struct { //综艺返回的json,取tvid
	Code string `json:"code"`
	Data struct {
		AlbumId        int    `json:"albumId"`
		SourceId       int    `json:"sourceId"`
		Name           string `json:"name"`
		Url            string `json:"url"`
		ImageUrl       string `json:"imageUrl"`
		ChannelId      int    `json:"channelId"`
		Description    string `json:"description"`
		VideoCount     int    `json:"videoCount"`
		LatestOrder    int    `json:"latestOrder"`
		Period         string `json:"period"`
		ContentKeyWord string `json:"contentKeyWord"`
		People         struct {
			Host []struct {
				Id            int      `json:"id"`
				Name          string   `json:"name"`
				ImageUrl      string   `json:"image_url"`
				Character     []string `json:"character"`
				PaopaoSummary struct {
					CircleId int `json:"circle_id"`
				} `json:"paopao_summary"`
			} `json:"host"`
		} `json:"people"`
		Categories []struct {
			Id             int    `json:"id"`
			QipuId         int    `json:"qipuId"`
			Name           string `json:"name"`
			SubType        int    `json:"subType"`
			LastUpdateTime int64  `json:"lastUpdateTime"`
			Display        int    `json:"display"`
			SubName        string `json:"subName"`
			ParentId       int    `json:"parentId"`
			Url            string `json:"url"`
		} `json:"categories"`
		UpdateStrategy  string        `json:"updateStrategy"`
		CommentCount    int           `json:"commentCount"`
		CommentAllowed  int           `json:"commentAllowed"`
		PublishTime     int64         `json:"publishTime"`
		IssueTime       int64         `json:"issueTime"`
		Source          int           `json:"source"`
		HeadImages      []interface{} `json:"headImages"`
		ShortTitle      string        `json:"shortTitle"`
		Subtitle        string        `json:"subtitle"`
		Exclusive       int           `json:"exclusive"`
		QiyiProduced    int           `json:"qiyiProduced"`
		PageStatus      int           `json:"pageStatus"`
		CopyrightStatus int           `json:"copyrightStatus"`
		Effective       int           `json:"effective"`
		Focus           string        `json:"focus"`
		PayMark         int           `json:"payMark"`
		PayMarkUrl      string        `json:"payMarkUrl"`
		Areas           []string      `json:"areas"`
		SuperAlbumId    []struct {
			Id                int    `json:"id"`
			CollectionType    string `json:"collection_type"`
			CollectionSubType string `json:"collection_sub_type"`
			Region            string `json:"region"`
		} `json:"superAlbumId"`
		Circle struct {
			CircleId   int    `json:"circle_id"`
			CircleType string `json:"circle_type"`
		} `json:"circle"`
		FirstVideoId   int64  `json:"firstVideoId"`
		LatestVideoId  int64  `json:"latestVideoId"`
		FirstVideoUrl  string `json:"firstVideoUrl"`
		LatestVideoUrl string `json:"latestVideoUrl"`
		FirstVideo     struct {
			TvId        int64  `json:"tvId"`
			Description string `json:"description"`
			Subtitle    string `json:"subtitle"`
			Circle      struct {
				CircleId   int    `json:"circle_id"`
				CircleType string `json:"circle_type"`
			} `json:"circle"`
			RewardAllowed bool   `json:"rewardAllowed"`
			Vid           string `json:"vid"`
			Name          string `json:"name"`
			PlayUrl       string `json:"playUrl"`
			IssueTime     int64  `json:"issueTime"`
			PublishTime   int64  `json:"publishTime"`
			PayMark       int    `json:"payMark"`
			PayMarkUrl    string `json:"payMarkUrl"`
			ImageUrl      string `json:"imageUrl"`
			Duration      string `json:"duration"`
			Period        string `json:"period"`
			Order         int    `json:"order"`
			Focus         string `json:"focus"`
			ShortTitle    string `json:"shortTitle"`
			QitanId       int    `json:"qitanId"`
			People        struct {
				Host []struct {
					Id            int      `json:"id"`
					Name          string   `json:"name"`
					ImageUrl      string   `json:"image_url"`
					Character     []string `json:"character"`
					PaopaoSummary struct {
						CircleId int `json:"circle_id"`
					} `json:"paopao_summary"`
				} `json:"host"`
				Guest []struct {
					Id            int      `json:"id"`
					Name          string   `json:"name"`
					ImageUrl      string   `json:"image_url"`
					Character     []string `json:"character"`
					PaopaoSummary struct {
						CircleId int `json:"circle_id"`
					} `json:"paopao_summary,omitempty"`
				} `json:"guest"`
			} `json:"people"`
			ImageSize           []string `json:"imageSize"`
			ImageProductionType []string `json:"imageProductionType"`
		} `json:"firstVideo"`
		LatestVideo struct {
			TvId        int64  `json:"tvId"`
			Description string `json:"description"`
			Subtitle    string `json:"subtitle"`
			Circle      struct {
				CircleId   int    `json:"circle_id"`
				CircleType string `json:"circle_type"`
			} `json:"circle"`
			RewardAllowed bool   `json:"rewardAllowed"`
			Vid           string `json:"vid"`
			Name          string `json:"name"`
			PlayUrl       string `json:"playUrl"`
			IssueTime     int64  `json:"issueTime"`
			PublishTime   int64  `json:"publishTime"`
			PayMark       int    `json:"payMark"`
			PayMarkUrl    string `json:"payMarkUrl"`
			ImageUrl      string `json:"imageUrl"`
			Duration      string `json:"duration"`
			Period        string `json:"period"`
			Order         int    `json:"order"`
			Focus         string `json:"focus"`
			ShortTitle    string `json:"shortTitle"`
			QitanId       int    `json:"qitanId"`
			People        struct {
				Host []struct {
					Id            int      `json:"id"`
					Name          string   `json:"name"`
					ImageUrl      string   `json:"image_url"`
					Character     []string `json:"character"`
					PaopaoSummary struct {
						CircleId int `json:"circle_id"`
					} `json:"paopao_summary"`
				} `json:"host"`
			} `json:"people"`
			ImageSize           []string `json:"imageSize"`
			ImageProductionType []string `json:"imageProductionType"`
		} `json:"latestVideo"`
		FeatureKeyword string   `json:"featureKeyword"`
		BossStatus     string   `json:"bossStatus"`
		BossMixer      bool     `json:"bossMixer"`
		QitanId        int      `json:"qitanId"`
		Season         int      `json:"season"`
		Score          float64  `json:"score"`
		UserId         int      `json:"userId"`
		ImageSize      []string `json:"imageSize"`
	} `json:"data"`
}

type ListByNumber struct { //综艺最终返回的详情
	Code string `json:"code"`
	Data []struct {
		TvId            int64  `json:"tvId"`
		Description     string `json:"description"`
		Subtitle        string `json:"subtitle"`
		Vid             string `json:"vid"`
		Name            string `json:"name"`
		PlayUrl         string `json:"playUrl"`
		IssueTime       int64  `json:"issueTime"`
		ContentType     int    `json:"contentType"`
		PayMark         int    `json:"payMark"`
		PayMarkUrl      string `json:"payMarkUrl"`
		ImageUrl        string `json:"imageUrl"`
		Duration        string `json:"duration"`
		AlbumImageUrl   string `json:"albumImageUrl"`
		Period          string `json:"period"`
		Exclusive       bool   `json:"exclusive"`
		Order           int    `json:"order"`
		QiyiProduced    bool   `json:"qiyiProduced"`
		Focus           string `json:"focus"`
		ShortTitle      string `json:"shortTitle"`
		DownloadAllowed string `json:"downloadAllowed"`
		People          struct {
			Host []struct {
				Id            int      `json:"id"`
				Name          string   `json:"name"`
				ImageUrl      string   `json:"image_url"`
				Character     []string `json:"character"`
				PaopaoSummary struct {
					CircleId int `json:"circle_id"`
				} `json:"paopao_summary"`
			} `json:"host"`
			Guest []struct {
				Id            int      `json:"id"`
				Name          string   `json:"name"`
				ImageUrl      string   `json:"image_url"`
				Character     []string `json:"character"`
				PaopaoSummary struct {
					CircleId int `json:"circle_id"`
				} `json:"paopao_summary,omitempty"`
			} `json:"guest,omitempty"`
		} `json:"people"`
		VideoFocuses []struct {
			Id               int    `json:"id"`
			Description      string `json:"description"`
			Time             int    `json:"time"`
			ImageUrl         string `json:"image_url"`
			StartTimeSeconds int    `json:"start_time_seconds"`
		} `json:"videoFocuses"`
		ImageSize           []string `json:"imageSize"`
		ImageProductionType []string `json:"imageProductionType"`
	} `json:"data"`
}
