package util

import (
	"encoding/json"
	"fmt"
	"runtime"

	"goods/common/constant"
)

type ZtError struct {
	// 错误代码
	Code int `json:"code"`
	// 错误信息
	Message string `json:"message"`
	// 错误详情
	Detail string `json:"detail"`
	// 错误调用信息
	Caller string `json:"caller"`
	// 父错误
	Parent *ZtError `json:"parent"`
}

func New(code int, err ...error) *ZtError {
	return newNbError(code, "", 2, err...)
}

func Wrap(err error) *ZtError {
	return wrap(err, 2)
}

func WrapBySkip(err error, skip int) *ZtError {
	return wrap(err, skip)
}

func wrap(err error, skip int) *ZtError {
	var caller string
	if _, file, line, ok := runtime.Caller(skip); ok {
		caller = fmt.Sprintf("%s:%d", file, line)
	}
	ztErr := new(ZtError)
	ztErr.Code = 000000
	ztErr.Message = "System exception"
	ztErr.Caller = caller
	if err == nil {
		return ztErr
	} else {
		parent, ok := err.(*ZtError)
		if ok {
			ztErr.Code = parent.Code
			ztErr.Message = parent.Message
			ztErr.Parent = parent
		} else {
			ztErr.Detail = err.Error()
		}
	}
	return ztErr
}

func newNbError(code int, message string, skip int, errs ...error) *ZtError {
	var caller string
	if _, file, line, ok := runtime.Caller(skip); ok {
		caller = fmt.Sprintf("%s:%d", file, line)
	}
	ztErr := new(ZtError)
	if message != "" {
		ztErr.Message = message
	}
	ztErr.Code = code
	ztErr.Message = new(constant.Error).ErrorMessage(code)
	ztErr.Caller = caller
	if len(errs) == 0 || errs[0] == nil {
		return ztErr
	} else {
		parent, ok := errs[0].(*ZtError)
		if ok {
			ztErr.Parent = parent
		} else {
			ztErr.Detail = errs[0].Error()
		}

	}
	return ztErr
}

func (e *ZtError) Error() string {
	m, _ := json.Marshal(e)
	return string(m)
}
