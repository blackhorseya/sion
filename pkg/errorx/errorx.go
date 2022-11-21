package errorx

import (
	"fmt"
)

// Error declare custom error
type Error struct {
	Status int         `json:"-"`
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}

func (e *Error) Error() string {
	return e.Msg
}

// ReplaceMsg replace message into Error
func (e *Error) ReplaceMsg(msg string) *Error {
	return &Error{
		Status: e.Status,
		Code:   e.Code,
		Msg:    msg,
		Data:   e.Data,
	}
}

// AppendMsg append message into Error
func (e *Error) AppendMsg(msg string) *Error {
	return &Error{
		Status: e.Status,
		Code:   e.Code,
		Msg:    e.Msg + " " + msg,
		Data:   e.Data,
	}
}

// WithArgs formats according to a format specifier and returns the resulting string.
func (e *Error) WithArgs(a ...interface{}) *Error {
	return &Error{
		Status: e.Status,
		Code:   e.Code,
		Msg:    fmt.Sprintf(e.Msg, a...),
		Data:   e.Data,
	}
}

// WithData append data into resp
func (e *Error) WithData(data interface{}) *Error {
	return &Error{
		Status: e.Status,
		Code:   e.Code,
		Msg:    e.Msg,
		Data:   data,
	}
}

// New create an app error
func New(status int, code int, msg string) *Error {
	return &Error{Status: status, Code: code, Msg: msg}
}
