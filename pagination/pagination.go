package pagination

import (
	"net/http"
	"strconv"
)

type Pagination struct {
	Total   int `json:"total"`
	PerPage int `json:"per_page"`
	PageNo  int `json:"page_no"`
}

func Paginate(r *http.Request) (int, int, int) {
	pageNo := 1
	limitNo := 5
	page := r.FormValue("page")
	limit := r.FormValue("limit")
	if limit != "" {
		i, _ := strconv.Atoi(limit)
		limitNo = i
	}

	if page != "" {
		i, _ := strconv.Atoi(page)
		pageNo = i
	}

	offset := (pageNo - 1) * limitNo
	return offset, limitNo, pageNo
}
