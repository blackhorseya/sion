package errorx

import (
	"fmt"
)

// APPError declare custom error
type APPError struct {
	Status int         `json:"-"`
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}

func (e *APPError) Error() string {
	return e.Msg
}

// ReplaceMsg replace message into APPError
func (e *APPError) ReplaceMsg(msg string) *APPError {
	return &APPError{
		Status: e.Status,
		Code:   e.Code,
		Msg:    msg,
		Data:   e.Data,
	}
}

// AppendMsg append message into APPError
func (e *APPError) AppendMsg(msg string) *APPError {
	return &APPError{
		Status: e.Status,
		Code:   e.Code,
		Msg:    e.Msg + " " + msg,
		Data:   e.Data,
	}
}

// WithArgs formats according to a format specifier and returns the resulting string.
func (e *APPError) WithArgs(a ...interface{}) *APPError {
	return &APPError{
		Status: e.Status,
		Code:   e.Code,
		Msg:    fmt.Sprintf(e.Msg, a...),
		Data:   e.Data,
	}
}

// WithData append data into resp
func (e *APPError) WithData(data interface{}) *APPError {
	return &APPError{
		Status: e.Status,
		Code:   e.Code,
		Msg:    e.Msg,
		Data:   data,
	}
}

// NewAPPError new *APPError
func NewAPPError(status int, code int, msg string) *APPError {
	return &APPError{Status: status, Code: code, Msg: msg}
}
