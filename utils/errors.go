package utils

import "fmt"

type YunxiaoError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Details string `json:"details,omitempty"`
}

func (e *YunxiaoError) Error() string {
	return fmt.Sprintf("[%s] %s (code: %d) %s", e.Type, e.Message, e.Code, e.Details)
}

func NewNetworkError(statusCode int, err error) *YunxiaoError {
	return &YunxiaoError{
		Type:    "yunxiao",
		Message: "call api error",
		Code:    statusCode,
		Details: err.Error(),
	}
}
