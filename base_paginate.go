package go_rsp

// 分页
type BasePaginate struct {
	CurrentPage int `json:"current_page"`
	TotalPages int `json:"total_pages"`
	TotalCount int `json:"total_count"`
	PerPage int `json:"per_page"`
}