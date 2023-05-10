package handler

import (
	"go_todolist/common/define"
	"sync"
)

var onceHandler sync.Once
var instance *Handler

type Handler struct {
}

func GetUserHandler() *Handler {
	onceHandler.Do(func() {
		instance = &Handler{}
	})
	return instance
}

func (handler *Handler) Login(request *define.Request) *define.Response {
	res := &define.Response{
		Code: 500,
	}
	return res
}
