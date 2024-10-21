package utils

const (
	OK            uint32 = 10000
	ParamInvalid         = 10001
	BusinessError        = 10002
)

const (
	Success = "Success"
)

type BaseResponse struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}
