package response

type CoverInfo struct {
	Id                     string      `json:"id"`
	Typeid                 int         `json:"typeid"`
	DirectorId             []string    `json:"director_id"`
	CommentShowType        int         `json:"comment_show_type"`
	CoverCheckupGrade      int         `json:"cover_checkup_grade"`
	Brief                  interface{} `json:"brief"`
	CopyrightNegotiationId interface{} `json:"copyright_negotiation_id"`
	PrePlayTimelongId      interface{} `json:"pre_play_timelong_id"`
	Subtype                []string    `json:"subtype"`
	Score                  struct {
		Hot       string `json:"hot"`
		Score     string `json:"score"`
		CMixScore string `json:"c_mix_score"`
	} `json:"score"`
	AreaName          string        `json:"area_name"`
	PcVipPullUrl      interface{}   `json:"pc_vip_pull_url"`
	Langue            string        `json:"langue"`
	Tag               interface{}   `json:"tag"`
	VideoIds          []string      `json:"video_ids"`
	VerticalPicUrl    interface{}   `json:"vertical_pic_url"`
	Keyword           interface{}   `json:"keyword"`
	PlotId            interface{}   `json:"plot_id"`
	RankGroupId       interface{}   `json:"rank_group_id"`
	FreeForPlatformId interface{}   `json:"free_for_platform_id"`
	Plid              interface{}   `json:"plid"`
	AreaId            string        `json:"area_id"`
	ClipsIds          []string      `json:"clips_ids"`
	DataFlag          int           `json:"data_flag"`
	CoverId           string        `json:"cover_id"`
	TypeName          string        `json:"type_name"`
	HorizontalPicUrl  interface{}   `json:"horizontal_pic_url"`
	CategoryMap       []interface{} `json:"category_map"`
	DoubanScore       interface{}   `json:"douban_score"`
	Director          []string      `json:"director"`
	VersionCid        interface{}   `json:"version_cid"`
	CostarId          interface{}   `json:"costar_id"`
	UpdateCalendar    interface{}   `json:"update_calendar"`
	CurrentTopic      interface{}   `json:"current_topic"`
	NomalIds          []struct {
		F int    `json:"F"`
		V string `json:"V"`
	} `json:"nomal_ids"`
	PromptText                  string      `json:"prompt_text"`
	PayfreeNum                  interface{} `json:"payfree_num"`
	DokiidIp                    string      `json:"dokiid_ip"`
	Area                        int         `json:"area"`
	CartoonAge                  interface{} `json:"cartoon_age"`
	PcVipPullText               interface{} `json:"pc_vip_pull_text"`
	HollywoodOnline             string      `json:"hollywood_online"`
	DigitalMode                 int         `json:"digital_mode"`
	PositiveTrailer             int         `json:"positive_trailer"`
	PcVipPullExpire             interface{} `json:"pc_vip_pull_expire"`
	SubGenre                    []string    `json:"sub_genre"`
	IsTrailer                   int         `json:"is_trailer"`
	FreeEndTime                 interface{} `json:"free_end_time"`
	ViewAllCount                int         `json:"view_all_count"`
	SecondTitle                 string      `json:"second_title"`
	Dialogue                    interface{} `json:"dialogue"`
	DoulieTags                  interface{} `json:"doulie_tags"`
	RecommendEpisodesIcon       interface{} `json:"recommend_episodes_icon"`
	EpisodeUpdated              string      `json:"episode_updated"`
	PublishDate                 string      `json:"publish_date"`
	CTipsCidConf                interface{} `json:"c_tips_cid_conf"`
	MtimeScore                  interface{} `json:"mtime_score"`
	Playright                   []string    `json:"playright"`
	RealExclusive               string      `json:"real_exclusive"`
	SeriesId                    interface{} `json:"series_id"`
	Type                        int         `json:"type"`
	CurrentNum                  int         `json:"current_num"`
	RecommendEpisodes           interface{} `json:"recommend_episodes"`
	ColumnId                    int         `json:"column_id"`
	NewPicHz                    string      `json:"new_pic_hz"`
	ChaseCalendarExpirationDate interface{} `json:"chase_calendar_expiration_date"`
	FreeBeginTime               interface{} `json:"free_begin_time"`
	IsCommodityShelvesName      interface{} `json:"is_commodity_shelves_name"`
	DataCheckupGrade            int         `json:"data_checkup_grade"`
	Web20Imgtag                 interface{} `json:"web20_imgtag"`
	PrePlayTimepoint            interface{} `json:"pre_play_timepoint"`
	Alias                       []string    `json:"alias"`
	CartoonAgeId                interface{} `json:"cartoon_age_id"`
	LeadingActor                []string    `json:"leading_actor"`
	UpdateDesc                  string      `json:"update_desc"`
	ListShowStyle               string      `json:"list_show_style"`
	LeadingActorId              []string    `json:"leading_actor_id"`
	DoulistId                   interface{} `json:"doulist_id"`
	Description                 string      `json:"description"`
	CopyrightId                 int         `json:"copyright_id"`
	Year                        string      `json:"year"`
	ImgtagVer                   string      `json:"imgtag_ver"`
	Title                       string      `json:"title"`
	InteractionTypeId           interface{} `json:"interaction_type_id"`
	PayForKnowledgeName         interface{} `json:"pay_for_knowledge_name"`
	PrePlayPlatsId              interface{} `json:"pre_play_plats_id"`
	SeriesName                  string      `json:"series_name"`
	VipIds                      []struct {
		F int    `json:"F"`
		V string `json:"V"`
	} `json:"vip_ids"`
	SeriesNum         string      `json:"series_num"`
	MainGenre         string      `json:"main_genre"`
	UpdateNotifyDesc  string      `json:"update_notify_desc"`
	MovieCommentSet   int         `json:"movie_comment_set"`
	ViewTodayCount    int         `json:"view_today_count"`
	EmotionIds        interface{} `json:"emotion_ids"`
	ActivityShowId    interface{} `json:"activity_show_id"`
	ImdbScore         interface{} `json:"imdb_score"`
	PayStatus         int         `json:"pay_status"`
	OutsiteFlag       int         `json:"outsite_flag"`
	Downright         []string    `json:"downright"`
	EpisodeAll        string      `json:"episode_all"`
	VipPage           bool        `json:"vipPage"`
	HasVplusRank      bool        `json:"hasVplusRank"`
	VideoTypePreCheck int         `json:"videoTypePreCheck"`
	PayInfo           struct {
		CVipPrice    interface{} `json:"c_vip_price"`
		CStatus      interface{} `json:"c_status"`
		CSinglePrice interface{} `json:"c_single_price"`
		CDiscount    interface{} `json:"c_discount"`
	} `json:"payInfo"`
	Theme struct {
	} `json:"theme"`
	CommentId struct {
		Result struct {
			Ret  int    `json:"ret"`
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		} `json:"result"`
		CommentId string `json:"comment_id"`
		SrcidType int    `json:"srcid_type"`
		Srcid     string `json:"srcid"`
	} `json:"commentId"`
	VideolistOld struct {
		List    []interface{} `json:"list"`
		Page    int           `json:"page"`
		Max     int           `json:"max"`
		Size    int           `json:"size"`
		Current int           `json:"current"`
	} `json:"videolistOld"`
	SellPoint struct {
		IError               int    `json:"iError"`
		StrMsg               string `json:"strMsg"`
		IDoubanHighScoreFlag int    `json:"iDoubanHighScoreFlag"`
		IWeiboHotCmtFlag     int    `json:"iWeiboHotCmtFlag"`
		StWeiboHotCmt        struct {
			StrTitle   string `json:"strTitle"`
			StrContent string `json:"strContent"`
		} `json:"stWeiboHotCmt"`
		IGoldenHorseAwardFlag int `json:"iGoldenHorseAwardFlag"`
	} `json:"sellPoint"`
	YearAndArea struct {
		Year string `json:"year"`
	} `json:"yearAndArea"`
}

type PiNia3 struct {
	Page struct {
		WujiConfig struct {
			AppConfig          string `json:"appConfig"`
			VersionKernelDemux string `json:"versionKernelDemux"`
			VersionKernelHlsjs string `json:"versionKernelHlsjs"`
			VersionKernelWasm  string `json:"versionKernelWasm"`
			VersionSuperPlayer string `json:"versionSuperPlayer"`
		} `json:"wujiConfig"`
		NavHtml       string `json:"navHtml"`
		FooterHtml    string `json:"footerHtml"`
		PlayerSideAds struct {
		} `json:"playerSideAds"`
		GrayConfig struct {
		} `json:"grayConfig"`
		RenderPage        bool `json:"renderPage"`
		IsPageReloading   bool `json:"isPageReloading"`
		WujiConfigFetched bool `json:"wujiConfigFetched"`
	} `json:"page"`
	Global struct {
		CurrentVid     string      `json:"currentVid"`
		CurrentCid     string      `json:"currentCid"`
		CurrentListKey interface{} `json:"currentListKey"`
		CoverInfo      struct {
			Type                   int           `json:"type"`
			CoverCheckupGrade      int           `json:"cover_checkup_grade"`
			CopyrightNegotiationId interface{}   `json:"copyright_negotiation_id"`
			Title                  string        `json:"title"`
			CurrentTopic           interface{}   `json:"current_topic"`
			AreaName               string        `json:"area_name"`
			PayForKnowledgeName    interface{}   `json:"pay_for_knowledge_name"`
			NewPicHz               string        `json:"new_pic_hz"`
			ColumnId               int           `json:"column_id"`
			DokiidIp               string        `json:"dokiid_ip"`
			DataCheckupGrade       int           `json:"data_checkup_grade"`
			Tag                    interface{}   `json:"tag"`
			VideoIds               []string      `json:"video_ids"`
			PositiveTrailer        int           `json:"positive_trailer"`
			Keyword                interface{}   `json:"keyword"`
			Alias                  []string      `json:"alias"`
			MovieCommentSet        int           `json:"movie_comment_set"`
			LeadingActor           []string      `json:"leading_actor"`
			BigHorizontalPicUrl    interface{}   `json:"big_horizontal_pic_url"`
			DataFlag               int           `json:"data_flag"`
			CoverId                string        `json:"cover_id"`
			SecondTitle            string        `json:"second_title"`
			TypeName               string        `json:"type_name"`
			AlbumClassificationId  interface{}   `json:"album_classification_id"`
			PlayPageStyleControlId interface{}   `json:"play_page_style_control_id"`
			Guests                 interface{}   `json:"guests"`
			CategoryMap            []interface{} `json:"category_map"`
			Downright              []string      `json:"downright"`
			PayForKnowledgeId      interface{}   `json:"pay_for_knowledge_id"`
			PublishDate            string        `json:"publish_date"`
			EpisodeAll             interface{}   `json:"episode_all"`
			CommentShowType        int           `json:"comment_show_type"`
			Playright              []string      `json:"playright"`
			Description            string        `json:"description"`
			CopyrightId            int           `json:"copyright_id"`
		} `json:"coverInfo"`
		VideoInfo struct {
			Title                  string        `json:"title"`
			Pic160X90              string        `json:"pic160x90"`
			Desc                   interface{}   `json:"desc"`
			Danmu                  int           `json:"danmu"`
			TypeName               string        `json:"type_name"`
			Vid                    string        `json:"vid"`
			DanmuStatus            int           `json:"danmu_status"`
			PlayPageStyleControlId interface{}   `json:"play_page_style_control_id"`
			CategoryMap            []interface{} `json:"category_map"`
			TopicIdList            interface{}   `json:"topic_id_list"`
			HeadTime               int           `json:"head_time"`
			CoverList              []string      `json:"cover_list"`
			Playright              []string      `json:"playright"`
			PositiveTrailer        int           `json:"positive_trailer"`
			TailTime               int           `json:"tail_time"`
			State                  int           `json:"state"`
			Type                   int           `json:"type"`
			CTitleOutput           string        `json:"c_title_output"`
			PioneerTag             interface{}   `json:"pioneer_tag"`
		} `json:"videoInfo"`
		IsVideoInfoLoading bool `json:"isVideoInfoLoading"`
		VideoInfo1006      struct {
			VideoUploadTime int    `json:"video_uploadTime"`
			VideoStatus     int    `json:"video_status"`
			ReleaseTime     string `json:"release_time"`
			VideoType       int    `json:"video_type"`
		} `json:"videoInfo1006"`
		ColumnInfo struct {
		} `json:"columnInfo"`
		PlayInfoFetched bool   `json:"playInfoFetched"`
		ShouldFixUrl    bool   `json:"shouldFixUrl"`
		PageStyle       string `json:"pageStyle"`
		IsLoginReady    bool   `json:"isLoginReady"`
		PageType        int    `json:"pageType"`
		PageTypeFetched bool   `json:"pageTypeFetched"`
	} `json:"global"`
	TopList struct {
		Data []struct {
			Title       string `json:"title"`
			Timelong    string `json:"timelong"`
			SecondTitle string `json:"secondTitle"`
			Id          string `json:"id"`
			Pic         string `json:"pic"`
		} `json:"data"`
		Title   string `json:"title"`
		Link    string `json:"link"`
		Type    int    `json:"type"`
		Fetched bool   `json:"fetched"`
	} `json:"topList"`
	KnowledgeIntro struct {
		Data    interface{} `json:"data"`
		Fetched bool        `json:"fetched"`
	} `json:"knowledgeIntro"`
	Introduction struct {
		IntroData struct {
			List []struct {
				ItemId         string `json:"item_id"`
				ItemType       string `json:"item_type"`
				ItemSourceType string `json:"item_source_type"`
				ItemParams     struct {
					DetailInfo                  string `json:"detail_info"`
					Type                        string `json:"type"`
					UpdateCalendar              string `json:"update_calendar"`
					ReportModId                 string `json:"report.mod_id"`
					CoverDescription            string `json:"cover_description"`
					Title                       string `json:"title"`
					SubTitle                    string `json:"sub_title"`
					Score                       string `json:"score"`
					AreaName                    string `json:"area_name"`
					NewPicVt                    string `json:"new_pic_vt"`
					NewPicHz                    string `json:"new_pic_hz"`
					ReportLid                   string `json:"report.lid"`
					ChaseCalendarExpirationDate string `json:"chase_calendar_expiration_date"`
					ActionUrl                   string `json:"action_url"`
					TagText                     string `json:"tag_text"`
					ShowYear                    string `json:"show_year"`
					SubGenre                    string `json:"sub_genre"`
					UpdateNotifyDesc            string `json:"update_notify_desc"`
					MainGenres                  string `json:"main_genres"`
					AnimeUpdateStatusId         string `json:"anime_update_status_id"`
					ViewAllCount                string `json:"view_all_count"`
					ReportIsTag                 string `json:"report.is_tag"`
					BroadcastTime               string `json:"broadcast_time"`
					ReportCid                   string `json:"report.cid"`
					ReportModTitle              string `json:"report.mod_title"`
					CoverYear                   string `json:"cover_year"`
					YearInfo                    string `json:"year_info"`
					VideoDescription            string `json:"video_description"`
					ReportModIdx                string `json:"report.mod_idx"`
					EpisodeAll                  string `json:"episode_all"`
					ReportItemIdx               string `json:"report.item_idx"`
					Hotval                      string `json:"hotval"`
					ReportUnModId               string `json:"report.un_mod_id"`
					Year                        string `json:"year"`
					ImgtagVer                   string `json:"imgtag_ver"`
				} `json:"item_params"`
				SubItems struct {
					StarList struct {
						ItemDatas []struct {
							ItemId         string `json:"item_id"`
							ItemType       string `json:"item_type"`
							ItemSourceType string `json:"item_source_type"`
							ItemParams     struct {
								Vcuid      string `json:"vcuid"`
								Fansid     string `json:"fansid"`
								IsVisitId  string `json:"is_visit_id"`
								IsVisit    string `json:"is_visit"`
								Name       string `json:"name"`
								TagId      string `json:"tag_id"`
								NameId     string `json:"name_id"`
								StarImgurl string `json:"star_imgurl"`
							} `json:"item_params"`
							SubItems struct {
							} `json:"sub_items"`
							ComplexJson string `json:"complex_json"`
						} `json:"item_datas"`
					} `json:"star_list"`
				} `json:"sub_items"`
				ComplexJson string `json:"complex_json"`
			} `json:"list"`
			Fetched bool `json:"fetched"`
		} `json:"introData"`
	} `json:"introduction"`
	EpisodeClips struct {
		Fetched     bool   `json:"fetched"`
		Title       string `json:"title"`
		HasNextPage bool   `json:"hasNextPage"`
		ListType    int    `json:"listType"`
		ListData    []struct {
			ItemId         string `json:"item_id"`
			ItemType       string `json:"item_type"`
			ItemSourceType string `json:"item_source_type"`
			ItemParams     struct {
				Title                 string `json:"title"`
				InteractionTypeId     string `json:"interaction_type_id"`
				UiType                string `json:"ui_type"`
				ShowAttentProgress    string `json:"show_attent_progress"`
				VideoAspectRatio      string `json:"video_aspect_ratio"`
				FeedParamsFeedDataKey string `json:"feed_params.feed_data_key"`
				CommentKey            string `json:"comment_key"`
				CopyRight             string `json:"copy_right"`
				FollowDataKey         string `json:"follow_data_key"`
				ShareContentTail      string `json:"share_content_tail"`
				PlayTitle             string `json:"play_title"`
				ExtraUserType         string `json:"extra_user_type"`
				ExtraUserLabelUrl     string `json:"extra_user_label_url"`
				PositiveTrailer       string `json:"positive_trailer"`
				Key                   string `json:"key"`
				ReportRtype           string `json:"report.rtype"`
				IsNoStoreWatchHistory string `json:"is_no_store_watch_history"`
				TailTime              string `json:"tail_time"`
				ShareUrl              string `json:"share_url"`
				Drm                   string `json:"drm"`
				AttentKey             string `json:"attent_key"`
				ExtraAccountType      string `json:"extra_account_type"`
				Episode               string `json:"episode"`
				VidEncrypt            string `json:"vid_encrypt"`
				VideoOrcDatakey       string `json:"video_orc_datakey"`
				CommentCount          string `json:"comment_count"`
				ShareTotalTime        string `json:"share_total_time"`
				Vid                   string `json:"vid"`
				FloatTitle            string `json:"float_title"`
				PraiseCount           string `json:"praise_count"`
				ShareCaptionKey       string `json:"share_caption_key"`
				Vuid                  string `json:"vuid"`
				PayStatus             string `json:"pay_status"`
				ReportItemIdx         string `json:"report.item_idx"`
				CCategoryValue        string `json:"c_category_value"`
				ShareDatakey          string `json:"share_datakey"`
				PraiseDatakey         string `json:"praise_datakey"`
				State                 string `json:"state"`
				Pic640360             string `json:"pic_640_360"`
				Type                  string `json:"type"`
				SyncCover             string `json:"sync_cover"`
				ReportFlowFrom        string `json:"report.flow_from"`
				ErrorTips             string `json:"error_tips"`
				FlopCardNum           string `json:"flop_card_num"`
				ExtraUserName         string `json:"extra_user_name"`
				WebPlayUrl            string `json:"web_play_url"`
				ReportVid             string `json:"report.vid"`
				FloatType             string `json:"float_type"`
				Url                   string `json:"url"`
				ReportTypeId          string `json:"report.type_id"`
				ErrorUrl              string `json:"error_url"`
				ShareSubTitle         string `json:"share_sub_title"`
				ColumnId              string `json:"column_id"`
				ReportActionUrl       string `json:"report_action_url"`
				FlopCardDatakey       string `json:"flop_card_datakey"`
				Pic496X280            string `json:"pic496x280"`
				CTitleDetail          string `json:"c_title_detail"`
				Duration              string `json:"duration"`
				Cid                   string `json:"cid"`
				PayType               string `json:"pay_type"`
				CanPlay               string `json:"can_play"`
				ShareAspectRatio      string `json:"share_aspect_ratio"`
				HeadTime              string `json:"head_time"`
				VideoEndRecommendKey  string `json:"video_end_recommend_key"`
				ShareImgUrl           string `json:"share_img_url"`
				FlopCardStatus        string `json:"flop_card_status"`
				ExtraUserImageUrl     string `json:"extra_user_image_url"`
				InteractiveContentId  string `json:"interactive_content_id"`
				ReportAReqCid         string `json:"report.a_req_cid"`
				CidPayType            string `json:"cid_pay_type"`
				ReportParamsBags      string `json:"report.params_bags"`
				ShowState             string `json:"show_state"`
				CFull                 string `json:"c_full"`
				FeedParamsPageId      string `json:"feed_params.page_id"`
				CloudRes              string `json:"cloud_res"`
				ShareTitle            string `json:"share_title"`
				ImageUrl              string `json:"image_url"`
				PraiseStatus          string `json:"praise_status"`
				ActionWholeUrl        string `json:"action_whole_url"`
				VDownright            string `json:"v_downright"`
				ShareSingleTitle      string `json:"share_single_title"`
				Aspect                string `json:"aspect"`
				ImgtagVer             string `json:"imgtag_ver"`
				ImmParamsFeedDataKey  string `json:"imm_params.feed_data_key"`
				CTagsFlag             string `json:"c_tags_flag"`
				ExtraAccountId        string `json:"extra_account_id"`
				ImmParamsPageId       string `json:"imm_params.page_id"`
				TrytimeSecond         string `json:"trytime_second"`
			} `json:"item_params"`
			SubItems struct {
			} `json:"sub_items"`
			ComplexJson string `json:"complex_json"`
		} `json:"listData"`
		ListMeta []interface{} `json:"listMeta"`
	} `json:"episodeClips"`
	EpisodeEp struct {
		Fetched     bool          `json:"fetched"`
		Title       string        `json:"title"`
		HasNextPage bool          `json:"hasNextPage"`
		ListType    int           `json:"listType"`
		ListData    []interface{} `json:"listData"`
		ListMeta    []interface{} `json:"listMeta"`
	} `json:"episodeEp"`
	EpisodeDeriveStore struct {
		Fetched     bool          `json:"fetched"`
		Title       string        `json:"title"`
		HasNextPage bool          `json:"hasNextPage"`
		ListType    int           `json:"listType"`
		ListData    []interface{} `json:"listData"`
		ListMeta    []interface{} `json:"listMeta"`
	} `json:"episodeDeriveStore"`
	EpisodeSeries struct {
		Fetched     bool          `json:"fetched"`
		Title       string        `json:"title"`
		HasNextPage bool          `json:"hasNextPage"`
		ListType    int           `json:"listType"`
		ListData    []interface{} `json:"listData"`
		ListMeta    []interface{} `json:"listMeta"`
	} `json:"episodeSeries"`
	EpisodeMain struct {
		Tabs                    []interface{} `json:"tabs"`
		TabsSettled             bool          `json:"tabsSettled"`
		CurrentTabIndex         int           `json:"currentTabIndex"`
		UiStyle                 int           `json:"uiStyle"`
		LeftTextRightPicShowNum int           `json:"leftTextRightPicShowNum"`
		FirstPreloaded          bool          `json:"firstPreloaded"`
		Fetched                 bool          `json:"fetched"`
		Title                   string        `json:"title"`
		HasNextPage             bool          `json:"hasNextPage"`
		ListType                int           `json:"listType"`
		ListData                [][]struct {
			ItemId         string `json:"item_id"`
			ItemType       string `json:"item_type"`
			ItemSourceType string `json:"item_source_type"`
			ItemParams     struct {
				CFull        string `json:"c_full"`
				UiType       string `json:"ui_type"`
				UnionTitle   string `json:"union_title"`
				Vid          string `json:"vid"`
				ImageUrl     string `json:"image_url"`
				ImgtagAll    string `json:"imgtag_all"`
				Duration     string `json:"duration"`
				Cid          string `json:"cid"`
				PublishDate  string `json:"publish_date"`
				PlayTitle    string `json:"play_title"`
				RefreshPage  string `json:"refresh_page"`
				CTitleOutput string `json:"c_title_output"`
				IsTrailer    string `json:"is_trailer"`
				Title        string `json:"title"`
			} `json:"item_params"`
			SubItems struct {
			} `json:"sub_items"`
			ComplexJson string `json:"complex_json"`
		} `json:"listData"`
		ListMeta []interface{} `json:"listMeta"`
	} `json:"episodeMain"`
	EpisodeRecommend struct {
		Fetched     bool          `json:"fetched"`
		Title       string        `json:"title"`
		HasNextPage bool          `json:"hasNextPage"`
		ListType    int           `json:"listType"`
		ListData    []interface{} `json:"listData"`
		ListMeta    []interface{} `json:"listMeta"`
	} `json:"episodeRecommend"`
	Epiosode struct {
		PlayListDataReady bool `json:"playListDataReady"`
	} `json:"epiosode"`
	VidLocationMap struct {
		Priority struct {
			MainList           int `json:"mainList"`
			SinglePlayList     int `json:"singlePlayList"`
			InfiniteListShort  int `json:"infiniteListShort"`
			ClipsList          int `json:"clipsList"`
			EpList             int `json:"epList"`
			SeriesList         int `json:"seriesList"`
			RecommendListShort int `json:"recommendListShort"`
			DeriveList         int `json:"deriveList"`
			RecommendList      int `json:"recommendList"`
		} `json:"priority"`
		Map struct{} `json:"map"`
	} `json:"vidLocationMap"`
	EpisodeRecommendShort struct {
		NextPageContext struct {
		} `json:"nextPageContext"`
		Fetched     bool          `json:"fetched"`
		Title       string        `json:"title"`
		HasNextPage bool          `json:"hasNextPage"`
		ListType    int           `json:"listType"`
		ListData    []interface{} `json:"listData"`
		ListMeta    []interface{} `json:"listMeta"`
	} `json:"episodeRecommendShort"`
	EpisodeInfiniteShort struct {
		NextPageContext struct {
		} `json:"nextPageContext"`
		Fetched     bool          `json:"fetched"`
		Title       string        `json:"title"`
		HasNextPage bool          `json:"hasNextPage"`
		ListType    int           `json:"listType"`
		ListData    []interface{} `json:"listData"`
		ListMeta    []interface{} `json:"listMeta"`
	} `json:"episodeInfiniteShort"`
	EpisodeSinglePlay struct {
		Fetched     bool          `json:"fetched"`
		Title       string        `json:"title"`
		HasNextPage bool          `json:"hasNextPage"`
		ListType    int           `json:"listType"`
		ListData    []interface{} `json:"listData"`
		ListMeta    []interface{} `json:"listMeta"`
	} `json:"episodeSinglePlay"`
	Account struct {
		CommentNum int    `json:"commentNum"`
		CurrentId  string `json:"currentId"`
		CachedData struct {
		} `json:"cachedData"`
	} `json:"account"`
	PayVipStatus struct {
		NewPayStatus struct {
		} `json:"newPayStatus"`
	} `json:"payVipStatus"`
	Panel struct {
		PanelAnthology struct {
			IsShow bool   `json:"isShow"`
			Action string `json:"action"`
			Args   string `json:"args"`
			Scene  string `json:"scene"`
		} `json:"panelAnthology"`
		PanelTipPay struct {
			IsShow bool   `json:"isShow"`
			Scene  string `json:"scene"`
		} `json:"panelTipPay"`
		PanelDiamondPay struct {
			IsShow       bool   `json:"isShow"`
			Action       string `json:"action"`
			Vid          string `json:"vid"`
			PurchaseType string `json:"purchaseType"`
			PayableType  string `json:"payableType"`
			Scene        string `json:"scene"`
		} `json:"panelDiamondPay"`
		PanelFloatWarning struct {
			IsShow bool `json:"isShow"`
		} `json:"panelFloatWarning"`
		PanelPurchase struct {
			IsShow       bool   `json:"isShow"`
			Action       string `json:"action"`
			Args         string `json:"args"`
			Vid          string `json:"vid"`
			PurchaseType string `json:"purchaseType"`
			PayableType  string `json:"payableType"`
			Scene        string `json:"scene"`
		} `json:"panelPurchase"`
		PanelMinipay struct {
			IsShow bool   `json:"isShow"`
			Action string `json:"action"`
			Args   string `json:"args"`
			Scene  string `json:"scene"`
		} `json:"panelMinipay"`
	} `json:"panel"`
}
type PiNia2 struct {
	Page struct {
		NavHtml       string `json:"navHtml"`
		FooterHtml    string `json:"footerHtml"`
		PlayerSideAds struct {
		} `json:"playerSideAds"`
		GrayConfig struct {
		} `json:"grayConfig"`
		RenderPage      bool `json:"renderPage"`
		IsPageReloading bool `json:"isPageReloading"`
		WujiConfig      struct {
			Player struct {
				VersionSuperPlayer string `json:"versionSuperPlayer"`
				VersionKernelWasm  string `json:"versionKernelWasm"`
				VersionKernelHlsjs string `json:"versionKernelHlsjs"`
				VersionKernelDemux string `json:"versionKernelDemux"`
				AppConfig          string `json:"appConfig"`
			} `json:"player"`
			PageConfig struct {
				DownloadBanner struct {
					JsonData struct {
						Title    string `json:"title"`
						SubTitle string `json:"subTitle"`
					} `json:"jsonData"`
					Value string `json:"value"`
					Image string `json:"image"`
				} `json:"download_banner"`
				TopListItemNum struct {
					JsonData struct {
						HasTopic int `json:"hasTopic"`
						NoTopic  int `json:"noTopic"`
					} `json:"jsonData"`
					Value string `json:"value"`
					Image string `json:"image"`
				} `json:"top_list_item_num"`
			} `json:"pageConfig"`
		} `json:"wujiConfig"`
		WujiConfigFetched bool `json:"wujiConfigFetched"`
		TabExperiment     struct {
			Debug     string        `json:"debug"`
			AllExpIds []interface{} `json:"allExpIds"`
			Statuses  []string      `json:"statuses"`
			ExpIds    []interface{} `json:"expIds"`
		} `json:"tabExperiment"`
		TabExperimentEnable          bool   `json:"tabExperimentEnable"`
		TabExperimentCookieCacheTime int    `json:"tabExperimentCookieCacheTime"`
		LastPlaylistUIStyle          string `json:"lastPlaylistUIStyle"`
		IsV2                         bool   `json:"isV2"`
		IsLoginReady                 bool   `json:"isLoginReady"`
	} `json:"page"`
	Global struct {
		CurrentVid     string      `json:"currentVid"`
		CurrentCid     string      `json:"currentCid"`
		CurrentListKey interface{} `json:"currentListKey"`
		CoverInfo      struct {
			MovieCommentSet        int           `json:"movie_comment_set"`
			BigHorizontalPicUrl    interface{}   `json:"big_horizontal_pic_url"`
			LeadingActor           []string      `json:"leading_actor"`
			DataFlag               int           `json:"data_flag"`
			CoverId                string        `json:"cover_id"`
			SecondTitle            string        `json:"second_title"`
			TypeName               string        `json:"type_name"`
			AlbumClassificationId  interface{}   `json:"album_classification_id"`
			PlayPageStyleControlId interface{}   `json:"play_page_style_control_id"`
			Guests                 interface{}   `json:"guests"`
			CategoryMap            []interface{} `json:"category_map"`
			PayStatus              int           `json:"pay_status"`
			Downright              []string      `json:"downright"`
			PayForKnowledgeId      interface{}   `json:"pay_for_knowledge_id"`
			PublishDate            string        `json:"publish_date"`
			EpisodeAll             string        `json:"episode_all"`
			Playright              []string      `json:"playright"`
			Description            string        `json:"description"`
			CommentShowType        int           `json:"comment_show_type"`
			CopyrightId            int           `json:"copyright_id"`
			Type                   int           `json:"type"`
			CoverCheckupGrade      int           `json:"cover_checkup_grade"`
			CopyrightNegotiationId interface{}   `json:"copyright_negotiation_id"`
			Title                  string        `json:"title"`
			CurrentTopic           interface{}   `json:"current_topic"`
			AreaName               string        `json:"area_name"`
			PayForKnowledgeName    interface{}   `json:"pay_for_knowledge_name"`
			NewPicHz               string        `json:"new_pic_hz"`
			ColumnId               int           `json:"column_id"`
			DokiidIp               string        `json:"dokiid_ip"`
			CartoonAge             interface{}   `json:"cartoon_age"`
			DataCheckupGrade       int           `json:"data_checkup_grade"`
			Tag                    interface{}   `json:"tag"`
			VideoIds               []string      `json:"video_ids"`
			PositiveTrailer        int           `json:"positive_trailer"`
			Keyword                interface{}   `json:"keyword"`
			Alias                  []string      `json:"alias"`
		} `json:"coverInfo"`
		VideoInfo struct {
		} `json:"videoInfo"`
		RelatedCoverInfo   interface{} `json:"relatedCoverInfo"`
		IsVideoInfoLoading bool        `json:"isVideoInfoLoading"`
		VideoInfo1006      struct {
		} `json:"videoInfo1006"`
		ColumnInfo struct {
		} `json:"columnInfo"`
		PlayInfoFetched   bool   `json:"playInfoFetched"`
		ShouldFixUrl      bool   `json:"shouldFixUrl"`
		PageStyle         string `json:"pageStyle"`
		VideoTypePrecheck int    `json:"videoTypePrecheck"`
		PageType          int    `json:"pageType"`
		PageTypeFetched   bool   `json:"pageTypeFetched"`
		HasTopic          bool   `json:"hasTopic"`
	} `json:"global"`
	Player struct {
		SubVidMod struct {
			IsSubVidMod     bool   `json:"isSubVidMod"`
			MainVid         string `json:"mainVid"`
			MainVidPlaytime int    `json:"mainVidPlaytime"`
		} `json:"subVidMod"`
		ChangeVideoType  string        `json:"changeVideoType"`
		HistoryStartTime int           `json:"historyStartTime"`
		IsPlayCalled     bool          `json:"isPlayCalled"`
		ScrollToVid      bool          `json:"scrollToVid"`
		HiddenVideoList  []interface{} `json:"hiddenVideoList"`
	} `json:"player"`
	TopList struct {
		Data []struct {
			Title       string `json:"title"`
			Timelong    string `json:"timelong"`
			SecondTitle string `json:"secondTitle"`
			Id          string `json:"id"`
			Pic         string `json:"pic"`
		} `json:"data"`
		Title   string `json:"title"`
		Link    string `json:"link"`
		Type    int    `json:"type"`
		Fetched bool   `json:"fetched"`
	} `json:"topList"`
	KnowledgeIntro struct {
		Data    string `json:"data"`
		Fetched bool   `json:"fetched"`
	} `json:"knowledgeIntro"`
	Introduction struct {
		IntroData struct {
			List []struct {
				ItemId         string `json:"item_id"`
				ItemType       string `json:"item_type"`
				ItemSourceType string `json:"item_source_type"`
				ItemParams     struct {
					UpdateNotifyDesc            string `json:"update_notify_desc"`
					MainGenres                  string `json:"main_genres"`
					AnimeUpdateStatusId         string `json:"anime_update_status_id"`
					ViewAllCount                string `json:"view_all_count"`
					ReportIsTag                 string `json:"report.is_tag"`
					BroadcastTime               string `json:"broadcast_time"`
					ReportCid                   string `json:"report.cid"`
					ReportModTitle              string `json:"report.mod_title"`
					CoverYear                   string `json:"cover_year"`
					YearInfo                    string `json:"year_info"`
					ReportModIdx                string `json:"report.mod_idx"`
					VideoDescription            string `json:"video_description"`
					EpisodeAll                  string `json:"episode_all"`
					ReportItemIdx               string `json:"report.item_idx"`
					Hotval                      string `json:"hotval"`
					ReportUnModId               string `json:"report.un_mod_id"`
					ImgtagVer                   string `json:"imgtag_ver"`
					Year                        string `json:"year"`
					Type                        string `json:"type"`
					DetailInfo                  string `json:"detail_info"`
					UpdateCalendar              string `json:"update_calendar"`
					CoverDescription            string `json:"cover_description"`
					ReportModId                 string `json:"report.mod_id"`
					Title                       string `json:"title"`
					SubTitle                    string `json:"sub_title"`
					Score                       string `json:"score"`
					AreaName                    string `json:"area_name"`
					TitleImageUrl               string `json:"title_image_url"`
					NewPicHz                    string `json:"new_pic_hz"`
					NewPicVt                    string `json:"new_pic_vt"`
					ReportLid                   string `json:"report.lid"`
					ChaseCalendarExpirationDate string `json:"chase_calendar_expiration_date"`
					TagText                     string `json:"tag_text"`
					ActionUrl                   string `json:"action_url"`
					ShowYear                    string `json:"show_year"`
					SubGenre                    string `json:"sub_genre"`
				} `json:"item_params"`
				SubItems struct {
					StarList struct {
						ItemDatas []struct {
							ItemId         string `json:"item_id"`
							ItemType       string `json:"item_type"`
							ItemSourceType string `json:"item_source_type"`
							ItemParams     struct {
								IsVisit    string `json:"is_visit"`
								IsVisitId  string `json:"is_visit_id"`
								Name       string `json:"name"`
								TagId      string `json:"tag_id"`
								StarImgurl string `json:"star_imgurl"`
								NameId     string `json:"name_id"`
								Vcuid      string `json:"vcuid"`
								Fansid     string `json:"fansid"`
							} `json:"item_params"`
							SubItems struct {
							} `json:"sub_items"`
							ComplexJson string `json:"complex_json"`
						} `json:"item_datas"`
					} `json:"star_list"`
				} `json:"sub_items"`
				ComplexJson string `json:"complex_json"`
			} `json:"list"`
			Fetched bool `json:"fetched"`
		} `json:"introData"`
	} `json:"introduction"`
	EpisodeMain struct {
		EpTabs                  []interface{} `json:"epTabs"`
		EpTabsSettled           bool          `json:"epTabsSettled"`
		CurrentEpTabIndex       int           `json:"currentEpTabIndex"`
		DefaultEpTabIndex       int           `json:"defaultEpTabIndex"`
		LeftTextRightPicShowNum int           `json:"leftTextRightPicShowNum"`
		FirstPreloaded          bool          `json:"firstPreloaded"`
		Refreshed               bool          `json:"refreshed"`
		Fetched                 bool          `json:"fetched"`
		Title                   string        `json:"title"`
		HasNextPage             bool          `json:"hasNextPage"`
		ListType                int           `json:"listType"`
		ListData                []struct {
			List [][]struct {
				Cid         string `json:"cid"`
				Index       int    `json:"index"`
				Pic         string `json:"pic"`
				PicVertial  string `json:"picVertial"`
				ImgTag      string `json:"imgTag"`
				Title       string `json:"title"`
				Vid         string `json:"vid"`
				Lid         string `json:"lid"`
				Duration    string `json:"duration"`
				PlayTitle   string `json:"playTitle"`
				FullTitle   string `json:"fullTitle"`
				PublishDate string `json:"publishDate"`
				ItemType    string `json:"itemType"`
				ListKey     string `json:"listKey"`
				Page        int    `json:"page"`
				FocusOn     string `json:"focusOn"`
				DtParams    struct {
					TabId    string `json:"tab_id"`
					TabIdx   int    `json:"tab_idx"`
					ModId    string `json:"mod_id"`
					ModIdx   int    `json:"mod_idx"`
					ModTitle string `json:"mod_title"`
				} `json:"dtParams"`
				RefreshPage           string `json:"refreshPage"`
				IsNoStoreWatchHistory bool   `json:"isNoStoreWatchHistory"`
			} `json:"list"`
			Tabs       []interface{} `json:"tabs"`
			TabIndex   int           `json:"tabIndex"`
			Style      int           `json:"style"`
			TabSettled bool          `json:"tabSettled"`
		} `json:"listData"`
		ListMeta []interface{} `json:"listMeta"`
	} `json:"episodeMain"`
	EpisodeCut struct {
		Fetched     bool          `json:"fetched"`
		Title       string        `json:"title"`
		HasNextPage bool          `json:"hasNextPage"`
		ListType    int           `json:"listType"`
		ListData    []interface{} `json:"listData"`
		ListMeta    []interface{} `json:"listMeta"`
	} `json:"episodeCut"`
	EpisodeEp struct {
		Fetched     bool          `json:"fetched"`
		Title       string        `json:"title"`
		HasNextPage bool          `json:"hasNextPage"`
		ListType    int           `json:"listType"`
		ListData    []interface{} `json:"listData"`
		ListMeta    []interface{} `json:"listMeta"`
	} `json:"episodeEp"`
	EpisodeClips struct {
		Fetched     bool   `json:"fetched"`
		Title       string `json:"title"`
		HasNextPage bool   `json:"hasNextPage"`
		ListType    int    `json:"listType"`
		ListData    []struct {
			ItemId         string `json:"item_id"`
			ItemType       string `json:"item_type"`
			ItemSourceType string `json:"item_source_type"`
			ItemParams     struct {
				CFull                    string `json:"c_full"`
				FeedParamsPageId         string `json:"feed_params.page_id"`
				CloudRes                 string `json:"cloud_res"`
				ShareTitle               string `json:"share_title"`
				ImageUrl                 string `json:"image_url"`
				PraiseStatus             string `json:"praise_status"`
				ActionWholeUrl           string `json:"action_whole_url"`
				VDownright               string `json:"v_downright"`
				ShareSingleTitle         string `json:"share_single_title"`
				Aspect                   string `json:"aspect"`
				ImmParamsFeedDataKey     string `json:"imm_params.feed_data_key"`
				ImgtagVer                string `json:"imgtag_ver"`
				CTagsFlag                string `json:"c_tags_flag"`
				ExtraAccountId           string `json:"extra_account_id"`
				TrytimeSecond            string `json:"trytime_second"`
				ImmParamsPageId          string `json:"imm_params.page_id"`
				Title                    string `json:"title"`
				InteractionTypeId        string `json:"interaction_type_id"`
				UiType                   string `json:"ui_type"`
				ShowAttentProgress       string `json:"show_attent_progress"`
				VideoAspectRatio         string `json:"video_aspect_ratio"`
				FeedParamsFeedDataKey    string `json:"feed_params.feed_data_key"`
				ImgtagAll                string `json:"imgtag_all"`
				CommentKey               string `json:"comment_key"`
				FollowDataKey            string `json:"follow_data_key"`
				CopyRight                string `json:"copy_right"`
				ShareContentTail         string `json:"share_content_tail"`
				ExtraUserType            string `json:"extra_user_type"`
				PlayTitle                string `json:"play_title"`
				ExtraUserLabelUrl        string `json:"extra_user_label_url"`
				Key                      string `json:"key"`
				PositiveTrailer          string `json:"positive_trailer"`
				ReportRtype              string `json:"report.rtype"`
				IsNoStoreWatchHistory    string `json:"is_no_store_watch_history"`
				TailTime                 string `json:"tail_time"`
				ShareUrl                 string `json:"share_url"`
				Drm                      string `json:"drm"`
				AttentKey                string `json:"attent_key"`
				Episode                  string `json:"episode"`
				ExtraAccountType         string `json:"extra_account_type"`
				IsTrailer                string `json:"is_trailer"`
				VidEncrypt               string `json:"vid_encrypt"`
				VideoOrcDatakey          string `json:"video_orc_datakey"`
				CommentCount             string `json:"comment_count"`
				ShareTotalTime           string `json:"share_total_time"`
				Vid                      string `json:"vid"`
				FloatTitle               string `json:"float_title"`
				PraiseCount              string `json:"praise_count"`
				Vuid                     string `json:"vuid"`
				ShareCaptionKey          string `json:"share_caption_key"`
				PayStatus                string `json:"pay_status"`
				ReportItemIdx            string `json:"report.item_idx"`
				CCategoryValue           string `json:"c_category_value"`
				ShareDatakey             string `json:"share_datakey"`
				PraiseDatakey            string `json:"praise_datakey"`
				Pic640360                string `json:"pic_640_360"`
				State                    string `json:"state"`
				Type                     string `json:"type"`
				SyncCover                string `json:"sync_cover"`
				FlopCardNum              string `json:"flop_card_num"`
				ReportFlowFrom           string `json:"report.flow_from"`
				WebPlayUrl               string `json:"web_play_url"`
				ReportVid                string `json:"report.vid"`
				FloatType                string `json:"float_type"`
				Url                      string `json:"url"`
				ReportTypeId             string `json:"report.type_id"`
				ShareSubTitle            string `json:"share_sub_title"`
				ColumnId                 string `json:"column_id"`
				ReportActionUrl          string `json:"report_action_url"`
				FlopCardDatakey          string `json:"flop_card_datakey"`
				Pic496X280               string `json:"pic496x280"`
				Duration                 string `json:"duration"`
				CTitleDetail             string `json:"c_title_detail"`
				Cid                      string `json:"cid"`
				PayType                  string `json:"pay_type"`
				HeadTime                 string `json:"head_time"`
				ShareAspectRatio         string `json:"share_aspect_ratio"`
				VideoEndRecommendKey     string `json:"video_end_recommend_key"`
				FlopCardStatus           string `json:"flop_card_status"`
				ShareImgUrl              string `json:"share_img_url"`
				ExtraUserImageUrl        string `json:"extra_user_image_url"`
				RecallItemItemSourceType string `json:"@recall_item.item_source_type"`
				ReportAReqCid            string `json:"report.a_req_cid"`
				CidPayType               string `json:"cid_pay_type"`
				ReportParamsBags         string `json:"report.params_bags"`
				ShowState                string `json:"show_state"`
			} `json:"item_params"`
			SubItems struct {
			} `json:"sub_items"`
			ComplexJson string `json:"complex_json"`
		} `json:"listData"`
		ListMeta []interface{} `json:"listMeta"`
	} `json:"episodeClips"`
	EpisodeSeries struct {
		Fetched     bool          `json:"fetched"`
		Title       string        `json:"title"`
		HasNextPage bool          `json:"hasNextPage"`
		ListType    int           `json:"listType"`
		ListData    []interface{} `json:"listData"`
		ListMeta    []interface{} `json:"listMeta"`
	} `json:"episodeSeries"`
	EpisodeDeriveStore struct {
		Fetched     bool          `json:"fetched"`
		Title       string        `json:"title"`
		HasNextPage bool          `json:"hasNextPage"`
		ListType    int           `json:"listType"`
		ListData    []interface{} `json:"listData"`
		ListMeta    []interface{} `json:"listMeta"`
	} `json:"episodeDeriveStore"`
	EpisodeRecommend struct {
		Fetched     bool          `json:"fetched"`
		Title       string        `json:"title"`
		HasNextPage bool          `json:"hasNextPage"`
		ListType    int           `json:"listType"`
		ListData    []interface{} `json:"listData"`
		ListMeta    []interface{} `json:"listMeta"`
	} `json:"episodeRecommend"`
	Epiosode struct {
		PlayListDataReady bool   `json:"playListDataReady"`
		CurrentVideoInfo  string `json:"currentVideoInfo"`
		NextVideoInfo     string `json:"nextVideoInfo"`
		ShouldUseShortApi bool   `json:"shouldUseShortApi"`
		ShowOtherVideos   bool   `json:"showOtherVideos"`
	} `json:"epiosode"`
	VidLocationMap struct {
		Priority struct {
			MainList           int `json:"mainList"`
			SinglePlayList     int `json:"singlePlayList"`
			RelativeList       int `json:"relativeList"`
			CutList            int `json:"cutList"`
			EpList             int `json:"epList"`
			ClipsList          int `json:"clipsList"`
			SeriesList         int `json:"seriesList"`
			InfiniteListShort  int `json:"infiniteListShort"`
			DeriveList         int `json:"deriveList"`
			RecommendListShort int `json:"recommendListShort"`
			RecommendList      int `json:"recommendList"`
		} `json:"priority"`
		Map struct {
			U0045U564C3 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"u0045u564c3"`
			N0045Jj07Gu []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"n0045jj07gu"`
			M0045Keote0 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"m0045keote0"`
			A00458Uz56T []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"a00458uz56t"`
			E0045H8Kl5A []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"e0045h8kl5a"`
			G0045Qfcvd6 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"g0045qfcvd6"`
			E0045K0Gvsw []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"e0045k0gvsw"`
			L0045Cbo32Q []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"l0045cbo32q"`
			I0045496Vu8 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"i0045496vu8"`
			P0045Nabcuq []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"p0045nabcuq"`
			T0045Df4I6F []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"t0045df4i6f"`
			M0045Coeqet []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"m0045coeqet"`
			J0045Mlif2P []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"j0045mlif2p"`
			O0045Nm9Agn []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"o0045nm9agn"`
			G0045D2Quuh []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"g0045d2quuh"`
			T0045W9542I []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"t0045w9542i"`
			U0045Pqj5F0 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"u0045pqj5f0"`
			J0045O3Imlh []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"j0045o3imlh"`
			H00457Oyzag []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"h00457oyzag"`
			G0045Bndbhq []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"g0045bndbhq"`
			S0045Ce8Agg []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"s0045ce8agg"`
			L0045L1Mopi []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"l0045l1mopi"`
			Q00451I7O14 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"q00451i7o14"`
			V0045J03K9K []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"v0045j03k9k"`
			H0045Dbwapf []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"h0045dbwapf"`
			F0045M2Nr04 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"f0045m2nr04"`
			K004559K5Uj []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"k004559k5uj"`
			V0045U3Uu6W []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"v0045u3uu6w"`
			W00458A2Xzp []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"w00458a2xzp"`
			F0045Efwm2E []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"f0045efwm2e"`
			J0045Yutboo []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"j0045yutboo"`
			X00457Vh8Wj []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"x00457vh8wj"`
			Z0045Dj8Yr6 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"z0045dj8yr6"`
			C0045Adk2J4 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"c0045adk2j4"`
			J0045B1Bx3O []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"j0045b1bx3o"`
			I0045Eweew1 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    int    `json:"page"`
			} `json:"i0045eweew1"`
			R00459C6Nqh []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"r00459c6nqh"`
			G00451Uohvb []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"g00451uohvb"`
			V00456N35H3 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"v00456n35h3"`
			L0045Jvijv7 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"l0045jvijv7"`
			Q0045Nva75X []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"q0045nva75x"`
			Y0045Nbra9K []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"y0045nbra9k"`
			O00457Wp790 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"o00457wp790"`
			D00451Jpzdx []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"d00451jpzdx"`
			K0045Qevswm []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"k0045qevswm"`
			Q0045Dtg6Db []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"q0045dtg6db"`
			B0045Qt4Cf6 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"b0045qt4cf6"`
			Z004565Mz6K []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"z004565mz6k"`
			V0045Lr2P7U []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"v0045lr2p7u"`
			M0045G79X53 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"m0045g79x53"`
			I0045Vkkcn1 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"i0045vkkcn1"`
			Z0045A74R1B []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"z0045a74r1b"`
			A0045Isrg1V []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"a0045isrg1v"`
			C0045Zp1Rkd []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"c0045zp1rkd"`
			L0045114Uzr []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"l0045114uzr"`
			F0045Oizkzo []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"f0045oizkzo"`
			L0045Bt0J7N []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"l0045bt0j7n"`
			N004574Ijrn []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"n004574ijrn"`
			U004588E9Xv []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"u004588e9xv"`
			J0045Jpa2Wv []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"j0045jpa2wv"`
			R0045Hbznag []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"r0045hbznag"`
			P0045Oydg7P []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"p0045oydg7p"`
			A0045Flxntu []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"a0045flxntu"`
			V0045D8X3X6 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"v0045d8x3x6"`
			O00450045Ag []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"o00450045ag"`
			S0045Asithl []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"s0045asithl"`
			Z00459Pfsno []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"z00459pfsno"`
			T0045Sk0L1V []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"t0045sk0l1v"`
			J0045Rilt8I []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"j0045rilt8i"`
			C0045Ba6Dua []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"c0045ba6dua"`
			N0045Si1Vs2 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"n0045si1vs2"`
			G0045Enmdl6 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"g0045enmdl6"`
			R00452R9M3P []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"r00452r9m3p"`
			J0045D8Ngoh []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"j0045d8ngoh"`
			R0045Te6Oic []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"r0045te6oic"`
			J00456Wuyqz []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"j00456wuyqz"`
			Z0045Sbxclm []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"z0045sbxclm"`
			M0045B1N77P []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"m0045b1n77p"`
			L0045X1Ivdn []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"l0045x1ivdn"`
			I00459Qs4Si []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"i00459qs4si"`
			P0045Bns16H []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"p0045bns16h"`
			I0045Qo9B3K []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"i0045qo9b3k"`
			V0045X72Xw1 []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"v0045x72xw1"`
			V004575Dw1P []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"v004575dw1p"`
			X0045Gwzl8T []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"x0045gwzl8t"`
			H00455Rbi4Y []struct {
				Index   int    `json:"index"`
				ListKey string `json:"listKey"`
				Page    string `json:"page"`
			} `json:"h00455rbi4y"`
		} `json:"map"`
	} `json:"vidLocationMap"`
	EpisodeRecommendShort struct {
		NextPageContext struct {
		} `json:"nextPageContext"`
		Fetched     bool          `json:"fetched"`
		Title       string        `json:"title"`
		HasNextPage bool          `json:"hasNextPage"`
		ListType    int           `json:"listType"`
		ListData    []interface{} `json:"listData"`
		ListMeta    []interface{} `json:"listMeta"`
		IsLoading   bool          `json:"isLoading"`
	} `json:"episodeRecommendShort"`
	EpisodeInfiniteShort struct {
		NextPageContext struct {
		} `json:"nextPageContext"`
		Fetched     bool          `json:"fetched"`
		Title       string        `json:"title"`
		HasNextPage bool          `json:"hasNextPage"`
		ListType    int           `json:"listType"`
		ListData    []interface{} `json:"listData"`
		ListMeta    []interface{} `json:"listMeta"`
		IsLoading   bool          `json:"isLoading"`
	} `json:"episodeInfiniteShort"`
	EpisodeSinglePlay struct {
		Fetched     bool          `json:"fetched"`
		Title       string        `json:"title"`
		HasNextPage bool          `json:"hasNextPage"`
		ListType    int           `json:"listType"`
		ListData    []interface{} `json:"listData"`
		ListMeta    []interface{} `json:"listMeta"`
	} `json:"episodeSinglePlay"`
	EpisodeRelative struct {
		NextPageContext struct {
		} `json:"nextPageContext"`
		Fetched     bool          `json:"fetched"`
		Title       string        `json:"title"`
		HasNextPage bool          `json:"hasNextPage"`
		ListType    int           `json:"listType"`
		ListData    []interface{} `json:"listData"`
		ListMeta    []interface{} `json:"listMeta"`
		IsLoading   bool          `json:"isLoading"`
	} `json:"episodeRelative"`
	Account struct {
		CommentNum int    `json:"commentNum"`
		CurrentId  string `json:"currentId"`
		CachedData struct {
		} `json:"cachedData"`
	} `json:"account"`
	PayVipStatus struct {
		NewPayStatus struct {
		} `json:"newPayStatus"`
		PayRemainTime string      `json:"payRemainTime"`
		PayCountdown  interface{} `json:"payCountdown"`
		CouponStatus  int         `json:"couponStatus"`
		CouponId      string      `json:"couponId"`
		Fetched       bool        `json:"fetched"`
	} `json:"payVipStatus"`
	Panel struct {
		TabName        string `json:"tabName"`
		PanelAnthology struct {
			IsShow bool   `json:"isShow"`
			Action string `json:"action"`
			Args   string `json:"args"`
			Scene  string `json:"scene"`
		} `json:"panelAnthology"`
		PanelTipPay struct {
			IsShow bool   `json:"isShow"`
			Scene  string `json:"scene"`
		} `json:"panelTipPay"`
		PanelDiamondPay struct {
			IsShow       bool   `json:"isShow"`
			Action       string `json:"action"`
			Vid          string `json:"vid"`
			PurchaseType string `json:"purchaseType"`
			PayableType  string `json:"payableType"`
			Scene        string `json:"scene"`
			IsSport      bool   `json:"isSport"`
		} `json:"panelDiamondPay"`
		PanelFloatWarning struct {
			IsShow bool `json:"isShow"`
		} `json:"panelFloatWarning"`
		PanelPurchase struct {
			IsShow       bool   `json:"isShow"`
			Action       string `json:"action"`
			Args         string `json:"args"`
			Vid          string `json:"vid"`
			PurchaseType string `json:"purchaseType"`
			PayableType  string `json:"payableType"`
			Scene        string `json:"scene"`
			GiftPackId   string `json:"giftPackId"`
			IsSport      bool   `json:"isSport"`
		} `json:"panelPurchase"`
		PanelMinipay struct {
			IsShow bool   `json:"isShow"`
			Action string `json:"action"`
			Args   string `json:"args"`
			Scene  string `json:"scene"`
		} `json:"panelMinipay"`
		PanelNfl struct {
			IsShow bool `json:"isShow"`
		} `json:"panelNfl"`
	} `json:"panel"`
}
