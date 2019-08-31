package go_rsp

// 消息
type BaseMessage struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func NewMessage(message string, msgType string) BaseMessage {
	baseMessage := BaseMessage{message, msgType}
	return baseMessage
}
