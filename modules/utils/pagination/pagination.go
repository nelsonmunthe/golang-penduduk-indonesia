package pagination

import (
	"example/data-penduduk-indonesia/entity"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	PreviousPage interface{} `json:"previous_page"`
	CurrentPage  int         `json:"current_page"`
	NextPage     interface{} `json:"next_page"`
	Total        int         `json:"total"`
	PerPage      int         `json:"per_page"`
}

func Paginate(page int, perPage int, count int) Pagination {
	var nextPage interface{}
	var previousPage interface{}
	if math.Ceil(float64(count)/float64(perPage)) > float64(page) {
		nextPage = page + 1
	}
	if page > 1 {
		previousPage = page - 1
	}
	return Pagination{
		PreviousPage: previousPage,
		CurrentPage:  page,
		NextPage:     nextPage,
		Total:        count,
		PerPage:      perPage,
	}
}

func GetPaginationQueryParam(context *gin.Context) entity.Paginate {
	paginate := entity.Paginate{}
	paginate.CurrentPage = 1
	paginate.PerPage = 3
	paginate.Keyword = ""

	if context.Query("currentPage") != "" {
		paginate.CurrentPage, _ = strconv.Atoi(context.Query("currentPage"))
	}
	if context.Query("perPage") != "" {
		paginate.PerPage, _ = strconv.Atoi(context.Query("perPage"))
	}
	if context.Query("keyword") != "" {
		paginate.Keyword = context.Query("keyword")
	}

	return paginate
}
