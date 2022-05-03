package service

import (
	"catvod/model"
	"catvod/service/iqytv"
	"catvod/service/mgtv"
	"catvod/service/qqtv"
	"catvod/service/zjmiao"
)

type TVProvider interface {
	GetHome() (res model.HomeContent)
	GetCategory(typeId string, page int) (res model.Category)
	GetFilter(string) string
	GetDetails([]string) (res []model.VodDetail)
	Search(string) (res []model.VodInfo)
}

func NewTVProvider(p string) TVProvider {
	switch p {
	case "qqtv":
		return &qqtv.QQTV{}
	case "mgtv":
		return &mgtv.MGTV{}
	case "iqiyi":
		return &iqytv.IQYTV{}
	case "zjmiao":
		return &zjmiao.ZJMIAO{}
	default:
		return &qqtv.QQTV{}
	}
}
