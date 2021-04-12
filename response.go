package ogs

// Rsp return code and message
func Rsp(code interface{}, message message) interface{} {
	return codeMessage{Code: code, Message: message}
}

// RspData return code, message and data
func RspData(code interface{}, message message, data interface{}) interface{} {
	r := codeMsgData{}
	r.Code = code
	r.Message = message
	r.Data = data
	return r
}

// RspDataPag return code, message and data with pagination
func RspDataPag(code interface{}, message message, data interface{}, pag paginate) interface{} {
	r := codeMsgDataWithPag{}
	r.Code = code
	r.Message = message
	r.Data = data
	r.Paginate = pag
	return r
}

// RspOK return +CodeOK+ code and message
func RspOK(msgContent string) interface{} {
	return Rsp(CodeOK, blankOrSuccessMsg(msgContent))
}

// RspError return error code and message
func RspError(code interface{}, msgContent string) interface{} {
	return Rsp(code, ErrorMsg(msgContent))
}

// RspDataOK return +CodeOK+ code, message and data
func RspDataOK(msgContent string, data interface{}) interface{} {
	return RspData(
		CodeOK,
		blankOrSuccessMsg(msgContent),
		data,
	)
}

// RspDataPagOK return +CodeOK+ code, message and data with pagination
func RspDataPagOK(msgContent string, data interface{}, pag paginate) interface{} {
	return RspDataPag(
		CodeOK,
		blankOrSuccessMsg(msgContent),
		data,
		pag,
	)
}
