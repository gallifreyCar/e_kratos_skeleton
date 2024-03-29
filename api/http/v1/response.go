package v1

type CommonResponse struct {
	// Code 状态码
	Code int `json:"code"`
	// Message 消息
	Message string `json:"message"`
	// Data 数据
	Data interface{} `json:"data"`
}
