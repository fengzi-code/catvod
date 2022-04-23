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
