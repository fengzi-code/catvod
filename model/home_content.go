package model

type HomeContent struct {
	Code     int        `json:"code"`
	Msg      string     `json:"msg"`
	VodClass []VodClass `json:"class"`
	// Filters interface{} `json:"filters"`
	VodList []VodInfo `json:"list"`
}
