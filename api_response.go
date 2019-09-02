package gorsp

/**
{
  message: {
    message: "",
    type: ""
  },
  code: 0,
  data: null
}
*/
type BaseResponse struct {
	Message BaseMessage `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

/**
{
  message: {
    message: "",
    type: ""
  },
  code: 0,
  data: null,
  paginate: {
    current_page: 0,
    total_pages: 0,
    total_count: 0,
    per_page: 0
  }
}
**/
type PaginateResponse struct {
	BaseResponse
	Paginate BasePaginate `json:"paginate"`
}

// response base format
func RspBaseData(data interface{}, code int, baseMessage BaseMessage) interface{} {
	baseResponse := BaseResponse{}
	baseResponse.Data = data
	baseResponse.Code = code
	baseResponse.Message = baseMessage
	return baseResponse
}

// response base format with paginate
func RspPagData(data interface{}, code int, baseMessage BaseMessage, basePaginate BasePaginate) interface{} {
	paginateResponse := PaginateResponse{}
	paginateResponse.Data = data
	paginateResponse.Code = code
	paginateResponse.Message = baseMessage
	paginateResponse.Paginate = basePaginate
	return paginateResponse
}
