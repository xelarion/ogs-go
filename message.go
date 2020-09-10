package ogs

type BaseMessage struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

func NewMessage(content string, msgType string) BaseMessage {
	return BaseMessage{content, msgType}
}

func BlankMessage() BaseMessage {
	return BaseMessage{"", ""}
}

func SuccessMessage(content string) BaseMessage {
	return NewMessage(content, "success")
}

func ErrorMessage(content string) BaseMessage {
	return NewMessage(content, "error")
}

func WarningMessage(content string) BaseMessage {
	return NewMessage(content, "warning")
}

func InfoMessage(content string) BaseMessage {
	return NewMessage(content, "info")
}

func blankOrSuccessMessage(content string) BaseMessage {
	if content == "" {
		return BlankMessage()
	} else {
		return SuccessMessage(content)
	}
}
