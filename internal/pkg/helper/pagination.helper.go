package helper

import (
	"net/http"
	"strconv"
)

type PaginationParam struct {
	Page  int
	Limit int
}

func GetPagination(r *http.Request) PaginationParam {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	return PaginationParam{
		Page:  page,
		Limit: limit,
	}
}

type PaginationMeta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

type PaginationResult struct {
	Data any            `json:"data"`
	Meta PaginationMeta `json:"meta"`
}
