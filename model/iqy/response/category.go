package response

type IqiyiCategory struct {
	Code string `json:"code"`
	Data struct {
		List []struct {
			AlbumId     int64  `json:"albumId"`
			Name        string `json:"name"`
			ImageUrl    string `json:"imageUrl"`
			ChannelId   int    `json:"channelId"`
			Description string `json:"description"`
			LatestOrder int    `json:"latestOrder"`
			Period      string `json:"period"`
			Title       string `json:"title"`
		} `json:"list"`
	} `json:"data"`
}
