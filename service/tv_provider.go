package service

import (
	"catvod/model"
	"catvod/service/cokemv"
	"catvod/service/dy555"
	"catvod/service/iqytv"
	"catvod/service/lgyy"
	"catvod/service/mgtv"
	"catvod/service/qqtv"
	"catvod/service/rxlg"
	"catvod/service/rxzh"
	"catvod/service/youku"
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
	case "dy555":
		return &dy555.DY555{}
	case "youku":
		return &youku.YOUKU{}
	case "lgyy":
		return &lgyy.LGYY{}
	case "cokemv":
		return &cokemv.COKEMV{}
	case "rxzh":
		return &rxzh.RXZH{}
	case "rxlg":
		return &rxlg.RXLG{}
	default:
		return &dy555.DY555{}
	}
}
