package ogs

import "math"

type BasePaginate struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
	TotalCount  int `json:"total_count"`
	PerPage     int `json:"per_page"`
}

func NewPaginate(currentPage int, totalCount int, perPage int) BasePaginate {
	basePaginate := BasePaginate{}
	basePaginate.CurrentPage = currentPage
	basePaginate.TotalCount = totalCount
	basePaginate.PerPage = perPage
	basePaginate.TotalPages = int(math.Ceil(float64(totalCount) / float64(perPage)))
	return basePaginate
}
