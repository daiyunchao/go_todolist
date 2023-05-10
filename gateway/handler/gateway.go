package handler

import (
	"context"
	"go_todolist/common/define"
	"sync"
)

var gatewayOnce sync.Once
var instance *Handler

type Handler struct {
}

func GetGateway() *Handler {
	gatewayOnce.Do(func() {
		instance = &Handler{}
	})
	return instance
}

func (handler *Handler) Login(ctx context.Context, request *define.Request) *define.Response {
	res := &define.Response{
		Code: 200,
	}
	userData := make(map[string]any)
	userData["nickname"] = "Tom"
	res.Data = userData
	return res
}
