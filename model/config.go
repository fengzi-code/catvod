package model

// ServerConfig 这个结构体是用来存放1.json里的内容的
type ServerConfig struct {
	Sites []struct {
		Key         string `json:"key"`
		Name        string `json:"name"`
		Type        int    `json:"type"`
		Api         string `json:"api"`
		Searchable  int    `json:"searchable"`
		QuickSearch int    `json:"quickSearch"`
		Filterable  int    `json:"filterable"`
	} `json:"sites"`
	Parses []struct {
		Name string `json:"name"`
		Type int    `json:"type"`
		Url  string `json:"url"`
		Ext  struct {
			Flag []string `json:"flag"`
		} `json:"ext,omitempty"`
	} `json:"parses"`
	Flags []string `json:"flags"`
	Ijk   []struct {
		Group   string `json:"group"`
		Options []struct {
			Category int    `json:"category"`
			Name     string `json:"name"`
			Value    string `json:"value"`
		} `json:"options"`
	} `json:"ijk"`
	Ads       []string `json:"ads"`
	Wallpaper string   `json:"wallpaper"`
	Spider    string   `json:"spider"`
}
