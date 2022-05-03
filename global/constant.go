package global

const (
	ContentType     = "application/x-www-form-urlencoded;charset=UTF-8"
	UserAgent       = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"
	MgTvCategoryApi = "https://pianku.api.mgtv.com/rider/list/pcweb/v3?platform=pcweb&channelId=%s&pn=%s%s"
	// 第1个%s是频道id，第2个%s是页码，第3个%s是filter
	DefaultPic            = "https://www.hjunkel.com/images/nopic2.gif"
	JsonType              = "application/json"
	MgHomeApi             = "https://rc.mgtv.com/pc/rank?allowedRC=0&t=day&c=2,3&limit=20&_support=10000000"
	MgStaticDir           = "static/mgtv"
	QQStaticDir           = "static/qqtv"
	IQYStaticDir          = "static/iqy"
	ZJStaticDir           = "static/zjmiao"
	MgBaseUrl             = "https://www.mgtv.com"
	IqiyiFilterBaseurl    = "https://list.iqiyi.com/www/"
	IqiyiFilterbaseEndUrl = "/19-----------1964_1979--24-1-1-iqiyi--.html?s_source=PCW_SC"
	IqiyiCatBaseUrl       = "https://pcw-api.iqiyi.com/search/recommend/list?data_type=1&ret_num=10&channel_id="
	IqiyiSobaseUrl        = "https://pcw-api.iqiyi.com/strategy/pcw/data/soBaseCardLeftSide?pageNum=1&site=iqiyi&key="
	MgtvOrigin            = "https://so.mgtv.com"
	MgtvReferer           = "https://so.mgtv.com"
)
