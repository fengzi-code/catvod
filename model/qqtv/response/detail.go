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

type PiNia2 struct {
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
