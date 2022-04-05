package model

/*
	详情分两种情况，一种是
*/

type VodDetail struct {
	VodInfo            // 基本信息
	TypeName    string `json:"type_name"`     // 类型名称: 如剧情片
	VodYear     string `json:"vod_year"`      // 年份	: 如2017
	VodArea     string `json:"vod_area"`      // 地区	: 如大陆
	VodActor    string `json:"vod_actor"`     // 演员	: 如陈凯歌
	VodDirector string `json:"vod_director"`  // 导演	: 如陈凯歌
	VodContent  string `json:"vod_content"`   // 影视简介
	VodPlayFrom string `json:"vod_play_from"` // 播放来源
	VodPlayUrl  string `json:"vod_play_url"`  // 播放地址
}
