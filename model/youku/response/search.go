package response

type YoukuSearch1 struct {
	Data struct {
		Nodes []struct {
			Nodes []struct {
				Nodes []struct {
					Data struct {
						VideoId string `json:"videoId,omitempty"`
					} `json:"data"`
				} `json:"nodes"`
				Data struct {
					FeatureDTO struct {
						Color string `json:"color"`
						Text  string `json:"text"`
					} `json:"featureDTO"`
					ThumbUrl       string `json:"thumbUrl"`
					RightButtonDTO struct {
						DisplayName string `json:"displayName"`
					} `json:"rightButtonDTO"`
					PosterDTO struct {
						VThumbUrl string `json:"vThumbUrl"`
					} `json:"posterDTO"`
					TempTitle string `json:"tempTitle"`
				} `json:"data,omitempty"`
			} `json:"nodes"`
		} `json:"nodes"`
	} `json:"data"`
}

type YoukuSearch struct {
	Api  string `json:"api"`
	Data struct {
		Data struct {
			DisplaySearchTerms string `json:"displaySearchTerms"`
			Pz                 int    `json:"pz"`
			Rstate             int    `json:"rstate"`
			IsEnd              int    `json:"isEnd"`
			IsModule           bool   `json:"isModule"`
			HighlightWord      string `json:"highlightWord"`
			Filter             struct {
				Duration []struct {
					Title string `json:"title"`
					Value string `json:"value"`
				} `json:"duration"`
				Cate []struct {
					Id    string `json:"id"`
					Title string `json:"title"`
				} `json:"cate"`
				NewType bool `json:"newType"`
				Tab     []struct {
					Title  string `json:"title"`
					Orders []struct {
						Title string `json:"title"`
						Value int    `json:"value"`
					} `json:"orders,omitempty"`
					Id string `json:"id"`
				} `json:"tab"`
				Format []struct {
					Title string `json:"title"`
					Value int    `json:"value"`
				} `json:"format"`
			} `json:"filter"`
			Total       int    `json:"total"`
			TypeVersion string `json:"typeVersion"`
			Pg          int    `json:"pg"`
			Action      struct {
				Report struct {
					TrackInfo struct {
						Searchtab  string `json:"searchtab"`
						SourceFrom string `json:"source_from"`
						SearchFrom string `json:"search_from"`
						Engine     string `json:"engine"`
						Aaid       string `json:"aaid"`
						K          string `json:"k"`
					} `json:"trackInfo"`
					PageName string `json:"pageName"`
				} `json:"report"`
			} `json:"action"`
			Status string `json:"status"`
		} `json:"data"`
		Level int  `json:"level"`
		More  bool `json:"more"`
		Type  int  `json:"type"`
		Nodes []struct {
			Data struct {
				DocSource int `json:"docSource"`
			} `json:"data"`
			Level int  `json:"level"`
			More  bool `json:"more"`
			Type  int  `json:"type"`
			Nodes []struct {
				Level int  `json:"level"`
				More  bool `json:"more"`
				Type  int  `json:"type"`
				Nodes []struct {
					Data struct {
						SourceId       int         `json:"sourceId,omitempty"`
						HasYouku       int         `json:"hasYouku,omitempty"`
						RealShowId     string      `json:"realShowId,omitempty"`
						DocId          interface{} `json:"docId,omitempty"`
						EpisodeTotal   int         `json:"episodeTotal,omitempty"`
						EpisodeCollect int         `json:"episodeCollect,omitempty"`
						Reputation     float64     `json:"reputation,omitempty"`
						FeatureDTO     struct {
							Color string `json:"color"`
							Text  string `json:"text"`
						} `json:"featureDTO,omitempty"`
						UgcSupply   int    `json:"ugcSupply,omitempty"`
						EpisodeType int    `json:"episodeType,omitempty"`
						ItemLog     string `json:"itemLog,omitempty"`
						ShowId      string `json:"showId,omitempty"`
						Exclusive   int    `json:"exclusive,omitempty"`
						Action      struct {
							Type   string `json:"type"`
							Report struct {
								TrackInfo struct {
									ObjectType      int         `json:"object_type"`
									GroupNum        int         `json:"group_num"`
									RecommendSource string      `json:"recommend_source,omitempty"`
									ObjectTitle     string      `json:"object_title"`
									GroupType       int         `json:"group_type"`
									ItemLog         string      `json:"item_log"`
									ObjectUrl       string      `json:"object_url"`
									GroupId         string      `json:"group_id"`
									RecommendTitle  string      `json:"recommend_title,omitempty"`
									SourceId        int         `json:"source_id,omitempty"`
									Cid             int64       `json:"cid,omitempty"`
									SourceFrom      string      `json:"source_from,omitempty"`
									K               string      `json:"k,omitempty"`
									ResultProjectid interface{} `json:"result_projectid"`
									SearchFrom      string      `json:"search_from,omitempty"`
									Aaid            string      `json:"aaid,omitempty"`
									ObjectNum       int         `json:"object_num,omitempty"`
								} `json:"trackInfo"`
								Arg1 string `json:"arg1,omitempty"`
								Scm  string `json:"scm,omitempty"`
								Spm  string `json:"spm,omitempty"`
							} `json:"report"`
							Value string `json:"value"`
						} `json:"action"`
						NeedLess       bool   `json:"needLess,omitempty"`
						ThumbUrl       string `json:"thumbUrl,omitempty"`
						RightButtonDTO struct {
							IsDownload  bool   `json:"isDownload"`
							ButtonType  int    `json:"buttonType"`
							DisplayName string `json:"displayName"`
							Action      struct {
								Report struct {
									Arg1      string `json:"arg1"`
									Scm       string `json:"scm"`
									Spm       string `json:"spm"`
									TrackInfo struct {
										ObjectUrl   string `json:"object_url"`
										GroupId     string `json:"group_id"`
										ObjectType  int    `json:"object_type"`
										GroupNum    int    `json:"group_num"`
										ObjectTitle string `json:"object_title"`
										Isplay      int    `json:"isplay"`
										GroupType   int    `json:"group_type"`
									} `json:"trackInfo"`
								} `json:"report"`
								Type  string `json:"type"`
								Value string `json:"value"`
							} `json:"action"`
						} `json:"rightButtonDTO,omitempty"`
						Notice        string `json:"notice,omitempty"`
						Info          string `json:"info,omitempty"`
						AllVipEpisode string `json:"allVipEpisode,omitempty"`
						LeftButtonDTO struct {
							ButtonType  string `json:"buttonType"`
							DisplayName string `json:"displayName"`
							Action      struct {
								Report struct {
									Arg1      string `json:"arg1"`
									Scm       string `json:"scm"`
									Spm       string `json:"spm"`
									TrackInfo struct {
										ObjectTitle string `json:"object_title"`
									} `json:"trackInfo"`
								} `json:"report"`
								Type  string `json:"type"`
								Value string `json:"value"`
							} `json:"action"`
						} `json:"leftButtonDTO,omitempty"`
						LogCate      string `json:"logCate,omitempty"`
						RecommendDTO struct {
							RecommendLeftIcon string `json:"recommendLeftIcon"`
							ShowType          int    `json:"showType"`
							Recommends        []struct {
								DisplayName string `json:"displayName"`
								Action      struct {
									Type   string `json:"type"`
									Report struct {
										TrackInfo struct {
											ObjectType  int    `json:"object_type"`
											ObjectTitle string `json:"object_title"`
											GroupType   int    `json:"group_type"`
											ItemLog     string `json:"item_log"`
											ObjectNum   int    `json:"object_num"`
										} `json:"trackInfo"`
										Spm  string `json:"spm"`
										Arg1 string `json:"arg1"`
										Scm  string `json:"scm"`
									} `json:"report"`
								} `json:"action"`
							} `json:"recommends"`
						} `json:"recommendDTO,omitempty"`
						Director  string `json:"director,omitempty"`
						PosterDTO struct {
							IconCorner struct {
								TagText string `json:"tagText"`
								TagType int    `json:"tagType"`
							} `json:"iconCorner"`
							VThumbUrl  string  `json:"vThumbUrl"`
							PosterType string  `json:"posterType"`
							Reputation float64 `json:"reputation"`
							Action     struct {
								Type   string `json:"type"`
								Report struct {
									Spm  string `json:"spm"`
									Arg1 string `json:"arg1"`
									Scm  string `json:"scm"`
								} `json:"report"`
								Value string `json:"value"`
							} `json:"action"`
						} `json:"posterDTO,omitempty"`
						Completed            int    `json:"completed,omitempty"`
						RecommendDisplayType string `json:"recommendDisplayType,omitempty"`
						TitleDTO             struct {
							DisplayName string `json:"displayName"`
							Action      struct {
								Type   string `json:"type,omitempty"`
								Report struct {
									Spm  string `json:"spm"`
									Arg1 string `json:"arg1"`
									Scm  string `json:"scm"`
								} `json:"report"`
								Value string `json:"value,omitempty"`
							} `json:"action"`
						} `json:"titleDTO,omitempty"`
						DocSource          int    `json:"docSource,omitempty"`
						Cats               string `json:"cats,omitempty"`
						TempTitle          string `json:"tempTitle,omitempty"`
						AnimeEdition       int    `json:"animeEdition,omitempty"`
						CateId             int    `json:"cateId,omitempty"`
						Paid               int    `json:"paid,omitempty"`
						IsYouku            int    `json:"isYouku,omitempty"`
						SourceName         string `json:"sourceName,omitempty"`
						ShowIdValid        int    `json:"showIdValid,omitempty"`
						IsTrailer          int    `json:"isTrailer,omitempty"`
						ProgramId          string `json:"programId,omitempty"`
						MediaCompleted     int    `json:"mediaCompleted,omitempty"`
						StripeBottom       string `json:"stripeBottom,omitempty"`
						DisplayName        string `json:"displayName,omitempty"`
						DownloadStatus     int    `json:"downloadStatus,omitempty"`
						Limit              []int  `json:"limit,omitempty"`
						PeripheryVideoType int    `json:"peripheryVideoType,omitempty"`
						Pic                bool   `json:"pic,omitempty"`
						Seconds            string `json:"seconds,omitempty"`
						ShowVideoStage     string `json:"showVideoStage,omitempty"`
						TagType            int    `json:"tagType,omitempty"`
						Title              string `json:"title,omitempty"`
						Type               int    `json:"type,omitempty"`
						Vid                string `json:"vid,omitempty"`
						VideoId            string `json:"videoId,omitempty"`
						JumpType           int    `json:"jumpType,omitempty"`
						PublishTime        string `json:"publishTime,omitempty"`
						DownloadLimit      int    `json:"downloadLimit,omitempty"`
						PublicType         int    `json:"publicType,omitempty"`
						ScreenShotDTO      struct {
							IconCorner struct {
								TagType int `json:"tagType"`
							} `json:"iconCorner"`
							RightBottomText string `json:"rightBottomText"`
							Action          struct {
								Report struct {
									Spm  string `json:"spm"`
									Arg1 string `json:"arg1"`
									Scm  string `json:"scm"`
								} `json:"report"`
							} `json:"action"`
							ThumbUrl   string `json:"thumbUrl"`
							NewImgInfo struct {
								IsCut    bool   `json:"isCut"`
								Log      string `json:"log"`
								Scale    string `json:"scale"`
								ThumbUrl string `json:"thumbUrl"`
							} `json:"newImgInfo"`
						} `json:"screenShotDTO,omitempty"`
						SourceImg string `json:"sourceImg,omitempty"`
						UserName  string `json:"userName,omitempty"`
						UserId    string `json:"userId,omitempty"`
					} `json:"data"`
					Level int  `json:"level"`
					More  bool `json:"more"`
					Type  int  `json:"type"`
					Style []struct {
						Type int `json:"type"`
					} `json:"style"`
					Id int `json:"id"`
				} `json:"nodes"`
				Style []struct {
					Type int `json:"type"`
				} `json:"style"`
				Id   int `json:"id"`
				Data struct {
					SourceId       int     `json:"sourceId"`
					HasYouku       int     `json:"hasYouku"`
					RealShowId     string  `json:"realShowId"`
					DocId          string  `json:"docId"`
					EpisodeTotal   int     `json:"episodeTotal"`
					EpisodeCollect int     `json:"episodeCollect"`
					Reputation     float64 `json:"reputation"`
					FeatureDTO     struct {
						Color string `json:"color"`
						Text  string `json:"text"`
					} `json:"featureDTO"`
					UgcSupply   int    `json:"ugcSupply"`
					EpisodeType int    `json:"episodeType"`
					ItemLog     string `json:"itemLog"`
					ShowId      string `json:"showId"`
					Exclusive   int    `json:"exclusive"`
					Action      struct {
						Type   string `json:"type"`
						Report struct {
							TrackInfo struct {
								ObjectType      int    `json:"object_type"`
								GroupNum        int    `json:"group_num"`
								RecommendSource string `json:"recommend_source"`
								ObjectTitle     string `json:"object_title"`
								GroupType       int    `json:"group_type"`
								ItemLog         string `json:"item_log"`
								ObjectUrl       string `json:"object_url"`
								GroupId         string `json:"group_id"`
								RecommendTitle  string `json:"recommend_title"`
								SourceId        int    `json:"source_id"`
								Cid             int64  `json:"cid"`
							} `json:"trackInfo"`
						} `json:"report"`
						Value string `json:"value"`
					} `json:"action"`
					NeedLess       bool   `json:"needLess"`
					ThumbUrl       string `json:"thumbUrl"`
					RightButtonDTO struct {
						IsDownload  bool   `json:"isDownload"`
						ButtonType  int    `json:"buttonType"`
						DisplayName string `json:"displayName"`
						Action      struct {
							Report struct {
								Arg1      string `json:"arg1"`
								Scm       string `json:"scm"`
								Spm       string `json:"spm"`
								TrackInfo struct {
									ObjectUrl   string `json:"object_url"`
									GroupId     string `json:"group_id"`
									ObjectType  int    `json:"object_type"`
									GroupNum    int    `json:"group_num"`
									ObjectTitle string `json:"object_title"`
									Isplay      int    `json:"isplay"`
									GroupType   int    `json:"group_type"`
								} `json:"trackInfo"`
							} `json:"report"`
							Type  string `json:"type"`
							Value string `json:"value"`
						} `json:"action"`
					} `json:"rightButtonDTO"`
					Info          string `json:"info"`
					AllVipEpisode string `json:"allVipEpisode"`
					LeftButtonDTO struct {
						ButtonType  string `json:"buttonType"`
						DisplayName string `json:"displayName"`
						Action      struct {
							Report struct {
								Arg1      string `json:"arg1"`
								Scm       string `json:"scm"`
								Spm       string `json:"spm"`
								TrackInfo struct {
									ObjectTitle string `json:"object_title"`
								} `json:"trackInfo"`
							} `json:"report"`
							Type  string `json:"type"`
							Value string `json:"value"`
						} `json:"action"`
					} `json:"leftButtonDTO"`
					LogCate      string `json:"logCate"`
					RecommendDTO struct {
						RecommendLeftIcon string `json:"recommendLeftIcon"`
						ShowType          int    `json:"showType"`
						Recommends        []struct {
							DisplayName string `json:"displayName"`
							Action      struct {
								Type   string `json:"type"`
								Report struct {
									TrackInfo struct {
										ObjectType  int    `json:"object_type"`
										ObjectTitle string `json:"object_title"`
										GroupType   int    `json:"group_type"`
										ItemLog     string `json:"item_log"`
										ObjectNum   int    `json:"object_num"`
									} `json:"trackInfo"`
									Spm  string `json:"spm"`
									Arg1 string `json:"arg1"`
									Scm  string `json:"scm"`
								} `json:"report"`
							} `json:"action"`
						} `json:"recommends"`
					} `json:"recommendDTO"`
					Director     string `json:"director"`
					StripeBottom string `json:"stripeBottom"`
					PosterDTO    struct {
						IconCorner struct {
							TagText string `json:"tagText"`
							TagType int    `json:"tagType"`
						} `json:"iconCorner"`
						VThumbUrl  string  `json:"vThumbUrl"`
						PosterType string  `json:"posterType"`
						Reputation float64 `json:"reputation"`
						Action     struct {
							Type   string `json:"type"`
							Report struct {
								Spm  string `json:"spm"`
								Arg1 string `json:"arg1"`
								Scm  string `json:"scm"`
							} `json:"report"`
							Value string `json:"value"`
						} `json:"action"`
					} `json:"posterDTO"`
					Completed            int    `json:"completed"`
					RecommendDisplayType string `json:"recommendDisplayType"`
					TitleDTO             struct {
						DisplayName string `json:"displayName"`
						Action      struct {
							Type   string `json:"type"`
							Report struct {
								Spm  string `json:"spm"`
								Arg1 string `json:"arg1"`
								Scm  string `json:"scm"`
							} `json:"report"`
							Value string `json:"value"`
						} `json:"action"`
					} `json:"titleDTO"`
					DocSource      int    `json:"docSource"`
					Cats           string `json:"cats"`
					TempTitle      string `json:"tempTitle"`
					AnimeEdition   int    `json:"animeEdition"`
					CateId         int    `json:"cateId"`
					Paid           int    `json:"paid"`
					IsYouku        int    `json:"isYouku"`
					SourceName     string `json:"sourceName"`
					ShowIdValid    int    `json:"showIdValid"`
					IsTrailer      int    `json:"isTrailer"`
					ProgramId      string `json:"programId"`
					MediaCompleted int    `json:"mediaCompleted"`
				} `json:"data,omitempty"`
			} `json:"nodes"`
			Style []struct {
				Type int `json:"type"`
			} `json:"style"`
			Id int `json:"id"`
		} `json:"nodes"`
		Id int `json:"id"`
	} `json:"data"`
	Ret []string `json:"ret"`
	V   string   `json:"v"`
}
