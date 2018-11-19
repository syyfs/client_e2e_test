package util

import "encoding/json"

const (
	SUCCESS = "success"

	ERROR = "error"
)

type Response struct {
	Code    int
	Message string
	Data    []byte
}

func GetResponse(code int, message string, data []byte) []byte {
	response := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	responseByte, _ := json.Marshal(response)
	return responseByte
}
