package handler

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"go_todolist/gateway/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	etcdUrl := "http://127.0.0.1:2379"
	serverName := "todolist/user"
	cli, err := clientv3.NewFromURL(etcdUrl)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	etcdResolver, err := resolver.NewBuilder(cli)
	grpcEtcdUrl := fmt.Sprintf("etcd:///%s", serverName)
	conn, err := grpc.Dial(grpcEtcdUrl, grpc.WithResolvers(etcdResolver), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := proto.NewRemoteUserClient(conn)
	return c, conn
}
