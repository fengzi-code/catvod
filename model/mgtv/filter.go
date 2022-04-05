package mgtv

type FilterIds struct {
	Id1
	Id2
	Id3
	Id10
	Id50
	Id51
	Id116
}

type Id2 struct {
	Id []struct {
		Key   string `json:"key"`
		Name  string `json:"name"`
		Value []struct {
			N string `json:"n"`
			V string `json:"v"`
		} `json:"value"`
	} `json:"2"`
}

type Id3 struct {
	Field1 []struct {
		Key   string `json:"key"`
		Name  string `json:"name"`
		Value []struct {
			N string `json:"n"`
			V string `json:"v"`
		} `json:"value"`
	} `json:"3"`
}

type Id1 struct {
	Field1 []struct {
		Key   string `json:"key"`
		Name  string `json:"name"`
		Value []struct {
			N string `json:"n"`
			V string `json:"v"`
		} `json:"value"`
	} `json:"1"`
}

type Id10 struct {
	Field1 []struct {
		Key   string `json:"key"`
		Name  string `json:"name"`
		Value []struct {
			N string `json:"n"`
			V string `json:"v"`
		} `json:"value"`
	} `json:"10"`
}

type Id50 struct {
	Field1 []struct {
		Key   string `json:"key"`
		Name  string `json:"name"`
		Value []struct {
			N string `json:"n"`
			V string `json:"v"`
		} `json:"value"`
	} `json:"50"`
}

type Id51 struct {
	Field1 []struct {
		Key   string `json:"key"`
		Name  string `json:"name"`
		Value []struct {
			N string `json:"n"`
			V string `json:"v"`
		} `json:"value"`
	} `json:"51"`
}

type Id116 struct {
	Field1 []struct {
		Key   string `json:"key"`
		Name  string `json:"name"`
		Value []struct {
			N string `json:"n"`
			V string `json:"v"`
		} `json:"value"`
	} `json:"116"`
}
