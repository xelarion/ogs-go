package ogs

import "math"

type paginate struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
	TotalCount  int `json:"total_count"`
	PerPage     int `json:"per_page"`
}

func NewPag(currentPage int, totalCount int, perPage int) paginate {
	return paginate{
		CurrentPage: currentPage,
		TotalCount:  totalCount,
		PerPage:     perPage,
		TotalPages:  int(math.Ceil(float64(totalCount) / float64(perPage))),
	}
}
