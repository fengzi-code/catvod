package response

type YoukuSearch1 struct {
	Data struct {
		Nodes []struct {
			Nodes []struct {
				Nodes []struct {
					Data struct {
						VideoId string `json:"videoId,omitempty"`
					} `json:"data"`
				} `json:"nodes"`
				Data struct {
					FeatureDTO struct {
						Color string `json:"color"`
						Text  string `json:"text"`
					} `json:"featureDTO"`
					ThumbUrl       string `json:"thumbUrl"`
					RightButtonDTO struct {
						DisplayName string `json:"displayName"`
					} `json:"rightButtonDTO"`
					PosterDTO struct {
						VThumbUrl string `json:"vThumbUrl"`
					} `json:"posterDTO"`
					TempTitle string `json:"tempTitle"`
				} `json:"data,omitempty"`
			} `json:"nodes"`
		} `json:"nodes"`
	} `json:"data"`
}
