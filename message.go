package ogs

type message struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

func NewMsg(content string, msgType string) message {
	return message{content, msgType}
}

func BlankMsg() message {
	return message{"", ""}
}

func InfoMsg(content string) message {
	return NewMsg(content, "info")
}

func SuccessMsg(content string) message {
	return NewMsg(content, "success")
}

func WarningMsg(content string) message {
	return NewMsg(content, "warning")
}

func ErrorMsg(content string) message {
	return NewMsg(content, "error")
}

func blankOrSuccessMsg(content string) message {
	if content == "" {
		return BlankMsg()
	} else {
		return SuccessMsg(content)
	}
}
