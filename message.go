package ogs

type BaseMessage struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

func NewMessage(content string, msgType string) BaseMessage {
	baseMessage := BaseMessage{content, msgType}
	return baseMessage
}

func BlankMessage() BaseMessage {
	baseMessage := BaseMessage{"", ""}
	return baseMessage
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
