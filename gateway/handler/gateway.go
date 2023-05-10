package handler

import (
	"context"
	"encoding/json"
	"go_todolist/common/define"
	"go_todolist/common/lib"
	"go_todolist/gateway/gwDefine"
	"go_todolist/gateway/proto"
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

func decodeRequest(req define.Request, dist any) {
	bytes, _ := json.Marshal(req.Data)
	json.Unmarshal(bytes, dist)
}
func (handler *Handler) Login(ctx context.Context, request *define.Request) *define.Response {
	conn, client := GetRpcServer().getUserConn()
	defer client.Close()
	req := gwDefine.RequestLogin{}
	decodeRequest(*request, &req)
	rpcReq := &proto.RequestGetUserInfoByNickname{
		Nickname: req.Nickname,
	}
	rpcRes, err := conn.GetUserInfoByNickname(ctx, rpcReq)
	if err != nil {
		return lib.CreateResponseError(define.ServerError, define.ServerErrorMsg)
	}
	return lib.CreateResponseSuccess(rpcRes)
}
