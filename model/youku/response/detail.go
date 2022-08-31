package response

type YoukuDescribe struct {
	Api  string `json:"api"`
	Data struct {
		Field1 struct {
			Data struct {
				Data struct {
					Session string `json:"session"`
					Extra   struct {
						VideoCategory      string        `json:"videoCategory"`
						Audiolang          []interface{} `json:"audiolang"`
						VideoId            string        `json:"videoId"`
						InteractiveChapter int           `json:"interactiveChapter"`
						ShowReleaseTime    string        `json:"showReleaseTime"`
						EpisodeFinalStage  int           `json:"episodeFinalStage"`
						ShowCategory       string        `json:"showCategory"`
						VideoType          string        `json:"videoType"`
						AllowLike          bool          `json:"allowLike"`
						IsNewReservation   bool          `json:"isNewReservation"`
						Completed          bool          `json:"completed"`
						AllowComment       bool          `json:"allowComment"`
						StreamTypes        []string      `json:"streamTypes"`
						TotalUp            int           `json:"totalUp"`
						VideoCategoryId    int           `json:"videoCategoryId"`
						ShowVideoStageDesc string        `json:"showVideoStageDesc"`
						ShowImg            string        `json:"showImg"`
						IsFavorite         bool          `json:"isFavorite"`
						TotalComment       int           `json:"totalComment"`
						EpisodeTotal       int           `json:"episodeTotal"`
						PageKey            string        `json:"pageKey"`
						ShowVideoStage     int           `json:"showVideoStage"`
						AllowShare         bool          `json:"allowShare"`
						Duration           float64       `json:"duration"`
						ShowSubtitle       string        `json:"showSubtitle"`
						ShowId             string        `json:"showId"`
						VerticalStyle      int           `json:"verticalStyle"`
						ShowImgV           string        `json:"showImgV"`
						IsPugv             bool          `json:"isPugv"`
						VideoPublishTime   string        `json:"videoPublishTime"`
						ShowCategoryId     int           `json:"showCategoryId"`
						AllowDownload      bool          `json:"allowDownload"`
						VideoImgV          string        `json:"videoImgV"`
						ShowName           string        `json:"showName"`
						ShowLongId         int           `json:"showLongId"`
						VideoImg           string        `json:"videoImg"`
						ReportUrl          string        `json:"reportUrl"`
						VideoTitle         string        `json:"videoTitle"`
						IsTracking         bool          `json:"isTracking"`
						EpisodeLast        int           `json:"episodeLast"`
						Paid               int           `json:"paid"`
						AdRestricted       int           `json:"adRestricted"`
						DownloadStatus     []interface{} `json:"downloadStatus"`
						VideoLongId        int           `json:"videoLongId"`
					} `json:"extra"`
					Title string `json:"title"`
				} `json:"data"`
				Id    int  `json:"id"`
				Level int  `json:"level"`
				More  bool `json:"more"`
				Nodes []struct {
					Data struct {
						Session string `json:"session"`
					} `json:"data"`
					Id    int  `json:"id"`
					Level int  `json:"level"`
					More  bool `json:"more"`
					Nodes []struct {
						Data struct {
							Session   string `json:"session"`
							AllowPlay int    `json:"allowPlay"`
							Action    struct {
								Report struct {
									SpmD      string `json:"spmD"`
									ScmD      string `json:"scmD"`
									TrackInfo struct {
										ComponentId         string `json:"component_id"`
										Drawerid            string `json:"drawerid"`
										PvvVid              string `json:"pvv_vid"`
										ComponentInstanceId int    `json:"component_instance_id"`
										Servertime          int64  `json:"servertime"`
										Pageid              string `json:"pageid"`
										PvvSid              string `json:"pvv_sid"`
										CmsReqId            string `json:"cms_req_id"`
									} `json:"trackInfo"`
									ScmC     string `json:"scmC"`
									Arg1     string `json:"arg1"`
									SpmC     string `json:"spmC"`
									SpmAB    string `json:"spmAB"`
									Index    int    `json:"index"`
									PageName string `json:"pageName"`
									ScmAB    string `json:"scmAB"`
								} `json:"report"`
								Type  string `json:"type"`
								Extra struct {
									Text string `json:"text"`
								} `json:"extra,omitempty"`
							} `json:"action"`
							AllowRefresh      int    `json:"allowRefresh"`
							Title             string `json:"title"`
							PageIndex         int    `json:"pageIndex,omitempty"`
							Subtitle          string `json:"subtitle,omitempty"`
							AllowUnionRefresh int    `json:"allowUnionRefresh,omitempty"`
							PageSize          int    `json:"pageSize,omitempty"`
							PositionStyle     int    `json:"positionStyle,omitempty"`
							AllowDownload     bool   `json:"allowDownload,omitempty"`
						} `json:"data"`
						Id    int  `json:"id"`
						Level int  `json:"level"`
						More  bool `json:"more"`
						Nodes []struct {
							Data struct {
								UpdateInfo  string   `json:"updateInfo,omitempty"`
								Heat        string   `json:"heat,omitempty"`
								ScorePrefix string   `json:"scorePrefix,omitempty"`
								Area        []string `json:"area,omitempty"`
								Img         string   `json:"img"`
								SubTitles   []struct {
									Subtitle string `json:"subtitle"`
								} `json:"subTitles,omitempty"`
								Title           string `json:"title"`
								TotalUp         int    `json:"totalUp,omitempty"`
								ShowGenre       string `json:"showGenre,omitempty"`
								ShowReleaseYear int    `json:"showReleaseYear,omitempty"`
								HeatDesc        string `json:"heatDesc,omitempty"`
								IntroTitle      string `json:"introTitle,omitempty"`
								Mark            struct {
									Data struct {
										Color string `json:"color"`
										Text  string `json:"text"`
									} `json:"data"`
									MediaMarkType string `json:"mediaMarkType"`
									MediaMarkEnum struct {
										Name string `json:"name"`
									} `json:"mediaMarkEnum"`
									Type string `json:"type"`
								} `json:"mark,omitempty"`
								ShowImg      bool   `json:"showImg,omitempty"`
								Desc         string `json:"desc,omitempty"`
								SubtitleType string `json:"subtitleType,omitempty"`
								Subtitle     string `json:"subtitle,omitempty"`
								Action       struct {
									Extra struct {
										Type        string `json:"type,omitempty"`
										ShowId      string `json:"showId,omitempty"`
										ExtraParams struct {
											AllowMulti bool `json:"allowMulti"`
										} `json:"extraParams,omitempty"`
									} `json:"extra"`
									Report struct {
										SpmD      string `json:"spmD"`
										ScmD      string `json:"scmD"`
										TrackInfo struct {
											ComponentId         string `json:"component_id"`
											Drawerid            string `json:"drawerid"`
											PvvVid              string `json:"pvv_vid"`
											ComponentInstanceId int    `json:"component_instance_id"`
											Servertime          int64  `json:"servertime"`
											Pageid              string `json:"pageid"`
											PvvSid              string `json:"pvv_sid"`
											CmsReqId            string `json:"cms_req_id"`
											Videoid             int    `json:"videoid,omitempty"`
											MaterialId          string `json:"material_id,omitempty"`
											VideoId             int    `json:"video_id,omitempty"`
										} `json:"trackInfo"`
										ScmC     string `json:"scmC"`
										Arg1     string `json:"arg1"`
										SpmC     string `json:"spmC"`
										SpmAB    string `json:"spmAB"`
										Index    int    `json:"index"`
										PageName string `json:"pageName"`
										ScmAB    string `json:"scmAB"`
									} `json:"report"`
									Type  string `json:"type"`
									Value string `json:"value"`
								} `json:"action,omitempty"`
								PersonId    int    `json:"personId,omitempty"`
								IsAliStar   bool   `json:"isAliStar,omitempty"`
								VersionType []int  `json:"versionType,omitempty"`
								TitleLine   int    `json:"titleLine,omitempty"`
								Stage       int    `json:"stage,omitempty"`
								VideoType   string `json:"videoType,omitempty"`
								SummaryType string `json:"summaryType,omitempty"`
								Paid        int    `json:"paid,omitempty"`
							} `json:"data"`
							Id       int    `json:"id"`
							Level    int    `json:"level"`
							More     bool   `json:"more"`
							Type     int    `json:"type"`
							TypeName string `json:"typeName"`
						} `json:"nodes"`
						Render []struct {
							Data struct {
								BorderBottomWidth int `json:"borderBottomWidth"`
							} `json:"data"`
							Type string `json:"type"`
						} `json:"render"`
						Type     int    `json:"type"`
						TypeName string `json:"typeName"`
					} `json:"nodes"`
					Type     int    `json:"type"`
					TypeName string `json:"typeName"`
				} `json:"nodes"`
				Type     int    `json:"type"`
				TypeName string `json:"typeName"`
			} `json:"data"`
			Headers struct {
				ManufacturingDate string `json:"manufacturingDate"`
				StatC10013        string `json:"stat_c_10013"`
			} `json:"headers"`
			MsCode  string `json:"msCode"`
			Success bool   `json:"success"`
		} `json:"2019030100"`
	} `json:"data"`
	Ret []string `json:"ret"`
	V   string   `json:"v"`
}

type YoukuPlay struct {
	Api  string `json:"api"`
	Data struct {
		Field1 struct {
			Data struct {
				Data struct {
					PageIndex         int    `json:"pageIndex"`
					Session           string `json:"session"`
					Subtitle          string `json:"subtitle"`
					AllowPlay         int    `json:"allowPlay"`
					AllowUnionRefresh int    `json:"allowUnionRefresh"`
					PageSize          int    `json:"pageSize"`
					Action            struct {
						Report struct {
							SpmD      string `json:"spmD"`
							ScmD      string `json:"scmD"`
							TrackInfo struct {
								ComponentId         string `json:"component_id"`
								PvvVid              string `json:"pvv_vid"`
								ComponentInstanceId int    `json:"component_instance_id"`
								Servertime          int64  `json:"servertime"`
								Pageid              string `json:"pageid"`
								PvvSid              string `json:"pvv_sid"`
								CmsReqId            string `json:"cms_req_id"`
							} `json:"trackInfo"`
							ScmC     string `json:"scmC"`
							Arg1     string `json:"arg1"`
							SpmC     string `json:"spmC"`
							SpmAB    string `json:"spmAB"`
							Index    int    `json:"index"`
							PageName string `json:"pageName"`
							ScmAB    string `json:"scmAB"`
						} `json:"report"`
						Type string `json:"type"`
					} `json:"action"`
					PositionStyle int    `json:"positionStyle"`
					AllowRefresh  int    `json:"allowRefresh"`
					Title         string `json:"title"`
					AllowDownload bool   `json:"allowDownload"`
				} `json:"data"`
				Id    int  `json:"id"`
				Level int  `json:"level"`
				More  bool `json:"more"`
				Nodes []struct {
					Data struct {
						VersionType  []int  `json:"versionType"`
						SubtitleType string `json:"subtitleType"`
						Img          string `json:"img"`
						TitleLine    int    `json:"titleLine"`
						Stage        int    `json:"stage"`
						VideoType    string `json:"videoType"`
						SummaryType  string `json:"summaryType"`
						Subtitle     string `json:"subtitle"`
						Paid         int    `json:"paid"`
						Action       struct {
							Extra struct {
								ShowId      string `json:"showId"`
								ExtraParams struct {
									AllowMulti bool `json:"allowMulti"`
								} `json:"extraParams"`
							} `json:"extra"`
							Report struct {
								SpmD      string `json:"spmD"`
								ScmD      string `json:"scmD"`
								TrackInfo struct {
									ComponentId         string `json:"component_id"`
									PvvVid              string `json:"pvv_vid"`
									ComponentInstanceId int    `json:"component_instance_id"`
									Servertime          int64  `json:"servertime"`
									Videoid             int    `json:"videoid"`
									MaterialId          string `json:"material_id"`
									Pageid              string `json:"pageid"`
									PvvSid              string `json:"pvv_sid"`
									VideoId             int    `json:"video_id"`
									CmsReqId            string `json:"cms_req_id"`
								} `json:"trackInfo"`
								ScmC     string `json:"scmC"`
								Arg1     string `json:"arg1"`
								SpmC     string `json:"spmC"`
								SpmAB    string `json:"spmAB"`
								Index    int    `json:"index"`
								PageName string `json:"pageName"`
								ScmAB    string `json:"scmAB"`
							} `json:"report"`
							Type  string `json:"type"`
							Value string `json:"value"`
						} `json:"action"`
						Title string `json:"title"`
						Mark  struct {
							Data struct {
								Color string `json:"color"`
								Text  string `json:"text"`
							} `json:"data"`
							MediaMarkType string `json:"mediaMarkType"`
							MediaMarkEnum struct {
								Name string `json:"name"`
							} `json:"mediaMarkEnum"`
							Type string `json:"type"`
						} `json:"mark,omitempty"`
					} `json:"data"`
					Id       int    `json:"id"`
					Level    int    `json:"level"`
					More     bool   `json:"more"`
					Type     int    `json:"type"`
					TypeName string `json:"typeName"`
				} `json:"nodes"`
				Render []struct {
					Data struct {
						BorderBottomWidth int `json:"borderBottomWidth"`
					} `json:"data"`
					Type string `json:"type"`
				} `json:"render"`
				Type     int    `json:"type"`
				TypeName string `json:"typeName"`
			} `json:"data"`
			Headers struct {
				ManufacturingDate string `json:"manufacturingDate"`
				StatC10013        string `json:"stat_c_10013"`
			} `json:"headers"`
			MsCode  string `json:"msCode"`
			Success bool   `json:"success"`
		} `json:"2019030100"`
	} `json:"data"`
	Ret []string `json:"ret"`
	V   string   `json:"v"`
}
