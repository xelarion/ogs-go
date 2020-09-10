package ogs

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
    content: "",
    type: ""
  },
  code: 10001
}
*/
func RspBase(code int, message BaseMessage) interface{} {
	return CodeMessage{Code: code, Message: message}
}

/**
{
  message: {
    content: "Invalid Token",
    type: "error"
  },
  code: 10003
}
*/
func RspError(code int, messageContent string) interface{} {
	return RspBase(code, ErrorMessage(messageContent))
}

/**
{
  message: {
    content: "",
    type: ""
  },
  code: 0
}
*/
func RspOK(messageContent string) interface{} {
	message := blankOrSuccessMessage(messageContent)
	return RspBase(StatusOK, message)
}

/**
{
  message: {
    content: "",
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
    content: "",
    type: ""
  },
  code: 0,
  data: null
}
**/
func RspOKWithData(messageContent string, data interface{}) interface{} {
	message := blankOrSuccessMessage(messageContent)
	return RspBaseWithData(StatusOK, message, data)
}

/**
{
  message: {
    content: "",
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
    content: "",
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
func RspOKWithPaginate(messageContent string, data interface{}, basePaginate BasePaginate) interface{} {
	message := blankOrSuccessMessage(messageContent)
	return RspBaseWithPaginate(StatusOK, message, data, basePaginate)
}
