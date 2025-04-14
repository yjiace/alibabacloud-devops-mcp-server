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

func NewNetworkError(err error) *YunxiaoError {
	return &YunxiaoError{
		Type:    "yunxiao",
		Message: "call api error",
		Details: err.Error(),
	}
}
