package ogs

type Message struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

func NewMsg(content string, msgType string) Message {
	return Message{content, msgType}
}

func BlankMsg() Message {
	return Message{"", ""}
}

func InfoMsg(content string) Message {
	return NewMsg(content, "info")
}

func SuccessMsg(content string) Message {
	return NewMsg(content, "success")
}

func WarningMsg(content string) Message {
	return NewMsg(content, "warning")
}

func ErrorMsg(content string) Message {
	return NewMsg(content, "error")
}

func blankOrSuccessMsg(content string) Message {
	if content == "" {
		return BlankMsg()
	} else {
		return SuccessMsg(content)
	}
}
