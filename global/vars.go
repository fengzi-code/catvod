package global

var (
	// Headers http请求头
	Headers = map[string]string{
		"User-Agent":   UserAgent,
		"Content-Type": ContentType,
		"origin":       "",
		"referer":      "",
		"Host":         "",
	}
	BindAddr string
	//  AppMode  app运行模式
	AppMode       string
	IsRongxin     bool
	RongXinKey    string
	RongXinJxUrl  string
	RongXingVrUrl string
)
