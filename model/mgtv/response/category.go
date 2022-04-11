package response

type Category struct { // mgtv返回的分类json
	Code int `json:"code"`
	Data struct {
		Template  string `json:"template"`
		TotalHits int    `json:"totalHits"`
		Advs      []struct {
			Row string `json:"row"`
			Id  string `json:"id"`
		} `json:"advs"`
		HitDocs []struct {
			Schema      string `json:"schema"`
			UpdateInfo  string `json:"updateInfo"`
			JumpKind    string `json:"jumpKind"`
			Img         string `json:"img"`
			RightCorner struct {
				Color string `json:"color,omitempty"`
				Text  string `json:"text,omitempty"`
			} `json:"rightCorner"`
			Year         string   `json:"year"`
			PartId       string   `json:"partId"`
			Kind         []string `json:"kind"`
			Prefix       string   `json:"prefix"`
			FstlvlId     string   `json:"fstlvlId"`
			Title        string   `json:"title"`
			Params       string   `json:"params"`
			SeUpdateTime string   `json:"se_updateTime"`
			PlayPartId   string   `json:"playPartId"`
			Subtitle     string   `json:"subtitle"`
			Rpt          string   `json:"rpt"`
			ClipId       string   `json:"clipId"`
			Ic           string   `json:"ic"`
			ZhihuScore   string   `json:"zhihuScore"`
			Views        string   `json:"views"`
			ImgUrlH      string   `json:"imgUrlH"`
			Story        string   `json:"story"`
		} `json:"hitDocs"`
		SearchKeys []interface{} `json:"searchKeys"`
	} `json:"data"`
	Msg   string `json:"msg"`
	Seqid string `json:"seqid"`
}
