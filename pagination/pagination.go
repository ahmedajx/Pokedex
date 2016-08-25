package pagination

import (
	"log"
	"net/http"
	"strconv"
)

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
	log.Println(offset)
	return offset, limitNo, pageNo
}
