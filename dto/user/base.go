package user

type Page struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
	Limit     int `json:"limit"`
	Offset    int `json:"offset"`
}

type PageOutput struct {
	Total     int64 `json:"total"`
	PageIndex int   `json:"page_index"`
}
