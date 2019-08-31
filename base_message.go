package api_response

// 消息
type BaseMessage struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

var msgTypes = [4]string{"info", "success", "warning", "error"}

func NewBaseMessage(message string, msg_type string) BaseMessage {
	baseMessage := BaseMessage{message, msg_type}
	return baseMessage
}

// 验证 message type 是否合法
func verifyMsgType(msg_type string) {
	msgTypes
}
