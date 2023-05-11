package proto

import (
	"context"
	"fmt"
	clientV3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	UnimplementedRemoteUserServer
}

func (s *Server) GetUserInfo(ctx context.Context, request *RequestGetUserInfo) (*ResponseGetUserInfo, error) {
	return nil, nil
}

func (s *Server) GetUserInfoByNickname(ctx context.Context, request *RequestGetUserInfoByNickname) (*ResponseGetUserInfoByNickname, error) {
	fmt.Printf("in GetUserInfoByNickname\n")
	return &ResponseGetUserInfoByNickname{
		Nickname: request.Nickname,
		Id:       "10001",
		Password: "11223344",
	}, nil
}

func (s *Server) RegisterServer(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	rpcServer := grpc.NewServer()
	RegisterRemoteUserServer(rpcServer, &Server{})
	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *Server) RegisterEtcdServer(address string) error {
	etcdUrl := "http://127.0.0.1:2379"
	serverName := "todolist/user"
	var ttl int64 = 100
	etcdClient, err := clientV3.NewFromURL(etcdUrl)
	if err != nil {
		return err
	}
	em, err := endpoints.NewManager(etcdClient, serverName)
	if err != nil {
		return err
	}
	lease, _ := etcdClient.Grant(context.TODO(), ttl)
	err = em.AddEndpoint(context.TODO(), fmt.Sprintf("%s/%s", serverName, address), endpoints.Endpoint{Addr: address}, clientV3.WithLease(lease.ID))
	if err != nil {
		return err
	}
	etcdClient.KeepAlive(context.TODO(), lease.ID)
	return err
}
