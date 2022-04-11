package service

import (
	"catvod/model"
	"catvod/service/mgtv"
	"catvod/service/qqtv"
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
	case "mgtest":
		return &mgtv.MGTV{}
	case "mgtv":
		return &mgtv.MGTV{}
	default:
		return &qqtv.QQTV{}
	}
}
