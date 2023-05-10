package handler

import (
	"go_todolist/gateway/proto"
	"google.golang.org/grpc"
	"log"
	"sync"
)

var rpcServerOnce sync.Once
var rpcInstance *RpcServer

type RpcServer struct {
}

func GetRpcServer() *RpcServer {
	rpcServerOnce.Do(func() {
		rpcInstance = &RpcServer{}
	})
	return rpcInstance
}
func (s *RpcServer) getUserConn() (proto.RemoteUserClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(":50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := proto.NewRemoteUserClient(conn)
	return c, conn
}
