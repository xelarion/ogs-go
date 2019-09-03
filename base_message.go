package ogs

type BaseMessage struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func NewMessage(message string, msgType string) BaseMessage {
	baseMessage := BaseMessage{message, msgType}
	return baseMessage
}

func BlankMessage() BaseMessage {
	baseMessage := BaseMessage{"", ""}
	return baseMessage
}

func SuccessMessage(message string) BaseMessage {
	return NewMessage(message, "success")
}

func ErrorMessage(message string) BaseMessage {
	return NewMessage(message, "error")
}

func WarningMessage(message string) BaseMessage {
	return NewMessage(message, "warning")
}

func InfoMessage(message string) BaseMessage {
	return NewMessage(message, "info")
}