package service

import (
	"catvod/model"
	"catvod/service/mgtv"
	"catvod/service/qqtv"
)

type TVProvider interface {
	GetHome() (res model.HomeContent)
	GetCategory(typeid string, page int) (res model.Category)
	// GetFilters(string) string
	// GetCategory(string, string, string)
	// GetDetails([]string)
	// Search(string)
}

func NewTVProvider(p string) TVProvider {
	switch p {
	case "qqtv":
		return &qqtv.QQTV{}
	case "mgtest":
		return &mgtv.MGTV{}
	default:
		return &qqtv.QQTV{}
	}
}
