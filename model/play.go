package model

/*
PlayResponse 是前端点播放时，后端响应给前端的数据结构
@Code int 状态码, 传0表示成功, 传不传都可以
@PlayUrl string 播放地址, 这个地址是需要解析的播放地址，当播放地址是直链时，这个地址必须为空
@Url string 当parse为0时，这里传的是视频解析后的直链地址，当parse为1时，这里传的是解析接口地址
@Parse int 为0时，表示播放地址是直链，url设置为解析后的视频直链，playurl需设置为空。为1时，表示由前端解析播放，需把PlayUrl和Url都传给前端
直链示例：
PlayResponse{
	Code: 0,
	PlayUrl: "",
	Url: "https://xx.xx.xx.xx/xx/xx/xx.mp4",	// 也有可能是m3u8格式的
	Parse: 0,
}
由前端解析示例：
PlayResponse{
	Code: 0,
	PlayUrl: "https://v.qq.com/x/cover/mzc00200grqdp0c.html",
	Url: "你的解析接口地址",
	Parse: 1,
}
*/
type PlayResponse struct {
	Code    int               `json:"code"`
	PlayUrl string            `json:"playUrl"`
	Url     string            `json:"url"`
	Parse   int               `json:"parse"` // 是否由前端APP解析，0: 不解析，1: 解析
	Header  map[string]string `json:"header"`
}
