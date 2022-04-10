package service

import (
	"catvod/model"
	"catvod/service/qqtv"
)

type TVProvider interface {
	GetHome() (res model.HomeContent)
	GetCategory(typeid string, page int) (res model.Category)
	GetFilter(string) string
	GetDetails([]string) (res []model.VodDetail)
	Search(string) (res []model.VodInfo)
}

func NewTVProvider(p string) TVProvider {
	switch p {
	case "qqtv":
		return &qqtv.QQTV{}
	default:
		return &qqtv.QQTV{}
	}
}
