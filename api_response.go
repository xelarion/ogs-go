package api_response

// 基础返回格式
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
	Code int `json:"code"`
	Data interface{} `json:"data"`
}

// 带分页的返回格式
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

func (br BaseResponse) Success() bool {
	return br.Code == 0 && br.Message.Type == "success"
}

// 返回基础数据格式
func RspBaseData(data interface{}, code int, baseMessage BaseMessage) interface{} {
	baseResponse := BaseResponse{}
	baseResponse.Data = data
	baseResponse.Code = code
	baseResponse.Message = baseMessage
	return baseResponse
}

// 返回分页数据格式
func RspPagData(data interface{}, ) interface{} {
	paginateResponse := PaginateResponse{}
	paginateResponse.Data = data
	return paginateResponse
}