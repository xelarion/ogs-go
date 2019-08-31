package go_rsp

// 消息
type BaseMessage struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func NewBaseMessage(message string, msg_type string) BaseMessage {
	baseMessage := BaseMessage{message, msg_type}
	return baseMessage
}
