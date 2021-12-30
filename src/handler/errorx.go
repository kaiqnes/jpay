package handler

import (
	"fmt"
)

type Errorx interface {
	Message() string
	Status() int
	JSON() interface{}
	ToString() string
}

type err struct {
	status  int
	message string
}

func (e *err) Message() string {
	return e.message
}

func (e *err) Status() int {
	return e.status
}

func (e *err) JSON() interface{} {
	return &struct {
		Status  int    `json:"status_code"`
		Message string `json:"message"`
	}{
		e.status,
		e.message,
	}
}

func (e *err) ToString() string {
	return fmt.Sprintf("Status:%d | Message: %s", e.status, e.message)
}

func NewError(status int, msg string) Errorx {
	return &err{
		status:  status,
		message: msg,
	}
}
