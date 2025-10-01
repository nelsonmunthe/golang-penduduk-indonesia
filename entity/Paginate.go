package entity

type Paginate struct {
	CurrentPage int    `json:"currentPage"`
	PerPage     int    `json:"perPage"`
	Keyword     string `json:"keyword"`
}
