package model

type DateTimeFilter struct {
	Since string `json:"since"`
	Till  string `json:"till"`
}

type Pagination struct {
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
}
