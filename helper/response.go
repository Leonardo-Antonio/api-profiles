package helper

type Response struct {
	MessageType string      `json:"message_type" xml:"message_type"`
	Message     string      `json:"message" xml:"message"`
	Error       bool        `json:"error" xml:"error"`
	Data        interface{} `json:"data" xml:"data"`
}

func ResponseJSON(
	messageType, message string,
	err bool,
	data interface{},
) *Response {
	return &Response{
		MessageType: messageType,
		Message:     message,
		Error:       err,
		Data:        data,
	}
}
