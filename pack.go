package ogs

type CodeMsg struct {
	Code    interface{} `json:"code"`
	Message Message     `json:"message"`
}

type CodeMsgData struct {
	CodeMsg
	Data interface{} `json:"data"`
}

type CodeMsgDataWithPag struct {
	CodeMsgData
	Pagination Pagination `json:"pagination"`
}
