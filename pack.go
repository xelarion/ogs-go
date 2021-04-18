package ogs

// code, message
type codeMessage struct {
	Code    interface{} `json:"code"`
	Message message     `json:"message"`
}

// code, message, data
type codeMsgData struct {
	codeMessage
	Data interface{} `json:"data"`
}

// code, message, data, pagination
type codeMsgDataWithPag struct {
	codeMsgData
	Pagination pagination `json:"pagination"`
}
