package gorsp

type CodeMessage struct {
	Code    int         `json:"code"`
	Message BaseMessage `json:"message"`
}

type CodeMessageData struct {
	CodeMessage
	Data interface{} `json:"data"`
}

type CodeMessageDataWithPag struct {
	CodeMessageData
	Paginate BasePaginate `json:"paginate"`
}

/**
{
  message: {
    message: "",
    type: ""
  },
  code: 10001
}
*/
func RspBase(code int, message BaseMessage) interface{} {
	codeMessage := CodeMessage{}
	codeMessage.Code = code
	codeMessage.Message = message
	return codeMessage
}

/**
{
  message: {
    message: "",
    type: ""
  },
  code: 0
}
*/
func RspOK(message BaseMessage) interface{} {
	return RspBase(StatusOK, message)
}

/**
{
  message: {
    message: "",
    type: ""
  },
  code: 30001,
  data: null
}
**/
func RspBaseWithData(code int, message BaseMessage, data interface{}) interface{} {
	codeMessageData := CodeMessageData{}
	codeMessageData.Code = code
	codeMessageData.Message = message
	codeMessageData.Data = data
	return codeMessageData
}

/**
{
  message: {
    message: "",
    type: ""
  },
  code: 0,
  data: null
}
**/
func RspOKWithData(message BaseMessage, data interface{}) interface{} {
	return RspBaseWithData(StatusOK, message, data)
}

// response base format with paginate
/**
{
  message: {
    message: "",
    type: ""
  },
  code: 30001,
  data: null,
  paginate: {
    current_page: 0,
    total_pages: 0,
    total_count: 0,
    per_page: 0
  }
}
**/
func RspBaseWithPaginate(code int, message BaseMessage, data interface{}, basePaginate BasePaginate) interface{} {
	codeMessageDataWithPag := CodeMessageDataWithPag{}
	codeMessageDataWithPag.Code = code
	codeMessageDataWithPag.Message = message
	codeMessageDataWithPag.Data = data
	codeMessageDataWithPag.Paginate = basePaginate
	return codeMessageDataWithPag
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
func RspOKWithPaginate(message BaseMessage, data interface{}, basePaginate BasePaginate) interface{} {
	return RspBaseWithPaginate(StatusOK, message, data, basePaginate)
}
